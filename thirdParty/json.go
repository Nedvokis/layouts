package thirdparty

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Build struct {
	BitrixID int32  `json:"ID,string"`
	Name     string `json:"NAME"`
}
type Litter struct {
	BitrixID int32  `json:"ID,string"`
	Name     string `json:"NAME"`
	Parent   int64  `json:"PARENT,string"`
}
type Layout struct {
	BitrixID    int32   `json:"ID,string"`
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
	SvgPath     string  `json:"-"`
}
type itemsStatuses struct {
	StaRooms    []string `json:"ROOMS"`
	StaStatuses []string `json:"STATUSES"`
	StaTypes    []string `json:"TYPES"`
}

type Data struct {
	Builds  []Build  `json:"BUILDS"`
	Litters []Litter `json:"LITERS"`
	Layouts []Layout `json:"OBJECTS"`
	// itemsStatuses `json:"VALUES"`
}

func (d *Data) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func GetLayouts() {
	resp, err := http.Get("https://bitrix.1dogma.ru/shahmatki/json.php")
	if err != nil {
		fmt.Printf("Error %v", err)
		return
	}
	defer resp.Body.Close()

	prod := &Data{}
	err = json.NewDecoder(resp.Body).Decode(prod)
	if err != nil {
		fmt.Printf("Error %v", err)
		return
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Result: %v\n", prod.Layouts[i].Area)
		fmt.Printf("Result: %v\n", prod.Layouts[i].LayoutID)
	}

	fmt.Printf("Result: %v\n", len(prod.Layouts))
}
