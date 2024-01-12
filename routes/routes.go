// routes/routes.go
package routes

import (
	"myapp/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *database.DB, jwtSecret string) {
	router.POST("/login", func(c *gin.Context) {
		var loginUser database.User
		if err := c.ShouldBindJSON(&loginUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		// Authenticate user
		existingUser, err := db.GetUserByUsername(loginUser.Username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
	
		// Check password
		if !database.CheckPasswordHash(loginUser.Password, existingUser.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
	
		// Generate JWT token
		token, err := database.GenerateJWTToken(existingUser.Id, jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{"token": token, "message": "Login successful"})
	})
	

	router.GET("/users/:id", func(c *gin.Context) {
		userID := c.Params.ByName("id")
		id, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		user, err := db.GetUserByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	router.POST("/users", func(c *gin.Context) {
		var newUser database.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		userID, err := db.CreateUser(newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	
		c.JSON(http.StatusCreated, gin.H{"id": userID, "message": "User created successfully"})
	})
	
}
