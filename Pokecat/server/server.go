package main

import (
	"bufio"
	// "encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
	".net-centric_pokemons/Db"
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
	p := Rand50Pokemon(pokemons)           //50 random pokemon
	username := conn.RemoteAddr().String() //Lấy địa chỉ IP của client và hiển thị
	fmt.Println("New client connected:", username)
	//
	saveLinksToJSON(p, "rand50Pokemons.json")
	//
	reader := bufio.NewReader(conn)
	player, _ := reader.ReadString('\n')
	player = strings.TrimSpace(player)
	conn.Write([]byte("AUTH_SUCCESS: Welcome to pokemons! \n"))
	for {
		conn.Write([]byte("AUTH: Enter player coordinate x:\n"))
		playerXStr, _ := reader.ReadString('\n')
		playerXStr = strings.TrimSpace(playerXStr)
		conn.Write([]byte("AUTH: Enter player coordinate y:\n"))
		playerYStr, _ := reader.ReadString('\n')
		playerYStr = strings.TrimSpace(playerYStr)
		//stop condition
		if playerXStr == "stop" || playerYStr == "stop" {
			break
		}
		//
		playerX, err := strconv.Atoi(playerXStr)
		if err != nil {
			conn.Write([]byte("ERROR: Invalid input X, please send a coordinate X number.\n"))
			continue
		}
		playerY, err := strconv.Atoi(playerYStr)
		if err != nil {
			conn.Write([]byte("ERROR: Invalid input X, please send a coordinate X number.\n"))
			continue
		}
		//func
		for i := 0; i < 50; i++ {
			if playerX == p[i].Coordinates.X && playerY == p[i].Coordinates.Y {
				conn.Write([]byte("RESPONSE: You have catch pokemon "))
				conn.Write([]byte(p[i].Name))
			}
		}
	}

}
