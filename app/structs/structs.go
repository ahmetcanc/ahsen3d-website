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

type Products struct {
	ID          int       `json:"id"`
	URL         string    `json:"url"`
	Category    string    `json:"category"`
	Title       *string   `json:"title,omitempty"`
	Description *string   `json:"description,omitempty"`
	Product_Url *string   `json:"product_url,omitempty"`
	UploadedAt  time.Time `json:"uploaded_at"`
}

type Home struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
