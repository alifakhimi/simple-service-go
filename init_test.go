package simple_test

// type testSvc struct {
// 	simple.Service
// }

// func (svc *testSvc) Init() error {
// 	fmt.Println("Test service initialize")
// 	return nil
// }

// func NewService() simple.Interface {
// 	return &testSvc{
// 		Service: *simple.New("./config.sample.json"),
// 	}
// }

// func TestRun(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		svc     *testSvc
// 		wantErr bool
// 	}{
// 		{
// 			name:    "Run service",
// 			svc:     NewService(),
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := simple.Run(tt.svc); (err != nil) != tt.wantErr {
// 				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
