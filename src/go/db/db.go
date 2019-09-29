//package db
package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"io/ioutil"
	_ "log"
	_ "os"
	_ "time"
)

// DatabaseInfo contains the information that will create a database string
type DatabaseInfo struct {
	DatabaseName string `json:"databaseName"`
	Username     string `json:"username"`
	Port         string `json:"port"`
}

// createConnectionUrl creates a postgres url that can be used to access a database
func (databaseInfo DatabaseInfo) createConnectionUrl() string {
	return fmt.Sprintf("postgres://%s@localhost:%s/%s", databaseInfo.Username, databaseInfo.Port, databaseInfo.DatabaseName)
}

const SYSTEMFILE = "./system.json"

var ConnPool *pgxpool.Pool

// Creates a connection pool that other files can use
func init() {

	var databaseInfo DatabaseInfo

	// Reads from the SYSTEMFILE file
	fileByteData, err := ioutil.ReadFile(SYSTEMFILE)
	if err != nil {
		fmt.Println(err)
	}

	// Turns the []byte into a databaseInfo object
	err = json.Unmarshal(fileByteData, &databaseInfo)
	if err != nil {
		fmt.Println(err)
	}

	// Creates the database Connection pool
	ConnPool, err = pgxpool.Connect(context.Background(), databaseInfo.createConnectionUrl())
	if err != nil {
		fmt.Println(err)
	}
}
