package simple

import (
	"errors"
	"fmt"

	"github.com/common-nighthawk/go-figure"

	"github.com/alifakhimi/simple-service-go/database/connection"
)

// Run starts Service
func (s *Service) Run(svc Interface) Simple {
	// if sss, ok := svc.(*Service); ok {
	// 	_ = sss
	// }

	// Print banners
	for _, b := range s.Config.Banners {
		figure.NewColorFigure(b.Text, b.Font, b.Color, true).Print()
	}
	if s.Config.Version != "" {
		fmt.Println(s.Config.Version)
	}
	if s.Config.Description != "" {
		fmt.Printf("%s \n\n", s.Config.Description)
	}

	// Connect to databases
	if err := connection.ConnectDBs(s.Config.Databases); err != nil {
		s.err = errors.Join(s.err, err)
		return s
	}

	// Initialize service package
	if err := svc.Init(); err != nil {
		s.err = errors.Join(s.err, err)
		return s
	}

	// Start http servers
	if err := s.Config.HttpServers.RunAll(); err != nil {
		s.err = errors.Join(s.err, err)
		return s
	}

	return s
}

func (s *Service) Error() (err error) {
	return s.err
}
