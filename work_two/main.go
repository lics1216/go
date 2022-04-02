package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
)

func StartHttpServer(srv *http.Server) error {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request){
		io.WriteString(w, "hello, world!\n")
	})
	fmt.Println("http server start")
	err := srv.ListenAndServe()

	return err
}

// 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	group, errCtx := errgroup.WithContext(ctx)

	// http server
	srv := &http.Server{Addr: ":8080"}

	group.Go(func() error {
		return StartHttpServer(srv)
	})

	group.Go(func() error {
		//阻塞， 因为 cancel、timeout、deadline 都可能导致 Done 被 close
		<-errCtx.Done()
		fmt.Println("http server stop")
		return srv.Shutdown(errCtx)
	})
    // 利用channel 来缓存linux signal 信号
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-signalChanel: // 因为 kill -9 或其他而终止
				cancel()
			}
		}

		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}

	fmt.Println("all group done!")
}
