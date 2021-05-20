package cronFunc

import (
	"context"
	"crypto/tls"
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
	LINK         = "https://bitrix.1dogma.ru/shahmatki/json.php"
)

func (d *Data) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func UpdateLayouts() error {
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
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(LINK)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	prod := &Data{}
	err = prod.FromJSON(resp.Body)
	if err != nil {
		return err
	}

	for i := 0; i < len(prod.Layouts); i++ {
		arg := db.UpdateLayoutParams{
			Area: sql.NullFloat64{
				Float64: prod.Layouts[i].Area,
				Valid:   true,
			},
			CitchenArea: sql.NullFloat64{
				Float64: prod.Layouts[i].CitchenArea,
				Valid:   true,
			},
			Door: sql.NullInt32{
				Int32: prod.Layouts[i].Door,
				Valid: true,
			},
			Floor: sql.NullInt32{
				Int32: prod.Layouts[i].Floor,
				Valid: true,
			},
			BitrixID: sql.NullInt32{
				Int32: int32(prod.Layouts[i].BitrixID),
				Valid: true,
			},
			LivingArea: sql.NullFloat64{
				Float64: prod.Layouts[i].LivingArea,
				Valid:   true,
			},
			Num: sql.NullString{
				String: prod.Layouts[i].Num,
				Valid:  true,
			},
			Price: sql.NullInt32{
				Int32: prod.Layouts[i].Price,
				Valid: true,
			},
			Room: sql.NullInt32{
				Int32: prod.Layouts[i].Room,
				Valid: true,
			},
			Status: sql.NullInt32{
				Int32: prod.Layouts[i].Status,
				Valid: true,
			},
			LayoutsUrl: sql.NullString{
				String: prod.Layouts[i].LayoutsURL,
				Valid:  true,
			},
			Type: sql.NullInt32{
				Int32: prod.Layouts[i].Type,
				Valid: true,
			},
		}

		_, err := store.UpdateLayout(context.Background(), arg)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return err
		}
	}
	return nil
}
