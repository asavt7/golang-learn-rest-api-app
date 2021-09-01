package services

import (
	"crypto/sha1"
	"fmt"
	"github.com/asavt7/todo/pkg/domain"
	"github.com/asavt7/todo/pkg/repos"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

const salt = "sadqwevsdvweqwcqwd"
const tokenTimeToLive = 12 * time.Hour
var signingKey = []byte("saasasasasdqwdasdqw")

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

type AuthService struct {
	repo repos.Authorization
}

func (a *AuthService) GenerateToken(username string, password string) (string, error) {

	user, err := a.repo.GetUser(username, a.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTimeToLive).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: strconv.Itoa(user.Id),
	})
	return token.SigningString()
}

func NewAuthService(repo repos.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = a.generatePasswordHash(user.Password)

	return a.repo.CreateUser(user)
}

func (a *AuthService) generatePasswordHash(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
