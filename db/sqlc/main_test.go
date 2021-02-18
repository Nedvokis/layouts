package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	dbDriver = "pgx"
	dbSource = "postgresql://root:WEBdeveloepr1452@localhost:5432/layouts?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot  connect to db: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
