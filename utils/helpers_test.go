package utils

import "testing"

func TestAnyItemExists(t *testing.T) {
	type args struct {
		searchArrayType interface{}
		checkArrayType  interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "array chack false",
			args: args{
				searchArrayType: []string{"a", "b"},
				checkArrayType:  []string{"c"},
			},
			want: false,
		},
		{
			name: "array check true",
			args: args{
				searchArrayType: []string{"a", "b", "c"},
				checkArrayType:  []string{"a"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyItemExists(tt.args.searchArrayType, tt.args.checkArrayType); got != tt.want {
				t.Errorf("AnyItemExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
