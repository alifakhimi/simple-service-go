package utils

import (
	"testing"

	"github.com/alifakhimi/simple-service-go/database"
)

func TestGetID(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want database.PID
	}{
		{
			name: "embedded database.Model",
			args: args{
				val: struct {
					database.Model
					Value string
				}{
					Model: database.Model{
						ID: database.PID(2),
					},
					Value: "value of struct",
				},
			},
			want: database.PID(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetID(tt.args.val); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTableName(t *testing.T) {
	type SampleTable struct{}
	type Sample struct{}

	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get table name of SampleTable",
			args: args{
				val: &SampleTable{},
			},
			want: "sample_tables",
		},
		{
			name: "get table name of Sample struct",
			args: args{
				val: &Sample{},
			},
			want: "samples",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTableName(tt.args.val); got != tt.want {
				t.Errorf("GetTableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
