package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TPoint struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type LocCompressed struct {
	Id     int `json:"id"`
	X      int `json:"x"`
	Y      int `json:"y"`
	Radius int `json:"radius"`
}

type EndSectorCompressed struct {
	IsEnd bool            `json:"isEnd"`
	Locs  []LocCompressed `json:"locs"`
}

type SectorCompressed interface {
}

type DividedSectorCompressed struct {
	Center TPoint           `json:"center"`
	NW     SectorCompressed `json:"NW"`
	NE     SectorCompressed `json:"NE"`
	SW     SectorCompressed `json:"SW"`
	SE     SectorCompressed `json:"SE"`
}

func main() {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	dataBytes := []byte(data)

	input := []interface{}{}
	json.Unmarshal(dataBytes, &input)
	fmt.Println(input)
}
