// +build integration

package main

import (
	"context"
	"fmt"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/zaur22/aura/pkg/config"
	"github.com/zaur22/aura/pkg/db"
	"github.com/zaur22/aura/pkg/rpc/api"
	"google.golang.org/grpc"
	"path"
	"testing"
)


func TestRPC(t *testing.T) {
	DBCloseFunc := initDB()
	defer DBCloseFunc()
	go Serv()

	<-servReady

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("integration:9090", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("Can`t to connect server: %s", err))
	}
	defer conn.Close()


	var (
		client = api.NewIncrementerClient(conn)
		ctx    = context.Background()
	)

	_, err = client.SetSettings(ctx, &api.SetSettingsRequest{
		MaxValue:      5,
		IncrementStep: 7,
	})
	if err != nil {
		t.Errorf("set settings operation error: %v", err)
	}

	_, err = client.IncrementNumber(ctx, &api.IncrementNumberRequest{})
	if err != nil {
		t.Errorf("increment operation error: %v", err)
	}

	res, err := client.GetNumber(ctx, &api.GetNumberRequest{})
	if err != nil {
		t.Errorf("get number operation error: %v", err)
		return
	}

	if res.Value != 2 {
		t.Errorf("bad result, expected 2, got %v", res.Value)
	}
}


func initDB() (closeDB func() error){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var configs = config.GetConfigs()

	dbDTO := db.ConnectDTO{
		Host:     configs.DBHost,
		Port:     configs.DBPort,
		UserName: configs.DBUsername,
		Password: configs.DBPassword,
		DBName:   configs.DBName,
		SSLMode:  configs.SSLMode,
	}

	dbConn := db.Connect(ctx, dbDTO)

	fixtures, err := testfixtures.New(
		testfixtures.Database(dbConn.DB),
		testfixtures.Dialect("postgres"),
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.Directory(path.Join("testdata")),
	)

	if err != nil {
		panic(fmt.Errorf("fixtures initing error: %v", err))
	}

	if err = fixtures.Load(); err != nil {
		panic(fmt.Errorf("fixtures loading error: %v", err))
	}

	return dbConn.Close
}