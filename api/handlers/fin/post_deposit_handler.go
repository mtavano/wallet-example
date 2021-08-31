package fin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mtavano/wallet-example/pkg/database"
)

type PostDepositHandler struct {
	query PostDepositGHandlerQuery
}

type PostDepositGHandlerQuery interface {
	CreateDeposit(input *database.CreateDepositInput) (int64, error)
	GetLedgerByUserID(userID string) (*database.Ledger, error)
}

func NewPostDepositHandler(
	query PostDepositGHandlerQuery,
) *PostDepositHandler {
	return &PostDepositHandler{
		query: query,
	}
}

type PostDepositRequest struct {
	UserID string `json:"user_id"`
	Amount uint64 `json:"amount"`
}

type PostDepositResponse struct {
	UserID     string `json:"user_id"`
	NewBalance int64  `json:"new_balance"`
	Currency   string `json:"currency"`
}

func (h *PostDepositHandler) Invoke(c *gin.Context) (interface{}, int, error) {
	var req PostDepositRequest
	err := c.BindJSON(&req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	_, err = h.query.GetLedgerByUserID(req.UserID)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	newBalance, err := h.query.CreateDeposit(&database.CreateDepositInput{
		UserID: req.UserID,
		Amount: req.Amount,
	})
	if err != nil && err.Error() == database.ErrUserNotFound.Error() {
		return nil, http.StatusNotFound, err
	}
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &PostDepositResponse{
		UserID:     req.UserID,
		NewBalance: newBalance,
		Currency:   database.CurrencyUSDC.String(),
	}, http.StatusCreated, nil
}
