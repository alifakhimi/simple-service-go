package simple

import (
	"flag"

	"github.com/sirupsen/logrus"
)

func ReadConfigFromFlag() {
	var (
		configPath string
	)

	flag.StringVar(&configPath, "c", "config.json", "config path with json extension")
	flag.Parse()

	if err := readConfig(configPath); err != nil {
		logrus.Panicln(err)
	}
}
