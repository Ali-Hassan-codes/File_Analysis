package routes

import (
	"net/http"

	"github.com/ali-hassan-Codes/file_analyzer_2/models"
	"github.com/gin-gonic/gin"
)

// SignupHandler handles user signup
func (r *Router) SignupHandler(c *gin.Context) {
	var user models.User

	// Bind JSON input
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Call UserService Signup method
	createdUser, err := r.UserService.Signup(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signup successful",
		"user": gin.H{
			"id":    createdUser.ID,
			"name":  createdUser.Name,
			"email": createdUser.Email,
		},
	})
}
