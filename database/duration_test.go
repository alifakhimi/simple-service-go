package database_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/alifakhimi/simple-service-go/database"
)

func TestDuration_MarshalJSON(t *testing.T) {
	type fields struct {
		Duration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "15m0s",
			fields: fields{
				Duration: 15 * time.Minute,
			},
			want:    []byte("15m0s"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := database.Duration{
				Duration: tt.fields.Duration,
			}
			got, err := d.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Duration.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Duration.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuration_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Duration time.Duration
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "15m0s",
			fields: fields{
				Duration: 15 * time.Minute,
			},
			args: args{
				b: []byte("15m0s"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &database.Duration{}
			if err := d.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Duration.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if d.Duration != tt.fields.Duration {
				t.Errorf("Duration Parse got %v, want %v", d.Duration, tt.fields.Duration)
			}
		})
	}
}