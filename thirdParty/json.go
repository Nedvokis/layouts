package thirdparty

import (
	"context"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	_ "github.com/jackc/pgx/v4/stdlib"
	db "github.com/layouts/db/sqlc"
)

type Build struct {
	BitrixID int64  `json:"ID,string"`
	Name     string `json:"NAME"`
}
type Litter struct {
	BitrixID int64  `json:"ID,string"`
	Name     string `json:"NAME"`
	Parent   int64  `json:"PARENT,string"`
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

type Room struct {
	Num1 string `json:"0"`
}

type itemsStatuses struct {
	StaRooms    []StaRooms             `json:"ROOMS"`
	StaStatuses map[string]interface{} `json:"STATUSES"`
	StaTypes    []string               `json:"TYPES"`
}
type StaRooms struct {
	BitrixID int64  `json:"id" form:"id"`
	TypeName string `json:"value" form:"value"`
}

type Data struct {
	Builds        []Build  `json:"BUILDS"`
	Litters       []Litter `json:"LITERS"`
	Layouts       []Layout `json:"OBJECTS"`
	itemsStatuses `json:"VALUES"`
}

func (d *Data) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func GetLayouts() error {
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

	complexes, err := store.GetListAllComplexes(context.Background())
	litters, err := store.GetListAllLitters(context.Background())
	layouts, err := store.GetAllListLayouts(context.Background())
	// staStatuse, err := store.GetListAllStaStatuse(context.Background())
	// staRoom, err := store.GetListAllStaRoom(context.Background())
	// staType, err := store.GetListAllStaType(context.Background())

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://bitrix.1dogma.ru/shahmatki/json.php")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	prod := &Data{}
	err = prod.FromJSON(resp.Body)
	if err != nil {
		return err
	}

	for _, v := range prod.StaRooms {
		arg := db.CreateStaRoomParams{
			BitrixID: v.BitrixID,
			TypeName: sql.NullString{
				String: v.TypeName,
				Valid:  true,
			},
		}
		store.CreateStaRoom(context.Background(), arg)

		// fmt.Printf("key: %v, value: %v \n", i, v)
	}
	for i, v := range prod.StaStatuses {
		key, err := strconv.Atoi(i)
		if err != nil {
			return err
		}
		arg := db.CreateStaStatuseParams{
			BitrixID: int64(key),
			TypeName: sql.NullString{
				String: v.(string),
				Valid:  true,
			},
		}
		store.CreateStaStatuse(context.Background(), arg)

		// fmt.Printf("key: %v, value: %v \n", i, v)
	}
	for i, v := range prod.StaTypes {
		arg := db.CreateStaTypeParams{
			BitrixID: int64(i),
			TypeName: sql.NullString{
				String: v,
				Valid:  true,
			},
		}
		store.CreateStaType(context.Background(), arg)

		// fmt.Printf("key: %v, value: %v \n", i, v)
	}

	mComplex := make(map[int64]bool)
	cComplex := []Build{}

	for _, item := range complexes {
		mComplex[item.BitrixID] = true
	}
	for _, item := range prod.Builds {
		if _, ok := mComplex[item.BitrixID]; !ok {
			cComplex = append(cComplex, item)
		}
	}
	for _, newComplex := range cComplex {
		arg := db.CreateComplexParams{
			BitrixID: newComplex.BitrixID,
			Name: sql.NullString{
				String: newComplex.Name,
				Valid:  true,
			},
		}
		_, err = store.CreateComplex(context.Background(), arg)
		if err != nil {
			return err
		}
	}

	mLitters := make(map[int64]bool)
	cLitters := []Litter{}

	for _, item := range litters {
		mLitters[item.BitrixID] = true
	}
	for _, item := range prod.Litters {
		if _, ok := mLitters[item.BitrixID]; !ok {
			cLitters = append(cLitters, item)
		}
	}
	for _, newLitter := range cLitters {
		arg := db.CreateLitterParams{
			BitrixID: newLitter.BitrixID,
			Name: sql.NullString{
				String: newLitter.Name,
				Valid:  true,
			},
			Parent: newLitter.Parent,
		}
		_, err := store.CreateLitter(context.Background(), arg)
		if err != nil {
			return err
		}
	}

	mLayouts := make(map[int64]bool)
	cLayouts := []Layout{}

	for _, item := range layouts {
		mLayouts[int64(item.BitrixID.Int32)] = true
	}
	for _, item := range prod.Layouts {
		if _, ok := mLayouts[item.BitrixID]; !ok {
			cLayouts = append(cLayouts, item)
		}
	}
	for _, newLayout := range cLayouts {
		arg := db.CreateLayoutParams{
			Parent: newLayout.Parent,
			Area: sql.NullFloat64{
				Float64: newLayout.Area,
				Valid:   true,
			},
			CitchenArea: sql.NullFloat64{
				Float64: newLayout.CitchenArea,
				Valid:   true,
			},
			Door: sql.NullInt32{
				Int32: newLayout.Door,
				Valid: true,
			},
			Floor: sql.NullInt32{
				Int32: newLayout.Floor,
				Valid: true,
			},
			BitrixID: sql.NullInt32{
				Int32: int32(newLayout.BitrixID),
				Valid: true,
			},
			LayoutID: sql.NullInt32{
				Int32: newLayout.LayoutID,
				Valid: true,
			},
			LivingArea: sql.NullFloat64{
				Float64: newLayout.LivingArea,
				Valid:   true,
			},
			Num: sql.NullString{
				String: newLayout.Num,
				Valid:  true,
			},
			Price: sql.NullInt32{
				Int32: newLayout.Price,
				Valid: true,
			},
			Room: sql.NullInt32{
				Int32: newLayout.Room,
				Valid: true,
			},
			Status: sql.NullInt32{
				Int32: newLayout.Status,
				Valid: true,
			},
		}
		if newLayout.Status != 0 {
			_, err := store.CreateLayout(context.Background(), arg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
