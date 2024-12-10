package main

import (
	"auth-service/middleware"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var jwtSecret = []byte("your_jwt_secret_key") // Secret key for JWT signing

// InitDB connects to the database
func initDB() {
	dsn := "host=localhost user=postgres password=Admin dbname=auth_service port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&Customer{})
}

// Register a new user
func registerUser(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var existingCustomer Customer
	if err := db.Where("email = ?", customer.Email).First(&existingCustomer).Error; err == nil {
		c.JSON(400, gin.H{"error": "Email already registered"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error hashing password"})
		return
	}

	customer.PasswordHash = string(hashedPassword)

	if err := db.Create(&customer).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully"})
}

// Login route to generate JWT token
func loginUser(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind input JSON to credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Find user by email
	var user Customer
	if err := db.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare the provided password with the stored password hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(credentials.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	// Send the token in response
	c.JSON(200, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}

func main() {
	// Initialize DB connection
	initDB()

	// Initialize Gin router
	r := gin.Default()

	// Register routes
	r.POST("/register", registerUser)
	r.POST("/login", loginUser)

	// Start the server

	r.GET("/customer/:id", middleware.AuthRequired, func(c *gin.Context) {
		userID := c.MustGet("userID").(string) // Extract userID from JWT token
		customerID := c.Param("id")            // Get the ID from the URL parameter

		// Ensure the user can only view their own customer data
		if userID != customerID {
			c.JSON(403, gin.H{"error": "Forbidden"})
			return
		}

		var customer Customer
		if err := db.Where("id = ?", customerID).First(&customer).Error; err != nil {
			c.JSON(404, gin.H{"error": "Customer not found"})
			return
		}

		customer.PasswordHash = "" // Don't send the password hash
		c.JSON(200, customer)
	})
	r.Run(":8080")
}
