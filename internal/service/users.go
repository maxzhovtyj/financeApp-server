package service

import (
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) Users {
	return &UserService{repo: repo}
}

func (s *UserService) SignUp(user models.User) (primitive.ObjectID, error) {
	return s.repo.SignUp(user)
}
