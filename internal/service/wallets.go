package service

import (
	"context"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
)

type WalletService struct {
	walletsRepo repository.Wallet
}

func NewWalletService(walletsRepo repository.Wallet) *WalletService {
	return &WalletService{
		walletsRepo: walletsRepo,
	}
}

func (w *WalletService) New(ctx context.Context, wallet models.Wallet) error {
	return w.walletsRepo.Create(ctx, wallet)
}

func (w *WalletService) NewOperation(ctx context.Context, operation models.Operation) error {
	return w.walletsRepo.NewOperation(ctx, operation)
}
