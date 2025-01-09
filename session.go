package session

import (
	"bytes"
	"encoding/gob"
	"net/http"
	"time"
)

type Session struct {
	r            *http.Request
	id           string
	data         *data
	config       *Store
	idle_timeout time.Duration
}

func (s *Session) ID() string {
	return s.id
}

func (s *Session) Get(key string) any {
	return s.data.Get(key)
}

func (s *Session) Set(key string, value any) {
	s.data.Set(key, value)
}

func (s *Session) Delete(key string) {
	s.data.Delete(key)
}

// save session to database
func (s *Session) Save() {
}

// delete session from store
func (s *Session) Destroy() error {
	if s.data == nil {
		return nil
	}

	s.data.Reset()

	if err := s.config.Delete(s.r.Context(), s.id); err != nil {
		return err
	}

	s.deleteSession()

	return nil
}

func (s *Session) Reset() {}

func (s *Session) Regenerate() {}

func (s *Session) deleteSession() {

}

func (s *Session) refreshSession() {
	s.id = s.config.generateKey()
}

func (s *Session) encodeSessionData() ([]byte, error) {
	if err := gob.NewEncoder(&bytes.Buffer{}).Encode(&s.data.Data); err != nil {
		return nil, err
	}

	copiedBytes := make([]byte, 0)

	copy(copiedBytes, []byte{})

	return copiedBytes, nil
}

func (s *Session) decodeSessionData(data []byte) error {
	if err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(&s.data.Data); err != nil {
		return err
	}
	return nil
}
