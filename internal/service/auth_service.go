package service

import (
	"crypto/rand"
	"equi_genea_auth_service/config"
	"equi_genea_auth_service/internal/models"
	"errors"
	"math/big"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultJWTExpireDuration = time.Hour * 24
	passwordLength           = 16
	chars                    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var ErrTokenInvalid = errors.New("invalid token")

type AuthService struct {
	secretKey string
}

func NewAuthService(cfg *config.AuthConfig) *AuthService {
	return &AuthService{secretKey: cfg.SecretKey}
}

func (s *AuthService) GeneratePassword() (string, error) {
	password := make([]byte, passwordLength)
	for i := 0; i < passwordLength; i++ {
		indexBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		password[i] = chars[indexBig.Int64()]
	}

	return string(password), nil
}

func (s *AuthService) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (s *AuthService) ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

func (s *AuthService) GenerateJWT(accountID string) (string, error) {
	claims := models.AuthClaims{
		AccountID: accountID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(DefaultJWTExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *AuthService) GetClaims(tokenString string) (*models.AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.AuthClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}
