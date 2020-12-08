## 学习笔记

```go
// 这道题目的意思我不是特别清楚，我的理解是启动多个 http 服务，然后注册一个信号，我触发信号后优雅关闭各个 http 服务。
package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

// 基于 errgroup 实现一个 http server 的启动和关闭,
// 以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

// 需要使用 context 来做goroutine的生命周期的管控，使用chan来通知状态，使用信号来触发停止服务。
func startServer(addr string, ctx context.Context) error {
	server := http.Server{
		Addr:    addr,
		Handler: nil,
	}

	go func() {
		<-ctx.Done()
		logrus.Infoln("recv stop")
		_ = server.Shutdown(context.Background())
	}()

	return server.ListenAndServe()
}

func main() {
	sig := make(chan os.Signal)

	// ctx 作为闭包参数传递到 startServer 中。
	g, ctx := errgroup.WithContext(context.Background())

	// 注册信号
	signal.Notify(sig, os.Interrupt)

	// 处理信号的服务
	g.Go(func() error {
		<-sig
		return errors.New("stop server")
	})

	// 启动 server1
	g.Go(func() error {
		if err := startServer("0.0.0.0:8080", ctx); err != nil {
			logrus.Infof("stop server1, %v", err)
			return err
		}
		return nil
	})

	// 启动 server2
	g.Go(func() error {
		if err := startServer("0.0.0.0:8081", ctx); err != nil {
			logrus.Infof("stop server2, %v", err)
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return
	}
}

```