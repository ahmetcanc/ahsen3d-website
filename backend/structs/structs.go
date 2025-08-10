package structs

import "time"

type Contact struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Subject   *string   `json:"subject,omitempty"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	ProductURL  string    `json:"product_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Photo struct {
	ID          int       `json:"id"`
	URL         string    `json:"url"`
	Category    string    `json:"category"`
	Title       *string   `json:"title,omitempty"`
	Description *string   `json:"description,omitempty"`
	UploadedAt  time.Time `json:"uploaded_at"`
}

type Home struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
