package common

import (
	"github.com/corey888773/ztp-shopping-cart/src/common/data"
	"github.com/corey888773/ztp-shopping-cart/src/common/util"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/api/v1/add_to_cart"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/api/v1/get_cart"
	cartcommands "github.com/corey888773/ztp-shopping-cart/src/features/carts/commands"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/data/repository"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/external/products/service"
	cartquerries "github.com/corey888773/ztp-shopping-cart/src/features/carts/queries"
	cart "github.com/corey888773/ztp-shopping-cart/src/features/carts/queries/get_cart"
	"github.com/gin-gonic/gin"
)

type Srv struct {
	Router             *gin.Engine
	PostgresConn       *data.PostgresConnector
	CartCommandHandler cartcommands.Handler
	CartQueryHandler   cartquerries.Handler
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

	cartCommandHandler := cartcommands.NewCommandHandler(writeCartRepository, productsService)
	cartQueryHandler := cartquerries.NewQueryHandler(readCartRepository, productsService, cart.ApplyEvents)

	return &Srv{
		Router:             gin.Default(),
		PostgresConn:       postgresConn,
		CartCommandHandler: cartCommandHandler,
		CartQueryHandler:   cartQueryHandler,
	}, nil
}

func (s *Srv) SetupRouter() {
	carts := s.Router.Group("/api/v1/carts")
	carts.POST("/", add_to_cart.AddToCart(s.CartCommandHandler))
	carts.GET("/:id", get_cart.GetCart(s.CartQueryHandler))
}

func (s *Srv) Start(httpAddress string) error {
	return s.Router.Run(httpAddress)
}

func (s *Srv) Stop() {
	_ = s.PostgresConn.Close()
}
