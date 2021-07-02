package thirdparty

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	_ "github.com/jackc/pgx/v4/stdlib"
	db "github.com/layouts/db/sqlc"
)

type JSONData []JSONLitter

type JSONLitter struct {
	ID       int64   `json:"id"`
	Entrance int32   `json:"entrance"`
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
	Number  int     `json:"number"`
}

type Numbers struct {
	StartNumber int `json:"start_number"`
	EndNumber   int `json:"end_number"`
	Step        int `json:"step"`
}

const (
	dbDriver     = "pgx"
	serverAdress = "0.0.0.0:8100"
	HOST         = "database"
	PORT         = 5432
)

func AddPathAndCreateSvgData() error {
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

	// Open our jsonFile
	jsonFile, err := os.Open("src/json/floors_svg.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var litters JSONData

	err = json.Unmarshal(byteValue, &litters)

	if err != nil {
		return err
	}

	for _, litter := range litters {
		arg := db.GetLayoutByLitterAndDoorParams{
			Parent: litter.ID,
			Door: sql.NullInt32{
				Int32: litter.Entrance,
				Valid: true,
			},
		}

		dbLayouts, err := store.GetLayoutByLitterAndDoor(context.Background(), arg)
		if err != nil {
			return err
		}
		if litter.ID == 137 {
			fmt.Println(arg)
			fmt.Println(dbLayouts)
		}
		for _, floor := range litter.Floors {
			for floorItt := floor.FloorNumber[0]; floorItt <= floor.FloorNumber[len(floor.FloorNumber)-1]; floorItt++ {
				for _, appartmentItt := range floor.Appartments {
					if appartmentItt.Number != 0 {
						number := appartmentItt.Number
						for dbLayoutItt := 0; dbLayoutItt < len(dbLayouts); dbLayoutItt++ {
							fmt.Println(int(dbLayouts[dbLayoutItt].Floor.Int32) == floorItt && dbLayouts[dbLayoutItt].Num.String == strconv.Itoa(number))
							if int(dbLayouts[dbLayoutItt].Floor.Int32) == floorItt && dbLayouts[dbLayoutItt].Num.String == strconv.Itoa(number) {
								fmt.Printf("dbLayoutsID: %v \n", dbLayouts[dbLayoutItt].ID)
								arr := db.UpdateSvgPathParams{
									ID: dbLayouts[dbLayoutItt].ID,
									SvgPath: sql.NullString{
										String: appartmentItt.Path,
										Valid:  true,
									},
								}
								err = store.UpdateSvgPath(context.Background(), arr)
								if err != nil {
									return err
								}
							}
						}
						continue
					}

					for number := appartmentItt.Numbers.StartNumber; number <= appartmentItt.Numbers.EndNumber; number += appartmentItt.Numbers.Step {
						for _, dbLayout := range dbLayouts {
							if int(dbLayout.Floor.Int32) == floorItt && dbLayout.Num.String == strconv.Itoa(number) {
								arr := db.UpdateSvgPathParams{
									ID: dbLayout.ID,
									SvgPath: sql.NullString{
										String: appartmentItt.Path,
										Valid:  true,
									},
								}
								err = store.UpdateSvgPath(context.Background(), arr)
								if err != nil {
									return err
								}
							}
						}
					}
				}
			}
		}
	}

	return nil
}
