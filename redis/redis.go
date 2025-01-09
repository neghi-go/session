package redis

import (
	"context"
	"errors"
	"time"

	"github.com/neghi-go/session"
	"github.com/redis/go-redis/v9"
)

type Options func(*RedisStorage)

func New(opts ...Options) (*RedisStorage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	rdb := &RedisStorage{}

	for _, opt := range opts {
		opt(rdb)
	}

	url, err := redis.ParseURL(rdb.url)
	if err != nil {
		return nil, err
	}

	if rdb.database != 0 {
		url.DB = rdb.database
	}

	db := redis.NewClient(url)

	if err := db.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	rdb.client = db

	if rdb.reset {
		err = rdb.Reset(ctx)
		if err != nil {
			return nil, err
		}
	}

	return rdb, nil
}

func WithDatabase(db int) Options {
	return func(rs *RedisStorage) {
		rs.database = db
	}
}

func WithURL(url string) Options {
	return func(rs *RedisStorage) {
		rs.url = url
	}
}

func WithReset() Options {
	return func(rs *RedisStorage) {
		rs.reset = true
	}
}

type RedisStorage struct {
	client *redis.Client

	database int

	url string

	reset bool
}

// Close implements sessions.Storage.
func (r *RedisStorage) Close(ctx context.Context) error {
	return r.client.Close()
}

// Delete implements sessions.Storage.
func (r *RedisStorage) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// Get implements sessions.Storage.
func (r *RedisStorage) Get(ctx context.Context, key string) ([]byte, error) {
	if key == "" {
		return []byte(""), nil
	}
	res, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return []byte(""), nil
		}
		return []byte(""), err
	}

	return res, nil
}

// Reset implements sessions.Storage.
func (r *RedisStorage) Reset(ctx context.Context) error {
	return r.client.FlushDB(ctx).Err()
}

// Set implements sessions.Storage.
func (r *RedisStorage) Set(ctx context.Context, key string, val []byte, ttl time.Duration) error {
	return r.client.Set(ctx, key, val, ttl).Err()
}

var _ session.Storage = (*RedisStorage)(nil)
