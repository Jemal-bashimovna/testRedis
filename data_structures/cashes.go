package datastructures

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func GetDataWithCache(ctx context.Context, db *redis.Client, userId string) (string, error) {

	cachedData, err := db.Get(ctx, "user:"+userId).Result()

	if err == redis.Nil {
		//  Если данные не найдены в кэше, получаем их из медленного источника
		fmt.Println("Data not found in Cache, query to DB...")
		data := GetDatafromDB(userId)

		//  Сохраняем данные в Redis с TTL (время жизни ключа, например, 5 минут)
		err = db.Set(ctx, "user:"+userId, data, 5*time.Minute).Err()
		if err != nil {
			return "", err
		}

		return data, nil
	}

	if err != nil {
		return "", err
	}

	// Если данные найдены в кэше, возвращаем их
	fmt.Println("Data found in Cache")
	return cachedData, nil
}

func GetDatafromDB(userId string) string {
	time.Sleep(2 * time.Second)
	return "John"
}
