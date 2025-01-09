package redis

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

var test_url string

func TestMain(m *testing.M) {
	req := testcontainers.ContainerRequest{
		Image:        "redis:alpine",
		ExposedPorts: []string{"6379/tcp"},
	}

	redis, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	test_url, _ = redis.Endpoint(context.Background(), "")
	testVal := m.Run()
	testcontainers.TerminateContainer(redis)
	os.Exit(testVal)
}

func Test_redis_Get(t *testing.T) {
	var (
		conn, _ = New(WithURL("redis://"+test_url), WithReset())
		key     = "id"
		val     = []byte("Jon Doe")
	)

	err := conn.Set(context.Background(), key, val, 0)
	require.NoError(t, err)

	data, err := conn.Get(context.Background(), key)
	require.NoError(t, err)
	require.Equal(t, val, data)
}

func Test_redis_Delete(t *testing.T) {
	var (
		conn, _ = New(WithURL("redis://"+test_url), WithReset())
		key     = "id"
		val     = []byte("Jon Doe")
	)

	err := conn.Set(context.Background(), key, val, 0)
	require.NoError(t, err)

	data, err := conn.Get(context.Background(), key)
	require.NoError(t, err)
	require.Equal(t, val, data)

	err = conn.Delete(context.Background(), key)
	require.NoError(t, err)

	res, err := conn.Get(context.Background(), key)
	require.NoError(t, err)
	require.Equal(t, []byte(""), res)
}
