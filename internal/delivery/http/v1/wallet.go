package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const walletUrl = "/wallet"

func (h *Handler) initWalletRoutes(group *echo.Group) {
	wallet := group.Group(walletUrl, h.userIdentity)
	{
		wallet.POST("", h.newWallet)
	}
}

type NewWalletInput struct {
	Name   string `json:"name" create:"required"`
	UserId string `json:"userId" create:"required"`
	Sum    string `json:"sum" create:"required"`
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

	err = h.service.Wallet.New(ctx.Request().Context(), models.Wallet{
		Name:   input.Name,
		UserId: userOid,
		Sum:    sumDecimal128,
	})
	if err != nil {
		return newErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return nil
}
