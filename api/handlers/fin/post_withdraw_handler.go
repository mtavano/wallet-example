package fin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mtavano/wallet-example/pkg/database"
)

type PostWithdrawHandler struct {
	query PostWithdrawGHandlerQuery
}

type PostWithdrawGHandlerQuery interface {
	CreateWithdraw(input *database.CreateWithdrawInput) (int64, error)
	GetLedgerByUserID(userID string) (*database.Ledger, error)
}

func NewPostWithdrawHandler(
	query PostWithdrawGHandlerQuery,
) *PostWithdrawHandler {
	return &PostWithdrawHandler{
		query: query,
	}
}

type PostWithdrawRequest struct {
	UserID string `json:"user_id"`
	Amount uint64 `json:"amount"`
}

type PostWithdrawResponse struct {
	UserID     string `json:"user_id"`
	NewBalance int64  `json:"new_balance"`
	Currency   string `json:"currency"`
}

func (h *PostWithdrawHandler) Invoke(c *gin.Context) (interface{}, int, error) {
	var req PostWithdrawRequest
	err := c.BindJSON(&req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	_, err = h.query.GetLedgerByUserID(req.UserID)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	newBalance, err := h.query.CreateWithdraw(&database.CreateWithdrawInput{
		UserID: req.UserID,
		Amount: req.Amount,
	})
	if err != nil && err.Error() == database.ErrInsuficientFunds.Error() {
		return nil, http.StatusPaymentRequired, err
	}
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &PostWithdrawResponse{
		UserID:     req.UserID,
		NewBalance: newBalance,
		Currency:   database.CurrencyUSDC.String(),
	}, http.StatusCreated, nil
}

