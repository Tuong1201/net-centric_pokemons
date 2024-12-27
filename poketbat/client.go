package main

import (
	"encoding/json"
	"fmt"
	"net"
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
	Turn            int       `json:"turn"`
	Winner          string    `json:"winner"`
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	for {
		// Receive game state
		var gameState GameState
		err := decoder.Decode(&gameState)
		if err != nil {
			fmt.Println("Error decoding game state:", err)
			return
		}

		if gameState.Winner != "" {
			fmt.Printf("Game Over! Winner: %s\n", gameState.Winner)
			break
		}

		fmt.Printf("Game State: %+v\n", gameState)
		if gameState.Turn == 0 {
			fmt.Println("Your turn! Enter action (attack/switch):")
		} else {
			fmt.Println("Waiting for the other player...")
			continue
		}

		// Get action from the user
		var action string
		fmt.Scanln(&action)

		// Send action to server
		err = encoder.Encode(action)
		if err != nil {
			fmt.Println("Error sending action:", err)
			return
		}
	}
}
