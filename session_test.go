package session

import (
	"reflect"
	"testing"
)

func TestSession_ID(t *testing.T) {
	tests := []struct {
		name string
		s    *Session
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ID(); got != tt.want {
				t.Errorf("Session.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		s    *Session
		args args
		want any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_Set(t *testing.T) {
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name string
		s    *Session
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestSession_Delete(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		s    *Session
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Delete(tt.args.key)
		})
	}
}

func TestSession_Save(t *testing.T) {
	tests := []struct {
		name string
		s    *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Save()
		})
	}
}

func TestSession_Destroy(t *testing.T) {
	tests := []struct {
		name    string
		s       *Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Destroy(); (err != nil) != tt.wantErr {
				t.Errorf("Session.Destroy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSession_Reset(t *testing.T) {
	tests := []struct {
		name string
		s    *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
		})
	}
}

func TestSession_Regenerate(t *testing.T) {
	tests := []struct {
		name string
		s    *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Regenerate()
		})
	}
}

func TestSession_deleteSession(t *testing.T) {
	tests := []struct {
		name string
		s    *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.deleteSession()
		})
	}
}

func TestSession_refreshSession(t *testing.T) {
	tests := []struct {
		name string
		s    *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.refreshSession()
		})
	}
}

func TestSession_encodeSessionData(t *testing.T) {
	tests := []struct {
		name    string
		s       *Session
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.encodeSessionData()
			if (err != nil) != tt.wantErr {
				t.Errorf("Session.encodeSessionData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session.encodeSessionData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_decodeSessionData(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		s       *Session
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.decodeSessionData(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Session.decodeSessionData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
