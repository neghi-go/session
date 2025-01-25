package server

import "github.com/neghi-go/session/store"

type Options func(*Server)

func WithStore(store store.Store) Options {
	return func(s *Server) {
		s.store = store
	}
}
