package fin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mtavano/wallet-example/pkg/database"
)

const userIDParam = "user_id"

type GetBalancesHandler struct {
	query GetBalancesQuery
}

type GetBalancesQuery interface {
	GetLedgerByUserID(userID string) (*database.Ledger, error)
}

func NewGetBalancesHandler(query GetBalancesQuery) *GetBalancesHandler {
	return &GetBalancesHandler{
		query: query,
	}
}

func (h *GetBalancesHandler) Invoke(c *gin.Context) (interface{}, int, error) {
	userID := c.Param(userIDParam)

	ledger, err := h.query.GetLedgerByUserID(userID)
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	return ledger, http.StatusOK, nil
}
