package common

import (
	cartcommands "github.com/corey888773/ztp-shopping-cart/src/common/commands"
	cartquerries "github.com/corey888773/ztp-shopping-cart/src/common/queries"
	"github.com/corey888773/ztp-shopping-cart/src/common/util"
	"github.com/corey888773/ztp-shopping-cart/src/data"
	"github.com/corey888773/ztp-shopping-cart/src/data/events/repository"
	"github.com/corey888773/ztp-shopping-cart/src/external/products/service"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/v1/add_to_cart"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/v1/checkout"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/v1/get_cart"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/v1/remove_from_cart"
	"github.com/gin-gonic/gin"
)

type Srv struct {
	Router         *gin.Engine
	PostgresConn   *data.PostgresConnector
	CartCommandBus cartcommands.Handler
	CartQueryBus   cartquerries.Handler
}

func NewServer(config util.Config) (*Srv, error) {
	postgresConn, err := data.NewPostgresConnector(data.PostgresConfig{
		Host:     config.PostgresHost,
		Port:     config.PostgresPort,
		Username: config.PostgresUsername,
		Password: config.PostgresPassword,
		SSLMode:  config.PostgresSslMode,
		DBName:   config.PostgresDbName,
	})
	if err != nil {
		return nil, err
	}

	productsService := service.ProductClientMock{}

	writeCartRepository := repository.NewWriteCartRepository(postgresConn.DB)
	readCartRepository := repository.NewReadCartRepository(postgresConn.DB)

	// Queries
	cartQueryBus := cartquerries.NewQueryBus()

	getCartHandler := get_cart.NewHandler(readCartRepository, productsService, get_cart.ApplyEvents)
	cartQueryBus.Register(&get_cart.Query{}, getCartHandler)

	// Commands
	cartCommandBus := cartcommands.NewCommandBus()

	addToCartHandler := add_to_cart.NewHandler(writeCartRepository, productsService)
	removeFromCartHandler := remove_from_cart.NewHandler(writeCartRepository, productsService)

	cartCommandBus.Register(&add_to_cart.Command{}, addToCartHandler)
	cartCommandBus.Register(&remove_from_cart.Command{}, removeFromCartHandler)

	return &Srv{
		Router:         gin.Default(),
		PostgresConn:   postgresConn,
		CartCommandBus: cartCommandBus,
		CartQueryBus:   cartQueryBus,
	}, nil
}

func (s *Srv) SetupRouter() {
	carts := s.Router.Group("/api/v1/carts")
	carts.POST("/", add_to_cart.AddToCart(s.CartCommandBus))
	carts.DELETE("/", remove_from_cart.RemoveFromCart(s.CartCommandBus))
	carts.GET("/:id", get_cart.GetCart(s.CartQueryBus))
	carts.POST("/checkout", checkout.Checkout(s.CartCommandBus))
}

func (s *Srv) Start(httpAddress string) error {
	return s.Router.Run(httpAddress)
}

func (s *Srv) Stop() {
	_ = s.PostgresConn.Close()
}
