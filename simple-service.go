package simple

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	simutils "github.com/alifakhimi/simple-utils-go"
	"github.com/alifakhimi/simple-utils-go/simrest"
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

func (s *Service) GetHttpServer(name string) (h *simutils.HttpServer, err error) {
	return s.Config.GetHttpServer(name)
}

func (s *Service) GetHttp(name string) (ech *echo.Echo, err error) {
	return s.Config.GetHttpServerEcho(name)
}

func (s *Service) GetClient(name string) (client *simrest.Client, err error) {
	return s.Config.GetClient(name)
}

func (s *Service) GetRestyClient(name string) (client *resty.Client, err error) {
	return s.Config.GetRestyClient(name)
}

func (s *Service) GetDB(name string) (db *simutils.DBConnection, err error) {
	return s.Config.GetDB(name)
}

func (s *Service) GetDBGorm(name string) (db *gorm.DB, err error) {
	return s.Config.GetDBGorm(name)
}
