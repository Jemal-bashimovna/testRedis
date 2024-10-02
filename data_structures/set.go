package datastructures

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func Set(ctx context.Context, db *redis.Client) {
	subscribers := []string{"user1@mail.ru", "user2@mail.ru", "user3@mail.ru", "user4@mail.ru"}

	for _, user := range subscribers {
		err := db.SAdd(ctx, "users", user).Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	result, err := db.SMembers(ctx, "users").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All subscribers: ", result)
	isExist := "user1@mail.ru"
	exist, err := db.SIsMember(ctx, "users", isExist).Result()
	if err != nil {
		log.Fatal(err)
	}
	println()
	fmt.Printf("Is exist member '%s' : %t\n", isExist, exist)

	if exist {
		fmt.Printf("%s subscribed \n", isExist)
	} else {
		fmt.Printf("%s is not subscribed \n", isExist)
	}
	println()
	err = db.SRem(ctx, "users", isExist).Err()
	if err != nil {
		log.Fatal(err)
	}

	result, err = db.SMembers(ctx, "users").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("All subscribers after removing %s: %s", isExist, result)
	println()

	count, err := db.SCard(ctx, "users").Result()
	if err != nil {
		log.Fatal(err)
	}
	println()
	fmt.Printf("Number of subscribers: %d\n", count)
	println()
}
