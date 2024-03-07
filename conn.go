package main

import (
	"github.com/go-redis/redis"
)

func getClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client

}

func saveReg(key string, value string) {
	client := getClient()
	err := client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func getReg(key string) string {
	client := getClient()
	val, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func delReg(key string) {
	client := getClient()

	err := client.Del(key).Err()
	if err != nil {
		panic(err)
	}
}

func getAllRegs() []string {
	client := getClient()

	val, err := client.Keys("*").Result()
	if err != nil {
		panic(err)
	}

	return val

}
