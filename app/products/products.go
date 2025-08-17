package products

import (
	"ahsen3d/db"
	"ahsen3d/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPhotos - Tüm photo kayıtlarını getirir
func GetProduct(c *gin.Context) {
	ctx := c.Request.Context()

	query := `
        SELECT id, url, category, title, description, uploaded_at
        FROM products
        ORDER BY uploaded_at DESC
    `
	rows, err := db.DB.Query(ctx, query)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var photos []structs.Products
	for rows.Next() {
		var p structs.Products
		if err := rows.Scan(&p.ID, &p.URL, &p.Category, &p.Title, &p.Description, &p.UploadedAt); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		photos = append(photos, p)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photos)
}

// CreatePhoto - Yeni photo kaydı oluşturur
func CreateProduct(c *gin.Context) {
	var input struct {
		URL         string  `json:"url" binding:"required,url"`
		Category    string  `json:"category" binding:"required"`
		Title       *string `json:"title"`
		Description *string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	query := `
        INSERT INTO products (url, category, title, description)
        VALUES ($1, $2, $3, $4)
        RETURNING id, uploaded_at
    `
	var photo structs.Products
	err := db.DB.QueryRow(ctx, query, input.URL, input.Category, input.Title, input.Description).Scan(&photo.ID, &photo.UploadedAt)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	photo.URL = input.URL
	photo.Category = input.Category
	photo.Title = input.Title
	photo.Description = input.Description

	c.JSON(http.StatusCreated, photo)
}

// UpdatePhoto - Var olan kaydı günceller
func UpdateProduct(c *gin.Context) {
	var input struct {
		ID          int     `json:"id" binding:"required"`
		URL         string  `json:"url" binding:"required,url"`
		Category    string  `json:"category" binding:"required"`
		Title       *string `json:"title"`
		Description *string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	query := `
        UPDATE products
        SET url = $1, category = $2, title = $3, description = $4
        WHERE id = $5
    `

	cmdTag, err := db.DB.Exec(ctx, query, input.URL, input.Category, input.Title, input.Description, input.ID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if cmdTag.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
}

// DeletePhoto - Kayıt siler
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID param is required"})
		return
	}

	ctx := c.Request.Context()

	query := `
        DELETE FROM products
        WHERE id = $1
    `

	cmdTag, err := db.DB.Exec(ctx, query, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if cmdTag.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
