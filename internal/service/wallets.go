package service

import (
	"context"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
)

type WalletService struct {
	walletsRepo    repository.Wallet
	operationsRepo repository.Operation
	transaction    repository.Transaction
}

func NewWalletService(
	walletsRepo repository.Wallet,
	operationsRepo repository.Operation,
	transaction repository.Transaction) *WalletService {
	return &WalletService{
		walletsRepo:    walletsRepo,
		operationsRepo: operationsRepo,
		transaction:    transaction,
	}
}

func (w *WalletService) New(ctx context.Context, wallet models.Wallet) error {
	return w.walletsRepo.Create(ctx, wallet)
}
