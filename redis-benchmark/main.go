package main
import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var client redis.UniversalClient
var ctx context.Context

const (
	ip   string = "127.0.0.1"
	port uint16 = 6383
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%v:%v", ip, port),
		Password:     "",
		DB:           0,
		PoolSize:     128,
		MinIdleConns: 100,
		MaxRetries:   5,
	})

	ctx = context.Background()
}

func main() {
	// 依次给redis 加入value 大小为10byte/1000byte 的键值对 1w、5w、50w 对
	write(10000, "value_len_10_10k", generateValue(10))
	//write(50000, "value_len_10_50k", generateValue(10))
	//write(500000, "value_len_10_500k", generateValue(10))

	//write(10000, "value_len_10k_10k", generateValue(1000))
	//write(50000, "value_len_10k_50k", generateValue(1000))
	//write(500000, "value_len_10k_500k", generateValue(1000))

}

func write(num int, key, value string) {
	for i := 0; i < num; i++ {
		k := fmt.Sprintf("%s:%v", key, i)
		cmd := client.Set(ctx, k, value, -1)
		err := cmd.Err()
		if err != nil {
			fmt.Println(cmd.String())
		}
	}
}

func generateValue(size int) string {
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		arr[i] = 'a'
	}
	return string(arr)
}