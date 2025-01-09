package session

import (
	"reflect"
	"testing"
	"time"
)

func TestNewCookie(t *testing.T) {
	type args struct {
		name  string
		value string
		opts  []Options
	}
	tests := []struct {
		name string
		args args
		want *Cookie
	}{
		{
			name: "Test New Cookie with Empty Options",
			args: args{
				name:  "cookie",
				value: "cookie",
				opts:  nil,
			},
			want: NewCookie("cookie", "cookie"),
		},
		{
			name: "Test New Cookie With Not Empty Options",
			args: args{
				name:  "cookie",
				value: "cookie",
				opts: []Options{
					SetSecure(),
					SetHTTPOnly(),
				},
			},
			want: NewCookie("cookie", "cookie", SetSecure(), SetHTTPOnly()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCookie(tt.args.name, tt.args.value, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetSecure(t *testing.T) {
	tests := []struct {
		name string
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetSecure(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSecure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetHTTPOnly(t *testing.T) {
	tests := []struct {
		name string
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetHTTPOnly(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetHTTPOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetDomain(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetDomain(tt.args.domain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetPath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetSameSite(t *testing.T) {
	type args struct {
		samesite SameSite
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetSameSite(tt.args.samesite); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSameSite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetMaxAge(t *testing.T) {
	type args struct {
		max_age time.Duration
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetMaxAge(tt.args.max_age); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetExpiry(t *testing.T) {
	type args struct {
		expiry time.Time
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetExpiry(tt.args.expiry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetExpiry() = %v, want %v", got, tt.want)
			}
		})
	}
}
