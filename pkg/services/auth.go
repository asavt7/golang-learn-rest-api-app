package services

import (
	"crypto/sha1"
	"fmt"
	"github.com/asavt7/todo/pkg/domain"
	"github.com/asavt7/todo/pkg/repos"
)

var salt = "sadqwevsdvweqwcqwd"

type AuthService struct {
	repo repos.Authorization
}

func NewAuthService(repo repos.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password =a.generatePasswordHash(user.Password)

	return a.repo.CreateUser(user)
}

func (a *AuthService) generatePasswordHash(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
