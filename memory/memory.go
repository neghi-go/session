package memory

import (
	"context"
	"sync"
	"time"
)

type Storage struct {
	data map[string][]byte
	sync.RWMutex
}

// Close implements sessions.Storage.
func (s *Storage) Close(ctx context.Context) error {
	panic("unimplemented")
}

// Delete implements sessions.Storage.
func (s *Storage) Delete(ctx context.Context, key string) error {
	s.Lock()
	defer s.Unlock()
	delete(s.data, key)
	return nil
}

// Get implements sessions.Storage.
func (s *Storage) Get(ctx context.Context, key string) ([]byte, error) {
	s.RLock()
	s.RUnlock()
	return s.data[key], nil
}

// Reset implements sessions.Storage.
func (s *Storage) Reset(ctx context.Context) error {
	panic("unimplemented")
}

// Set implements sessions.Storage.
func (s *Storage) Set(ctx context.Context, key string, val []byte, ttl time.Duration) error {
	panic("unimplemented")
}

var memory_pool = sync.Pool{
	New: func() any {
		d := new(Storage)
		d.data = make(map[string][]byte)
		return d
	},
}

func getStorage() *Storage {
	obj := memory_pool.Get().(*Storage)
	return obj
}

func New() *Storage {
	return getStorage()
}

//var _ sessions.Storage = (*Storage)(nil)
