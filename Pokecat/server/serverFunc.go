package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

func Rand50Pokemon(pokemons []Pokemons) []Pokemons {
	var rand50Pokemons []Pokemons
	for i := 0; i <= 50; i++ {
		randID := rand.Intn(len(pokemons))              //50 pokemons.Id random 0 -> length-1
		pokemons[randID].coordinate.x = rand.Intn(1000) //random coordinate of pokemons
		pokemons[randID].coordinate.y = rand.Intn(1000) //random coordinate of pokemons
		rand50Pokemons = append(rand50Pokemons, pokemons[randID])
	}
	return rand50Pokemons
}
func ReadJSONFile(filename string, target interface{}) error {
	// Mở file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	// Parse JSON
	decoder := json.NewDecoder(file)
	err = decoder.Decode(target)
	if err != nil {
		return fmt.Errorf("could not decode JSON: %v", err)
	}

	return nil
}
