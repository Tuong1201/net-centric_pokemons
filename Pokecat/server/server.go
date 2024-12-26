package main

import (
	"bufio"
	"math/rand/v2"

	// "encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Server listenting on port 8080")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	var pokemons []Pokemons
	err := ReadJSONFile("C:/Users/hp/Desktop/projects/net-centric/Pokedex/pokedex.json", &pokemons)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}
	// var x, y int
	p := Rand50Pokemon(pokemons)                //50 random pokemon
	ClientConnect := conn.RemoteAddr().String() //Lấy địa chỉ IP của client và hiển thị
	fmt.Println("New client connected:", ClientConnect)
	//
	saveLinksToJSON(p, "rand50Pokemons.json")
	//
	Player := NewPlayer{
		Username: "player1",
		Password: "123",
		PlayerCoordinate: NewPlayerCoordinate{
			PlayerX: rand.IntN(1000),
			PlayerY: rand.IntN(1000),
		},
	}
	reader := bufio.NewReader(conn)
	//clasify username and password
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username:", err)
		conn.Write([]byte("ERROR: Failed to read username.\n"))
		return
	}
	username = strings.TrimSpace(username)
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username:", err)
		conn.Write([]byte("ERROR: Failed to read username.\n"))
		return
	}
	password = strings.TrimSpace(password)
	if username != Player.Username || password != Player.Password {
		fmt.Printf("Invalid. Connection closed.\n")
		conn.Write([]byte("AUTH_FAIL: Invalid credentials. Connection closed.\n"))
		return
	}
	conn.Write([]byte("AUTH_SUCCESS: Welcome to pokemons! \n"))
	fmt.Println("User authenticated successfully.")

	fmt.Println(Player)
	conn.Write([]byte(strconv.Itoa(Player.PlayerCoordinate.PlayerX) + "\n"))
	conn.Write([]byte(strconv.Itoa(Player.PlayerCoordinate.PlayerY) + "\n"))

	for {
		conn.Write([]byte("INPUTX: Enter player coordinate x moves:\n"))
		playerXStr, _ := reader.ReadString('\n')
		playerXStr = strings.TrimSpace(playerXStr)
		//
		conn.Write([]byte("INPUTY: Enter player coordinate y moves:\n"))
		playerYStr, _ := reader.ReadString('\n')
		playerYStr = strings.TrimSpace(playerYStr)
		//stop condition
		if playerXStr == "stop" || playerYStr == "stop" {
			break
		}
		//convert string to int coordinate from clients
		playerX, err := strconv.Atoi(playerXStr)
		if err != nil {
			conn.Write([]byte("ERROR: Invalid input X, please send a coordinate X number.\n"))
			continue
		}
		playerY, err := strconv.Atoi(playerYStr)
		if err != nil {
			conn.Write([]byte("ERROR: Invalid input Y, please send a coordinate Y number.\n"))
			continue
		}
		//Condition
		if playerX > 1000 || playerX < 0 || playerY > 1000 || playerY < 0 {
			conn.Write([]byte("ERROR: Invalid input, please send a coordinate number less than 1000 and more than 0.\n"))
			continue
		}
		//func
		for i := 0; i < 50; i++ {
			if playerX == p[i].Coordinates.X && playerY == p[i].Coordinates.Y {
				conn.Write([]byte("RESPONSE: You have catch pokemon "))
				conn.Write([]byte(p[i].Name))
				Player.PlayerCoordinate.PlayerX = playerX
				Player.PlayerCoordinate.PlayerY = playerY
				Player.CapturedPokemons = append(Player.CapturedPokemons, p[i].Name)
				break
			}
		}
		conn.Write([]byte("Status: Your current pokemons." + Player.CapturedPokemons[0] + "\n"))
		fmt.Println(Player)

	}

}
