package session

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	type args struct {
		opts []StoreOpts
	}
	tests := []struct {
		name string
		args args
		want *Store
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithStorage(t *testing.T) {
	type args struct {
		storage Storage
	}
	tests := []struct {
		name string
		args args
		want StoreOpts
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithStorage(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetIdleTimeout(t *testing.T) {
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name string
		args args
		want StoreOpts
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetIdleTimeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIdleTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetAbsoluteTimeout(t *testing.T) {
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name string
		args args
		want StoreOpts
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetAbsoluteTimeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAbsoluteTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetSessionName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want StoreOpts
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetSessionName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSessionName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetSecret(t *testing.T) {
	type args struct {
		secret string
	}
	tests := []struct {
		name string
		args args
		want StoreOpts
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetSecret(tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_Get(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		want    *Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Get(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_getSession(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		want    *Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.getSession(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.getSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.getSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_getSessionID(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		s    *Store
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.getSessionID(tt.args.r); got != tt.want {
				t.Errorf("Store.getSessionID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Store.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStore_Reset(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Reset(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Store.Reset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStore_GetByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		want    *Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_generateKey(t *testing.T) {
	store := New()
	key := store.generateKey()
	require.NotEmpty(t, key)
}
