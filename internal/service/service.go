package service

import (
	"context"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
	"github.com/maxzhovtyj/financeApp-server/pkg/auth"
	"github.com/maxzhovtyj/financeApp-server/pkg/hash"
	"time"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Users interface {
	SignUp(ctx context.Context, user models.User) error
	SignIn(ctx context.Context, email, password string) (string, string, error)
}

type Service struct {
	Users Users
}

func New(
	repo *repository.Repository,
	tokenManager auth.TokenManager,
	accessTokenTTL, refreshTokenTTL time.Duration,
	hashing hash.PasswordHashing) *Service {
	return &Service{
		Users: NewUsersService(repo.Users, tokenManager, accessTokenTTL, refreshTokenTTL, hashing),
	}
}
