package simple

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/alifakhimi/simple-service-go/database/connection"
	"github.com/alifakhimi/simple-service-go/utils/httpserver"
	"github.com/alifakhimi/simple-service-go/utils/rest"
)

var ()

type (
	Interface interface {
		Init() error
	}

	Simple interface {
		Run(svc Interface) Simple
		Error() error
	}

	Service struct {
		*Config
		// err returns errors
		err error
	}
)

func New(configPath string) Simple {
	return &Service{
		Config: NewConfig(configPath),
	}
}

func NewWithConfig(conf *Config) Simple {
	return &Service{
		Config: conf,
	}
}

func (s *Service) GetHttpServer(name string) (h *httpserver.HttpServer, err error) {
	return s.Config.GetHttpServer(name)
}

func (s *Service) GetHttp(name string) (ech *echo.Echo, err error) {
	return s.Config.GetHttpServerEcho(name)
}

func (s *Service) GetClient(name string) (client *rest.Client, err error) {
	return s.Config.GetClient(name)
}

func (s *Service) GetRestyClient(name string) (client *resty.Client, err error) {
	return s.Config.GetRestyClient(name)
}

func (s *Service) GetDB(name string) (db *connection.DBConnection, err error) {
	return s.Config.GetDB(name)
}

func (s *Service) GetDBGorm(name string) (db *gorm.DB, err error) {
	return s.Config.GetDBGorm(name)
}
