package simple

import (
	"encoding/json"
	"errors"

	simutils "github.com/alifakhimi/simple-utils-go"
	"github.com/alifakhimi/simple-utils-go/simrest"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// error block
var (
	ErrHttpServerNotFound = errors.New("http server not found")
	ErrClientNotFound     = errors.New("client not found")
	ErrDBConnNotFound     = errors.New("db connection not found")
)

type (
	Config struct {
		// Name is name of service
		Name string `json:"name,omitempty"`
		// DisplayName displays name of service in human readable
		DisplayName string `json:"display_name,omitempty"`
		// Version show version in banner
		Version string `json:"version,omitempty"`
		// Description
		Description string `json:"description,omitempty"`
		// Website
		Website string `json:"website,omitempty"`
		// Address optionally specifies the TCP address for the server to listen on,
		HttpServers simutils.HttpServers `json:"http_servers,omitempty"`
		// Clients is a list of rest client
		Clients simrest.Clients `json:"clients,omitempty"`
		// Databases is a list of database connection
		Databases simutils.DBs `json:"databases,omitempty"`
		// Meta
		Meta any `json:"meta,omitempty"`
		// Banner will be displayed when the service starts
		Banners []*Banner `json:"banners,omitempty"`
		// viper is a config tools
		*viper.Viper
	}

	Banner struct {
		Text  string `json:"text,omitempty"`
		Font  string `json:"font,omitempty"`
		Color string `json:"color,omitempty"`
	}
)

func NewConfig(path ...string) *Config {
	c := Config{
		HttpServers: simutils.HttpServers{},
		Clients:     simrest.Clients{},
		Databases:   simutils.DBs{},
	}

	if len(path) > 0 && path[0] != "" {
		if err := ReadConfig(path[0], &c); err != nil {
			logrus.Panicln(err)
		}
	}

	return &c
}

func (conf *Config) GetHttpServer(name string) (h *simutils.HttpServer, err error) {
	if len(conf.HttpServers) == 0 {
		return nil, ErrHttpServerNotFound
	}

	if d, ok := conf.HttpServers[name]; !ok {
		return nil, ErrHttpServerNotFound
	} else {
		return d, nil
	}
}

func (conf *Config) GetHttpServerEcho(name string) (ech *echo.Echo, err error) {
	if d, err := conf.GetHttpServer(name); err != nil {
		return nil, err
	} else {
		return d.Echo(), nil
	}
}

func (conf *Config) GetClient(name string) (client *simrest.Client, err error) {
	if len(conf.Clients) == 0 {
		return nil, ErrClientNotFound
	}

	if c, ok := conf.Clients[name]; !ok {
		return nil, ErrClientNotFound
	} else {
		return c, nil
	}
}

func (conf *Config) GetRestyClient(name string) (client *resty.Client, err error) {
	if c, err := conf.GetClient(name); err != nil {
		return nil, err
	} else {
		return c.Client, nil
	}
}

func (conf *Config) GetDB(name string) (db *simutils.DBConnection, err error) {
	if len(conf.Databases) == 0 {
		return nil, ErrDBConnNotFound
	}

	if d, ok := conf.Databases[name]; !ok {
		return nil, ErrDBConnNotFound
	} else {
		return d, nil
	}
}

func (conf *Config) GetDBGorm(name string) (db *gorm.DB, err error) {
	if d, err := conf.GetDB(name); err != nil {
		return nil, err
	} else {
		return d.DB, nil
	}
}

func readConfig(path string) error {
	logrus.Infoln("using config file:", path)

	viper.SetConfigType("json")
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func ReadConfig(path string, conf any) (err error) {
	// Read config from path
	if err = readConfig(path); err != nil {
		return err
	}

	configMap := make(map[string]any)

	if err := viper.Unmarshal(&configMap); err != nil {
		return err
	}

	b, err := json.Marshal(configMap)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, conf); err != nil {
		return err
	}

	if c, ok := conf.(*Config); ok {
		c.Viper = viper.GetViper()
	}

	return nil
}
