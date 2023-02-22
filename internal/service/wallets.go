package service

import (
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
)

type WalletService struct {
	walletsRepo    repository.Wallet
	operationsRepo repository.Operation
	transaction    repository.Transaction
}

func NewWalletsService(
	walletsRepo repository.Wallet,
	operationsRepo repository.Operation,
	transaction repository.Transaction) *WalletService {
	return &WalletService{
		walletsRepo:    walletsRepo,
		operationsRepo: operationsRepo,
		transaction:    transaction,
	}
}
