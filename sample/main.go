package main

import (
	"fmt"

	"github.com/alifakhimi/simple-service-go"
	"github.com/sirupsen/logrus"
)

type sampleSvc struct {
	Simple simple.Simple
}

func (svc *sampleSvc) Init() error {
	fmt.Println("Initialize test service")
	return nil
}

func newService() simple.Interface {
	return &sampleSvc{
		Simple: simple.New("./config.sample.json"),
	}
}

func main() {
	var (
		service *sampleSvc
	)

	if svc, ok := newService().(*sampleSvc); !ok {
		logrus.Fatalln("error implementation")
	} else {
		service = svc
	}

	if err := service.Simple.Run(service); err != nil {
		logrus.Fatalln(err)
	}
}
