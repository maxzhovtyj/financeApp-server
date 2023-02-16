package service

import (
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
	"github.com/maxzhovtyj/financeApp-server/pkg/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Users interface {
	SignUp(user models.User) (primitive.ObjectID, error)
	SignIn(email, password string) (string, string, error)
}

type Service struct {
	Users Users
}

func New(repo *repository.Repository, tokenManager auth.TokenManager, accessTokenTTL, refreshTokenTTL time.Duration) *Service {
	return &Service{
		Users: NewUsersService(repo.Users, tokenManager, accessTokenTTL, refreshTokenTTL),
	}
}
