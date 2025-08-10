package products

import (
	"ahsen3d/db"
	"ahsen3d/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProducts - tüm ürünleri getirir
func GetProducts(c *gin.Context) {
	ctx := c.Request.Context()

	query := `
        SELECT id, name, description, price, stock, product_url, created_at, updated_at
        FROM products
        ORDER BY created_at DESC
    `
	rows, err := db.DB.Query(ctx, query)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []structs.Product
	for rows.Next() {
		var p structs.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.ProductURL, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// CreateProduct - yeni ürün ekler
func CreateProduct(c *gin.Context) {
	var input struct {
		Name        string  `json:"name" binding:"required"`
		Description *string `json:"description"`
		Price       float64 `json:"price" binding:"required,gt=0"`
		Stock       int     `json:"stock" binding:"required,gte=0"`
		ProductURL  string  `json:"product_url" binding:"required,url"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	query := `
        INSERT INTO products (name, description, price, stock, product_url)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, updated_at
    `
	var p structs.Product
	err := db.DB.QueryRow(ctx, query, input.Name, input.Description, input.Price, input.Stock, input.ProductURL).
		Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	p.Name = input.Name
	p.Description = input.Description
	p.Price = input.Price
	p.Stock = input.Stock
	p.ProductURL = input.ProductURL

	c.JSON(http.StatusCreated, p)
}

// UpdateProduct - mevcut ürünü günceller
func UpdateProduct(c *gin.Context) {
	var input struct {
		ID          int     `json:"id" binding:"required"`
		Name        string  `json:"name" binding:"required"`
		Description *string `json:"description"`
		Price       float64 `json:"price" binding:"required,gt=0"`
		Stock       int     `json:"stock" binding:"required,gte=0"`
		ProductURL  string  `json:"product_url" binding:"required,url"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	query := `
        UPDATE products
        SET name = $1, description = $2, price = $3, stock = $4, product_url = $5, updated_at = NOW()
        WHERE id = $6
    `
	cmdTag, err := db.DB.Exec(ctx, query, input.Name, input.Description, input.Price, input.Stock, input.ProductURL, input.ID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if cmdTag.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// DeleteProduct - ürün siler
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
