package fin

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mtavano/wallet-example/pkg/database"
)

type PostSellCryptoHandler struct {
	query  PostSellCryptoHandlerQuery
	crypto CryptoCurrencyWrapper
}

type PostSellCryptoHandlerQuery interface {
	SellCrypto(input *database.SellCryptoInput) (int64, int64, error)
	GetLedgerByUserID(userID string) (*database.Ledger, error)
}

func NewPostSellCryptoHandler(
	query PostSellCryptoHandlerQuery,
	crypto CryptoCurrencyWrapper,
) *PostSellCryptoHandler {
	return &PostSellCryptoHandler{
		query:  query,
		crypto: crypto,
	}
}

type PostSellCryptoRequest struct {
	UserID   string `json:"user_id"`
	Amount   uint64 `json:"amount"`
	Currency string `json:"currency"`
}

type PostSellCryptoResponse struct {
	UserID      string `json:"user_id"`
	NewBalances map[string]int64
}

func (h *PostSellCryptoHandler) Invoke(c *gin.Context) (interface{}, int, error) {
	var req PostSellCryptoRequest
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

	fiatBalance, cryptoBalance, err := h.query.SellCrypto(&database.SellCryptoInput{
		UserID:   req.UserID,
		Currency: crypto,
		Amount:   req.Amount,
		Price:    h.crypto.GetPrice(cryptoStr),
	})

	return &PostSellCryptoResponse{
		UserID: req.UserID,
		NewBalances: map[string]int64{
			currencyUSDC: fiatBalance,
			cryptoStr:    cryptoBalance,
		},
	}, http.StatusCreated, nil
}

func (h *PostSellCryptoHandler) getCryptoCurrency(currency string) (database.Currency, error) {
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
