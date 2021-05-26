package services

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gitlab.com/gunererd/dummy-challange/src/utils"
)

type TokenService interface {
	GenerateToken(username string) (string, error)
	ValidateToken(token string) (string, error)
}

type authCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type tokenService struct {
	secretKey string
	tokenTTL  int
}

func NewTokenService() TokenService {
	return &tokenService{
		secretKey: getSecretKey(),
		tokenTTL:  getTokenTTL(),
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func getTokenTTL() int {
	tokenTTL := os.Getenv("JWT_TOKEN_TTL")
	if tokenTTL == "" {
		tokenTTL = "5"
	}

	i32, _ := strconv.ParseInt(tokenTTL, 10, 32)
	ttl := int(i32)

	return ttl
}

func (service *tokenService) ValidateToken(givenToken string) (string, error) {

	givenTokenParts := strings.Split(givenToken, " ")

	if len(givenTokenParts) != 2 {

		err := &utils.FError{
			Code:      401,
			ErrorCode: "errors.badTokenForm",
			Message:   fmt.Sprintf("Token must be in form of 'Bearer .*'"),
		}

		return "", err
	}

	token, err := jwt.ParseWithClaims(givenTokenParts[1], &authCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(service.secretKey), nil
	})

	if claims, ok := token.Claims.(*authCustomClaims); ok && token.Valid {
		return claims.Username, nil
	}

	err = &utils.FError{
		Code:      401,
		ErrorCode: "errors.invalidToken",
		Message:   err.Error(),
	}

	return "", err
}

func (service *tokenService) GenerateToken(username string) (string, error) {
	claims := &authCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		err := &utils.FError{
			Code:      500,
			ErrorCode: "errors.tokenCouldNotBeGenerated",
			Message:   err.Error(),
		}

		return "", err
	}
	return tokenString, nil
}
