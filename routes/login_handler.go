package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("my_secret_key")

func (r *Router) LoginHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := r.LoginService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// ---- Generate JWT token ----
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
		"exp":   jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
	})

	tokenString, _ := token.SignedString(SecretKey)

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   tokenString,  // ⭐ The important part
		"user": gin.H{
			"id":       user.ID,
			"username": user.Name,
			"email":    user.Email,
		},
	})
}
