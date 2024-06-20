package simple_test

import (
	"fmt"
	"testing"

	simutils "github.com/alifakhimi/simple-utils-go"
	"github.com/alifakhimi/simple-utils-go/simrest"
	"github.com/spf13/viper"

	"github.com/alifakhimi/simple-service-go"
)

type testSvc struct {
	Simple simple.Simple
}

func (svc *testSvc) Init() error {
	fmt.Println("Test service initialize")
	return nil
}

func newService() simple.Interface {
	return &testSvc{
		Simple: simple.New("./config.sample.json"),
	}
}

func TestService_Run(t *testing.T) {
	type fields struct {
		Config *simple.Config
	}
	type args struct {
		svc simple.Interface
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Run service",
			fields: fields{
				Config: &simple.Config{
					Name:        "Test",
					Version:     "1.0.0",
					Description: "Test Desc",
					Website:     "https://example.com",
					HttpServers: map[string]*simutils.HttpServer{
						"test": {
							HttpServerConfig: simutils.HttpServerConfig{
								Address: ":8080",
								Prefix:  "/api/v1",
								Debug:   true,
							},
						},
					},
					Clients:   simrest.Clients{},
					Databases: simutils.DBs{},
					Meta:      nil,
					Banners: []*simple.Banner{
						{
							Text:  "SIKA",
							Font:  "mini",
							Color: "blue",
						},
					},
					Viper: viper.GetViper(),
				},
			},
			args: args{
				svc: newService(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &simple.Service{
				Config: tt.fields.Config,
			}
			if err := s.Run(tt.args.svc).Error(); (err != nil) != tt.wantErr {
				t.Errorf("Service.Run() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
