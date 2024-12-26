package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Pokemons struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Exp        string `json:"exp"`
	HP         string `json:"hp"`
	Attack     string `json:"attack"`
	Defense    string `json:"defense"`
	SpAttack   string `json:"sp_attack"`
	SpDefense  string `json:"sp_defense"`
	Speed      string `json:"speed"`
	TotalEVs   string `json:"total_evs"`
	coordinate Coordinate
}
type Coordinate struct {
	x int
	y int
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	//
	fmt.Println("Connected to server")
	reader := bufio.NewReader(os.Stdin)   //Read from user
	serverReader := bufio.NewReader(conn) //Read from Server
	fmt.Print("Which players do you want to play: ")
	player, _ := reader.ReadString('\n')
	conn.Write([]byte(strings.TrimSpace(player) + "\n"))

	authResponse, _ := serverReader.ReadString('\n')
	fmt.Print("Server: " + authResponse)

}
