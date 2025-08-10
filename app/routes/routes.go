package routes

import (
	"ahsen3d/admin"
	"ahsen3d/contact"
	"ahsen3d/db"
	"ahsen3d/home"
	"ahsen3d/photos"
	"ahsen3d/products"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	api := router.Group("/api")
	adminapi := router.Group("/saas")

	authHandler := admin.NewAuthHandler(db.DB)

	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)

	// Contact routes
	api.GET("/contact", contact.GetContacts)
	adminapi.POST("/contact", contact.CreateContact)
	adminapi.PUT("/contact", contact.UpdateContact)
	adminapi.DELETE("/contact/:id", contact.DeleteContact)

	// Home routes
	api.GET("/home", home.GetHomes)
	adminapi.POST("/home", home.CreateHome)
	adminapi.PUT("/home", home.UpdateHome)
	adminapi.DELETE("/home/:id", home.DeleteHome)

	// Photos routes
	api.GET("/photos", photos.GetPhotos)
	adminapi.POST("/photos", photos.CreatePhoto)
	adminapi.PUT("/photos", photos.UpdatePhoto)
	adminapi.DELETE("/photos/:id", photos.DeletePhoto)

	// Products routes
	api.GET("/products", products.GetProducts)
	adminapi.POST("/products", products.CreateProduct)
	adminapi.PUT("/products", products.UpdateProduct)
	adminapi.DELETE("/products/:id", products.DeleteProduct)
}
