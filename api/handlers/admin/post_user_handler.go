package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mtavano/wallet-example/pkg/database"
)

type PostUserHandler struct {
	query CreateUserQuery
}

func NewPostUserHandler(query CreateUserQuery) *PostUserHandler {
	return &PostUserHandler{
		query: query,
	}
}

type CreateUserQuery interface {
	CreateUser(userid string) error
}

type PostUserRequest struct {
	UserID string `json:"user_id"`
}

func (h *PostUserHandler) Invoke(c *gin.Context) (interface{}, int, error) {
	var req PostUserRequest
	err := c.BindJSON(&req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = h.query.CreateUser(req.UserID)
	if err != nil && err.Error() == database.ErrUserAlreadyExist.Error() {
		return nil, http.StatusBadRequest, err
	}
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return nil, http.StatusCreated, nil
}
