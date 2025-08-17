package contact

import (
	"ahsen3d/db"
	"ahsen3d/structs"
	_ "embed"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/resend/resend-go/v2"
)

//go:embed email_template.html
var emailTemplate string

func GetContacts(c *gin.Context) {
	ctx := c.Request.Context()

	query := `
        SELECT id, name, email, subject, message, created_at
        FROM contacts
        ORDER BY created_at DESC
    `
	rows, err := db.DB.Query(ctx, query)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var contacts []structs.Contact
	for rows.Next() {
		var contact structs.Contact
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.Subject, &contact.Message, &contact.CreatedAt); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		contacts = append(contacts, contact)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contacts)
}
func CreateContact(c *gin.Context) {
	var input struct {
		Name    string  `json:"name" binding:"required"`
		Email   string  `json:"email" binding:"required,email"`
		Subject *string `json:"subject"`
		Message string  `json:"message" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	query := `
        INSERT INTO contacts (name, email, subject, message)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `
	var id int
	var createdAt time.Time
	err := db.DB.QueryRow(ctx, query, input.Name, input.Email, input.Subject, input.Message).Scan(&id, &createdAt)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := resend.NewClient("re_dn9o3zo5_6c6uTosjirt7K7LBqmHEUqHJ")
	htmlContent := fmt.Sprintf(string(emailTemplate), input.Name, input.Email, input.Message)

	_, err = client.Emails.SendWithContext(ctx, &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{"ahmetcanceylann81@gmail.com"},
		ReplyTo: input.Email,
		Subject: fmt.Sprintf("Yeni mesaj: %s", getSubject(input.Subject)),
		Html:    htmlContent,
	})
	if err != nil {
		fmt.Println("Resend error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         id,
		"name":       input.Name,
		"email":      input.Email,
		"subject":    input.Subject,
		"message":    input.Message,
		"created_at": createdAt,
	})
}

func getSubject(s *string) string {
	if s != nil {
		return *s
	}
	return "(No Subject)"
}

func DeleteContact(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID param is required"})
		return
	}

	ctx := c.Request.Context()

	query := `
        DELETE FROM contacts
        WHERE id = $1
    `

	cmdTag, err := db.DB.Exec(ctx, query, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if cmdTag.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})
}
