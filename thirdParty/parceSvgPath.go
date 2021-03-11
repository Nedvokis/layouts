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
	Endnumber   int `json:"end_number"`
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

	for i := 0; i < len(litters); i++ {
		arg := db.GetLayoutByLitterAndDoorParams{
			Parent: litters[i].ID,
			Door: sql.NullInt32{
				Int32: litters[i].Entrance,
				Valid: true,
			},
		}

		dbLayouts, err := store.GetLayoutByLitterAndDoor(context.Background(), arg)
		if err != nil {
			return err
		}
		// fmt.Printf("Length of layouts array: %v \n", litters[i].ID)
		for fK := 0; fK < len(litters[i].Floors); fK++ {
			for floorItt := litters[i].Floors[fK].FloorNumber[0]; floorItt <= litters[i].Floors[fK].FloorNumber[len(litters[i].Floors[fK].FloorNumber)-1]; floorItt++ {
				for appartmentItt := 0; appartmentItt < len(litters[i].Floors[fK].Appartments); appartmentItt++ {
					if (litters[i].ID == 65 || litters[i].ID == 67 || litters[i].ID == 68 || litters[i].ID == 69 || litters[i].ID == 70) && litters[i].Floors[fK].Appartments[appartmentItt].Number != 0 {
						fmt.Printf("one number: %v \n", litters[i].Floors[fK].Appartments[appartmentItt].Number)
						number := litters[i].Floors[fK].Appartments[appartmentItt].Number
						for dbLayoutItt := 0; dbLayoutItt < len(dbLayouts); dbLayoutItt++ {
							if int(dbLayouts[dbLayoutItt].Floor.Int32) == floorItt && dbLayouts[dbLayoutItt].Num.String == strconv.Itoa(number) {
								fmt.Printf("dbLayoutsID: %v \n", dbLayouts[dbLayoutItt].ID)
								arr := db.UpdateSvgPathParams{
									ID: dbLayouts[dbLayoutItt].ID,
									SvgPath: sql.NullString{
										String: litters[i].Floors[fK].Appartments[appartmentItt].Path,
										Valid:  true,
									},
								}
								err = store.UpdateSvgPath(context.Background(), arr)
								if err != nil {
									fmt.Printf("Error ocured: %v", err)
								}
							}
						}
						continue
					}
					for number := litters[i].Floors[fK].Appartments[appartmentItt].Numbers.StartNumber; number < litters[i].Floors[fK].Appartments[appartmentItt].Numbers.Endnumber; number += litters[i].Floors[fK].Appartments[appartmentItt].Numbers.Step {
						for dbLayoutItt := 0; dbLayoutItt < len(dbLayouts); dbLayoutItt++ {
							if int(dbLayouts[dbLayoutItt].Floor.Int32) == floorItt && dbLayouts[dbLayoutItt].Num.String == strconv.Itoa(number) {
								arr := db.UpdateSvgPathParams{
									ID: dbLayouts[dbLayoutItt].ID,
									SvgPath: sql.NullString{
										String: litters[i].Floors[fK].Appartments[appartmentItt].Path,
										Valid:  true,
									},
								}
								err = store.UpdateSvgPath(context.Background(), arr)
								if err != nil {
									fmt.Printf("Error ocured: %v", err)
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
