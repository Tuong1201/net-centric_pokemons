package main

type Item struct {
	ItemID   string
	ItemName string
	Quantity int
}

type Location struct {
	X    int
	Y    int
	Zone string
}

type Pokemon struct {
	PokemonID string
	Name      string
	Level     int
	CP        int
	HP        int
	IsShiny   bool
}

type Player struct {
	PlayerID   string
	Username   string
	Level      int
	Health     int
	Mana       int
	Experience int
	IsAlive    bool
	Inventory  []Item
	Location   Location
	Pokemons   []Pokemon
}
