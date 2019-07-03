package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type RedisConn struct {
	Cluster bool
	*redis.Client
	*redis.ClusterClient
}

func NewConn(ip string, port int, password string) (*RedisConn, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", ip, port),
		Password: password,
		DB: 0,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return &RedisConn{false, client, nil}, nil
}

func convertToSecond(nano time.Duration) time.Duration {
	return nano * 1000000000
}

func (rc *RedisConn) Set(key string, value interface{}, expiration time.Duration) error {
	time := convertToSecond(expiration)
	fmt.Println(time)
	return rc.Client.Set(key, value, time).Err()
}

func (rc *RedisConn) Get(key string) (string, error) {
	return rc.Client.Get(key).Result()
}