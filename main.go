package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/layouts/cronFunc"
)

const (
	dbDriver     = "pgx"
	serverAdress = "0.0.0.0:8100"
	HOST         = "database"
	PORT         = 5432
)

func main() {
	err := cronFunc.UpdateLayouts()

	fmt.Println("Let's a go!")
	if err != nil {
		fmt.Println(err)
	}

	// c := cron.New()
	// c.AddFunc("@every 1m", func() { cronFunc.UpdateLayouts() })
	// c.Start()

	// dbUser, dbPassword, dbName :=
	// 	os.Getenv("POSTGRES_USER"),
	// 	os.Getenv("POSTGRES_PASSWORD"),
	// 	os.Getenv("POSTGRES_DB")
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	HOST, PORT, dbUser, dbPassword, dbName)

	// conn, err := sql.Open(dbDriver, dsn)
	// if err != nil {
	// 	log.Fatal("cannot connect to db: ", err)
	// }
	// store := db.NewStore(conn)
	// server := api.NewServer(store)

	// err = server.Start(serverAdress)
	// if err != nil {
	// 	log.Fatal("cannot start server: ", err)
	// }

}
