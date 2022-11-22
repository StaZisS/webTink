package handler

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func (h *Handler) userIndentity(c *gin.Context) {
	var access_token string
	cookie, err := c.Cookie("access_token")

	authorizationHeader := c.GetHeader("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		access_token = fields[1]
	} else if err == nil {
		access_token = cookie
	}
	if access_token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
		return
	}

	sub, err := ValidateToken(access_token, viper.GetString("publicKey.access"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.Set("userId", uuid.MustParse(fmt.Sprint(sub)))
	c.Next()
}

func getUserId(c *gin.Context) (uuid.UUID, error) {
	id, ok := c.Get("userId")
	if !ok {
		return uuid.Nil, errors.New("user id not found")
	}
	idUUID, ok := id.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("user id is not UUID")
	}
	return idUUID, nil
}

func ValidateToken(token string, publicKey string) (interface{}, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)

	if err != nil {
		return "", fmt.Errorf("validate: parse key: %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	return claims["sub"], nil
}
