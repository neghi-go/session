package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/neghi-go/session"
	"github.com/neghi-go/session/store"
	"github.com/neghi-go/session/store/memory"
	"github.com/neghi-go/utilities"
)

type Server struct {
	store store.Store

	absoluteTimeout time.Duration
	idleTimeout     time.Duration

	identifier string
	keyGenFunc func() string

	secure   bool
	httpOnly bool
	domain   string
	path     string
	sameSite http.SameSite

	session *Session
}

// DelField implements sessions.Session.
func (s *Server) DelField(key string) error {
	s.session.data.Del(key)
	return nil
}

// GetField implements sessions.Session.
func (s *Server) GetField(key string) interface{} {
	panic("unimplemented")
}

// SetField implements sessions.Session.
func (s *Server) SetField(key string, value interface{}) error {
	panic("unimplemented")
}

func NewServerSession(opts ...Options) *Server {
	cfg := &Server{
		store: memory.New(),
		keyGenFunc: func() string {
			return utilities.Generate(16)
		},
		identifier:      "sessions-key",
		sameSite:        http.SameSiteLaxMode,
		secure:          false,
		httpOnly:        true,
		idleTimeout:     0,
		absoluteTimeout: time.Hour,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// Generate implements sessions.Session.
func (s *Server) Generate(w http.ResponseWriter, subject string, params ...interface{}) error {
	//Generate Session
	session := s.generateSession()
	session.data.Set("subject", subject)
	for idx, data := range params {
		session.data.Set(fmt.Sprint(idx), data)
	}
	s.persistToStore(context.Background())
	//Send Cookie to
	http.SetCookie(w, &http.Cookie{
		Name:  s.identifier,
		Value: session.id,

		Expires:  time.Now().Add(s.absoluteTimeout * time.Second).UTC(),
		Secure:   s.secure,
		HttpOnly: s.httpOnly,

		Domain: s.domain,
		Path:   s.path,

		SameSite: s.sameSite,
	})

	return nil
}

// Validate implements sessions.Session.
func (s *Server) Validate(key string) error {
	data := make(map[string]interface{})
	d, err := s.store.Get(context.Background(), key)
	if err != nil {
		return err
	}
	err = utilities.GobDecode(d, &data)
	if err != nil {
		return err
	}
	s.session = &Session{
		id: key,
		data: &Data{
			mu:   &sync.RWMutex{},
			data: data,
		},
	}
	return nil
}

var _ sessions.Session = (*Server)(nil)

func (s *Server) generateSession() *Session {
	scs := &Session{
		id: s.keyGenFunc(),
		data: &Data{
			mu:   &sync.RWMutex{},
			data: make(map[string]interface{}),
		},
	}
	return scs
}

func (s *Server) persistToStore(ctx context.Context) error {
	//Persist session to Store
	d, err := utilities.GobEncode(s.session.data.data)
	if err != nil {
		return err
	}
	err = s.store.Set(ctx, s.session.id, d, s.absoluteTimeout)
	if err != nil {
		return err
	}
	return nil
}
