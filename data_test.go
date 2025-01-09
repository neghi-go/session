package session

import (
	"reflect"
	"testing"
)

func Test_data_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    *data
		args args
		want any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("data.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_data_Set(t *testing.T) {
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name string
		d    *data
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Set(tt.args.key, tt.args.value)
		})
	}
}

func Test_data_Delete(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    *data
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Delete(tt.args.key)
		})
	}
}

func Test_data_Reset(t *testing.T) {
	tests := []struct {
		name string
		d    *data
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Reset()
		})
	}
}
