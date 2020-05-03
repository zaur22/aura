package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type ConnectDTO struct {
	Host     string
	UserName string
	Password string
	DBName   string
	SSLMode  string
}

func Connect(dto ConnectDTO) *sqlx.DB {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"user=%v password=%v dbname=%v host=%v sslmode=%v",
			dto.UserName, dto.Password, dto.DBName, dto.Host, dto.SSLMode),
	)

	if err != nil {
		panic(fmt.Sprintf("could not connect to db: %v\n", err))
	}

	return db
}
