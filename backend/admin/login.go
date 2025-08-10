package admin

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	DB *pgxpool.Pool
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *AuthService) Register(username, password, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	_, err = a.DB.Exec(ctx, `
        INSERT INTO users (username, password_hash, email)
        VALUES ($1, $2, $3)
    `, username, hashedPassword, email)

	return err
}

func (a *AuthService) Login(username, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var storedHash string
	err := a.DB.QueryRow(ctx, `
        SELECT password_hash FROM users WHERE username = $1
    `, username).Scan(&storedHash)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("kullanıcı bulunamadı")
		}
		return err
	}

	if !checkPasswordHash(password, storedHash) {
		return errors.New("şifre hatalı")
	}

	return nil
}

type AuthHandler struct {
	Service *AuthService
}

func NewAuthHandler(db *pgxpool.Pool) *AuthHandler {
	return &AuthHandler{
		Service: &AuthService{DB: db},
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.Register(req.Username, req.Password, req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Kayıt başarılı"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.Login(req.Username, req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Giriş başarılı"})
}
