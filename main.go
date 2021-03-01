package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/layouts/api"
	db "github.com/layouts/db/sqlc"
)

const (
	dbDriver     = "pgx"
	dbSource     = "postgresql://root:WEBdeveloepr1452@localhost:5432/layouts?sslmode=disable"
	serverAdress = "0.0.0.0:8100"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
