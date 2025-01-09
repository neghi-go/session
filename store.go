package session

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/http"
	"time"

	"github.com/neghi-go/session/memory"
)

type Storage interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	Reset(ctx context.Context) error
}

type SessionKey int

const (
	SessionIDKey SessionKey = iota
)

type Store struct {
	storage Storage

	idle_timeout time.Duration

	absolute_timeout time.Duration

	sessionName string

	secret string
}

type StoreOpts func(*Store)

func New(opts ...StoreOpts) *Store {
	cfg := &Store{
		storage:          memory.New(),
		idle_timeout:     0,
		absolute_timeout: 0,
		sessionName:      "session",
		secret:           "test-secret",
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func WithStorage(storage Storage) StoreOpts {
	return func(s *Store) {
		s.storage = storage
	}
}

func SetIdleTimeout(timeout time.Duration) StoreOpts {
	return func(s *Store) {
		s.idle_timeout = timeout
	}
}

func SetAbsoluteTimeout(timeout time.Duration) StoreOpts {
	return func(s *Store) {
		s.absolute_timeout = timeout
	}
}

func SetSessionName(name string) StoreOpts {
	return func(s *Store) {
		s.sessionName = name
	}
}

func SetSecret(secret string) StoreOpts {
	return func(s *Store) {
		s.secret = secret
	}
}

func (s *Store) Get(r *http.Request) (*Session, error) {
	return s.getSession(r)
}

func (s *Store) getSession(r *http.Request) (*Session, error) {
	var data []byte
	var err error

	id, ok := r.Context().Value(SessionIDKey).(string)
	if !ok {
		id = s.getSessionID(r)
	}

	if id != "" {
		data, err = s.storage.Get(r.Context(), id)
		if err != nil {
			return nil, err
		}

		if data == nil {
			id = ""
		}
	}

	if id == "" {
		id = s.generateKey()
		r.WithContext(context.WithValue(r.Context(), SessionIDKey, id))
	}

	sess := &Session{}
	sess.id = id

	//decode and store values or what ever

	return sess, nil
}

func (s *Store) getSessionID(r *http.Request) string {
	id, _ := r.Cookie(s.sessionName)
	return id.Value
}

func (s *Store) Delete(ctx context.Context, id string) error {
	return s.storage.Delete(ctx, id)
}

func (s *Store) Reset(ctx context.Context) error {
	return s.storage.Reset(ctx)
}

func (s *Store) GetByID(ctx context.Context, id string) (*Session, error) {
	var data []byte
	var err error

	data, err = s.storage.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("invalid session id")
	}

	sess := &Session{}

	return sess, nil
}

func (s *Store) generateKey() string {
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	h := hmac.New(sha256.New, []byte(base64.RawURLEncoding.EncodeToString(buf)))
	_, err = h.Write([]byte(s.secret))
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}
