package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	Ctx = context.TODO()
)

type Database struct {
	Client *redis.Client
}

//NewDatabase creates a redis client on default ip:port
func NewDatabase() (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}

func (db *Database) SetToRedis(ctx context.Context, key, val string) {
	err := db.Client.Set(ctx, key, val, 0).Err()
	if err != nil {
		log.Printf("Could not set %v - %v in redis database", key, val)
	}
}

func (db *Database) GetFromRedis(ctx context.Context, key string) string {
	val, err := db.Client.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Could not get %v - %v in redis database", key, val)
	}

	return val
}

func (db *Database) GetAllKeyValues(ctx context.Context) map[string]string {
	var keys []string

	iter := db.Client.Scan(ctx, 0, "https:*", 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(keys)

	keyvalues := map[string]string{}

	for _, key := range keys {
		keyvalues[key] = db.GetFromRedis(ctx, key)
	}

	return keyvalues
}

func (db *Database) GetAllImages(ctx context.Context) []string {
	var keys []string

	iter := db.Client.Scan(ctx, 0, "https:*", 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		fmt.Println(err)
	}

	return keys
}
