package common

import (
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/commands"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/queries"
	data2 "github.com/corey888773/ztp-shopping-cart/products-api/src/data"
	"github.com/corey888773/ztp-shopping-cart/products-api/src/data/products"
	repository2 "github.com/corey888773/ztp-shopping-cart/products-api/src/data/products/repository"
	checkout2 "github.com/corey888773/ztp-shopping-cart/products-api/src/features/v1/checkout"
	get_all_products2 "github.com/corey888773/ztp-shopping-cart/products-api/src/features/v1/get_all_products"
	get_products2 "github.com/corey888773/ztp-shopping-cart/products-api/src/features/v1/get_products"
	lock_product2 "github.com/corey888773/ztp-shopping-cart/products-api/src/features/v1/lock_product"
	unlock_product2 "github.com/corey888773/ztp-shopping-cart/products-api/src/features/v1/unlock_product"
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
	postgresConn, err := data2.NewPostgresConnector(data2.PostgresConfig{
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

	unitOfWork := data2.NewUnitOfWork(postgresConn.DB)
	productsReadRepository := repository2.NewReadProductsRepository(postgresConn.DB)
	productsWriteRepository := repository2.NewWriteProductsRepository(postgresConn.DB)

	commandBus := commands.NewCommandBus()
	queryBus := queries.NewQueryBus()

	getProductsQueryHandler := get_products2.NewHandler(productsReadRepository)
	queryBus.Register(&get_products2.Query{}, getProductsQueryHandler)

	lockProductCommandHandler := lock_product2.NewHandler(unitOfWork, productsWriteRepository, productsReadRepository)
	commandBus.Register(&lock_product2.Command{}, lockProductCommandHandler)

	unlockProductCommandHandler := unlock_product2.NewHandler(unitOfWork, productsWriteRepository, productsReadRepository)
	commandBus.Register(&unlock_product2.Command{}, unlockProductCommandHandler)

	checkoutCommandHandler := checkout2.NewHandler(unitOfWork, productsReadRepository, productsWriteRepository)
	commandBus.Register(&checkout2.Command{}, checkoutCommandHandler)

	// Register query handler for get_all_products feature
	getAllProductsQueryHandler := get_all_products2.NewHandler(productsReadRepository)
	queryBus.Register(&get_all_products2.Query{}, getAllProductsQueryHandler)

	return &Srv{
		Router:     gin.Default(),
		CommandBus: commandBus,
		QueryBus:   queryBus,
	}, nil
}

func (s *Srv) SetupRouter() {
	productRoutes := s.Router.Group("api/v1/products")
	productRoutes.POST("/", get_products2.GetProducts(s.QueryBus))
	productRoutes.POST("/lock", lock_product2.LockProduct(s.CommandBus))
	productRoutes.POST("/unlock", unlock_product2.UnlockProduct(s.CommandBus))
	productRoutes.POST("/checkout", checkout2.Checkout(s.CommandBus))
	productRoutes.GET("/all", get_all_products2.GetAllProducts(s.QueryBus))
}

func (s *Srv) Start(httpAddress string) error {
	return s.Router.Run(httpAddress)
}

func (s *Srv) Stop() {

}
