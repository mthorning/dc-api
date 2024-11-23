package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
)

var (
        redisClient *redis.Client
        ctx         = context.Background()
)

func index(w http.ResponseWriter, req *http.Request) {
	val, err := redisClient.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "hello frens!\n foo: %s", val)
}

func setRedisVal() {
	ctx := context.Background()

	err := redisClient.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}
}

func main() {
        setRedisVal()
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)
}
