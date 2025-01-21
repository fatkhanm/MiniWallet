package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your_secret_key")

func GenerateToken(customerXid string) (string, error) {
	claims := jwt.MapClaims{
		"customer_xid": customerXid,
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ExtractUserIDFromToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		fmt.Println("missing authorization header")
		return "", errors.New("missing authorization header")
	}

	// Parse token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Printf("invalid token signing method: %v", token.Header["alg"])
			return nil, errors.New("invalid token signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		fmt.Println("invalid token")
		return "", errors.New("invalid token")
	}

	// Extract user ID from claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("invalid token claims")
		return "", errors.New("invalid token claims")
	}

	customerXid, ok := (claims["customer_xid"]).(string)
	if !ok {
		fmt.Println("customer_xid not found in token")
		return "", errors.New("customer_xid not found in token")
	}

	return customerXid, nil
}
