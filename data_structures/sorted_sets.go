package datastructures

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func SortedSets(ctx context.Context, db *redis.Client) {

	fmt.Println("Adding elements to Sorted Set")
	err := db.ZAdd(ctx, "game_scores", redis.Z{Score: 200, Member: "player1"},
		redis.Z{Score: 100, Member: "player2"},
		redis.Z{Score: 160, Member: "player3"}).Err()

	if err != nil {
		log.Fatalf("error to adding elements to sorted sets: %v", err)
	}
	println()

	fmt.Println("Getting all elements")

	result, err := db.ZRange(ctx, "game_scores", 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	println()

	fmt.Println("Getting all elements with scores:")
	resultScores, err := db.ZRangeWithScores(ctx, "game_scores", 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All players with their scores:")
	println()
	for _, player := range resultScores {
		fmt.Printf("%s : %.0f\n", player.Member, player.Score)
	}

	println()

	// Get index of element

	rank, err := db.ZRank(ctx, "game_scores", "player1").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("index of player1: ", rank)

	rank, err = db.ZRank(ctx, "game_scores", "player2").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("index of player2: ", rank)

	rank, err = db.ZRank(ctx, "game_scores", "player3").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("index of player3: ", rank)
	println()

	fmt.Println("Get score of player:")

	score, err := db.ZScore(ctx, "game_scores", "player1").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("score of player1: ", score)
	println()

	err = db.ZRem(ctx, "game_scores", "player1").Err()
	if err != nil {
		log.Fatal(err)
	}

	updatedScores, err := db.ZRangeWithScores(ctx, "game_scores", 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}
	for _, score := range updatedScores {
		fmt.Printf("%s : %.0f\n", score.Member, score.Score)
	}
	println()

	highScore, err := db.ZRangeByScoreWithScores(ctx, "game_scores", &redis.ZRangeBy{
		Min: "150", //
		Max: "+inf",
	}).Result()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("High scoring players:")
	for _, score := range highScore {
		fmt.Printf("%s : %.0f\n", score.Member, score.Score)
	}
	println()
}
