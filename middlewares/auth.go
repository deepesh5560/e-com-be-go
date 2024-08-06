package middlewares

import (
	initializers "e-com-be-go/initilizers"
	"e-com-be-go/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key")

func CreateToken(email string) (string, error) {
	if len(secretKey) == 0 {
		return "", fmt.Errorf("secret key is not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
		"role":  "admin",
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error signing token:", err) // Debug print
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"]
	var user models.User
	initializers.DB.First(&user, "Email = ?", email)

	role := claims["role"]

	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	c.Set("role", role)

	c.Next()
}
