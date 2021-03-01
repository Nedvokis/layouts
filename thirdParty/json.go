package thirdparty

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
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
	StaRooms    map[string]interface{} `json:"ROOMS"`
	StaStatuses map[string]interface{} `json:"STATUSES"`
	StaTypes    []string               `json:"TYPES"`
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

const (
	dbDriver = "pgx"
	dbSource = "postgresql://root:WEBdeveloepr1452@localhost:5432/layouts?sslmode=disable"
)

func GetLayouts() error {
	conn, err := sql.Open(dbDriver, dbSource)
	store := db.NewStore(conn)
	if err != nil {
		return err
	}

	litters, err := store.GetListAllLitters(context.Background())
	staStatuse, err := store.GetListAllStaStatuse(context.Background())
	staRoom, err := store.GetListAllStaRoom(context.Background())
	staType, err := store.GetListAllStaType(context.Background())

	// for i := 0; i < len(litters); i++ {
	// 	fmt.Printf("Литер: %v \n", litters[i])
	// }
	// for i := 0; i < len(staRoom); i++ {
	// 	fmt.Printf("Литер: %v \n", staRoom[i])
	// }
	// for i := 0; i < len(staStatuse); i++ {
	// 	fmt.Printf("Литер: %v \n", staStatuse[i])
	// }
	// for i := 0; i < len(staType); i++ {
	// 	fmt.Printf("Литер: %v \n", staType[i])
	// }

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

	for i, v := range prod.StaRooms {
		key, err := strconv.Atoi(i)
		if err != nil {
			return err
		}
		arg := db.CreateStaRoomParams{
			BitrixID: int64(key),
			TypeName: sql.NullString{
				String: v.(string),
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

	for i := 0; i < len(prod.Builds); i++ {
		arg := db.CreateComplexParams{
			BitrixID: prod.Builds[i].BitrixID,
			Name: sql.NullString{
				String: prod.Builds[i].Name,
				Valid:  true,
			},
		}
		store.CreateComplex(context.Background(), arg)
	}
	for i := 0; i < len(prod.Litters); i++ {
		arg := db.CreateLitterParams{
			BitrixID: prod.Litters[i].BitrixID,
			Name: sql.NullString{
				String: prod.Litters[i].Name,
				Valid:  true,
			},
			Parent: prod.Litters[i].Parent,
		}
		store.CreateLitter(context.Background(), arg)
	}

	for i := 0; i < len(prod.Layouts); i++ {
		litterExist := false
		for litterKey := 0; litterKey < len(litters); litterKey++ {
			if litters[litterKey].BitrixID == prod.Layouts[i].Parent {
				litterExist = true
			}
		}
		staStatuseExist := false
		for k := 0; k < len(staStatuse); k++ {
			if staStatuse[k].BitrixID == int64(prod.Layouts[i].Status) {
				staStatuseExist = true
			}
		}
		staRoomExist := false
		for k := 0; k < len(staRoom); k++ {
			if staRoom[k].BitrixID == int64(prod.Layouts[i].Room) {
				staRoomExist = true
			}
		}
		staTypeExist := false
		for k := 0; k < len(staType); k++ {
			if staType[k].BitrixID == int64(prod.Layouts[i].Type) {
				staTypeExist = true
			}
		}

		arg := db.CreateLayoutParams{
			Parent: prod.Layouts[i].Parent,
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
			LayoutID: sql.NullInt32{
				Int32: prod.Layouts[i].LayoutID,
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
			SvgPath: sql.NullString{
				String: prod.Layouts[i].SvgPath,
				Valid:  false,
			},
			Type: sql.NullInt32{
				Int32: prod.Layouts[i].Type,
				Valid: true,
			},
		}

		// fmt.Printf("WaT:%v, %v, %v, %v", !litterExist, !staStatuseExist, !staRoomExist, !staTypeExist)
		if !litterExist || !staStatuseExist || !staRoomExist || !staTypeExist {
			continue
		}

		_, err := store.CreateLayout(context.Background(), arg)
		if err != nil {
			return err
		}
		// fmt.Printf("Success %v \n", complex)
	}
	return nil
}
