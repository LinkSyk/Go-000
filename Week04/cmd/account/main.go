// +build wireinject

package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"week04/api/account/v1"
	"week04/internal/biz"
	"week04/internal/repo"
	"week04/internal/service"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var (
	actServiceSet = wire.NewSet(biz.NewBizImpl, repo.NewMysqlDBImpl, service.NewAccountService)
)

// 有关 accountService的依赖注入
func initAccountService(db *sqlx.DB) (*service.AccountService, error) {
	wire.Build(actServiceSet)
	return &service.AccountService{}, nil
}

func StartGrpc(ctx context.Context, address string, accountSvc *service.AccountService) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	account.RegisterAccountServer(s, accountSvc)

	go func() {
		<-ctx.Done()
		s.Stop()
	}()

	return s.Serve(l)
}

func main() {
	mysqlConn, err := repo.ConnectMysql("root", "pass", "localhost", "user", 3306)
	if err != nil {
		log.Fatalln("init db error")
	}

	// 利用 wire 进行依赖注入
	accountSvc, err := initAccountService(mysqlConn)
	if err != nil {
		log.Fatalln("init accountService Depenices failed")
	}

	g, ctx := errgroup.WithContext(context.Background())
	sig := make(chan os.Signal)

	signal.Notify(sig, os.Interrupt, os.Kill)

	g.Go(func() error {
		select {
		case <-ctx.Done():
			return errors.New("start grpc error")
		case <-sig:
			return errors.New("stop server")
		}
	})

	g.Go(func() error {
		return StartGrpc(ctx, "0.0.0.0:8080", accountSvc)
	})

	if err := g.Wait(); err != nil {
		log.Infoln("stop server")
	}

}
