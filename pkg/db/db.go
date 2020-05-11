package db

import (
	"context"
	"database/sql/driver"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/common/log"
	"time"
)

type ConnectDTO struct {
	Host     string
	Port     uint16
	UserName string
	Password string
	DBName   string
	SSLMode  string
}

func Connect(ctx context.Context, dto ConnectDTO) *sqlx.DB {

	config, _ := pgx.ParseConfig("")
	config.Host = dto.Host
	config.Port = dto.Port
	config.User = dto.UserName
	config.Password = dto.Password
	config.Database = dto.DBName

	db := stdlib.OpenDB(*config)

	pingCount := 0
	var err error
	for err = db.Ping(); err == driver.ErrBadConn; {
		log.Infof("Can't ping database. Waiting to retrying 200 ms...\n")
		time.Sleep(200 * time.Millisecond)
		pingCount++
		if pingCount == 3 {
			panic(fmt.Errorf("can't connect to DB, bad connection: %v", err))
		}
	}

	if err != nil {
		panic(fmt.Errorf("can't connect to DB: %v", err))
	}

	return sqlx.NewDb(db, "psx")
}
