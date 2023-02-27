package service

import (
	"context"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WalletService struct {
	walletsRepo repository.Wallet
}

func NewWalletService(walletsRepo repository.Wallet) *WalletService {
	return &WalletService{
		walletsRepo: walletsRepo,
	}
}

func (w *WalletService) GetAll(ctx context.Context, userOid primitive.ObjectID) ([]models.Wallet, error) {
	return w.walletsRepo.GetAllWallets(ctx, userOid)
}

func (w *WalletService) Get(ctx context.Context, walletOid primitive.ObjectID) (models.Wallet, []models.Operation, error) {
	return w.walletsRepo.GetWallet(ctx, walletOid)
}

func (w *WalletService) New(ctx context.Context, wallet models.Wallet) error {
	return w.walletsRepo.Create(ctx, wallet)
}

func (w *WalletService) NewOperation(ctx context.Context, operation models.Operation) error {
	return w.walletsRepo.NewOperation(ctx, operation)
}
