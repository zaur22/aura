package main

import (
	"context"
	"github.com/zaur22/aura/pkg/config"
	"github.com/zaur22/aura/pkg/db"
	"github.com/zaur22/aura/pkg/incrementor"
	"os"
)

func main() {
	os.Exit(Serve())
}

func Serve() (ret int) {
	defer func() {
		if r := recover(); r != nil {
			ret = 1
			//TODO: log
		}
	}()

	_, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	var configs = config.GetConfigs()

	var dbConnect = db.Connect(
		db.ConnectDTO{
			Host:     configs.DBHost,
			UserName: configs.DBUsername,
			Password: configs.DBPassword,
			DBName:   configs.DBName,
			SSLMode:  configs.SSLMode,
		})
	defer func() {
		if err := dbConnect.Close(); err != nil {
			//TODO: log
		}
	}()

	_, err := incrementer.NewIncrementer(
		incrementer.NewIncrementerDTO{
			DB: dbConnect,
		})

	if err != nil {
		panic(err)
	}

	return 0
}
