/**
### 第一次作业
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

我的答案：要往上抛，因为dao 自己处理不了就应该往上抛。可以调用errors.Wrap或者errors.New 方法往上抛，上层可以获得错误类型、错误根因、错误trace
 */
package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	errors2 "github.com/pkg/errors"
	"log"
	"os"
	"time"
)

// 这是一个唱片的模型字段
type Album struct {
	ID     int64
	Title  string // 名称
	Artist string // 作者
	Price  float32 // 价钱
}

// 有一个创建订单的方法
// 一个顾客想买 quantity 个唱片，唱片id是albumID，顾客id 是custID
func CreateOrder(ctx context.Context, albumID, quantity, custID int) (orderID int64, err error) {
	tx, err := db.BeginTx(ctx, nil)
	defer tx.Rollback()

	var enough bool
	// 判断albumID 这种唱片库存有没有quantity 个，也有可能压根没有这种唱片
	if err = tx.QueryRowContext(ctx, "SELECT (quantity >= ?) from album where id = ?",
		quantity, albumID).Scan(&enough); err != nil {
		if err == sql.ErrNoRows {
			//return 0, errors2.Wrap(err,"no such album")
			return 0, errors2.New("no such album")
		}
	}

	// other logic

}

//  In production, you’d avoid the global variable,
//  such as by passing the variable to functions that need it or by wrapping it in a struct
var db *sql.DB

func main() {

	cfg := mysql.Config{
		User: os.Getenv(DBUSER),
		Passwd: os.Getenv(DBPASS),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}
	// Get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected success !")

	// sql tx
	ctx, _ := context.WithCancel(context.Background())
	orderId, err := CreateOrder(ctx, 20, 2, 123)
	if err != nil {
		fmt.Printf("error trace: \n%+v\n", err)
		// %T 返回的是错误类型，可以用于err 错位类型断言处理
		fmt.Printf("origin error: %T \n", errors2.Cause(err))
		// 原始错误
		fmt.Printf("origin error: %v\n", errors2.Cause(err))
		log.Fatal(err)
	}
	fmt.Printf("order id %v", orderId)
}
