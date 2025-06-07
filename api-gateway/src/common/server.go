package common

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Config struct {
	ServerPort     any    `mapstructure:"SERVER_PORT"`
	CartsApiUrl    string `mapstructure:"CARTS_API_URL"`
	ProductsApiUrl string `mapstructure:"PRODUCTS_API_URL"`
}

// Srv represents the API Gateway server
type Srv struct {
	Router        *gin.Engine
	cartsProxy    *httputil.ReverseProxy
	productsProxy *httputil.ReverseProxy
}

func NewServer(cfg Config) (*Srv, error) {
	cartsURL, err := url.Parse(cfg.CartsApiUrl)
	if err != nil {
		return nil, err
	}
	productsURL, err := url.Parse(cfg.ProductsApiUrl)
	if err != nil {
		return nil, err
	}

	return &Srv{
		Router:        gin.Default(),
		cartsProxy:    httputil.NewSingleHostReverseProxy(cartsURL),
		productsProxy: httputil.NewSingleHostReverseProxy(productsURL),
	}, nil
}

func (s *Srv) SetupRouter() {
	s.Router.Any("/api/v1/carts/*proxyPath", func(c *gin.Context) {
		s.cartsProxy.ServeHTTP(c.Writer, c.Request)
	})
	s.Router.Any("/api/v1/products/*proxyPath", func(c *gin.Context) {
		s.productsProxy.ServeHTTP(c.Writer, c.Request)
	})
}

func (s *Srv) Start(addr string) error {
	return s.Router.Run(addr)
}

func (s *Srv) Stop() {
}
