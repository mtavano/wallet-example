package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mtavano/wallet-example/api/handlers/admin"
	"github.com/mtavano/wallet-example/pkg/database"
)

type handler func(*gin.Context) (interface{}, int, error)

type Server struct {
	store  *database.Store
	port   int
	router *gin.Engine
}

func NewServer(store *database.Store, port int) *Server {
	router := gin.Default()

	s := &Server{
		store:  store,
		port:   port,
		router: router,
	}
	s.routeEndpoints()

	return s
}

func (s *Server) routeEndpoints() {
	// handlers
	postUserHandler := admin.NewPostUserHandler(s.store)
	// admin routes
	s.router.POST("/admin/users", s.handleFunc(postUserHandler.Invoke))
}

func (s *Server) handleFunc(fn handler) func(*gin.Context) {
	return func(c *gin.Context) {
		payload, statusCode, err := fn(c)
		if err != nil {
			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(statusCode, payload)
	}
}

func (s *Server) Run() {
	s.router.Run(fmt.Sprintf(":%d", s.port))
}
