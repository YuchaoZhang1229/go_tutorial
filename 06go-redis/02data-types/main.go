package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var ctx = context.Background()

func connectRedis() *redis.Client {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Replace with your Redis server address
		Password: "",               // No password for local development
		DB:       0,                // Default DB
		Protocol: 2,                // Connection protocol
	})

	// Ping the Redis server to check the connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	return rdb
}

func strings(rdb *redis.Client) {
	// Set a key-value pair
	res1, err := rdb.Set(ctx, "greeting", "Hello, Redis!", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res1) // >>> OK

	res2, err := rdb.Get(ctx, "greeting").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key: greeting does not exist")
	case err != nil:
		fmt.Println("Get failed", err)
	case res2 == "":
		fmt.Println("value is empty")
	}
	fmt.Println(res2) // >>> Deimos
}

// Redis provides list data structures that can be useful for implementing queues.
func lists(rdb *redis.Client) {
	// LPush nsert values at the head of the list
	err := rdb.LPush(ctx, "tasks", "Task 1", "Task 2").Err() // Task 2 Task1
	if err != nil {
		log.Fatal(err)
	}

	// RPush insert values at the back/ tail of the list
	err = rdb.RPush(ctx, "tasks", "Task 1", "Task 2").Err() // Task 2 Task1 Task1 Task 2
	if err != nil {
		log.Fatal(err)
	}

	// LPop remove the first element in the list
	task, err := rdb.LPop(ctx, "tasks").Result() // Task1 Task1 Task 2
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Popped Task:", task)

	// RPopr remove the last element in the list
	task, err = rdb.RPop(ctx, "tasks").Result() // Task1 Task1
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Popped Task:", task)
}

// Hashes in Redis allow you to store multiple field-value pairs under a single key.
// In Redis >= 4.0, use HSet: rdb.HSet(ctx, "myhash", Per{Name: "hi", Age: 20})
// In Redis < 4.0, use HMSet: rdb.HMSet(ctx, "myhash", Per{Name: "hi", Age: 20})
func hashes(rdb *redis.Client) {
	// Set hash field-values
	err := rdb.HSet(ctx, "user:1", map[string]interface{}{
		"name":  "John Doe",
		"email": "john@example.com",
		"age":   25,
	}).Err()
	if err != nil {
		log.Fatal(err)
	}

	// Get hash field-values
	userInfo, err := rdb.HGetAll(ctx, "user:1").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User Info:", userInfo)
}

func main() {
	rdb := connectRedis()
	strings(rdb)
	lists(rdb)
	hashes(rdb)

}
