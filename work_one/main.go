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

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// CreateOrder creates an order for an album and returns the new order ID.
func CreateOrder(ctx context.Context, albumID, quantity, custID int) (orderID int64, err error) {
	tx, err := db.BeginTx(ctx, nil)
	defer tx.Rollback()

	var enough bool
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
