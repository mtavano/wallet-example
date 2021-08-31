package fin

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mtavano/wallet-example/pkg/database"
)

type PostTransferCurrencyHandler struct {
	query PostTransferCurrencyQuery
}

func NewPostTransferCurrencyHandler(query PostTransferCurrencyQuery) *PostTransferCurrencyHandler {
	return &PostTransferCurrencyHandler{
		query: query,
	}
}

type PostTransferCurrencyQuery interface {
	GetLedgerByUserID(userID string) (*database.Ledger, error)
	CreateTransfer(input *database.CreateTransferInput) (int64, int64, error)
}

type PostTransferCurrencyRequest struct {
	SourceID      string `json:"source_id"`
	DestinationID string `json:"destination_id"`
	Amount        uint64 `json:"amount"`
	Currency      string `json:"currency"`
}

type PostTransferCurrencyResponse struct {
	Source      *TransferDetails `json:"source"`
	Destination *TransferDetails `json:"destination"`
}

type TransferDetails struct {
	UserID     string `json:"user_id"`
	NewBalance int64  `json:"new_balance"`
	Currency   string `json:"currency"`
}

func (h *PostTransferCurrencyHandler) Invoke(c *gin.Context) (interface{}, int, error) {
	var req PostTransferCurrencyRequest
	err := c.BindJSON(&req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	currency, err := h.getCurrencyToTransfer(req.Currency)
	if err != nil {
		return nil, http.StatusBadRequest, ErrCurrencyNotSupported
	}

	newSrcBalance, newDestBalance, err := h.query.CreateTransfer(&database.CreateTransferInput{
		SourceID:      req.SourceID,
		DestinationID: req.DestinationID,
		Amount:        req.Amount,
		Currency:      currency,
	})
	if err != nil && err.Error() == database.ErrInsuficientFunds.Error() {
		return nil, http.StatusPaymentRequired, err
	}
	if err != nil && err.Error() == database.ErrUserNotFound.Error() {
		return nil, http.StatusNotFound, err
	}
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &PostTransferCurrencyResponse{
		Source: &TransferDetails{
			UserID:     req.SourceID,
			NewBalance: newSrcBalance,
		},
		Destination: &TransferDetails{
			UserID:     req.DestinationID,
			NewBalance: newDestBalance,
		},
	}, http.StatusCreated, nil
}

func (h *PostTransferCurrencyHandler) getCurrencyToTransfer(currency string) (database.Currency, error) {
	var err error
	c := strings.ToUpper(currency)

	switch c {
	case currencyBTC:
		return database.CurrencyBTC, nil
	case currencyETH:
		return database.CurrencyETH, nil
	case currencyUSDC:
		return database.CurrencyUSDC, nil
	}
	return "", err
}
