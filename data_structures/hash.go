package datastructures

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func HashData(ctx context.Context, db *redis.Client) {
	// Setting values ​​in Hash

	err := db.HSet(ctx, "user:10", "name", "John", "age", "30", "email", "john@mail.com").Err()
	if err != nil {
		log.Fatalf("error in setting values: %s", err)
	}

	// Getting all values in Hash

	result, err := db.HGetAll(ctx, "user:10").Result()
	if err != nil {
		log.Fatalf("error in getting values: %s", err)
	}
	fmt.Println("user:10 : ", result)

	// Getting a single field from Hash
	name, err := db.HGet(ctx, "user:10", "name").Result()
	if err != nil {
		log.Fatalf("error in getting values: %s", err)
	}
	fmt.Println("user:10 - name : ", name)

	age, err := db.HIncrBy(ctx, "user:10", "age", 6).Result()
	if err != nil {
		log.Fatalf("error in incrby age: %s", err)
	}
	fmt.Println("user:10 - age : ", age)

	res, err := db.HMGet(ctx, "user:10", "age", "email").Result()
	if err != nil {
		log.Fatalf("error in getting values: %s", err)
	}
	fmt.Println("user:10 - ", res)
}
