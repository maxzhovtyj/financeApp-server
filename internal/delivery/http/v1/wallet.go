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
		wallet.GET("/all", h.getAllWallets)
		wallet.POST("", h.newWallet)

		wallet.POST("/operation", h.newOperation)
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

func (h *Handler) getAllWallets(ctx echo.Context) error {
	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		return newErrorResponse(ctx, http.StatusBadRequest, models.ErrInvalidInputBody)
	}

	userOid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return newErrorResponse(ctx, http.StatusBadRequest, models.ErrInvalidInputBody)
	}

	wallets, err := h.service.Wallet.GetAll(ctx.Request().Context(), userOid)
	if err != nil {
		return newErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, wallets)
}

type NewOperationInput struct {
	WalletId    string `json:"walletId" bson:"walletId" create:"required"`
	Income      bool   `json:"income" bson:"income" create:"required"`
	Sum         string `json:"sum" bson:"sum" create:"required"`
	Description string `json:"description,omitempty" bson:"description"`
}

func (h *Handler) newOperation(ctx echo.Context) error {
	var input NewOperationInput

	if err := BindAndValidate(ctx, &input); err != nil {
		return newErrorResponse(ctx, http.StatusBadRequest, models.ErrInvalidInputBody)
	}

	walletOid, err := primitive.ObjectIDFromHex(input.WalletId)
	if err != nil {
		return newErrorResponse(ctx, http.StatusBadRequest, models.ErrInvalidInputBody)
	}

	if input.Income == false {
		input.Sum = "-" + input.Sum
	}

	sumDecimal128, err := primitive.ParseDecimal128(input.Sum)
	if err != nil {
		return newErrorResponse(ctx, http.StatusBadRequest, models.ErrInvalidInputBody)
	}

	err = h.service.Wallet.NewOperation(ctx.Request().Context(), models.Operation{
		Income:      input.Income,
		WalletId:    walletOid,
		Description: input.Description,
		Sum:         sumDecimal128,
	})
	if err != nil {
		return newErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return nil
}
