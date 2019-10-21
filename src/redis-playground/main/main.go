package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func ConnectNewClient() (*redis.Client, error) {
	cl := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := cl.Ping().Result()
	fmt.Println(pong, err)
	return cl, err
}

func main() {
	cl, _ := ConnectNewClient()
	err := cl.Set("key", "value", 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	val, err := cl.Get("key").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("key", val)
	val2, err := cl.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("key2", val2)
	}

	// publish to go-channel
	cl.Publish("go-channel", "Hey")
}
