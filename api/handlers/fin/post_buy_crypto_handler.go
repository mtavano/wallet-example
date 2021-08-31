package fin

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mtavano/wallet-example/pkg/database"
)

var ErrCurrencyNotSupported = errors.New("currency not supported")

type PostBuyCryptoHandler struct {
	query  PostBuyCryptoHandlerQuery
	crypto CryptoCurrencyWrapper
}

type PostBuyCryptoHandlerQuery interface {
	BuyCrypto(input *database.BuyCryptoInput) (int64, int64, error)
	GetLedgerByUserID(userID string) (*database.Ledger, error)
}

type CryptoCurrencyWrapper interface {
	GetPrice(string) uint64
}

func NewPostBuyCryptoHandler(
	query PostBuyCryptoHandlerQuery,
	crypto CryptoCurrencyWrapper,
) *PostBuyCryptoHandler {
	return &PostBuyCryptoHandler{
		query:  query,
		crypto: crypto,
	}
}

type PostBuyCryptoRequest struct {
	UserID   string `json:"user_id"`
	Amount   uint64 `json:"amount"`
	Currency string `json:"currency"`
}

type PostBuyCryptoResponse struct {
	UserID      string `json:"user_id"`
	NewBalances map[string]int64
}

func (h *PostBuyCryptoHandler) Invoke(c *gin.Context) (interface{}, int, error) {
	var req PostBuyCryptoRequest
	err := c.BindJSON(&req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	crypto, err := h.getCryptoCurrency(req.Currency)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	cryptoStr := crypto.String()

	_, err = h.query.GetLedgerByUserID(req.UserID)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	fiatBalance, cryptoBalance, err := h.query.BuyCrypto(&database.BuyCryptoInput{
		UserID:   req.UserID,
		Currency: crypto,
		Amount:   req.Amount,
		Price:    h.crypto.GetPrice(cryptoStr),
	})

	return &PostBuyCryptoResponse{
		UserID: req.UserID,
		NewBalances: map[string]int64{
			currencyUSDC: fiatBalance,
			cryptoStr:    cryptoBalance,
		},
	}, http.StatusCreated, nil
}

func (h *PostBuyCryptoHandler) getCryptoCurrency(currency string) (database.Currency, error) {
	var err error
	c := strings.ToUpper(currency)

	switch c {
	case currencyBTC:
		return database.CurrencyBTC, nil
	case currencyETH:
		return database.CurrencyETH, nil
	default:
		err = ErrCurrencyNotSupported
	}
	return "", err
}
