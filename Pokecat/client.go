package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Pokemons struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Exp        string     `json:"exp"`
	HP         string     `json:"hp"`
	Attack     string     `json:"attack"`
	Defense    string     `json:"defense"`
	SpAttack   string     `json:"sp_attack"`
	SpDefense  string     `json:"sp_defense"`
	Speed      string     `json:"speed"`
	TotalEVs   string     `json:"total_evs"`
	coordinate Coordinate `json:"coordinate"`
}
type Coordinate struct {
	x int `json:"x"`
	y int `json:"y"`
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to server")
	//
	var pokemon []Pokemons
	decoder := json.NewDecoder(conn)
	err = decoder.Decode(&pokemon)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	for _, msg := range pokemon {
		fmt.Printf("Id=%s", msg.Id)
	}
}
