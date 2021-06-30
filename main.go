package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/layouts/api"
	"github.com/layouts/cronFunc"
	db "github.com/layouts/db/sqlc"
	"github.com/robfig/cron/v3"
)

const (
	dbDriver     = "pgx"
	serverAdress = "0.0.0.0:8100"
	HOST         = "database"
	PORT         = 5432
)

func main() {
	err := cronFunc.UpdateLayouts()
	if err != nil {
		fmt.Println(err)
	}
	c := cron.New()
	c.AddFunc("@daily", func() {
		err := cronFunc.UpdateLayouts()
		if err != nil {
			fmt.Println(err)
		}
	})
	c.Start()

	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, dbUser, dbPassword, dbName)

	conn, err := sql.Open(dbDriver, dsn)
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
