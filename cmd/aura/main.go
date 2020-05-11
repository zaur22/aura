package main

import (
	"context"
	"github.com/oklog/run"
	"github.com/sirupsen/logrus"
	"github.com/zaur22/aura/pkg/config"
	"github.com/zaur22/aura/pkg/db"
	"github.com/zaur22/aura/pkg/incrementer"
	"github.com/zaur22/aura/pkg/rpc"
	"os"
	"os/signal"
	"syscall"
	"time"
)


//servReady канал закрывается, когда сервер готов к работе
//нужно для тестов
var servReady = make(chan struct{})

func main() {
	os.Exit(Serv())
}

func Serv() int {

	logrus.SetLevel(logrus.DebugLevel)
	c := config.GetConfigs()

	var (
		ctxDB, cancelDB   = context.WithCancel(context.Background())
		ctxRPC, cancelRPC = context.WithCancel(context.Background())
		ctxIncr           = context.Background()
	)

	conn := db.Connect(
		ctxDB,
		db.ConnectDTO{
			Host:     c.DBHost,
			Port:     c.DBPort,
			UserName: c.DBUsername,
			Password: c.DBPassword,
			DBName:   c.DBName,
			SSLMode:  c.SSLMode,
		},
	)

	var g run.Group
	{
		// Обработчик сисколов на прерывание процесса
		term := make(chan os.Signal, 1)
		signal.Notify(term, os.Interrupt, syscall.SIGTERM)
		cancel := make(chan struct{})
		g.Add(
			func() error {
				select {
				case <-term:
					logrus.Warningf("Received SIGTERM, exiting gracefully...")
				case <-cancel:
				}
				return nil
			},
			func(err error) {
				close(cancel)
			},
		)
	}
	{
		// Система завершения коннектов к БД
		g.Add(
			func() error {
				<-ctxDB.Done()
				err := conn.Close()
				return err
			},
			func(err error) {
				logrus.Warningf("Closing DB connection")
				cancelDB()
			},
		)
	}
	{
		// Запуск и завершение rpc сервера
		g.Add(
			func() error {
				defer logrus.Warningf("RPC server closed\n")
				incr, err := incrementer.NewIncrementer(
					ctxIncr,
					incrementer.NewIncrementerDTO{
						DBClient: conn,
					})
				if err != nil {
					return err
				}

				rpcServ := rpc.NewServer(rpc.NewServerDTO{
					Incrementer: incr,
				})

				if err := rpcServ.Run(ctxRPC, c.HttpAddress); err != nil {
					return err
				}
				return nil
			},
			func(err error) {
				logrus.Warningf("Closing RPC server\n")
				cancelRPC()
			})
	}

	{
		// указание готовности системы
		var stop = make(chan struct{})
		g.Add(func() error {
			time.Sleep(100 * time.Millisecond)
			close(servReady)
			logrus.Infof("Server started and listen, %v", c.HttpAddress)
			<-stop
			return nil
		},
		func(err error) {
			close(stop)
		})
	}

	if err := g.Run(); err != nil {
		logrus.Warningf("ERROR: %v\n", err)
		return 1
	}
	return 0
}
