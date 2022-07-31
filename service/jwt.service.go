package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateToken(userId string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService() JwtService {
	return &jwtService{
		issuer:    "appgo",
		secretKey: GetSecretKey(),
	}
}

func GetSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "verylongsecretkey"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(userId string) string {
	// Define expire time for 1 month
	today := time.Now()
	expiresAt := today.AddDate(0, 1, 0).Unix()

	claims := &jwtCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    j.issuer,
			IssuedAt:  today.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secretKey))

	if err != nil {
		log.Fatalf("Failed generate token %v", err)
	}

	return signedToken
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(_t *jwt.Token) (interface{}, error) {
		if _, ok := _t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", _t.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})
}
