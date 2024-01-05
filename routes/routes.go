// routes/routes.go
package routes

import (
	"net/http"
	"strconv"

	"myapp/database"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the API routes.
func SetupRoutes(router *gin.Engine, db *database.DB) {
	// API routes
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
}
