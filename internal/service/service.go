package service

import (
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users interface {
	SignUp(user models.User) (primitive.ObjectID, error)
}

type Service struct {
	Users Users
}

func New(repo *repository.Repository) *Service {
	return &Service{Users: NewUsersService(repo.Users)}
}
