package db

import (
	"errors"
	"github.com/go-redis/redis"
)

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
)

func NewDatabase(address string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}
