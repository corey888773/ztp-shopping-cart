package common

import (
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/commands"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/queries"
	"github.com/corey888773/ztp-shopping-cart/products-api/data"
	"github.com/corey888773/ztp-shopping-cart/products-api/data/products"
	"github.com/corey888773/ztp-shopping-cart/products-api/data/products/repository"
	"github.com/corey888773/ztp-shopping-cart/products-api/features/v1/get_products"
	"github.com/corey888773/ztp-shopping-cart/products-api/features/v1/lock_product"
	"github.com/corey888773/ztp-shopping-cart/products-api/features/v1/unlock_product"
	"github.com/gin-gonic/gin"
)

type Config struct {
	ServerPort       any    `mapstructure:"SERVER_PORT"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUsername string `mapstructure:"POSTGRES_USERNAME"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresSslMode  string `mapstructure:"POSTGRES_SSL_MODE"`
	PostgresDbName   string `mapstructure:"POSTGRES_DB_NAME"`
}

type Srv struct {
	Router     *gin.Engine
	CommandBus commands.Handler
	QueryBus   queries.Handler
}

func NewServer(config Config) (*Srv, error) {
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

	err = products.InitDbWithMockProducts(postgresConn.DB)
	if err != nil {
		return nil, err
	}

	unitOfWork := data.NewUnitOfWork(postgresConn.DB)
	productsReadRepository := repository.NewReadProductsRepository(postgresConn.DB)
	productsWriteRepository := repository.NewWriteProductsRepository(postgresConn.DB)

	commandBus := commands.NewCommandBus()
	queryBus := queries.NewQueryBus()

	getProductsQueryHandler := get_products.NewHandler(productsReadRepository)
	queryBus.Register(&get_products.Query{}, getProductsQueryHandler)

	lockProductCommandHandler := lock_product.NewHandler(unitOfWork, productsWriteRepository, productsReadRepository)
	commandBus.Register(&lock_product.Command{}, lockProductCommandHandler)

	unlockProductCommandHandler := unlock_product.NewHandler(unitOfWork, productsWriteRepository, productsReadRepository)
	commandBus.Register(&unlock_product.Command{}, unlockProductCommandHandler)

	return &Srv{
		Router:     gin.Default(),
		CommandBus: commandBus,
		QueryBus:   queryBus,
	}, nil
}

func (s *Srv) SetupRouter() {
	productRoutes := s.Router.Group("api/v1/products")
	productRoutes.POST("/", get_products.GetProducts(s.QueryBus))
	productRoutes.POST("/lock/:id", lock_product.LockProduct(s.CommandBus))
	productRoutes.POST("/unlock/:id", unlock_product.UnlockProduct(s.CommandBus))
}

func (s *Srv) Start(httpAddress string) error {
	return s.Router.Run(httpAddress)
}

func (s *Srv) Stop() {

}
