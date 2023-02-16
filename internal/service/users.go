package service

import (
	"fmt"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
	"github.com/maxzhovtyj/financeApp-server/pkg/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserService struct {
	repo         repository.Users
	tokenManager auth.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewUsersService(repo repository.Users, tokenManager auth.TokenManager, accessTokenTTL, refreshTokenTTL time.Duration) Users {
	return &UserService{
		repo:            repo,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (s *UserService) SignUp(user models.User) (primitive.ObjectID, error) {
	return s.repo.SignUp(user)
}

func (s *UserService) SignIn(email, password string) (string, string, error) {
	user, err := s.repo.GetUser(email)
	if err != nil {
		return "", "", err
	}

	// TODO check whether passwords are the same
	fmt.Println(password)

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
