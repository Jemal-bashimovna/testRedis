package datastructures

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Strings(ctx context.Context, client *redis.Client) {

	err := client.Set(ctx, "key1", "val1", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Значение для ключа 'key1' установлено")

	val, err := client.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Значение для ключа 'key1': %s\n", val)
	fmt.Println()

	err = client.Set(ctx, "number", 10, 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Значение для ключа 'number' установлено")

	val2, err := client.Get(ctx, "number").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Значение для ключа 'key1': %s\n", val2)
	fmt.Println()

	newVal, err := client.IncrBy(ctx, "number", 6).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("incrby 6 : ", newVal)

	newVal, err = client.DecrBy(ctx, "number", 10).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("decrby 10 : ", newVal)

	append, err := client.Append(ctx, "number", "00").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("append 00: ", append)

	number, err := client.Get(ctx, "number").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Current value of 'number': ", number)

	length, err := client.StrLen(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("length: ", length)

}
