package database

import "github.com/go-redis/redis/v9"

type RedisDB struct {
	Rdb *redis.Client
}

func NewRedisDB() *RedisDB {
	return &RedisDB{}
}

func (r *RedisDB) RedisConnect() {
	r.Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		// use default DB
	})
}
