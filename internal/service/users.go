package service

import (
	"context"
	"errors"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
	"github.com/maxzhovtyj/financeApp-server/pkg/auth"
	"github.com/maxzhovtyj/financeApp-server/pkg/hash"
	"time"
)

type UserService struct {
	repo         repository.Users
	tokenManager auth.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration

	hashing hash.PasswordHashing
}

func NewUsersService(
	repo repository.Users,
	tokenManager auth.TokenManager,
	accessTokenTTL, refreshTokenTTL time.Duration,
	hashing hash.PasswordHashing) *UserService {
	return &UserService{
		repo:            repo,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
		hashing:         hashing,
	}
}

func (s *UserService) SignUp(ctx context.Context, user models.User) (err error) {
	user.Password, err = s.hashing.Hash(user.Password)
	if err != nil {
		return err
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		if errors.Is(err, models.ErrUserAlreadyExists) {
			return err
		}

		return err
	}

	return nil
}

func (s *UserService) SignIn(ctx context.Context, email, password string) (string, string, error) {
	passwordHash, err := s.hashing.Hash(password)
	if err != nil {
		return "", "", err
	}

	user, err := s.repo.GetByCredentials(ctx, email, passwordHash)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			return "", "", err
		}

		return "", "", err
	}

	accessToken, err := s.tokenManager.NewJWT(user.Id.String(), s.accessTokenTTL)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
