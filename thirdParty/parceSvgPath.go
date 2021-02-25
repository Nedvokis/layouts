package thirdparty

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	db "github.com/layouts/db/sqlc"
)

type JSONData []JSONLitter

type JSONLitter struct {
	ID       int     `json:"id"`
	Entrance int     `json:"entrance"`
	Floors   []Floor `json:"floors"`
}

type Floor struct {
	FloorNumber  []int        `json:"floor_number"`
	HeightForSvg float32      `json:"height_for_svg"`
	WidthForSvg  float32      `json:"width_for_svg"`
	Appartments  []Appartment `json:"appartaments"`
}

type Appartment struct {
	Path    string  `json:"path"`
	Numbers Numbers `json:"numbers"`
}

type Numbers struct {
	StartNumber int `json:"start_number"`
	Endnumber   int `json:"end_number"`
	Step        int `json:"step"`
}

func AddPathAndCreateSvgData() {
	conn, err := sql.Open(dbDriver, dbSource)
	store := db.NewStore(conn)
	if err != nil {
		log.Fatal("cannot  connect to db: ", err)
	}

	dbLayouts, err := store.GetAllListLayouts(context.Background())

	fmt.Printf("lol: %v \n", dbLayouts[0])

	// Open our jsonFile
	jsonFile, err := os.Open("src/json/floors_svg.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var litters JSONData

	err = json.Unmarshal(byteValue, &litters)

	if err != nil {
		fmt.Printf("Error ocured: %v", err)
	}

	for i := 0; i < len(litters); i++ {
		for fK := 0; fK < len(litters[i].Floors); fK++ {
			fmt.Printf("Litter: %v \n", litters[i].ID)
			fmt.Printf("What: %v \n", litters[i].Floors[fK].FloorNumber[1])
		}
	}

}
