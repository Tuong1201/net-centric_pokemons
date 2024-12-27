package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

type Pokemon struct {
	Name     string `json:"name"`
	HP       int    `json:"hp"`
	Attack   int    `json:"attack"`
	Defense  int    `json:"defense"`
	Speed    int    `json:"speed"`
	IsActive bool   `json:"is_active"`
}

type GameState struct {
	Player1Pokemons []Pokemon `json:"player1_pokemons"`
	Player2Pokemons []Pokemon `json:"player2_pokemons"`
	Turn            int       `json:"turn"` // 0 for Player 1, 1 for Player 2
	Winner          string    `json:"winner"`
}

var (
	gameState = GameState{
		Player1Pokemons: []Pokemon{
			{"Pikachu", 100, 50, 30, 90, true},
			{"Charmander", 120, 60, 40, 80, false},
			{"Squirtle", 140, 40, 60, 70, false},
		},
		Player2Pokemons: []Pokemon{
			{"Bulbasaur", 130, 50, 50, 75, true},
			{"Eevee", 110, 55, 45, 85, false},
			{"Meowth", 90, 45, 40, 95, false},
		},
		Turn:   0,
		Winner: "",
	}
	mu sync.Mutex
)

func handlePlayer(conn net.Conn, playerID int) {
	defer conn.Close()
	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	for {
		mu.Lock()
		if gameState.Winner != "" {
			mu.Unlock()
			return
		}

		if gameState.Turn != playerID {
			mu.Unlock()
			continue
		}

		// Send current game state
		err := encoder.Encode(gameState)
		if err != nil {
			fmt.Println("Error sending game state:", err)
			mu.Unlock()
			return
		}

		// Receive action from the player
		var action string
		err = decoder.Decode(&action)
		if err != nil {
			fmt.Println("Error receiving action:", err)
			mu.Unlock()
			return
		}

		// Process action
		processAction(playerID, action)

		// Check if the game is over
		checkWinner()

		// Switch turn
		gameState.Turn = 1 - gameState.Turn
		mu.Unlock()
	}
}

func processAction(playerID int, action string) {
	attacker := &gameState.Player1Pokemons
	defender := &gameState.Player2Pokemons

	if playerID == 1 {
		attacker, defender = defender, attacker
	}

	activeAttacker := findActivePokemon(attacker)
	activeDefender := findActivePokemon(defender)

	if action == "attack" {
		damage := activeAttacker.Attack - activeDefender.Defense
		if damage < 0 {
			damage = 0
		}
		activeDefender.HP -= damage
		fmt.Printf("%s attacked %s for %d damage\n", activeAttacker.Name, activeDefender.Name, damage)
		if activeDefender.HP <= 0 {
			activeDefender.HP = 0
			activeDefender.IsActive = false
			fmt.Printf("%s fainted!\n", activeDefender.Name)
		}
	} else if action == "switch" {
		for i, pokemon := range *attacker {
			if !pokemon.IsActive && pokemon.HP > 0 {
				(*attacker)[i].IsActive = true
				activeAttacker.IsActive = false
				fmt.Printf("%s switched to %s\n", pokemon.Name, (*attacker)[i].Name)
				break
			}
		}
	}
}

func findActivePokemon(team *[]Pokemon) *Pokemon {
	for i := range *team {
		if (*team)[i].IsActive {
			return &(*team)[i]
		}
	}
	return nil
}

func checkWinner() {
	if allFainted(gameState.Player1Pokemons) {
		gameState.Winner = "Player 2"
		fmt.Println("Player 2 wins!")
	} else if allFainted(gameState.Player2Pokemons) {
		gameState.Winner = "Player 1"
		fmt.Println("Player 1 wins!")
	}
}

func allFainted(team []Pokemon) bool {
	for _, pokemon := range team {
		if pokemon.HP > 0 {
			return false
		}
	}
	return true
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is running on port 8080...")

	for i := 0; i < 2; i++ {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handlePlayer(conn, i)
	}
}
