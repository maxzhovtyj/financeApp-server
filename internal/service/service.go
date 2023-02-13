package service

import "github.com/maxzhovtyj/financeApp-server/internal/repository"

type Users interface {
}

type Service struct {
	Users Users
}

func New(repo *repository.Repository) *Service {
	return &Service{Users: NewUsersService(repo.Users)}
}
