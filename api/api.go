package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mtavano/wallet-example/api/handlers/admin"
	"github.com/mtavano/wallet-example/api/handlers/fin"
	"github.com/mtavano/wallet-example/pkg/cryptocurrency"
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
	// wrappers
	cryptoWrapper := cryptocurrency.NewWrapper()

	// handlers
	postUserHandler := admin.NewPostUserHandler(s.store)
	postCryptoPriceHandler := admin.NewPostCryptoPriceHandler(cryptoWrapper)
	getCryptoPriceHandler := admin.NewGetCryptoPriceHandler(cryptoWrapper)

	postDepositHandler := fin.NewPostDepositHandler(s.store)
	postWithdrawHandler := fin.NewPostWithdrawHandler(s.store)
	postBuyCryptoHandler := fin.NewPostBuyCryptoHandler(s.store, cryptoWrapper)
	postSellCryptoHandler := fin.NewPostSellCryptoHandler(s.store, cryptoWrapper)
	postTransferCurrencyHandler := fin.NewPostTransferCurrencyHandler(s.store)
	getBalancesHandler := fin.NewGetBalancesHandler(s.store)

	// admin routes
	s.router.POST("/admin/users", s.handleFunc(postUserHandler.Invoke))
	s.router.POST("/admin/crypto-price", s.handleFunc(postCryptoPriceHandler.Invoke))
	s.router.GET("/admin/crypto-price/:currency", s.handleFunc(getCryptoPriceHandler.Invoke))

	// api routes
	s.router.POST("/api/fin/deposits", s.handleFunc(postDepositHandler.Invoke))
	s.router.POST("/api/fin/withdrawals", s.handleFunc(postWithdrawHandler.Invoke))
	s.router.POST("/api/fin/buy-crypto", s.handleFunc(postBuyCryptoHandler.Invoke))
	s.router.POST("/api/fin/sell-crypto", s.handleFunc(postSellCryptoHandler.Invoke))
	s.router.POST("/api/fin/transfer-currency", s.handleFunc(postTransferCurrencyHandler.Invoke))
	s.router.GET("/api/fin/balances/:user_id", s.handleFunc(getBalancesHandler.Invoke))
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
