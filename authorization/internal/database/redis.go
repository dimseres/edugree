package database

import "github.com/redis/go-redis/v9"

type RedisConnectionConfig struct {
	Host     string
	Password string
	DB       int
}

var redisConn *redis.Client

func InitRedisConnection(config *RedisConnectionConfig) *redis.Client {
	redisConn = redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password, // no password set
		DB:       config.DB,       // use default DB
	})

	return redisConn
}

func GetRedisConnection() *redis.Client {
	return redisConn
}
