package Db

import (
	"math/rand/v2"
)

type Player struct {
	Username         string     `json:"username"`
	Password         string     `json:"password"`
	Coordinate       coordinate `json:"coordinate"`
	CapturedPokemons []Pokemons `json:"captured_pokemons"`
}
type Pokemons struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Exp       string `json:"exp"`
	HP        string `json:"hp"`
	Attack    string `json:"attack"`
	Defense   string `json:"defense"`
	SpAttack  string `json:"sp_attack"`
	SpDefense string `json:"sp_defense"`
	Speed     string `json:"speed"`
	TotalEVs  string `json:"total_evs"`
}
type coordinate struct {
	x int `json:"x"`
	y int `json:"y"`
}

func PlayerDb(player1 Player, player2 Player, player3 Player) (Player, Player, Player) {
	player1 = Player{
		Username: "player1",
		Password: "123",
		Coordinate: coordinate{
			x: rand.IntN(1000),
			y: rand.IntN(1000),
		},
	}
	player2 = Player{
		Username: "player2",
		Password: "123",
		Coordinate: coordinate{
			x: rand.IntN(1000),
			y: rand.IntN(1000),
		},
	}
	player3 = Player{
		Username: "player3",
		Password: "123",
		Coordinate: coordinate{
			x: rand.IntN(1000),
			y: rand.IntN(1000),
		},
	}
	return player1, player2, player3
}
