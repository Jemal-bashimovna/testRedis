package main

import (
	"context"
	"fmt"
	"log"
	datastructures "testproject/data_structures"
	"testproject/db"
)

var ctx = context.Background()

func main() {

	// 1. Подключение к Redis
	// opt, err := redis.ParseURL("redis://@localhost:6379/0")
	// if err != nil {
	// 	panic(err)
	// }

	// client := redis.NewClient(opt)

	client := db.NewRedisClient()

	// datastructures.Strings(ctx, client)
	// datastructures.HashData(ctx, client)
	// datastructures.Lists(ctx, client)
	// datastructures.Set(ctx, client)
	// datastructures.SortedSets(ctx, client)

	userId := "10"
	userData, err := datastructures.GetDataWithCache(ctx, client, userId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("user's data : %s\n", userData)

	defer client.Close()
}
