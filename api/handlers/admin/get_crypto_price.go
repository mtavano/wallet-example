package admin

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const currencyParam = "currency"

var ErrCurrencyNotSupported = errors.New("currency not supported")

type GetyCryptoPriceHandler struct {
	crypto GetCryptoCurrencyWrapper
}

type GetCryptoCurrencyWrapper interface {
	GetPrice(currency string) uint64
}

func NewGetCryptoPriceHandler(crypto GetCryptoCurrencyWrapper) *GetyCryptoPriceHandler {
	return &GetyCryptoPriceHandler{
		crypto: crypto,
	}
}

type GetCyCryptoPriceResponse struct {
	Currency string `json:"currency"`
	Price    uint64 `json:"price"`
}

func (h *GetyCryptoPriceHandler) Invoke(c *gin.Context) (interface{}, int, error) {
	currency := c.Param(currencyParam)
	if currency == "" {
		return nil, http.StatusBadRequest, ErrCurrencyNotSupported
	}

	price := h.crypto.GetPrice(currency)
	return &GetCyCryptoPriceResponse{
		Currency: currency,
		Price:    price,
	}, http.StatusOK, nil
}

