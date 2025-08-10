package contact

import (
	"ahsen3d/db"
	"ahsen3d/structs"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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

	c.JSON(http.StatusCreated, gin.H{
		"id":         id,
		"name":       input.Name,
		"email":      input.Email,
		"subject":    input.Subject,
		"message":    input.Message,
		"created_at": createdAt,
	})
}

func UpdateContact(c *gin.Context) {
	var input struct {
		ID      int     `json:"id" binding:"required"`
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
        UPDATE contacts
        SET name = $1, email = $2, subject = $3, message = $4
        WHERE id = $5
    `

	cmdTag, err := db.DB.Exec(ctx, query, input.Name, input.Email, input.Subject, input.Message, input.ID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if cmdTag.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact updated successfully"})
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
