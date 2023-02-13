package service

import "github.com/maxzhovtyj/financeApp-server/internal/repository"

type UserService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) Users {
	return &UserService{repo: repo}
}
