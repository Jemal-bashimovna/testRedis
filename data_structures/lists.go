package datastructures

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func Lists(ctx context.Context, db *redis.Client) {

	fmt.Println("List of tasks")

	err := db.RPush(ctx, "tasks", "task1", "task2", "task3").Err()
	if err != nil {
		log.Fatalf("error adding tasks: %s", err)
	}

	lists, err := db.LRange(ctx, "tasks", 0, -1).Result()
	if err != nil {
		log.Fatalf("error getting elements: %s", err)
	}

	fmt.Println(lists)

	firstElement, err := db.LPop(ctx, "tasks").Result()
	if err != nil {
		log.Fatalf("error removing element: %s", err)
	}
	fmt.Println("Removed element: ", firstElement)

	lists, err = db.LRange(ctx, "tasks", 0, -1).Result()
	if err != nil {
		log.Fatalf("error getting elements: %s", err)
	}

	fmt.Println(lists)

	movedEl, err := db.LMove(ctx, "tasks", "done", "left", "left").Result()
	if err != nil {
		log.Fatalf("%s", err)
	}
	doneList, err := db.LRange(ctx, "done", 0, -1).Result()
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("Element '%s' moved to '%s' \n", movedEl, doneList)

	length, err := db.LLen(ctx, "tasks").Result()
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Printf("%d elements in 'tasks' \n", length)

	// for i := 0; i <= len(lists); i++ {
	// 	res, err := db.LPop(ctx, "tasks").Result()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("removed element: ", res)
	// }

	fmt.Println("Car list: ")

	err = db.LPush(ctx, "cars", "BMW", "Toyota", "Kia").Err()
	if err != nil {
		log.Fatal(err)
	}

	carList, err := db.LRange(ctx, "cars", 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(carList)

	for i := 1; i <= len(carList); i++ {
		car, err := db.LPop(ctx, "cars").Result()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("removed elemet: ", car)
	}
}
