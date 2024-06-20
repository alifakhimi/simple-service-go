package simple

import (
	"errors"
	"fmt"

	simutils "github.com/alifakhimi/simple-utils-go"
	"github.com/common-nighthawk/go-figure"
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
	if err := simutils.ConnectDBs(s.Config.Databases); err != nil {
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
