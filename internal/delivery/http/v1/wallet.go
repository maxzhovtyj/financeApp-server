package v1

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type NewWalletInput struct {
	UserId string `json:"userId"`
	Sum    string `json:"sum"`
}

func (h *Handler) newWallet(ctx echo.Context) error {
	var input NewWalletInput

	err := BindAndValidate(ctx, &input)
	if err != nil {
		return newErrorResponse(ctx, http.StatusBadRequest, models.ErrInvalidInputBody)
	}

	userOid, err := primitive.ObjectIDFromHex(input.UserId)
	if err != nil {
		return newErrorResponse(ctx, http.StatusBadRequest, models.ErrInvalidInputBody)
	}

	sumDecimal128, err := primitive.ParseDecimal128(input.Sum)
	if err != nil {
		return newErrorResponse(ctx, http.StatusBadRequest, models.ErrInvalidInputBody)
	}

	err = h.service.Wallet.New(context.Background(), models.Wallet{
		UserId: userOid,
		Sum:    sumDecimal128,
	})

	if err != nil {
		return newErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return nil
}
