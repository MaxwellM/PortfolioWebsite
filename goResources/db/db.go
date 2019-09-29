package db

import (
	"log"
	"os"
	"time"

	"github.com/jackc/pgx"
)

var ConnPool *pgx.ConnPool

func init() {
	config, err := pgx.ParseEnvLibpq()
	if err != nil {
		log.Println("Unable to parse environment:", err)
		os.Exit(1)
	}

	ConnPool, err = pgx.NewConnPool(pgx.ConnPoolConfig{ConnConfig: config, AcquireTimeout: time.Second * 2, MaxConnections: 64})
	if err != nil {
		log.Printf("Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
}
