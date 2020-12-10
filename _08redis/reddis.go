package main
import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v8"
)
func main() {
	ExampleClient()
}

var ctx = context.Background()

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.103.113:6379",
		Password: "Qbkj@2020", // no password set
		DB:       1,  // use default DB
	})

	err := rdb.Set(ctx, "key", "valuetang", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
