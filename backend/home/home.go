package home

import (
	"ahsen3d/db"
	"ahsen3d/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHomes - Tüm home_contents kayıtlarını getirir
func GetHomes(c *gin.Context) {
	ctx := c.Request.Context()

	query := `
        SELECT id, title, description, created_at, updated_at
        FROM home_contents
        ORDER BY created_at DESC
    `
	rows, err := db.DB.Query(ctx, query)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var homes []structs.Home
	for rows.Next() {
		var h structs.Home
		if err := rows.Scan(&h.ID, &h.Title, &h.Description, &h.CreatedAt, &h.UpdatedAt); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		homes = append(homes, h)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, homes)
}

// CreateHome - Yeni home_contents kaydı oluşturur
func CreateHome(c *gin.Context) {
	var input struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	query := `
        INSERT INTO home_contents (title, description)
        VALUES ($1, $2)
        RETURNING id, created_at, updated_at
    `
	var home structs.Home
	err := db.DB.QueryRow(ctx, query, input.Title, input.Description).Scan(&home.ID, &home.CreatedAt, &home.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	home.Title = input.Title
	home.Description = input.Description

	c.JSON(http.StatusCreated, home)
}

// UpdateHome - Var olan kaydı günceller
func UpdateHome(c *gin.Context) {
	var input struct {
		ID          int    `json:"id" binding:"required"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	query := `
        UPDATE home_contents
        SET title = $1, description = $2, updated_at = NOW()
        WHERE id = $3
    `

	cmdTag, err := db.DB.Exec(ctx, query, input.Title, input.Description, input.ID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if cmdTag.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Home content not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Home content updated successfully"})
}

// DeleteHome - Kayıt siler
func DeleteHome(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID param is required"})
		return
	}

	ctx := c.Request.Context()

	query := `
        DELETE FROM home_contents
        WHERE id = $1
    `

	cmdTag, err := db.DB.Exec(ctx, query, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if cmdTag.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Home content not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Home content deleted successfully"})
}
