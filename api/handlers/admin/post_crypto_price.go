package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostyCryptoPriceHandler struct {
	crypto PostCryptoCurrencyWrapper
}

type PostCryptoCurrencyWrapper interface {
	SetPrice(currency string, price uint64)
}

func NewPostCryptoPriceHandler(crypto PostCryptoCurrencyWrapper) *PostyCryptoPriceHandler {
	return &PostyCryptoPriceHandler{
		crypto: crypto,
	}
}

type PostyCryptoPriceRequest struct {
	Currency string `json:"currency"`
	Price    uint64 `json:"price"`
}

func (h *PostyCryptoPriceHandler) Invoke(c *gin.Context) (interface{}, int, error) {
	var req PostyCryptoPriceRequest
	err := c.BindJSON(&req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	if req.Currency == "" {
		return nil, http.StatusBadRequest, ErrCurrencyNotSupported
	}

	h.crypto.SetPrice(req.Currency, req.Price)
	return nil, http.StatusOK, nil
}
