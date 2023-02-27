package service

import (
	"context"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
	"github.com/maxzhovtyj/financeApp-server/pkg/auth"
	"github.com/maxzhovtyj/financeApp-server/pkg/hash"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Users interface {
	SignUp(ctx context.Context, user models.User) error
	SignIn(ctx context.Context, email, password string) (string, string, error)
}

type Wallet interface {
	New(ctx context.Context, wallet models.Wallet) error
	GetAll(ctx context.Context, userOid primitive.ObjectID) ([]models.Wallet, error)
	Get(ctx context.Context, walletOid primitive.ObjectID) (models.Wallet, []models.Operation, error)
	NewOperation(ctx context.Context, operation models.Operation) error
}

type Service struct {
	Users  Users
	Wallet Wallet
}

func New(
	repo *repository.Repository,
	tokenManager auth.TokenManager,
	accessTokenTTL, refreshTokenTTL time.Duration,
	hashing hash.PasswordHashing) *Service {
	return &Service{
		Users:  NewUsersService(repo.Users, tokenManager, accessTokenTTL, refreshTokenTTL, hashing),
		Wallet: NewWalletService(repo.Wallet),
	}
}
