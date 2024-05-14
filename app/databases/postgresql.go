package databases

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	PostgreSQLInstance *sql.DB
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
	Params   string
}

const maxOpenConns, maxIdleConns = 20, 10

func ConnectPostgre() {
	dbConfig := DatabaseConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		Params:   os.Getenv("DB_PARAMS"),
	}

	strConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name, dbConfig.Params)

	db, err := sql.Open("postgres", strConnection)
	if err != nil {
		fmt.Println("Error creating database connection:", err)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		fmt.Println("Error creating database connection:", err)
	}

	PostgreSQLInstance = db

	// defer func() {
	// 	if err := PostgreSQLInstance.Close(); err != nil {
	// 		fmt.Println("Error closing database connection:", err)
	// 	}
	// }()
}
