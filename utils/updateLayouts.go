package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	db "github.com/layouts/db/sqlc"
)

type Data struct {
	Layouts []Layout `json:"OBJECTS"`
}

type Layout struct {
	BitrixID    int64   `json:"ID,string"`
	Parent      int64   `json:"PARENT,string"`
	Area        float64 `json:"AREA,string"`
	CitchenArea float64 `json:"CITCHEN_AREA,string"`
	Door        int32   `json:"DOOR,string"`
	Floor       int32   `json:"FLOOR,string"`
	LayoutID    int32   `json:"LAYOUT_ID,string"`
	LivingArea  float64 `json:"LIVING_AREA,string"`
	Num         string  `json:"NUM"`
	Price       int32   `json:"PRICE,string"`
	Room        int32   `json:"ROOM,string"`
	Status      int32   `json:"STATUS,string"`
	Type        int32   `json:"TYPE,string"`
	LayoutsURL  string  `json:"LAYOUTS_URL"`
	SvgPath     string  `json:"-"`
}

const (
	dbDriver     = "pgx"
	serverAdress = "0.0.0.0:8100"
	HOST         = "database"
	PORT         = 5432
)

func (d *Data) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func updateLayouts() error {
	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, dbUser, dbPassword, dbName)

	conn, err := sql.Open(dbDriver, dsn)
	store := db.NewStore(conn)
	if err != nil {
		return err
	}
	resp, err := http.Get("https://bitrix.1dogma.ru/shahmatki/json.php")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	prod := &Data{}
	err = prod.FromJSON(resp.Body)
	if err != nil {
		return err
	}
}
