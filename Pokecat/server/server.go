package main

import (
	"bufio"
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
	var pokeworld [1000][1000]string
	var pokemons []Pokemons
	err := ReadJSONFile("C:/Users/hp/Desktop/projects/net-centric/Pokedex/pokedex.json", &pokemons)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}
	// var x, y int
	p := Rand50Pokemon(pokemons) //50 random pokemon
	for i := 0; i <= 50; i++ {
		pokeworld[pokemons[i].coordinate.x][pokemons[i].coordinate.y] = pokemons[i].Id
	}
	username := conn.RemoteAddr().String() //Lấy địa chỉ IP của client và hiển thị
	fmt.Println("New client connected:", username)
	// jsonData, err := json.Marshal(pokemons)
	// if err != nil {
	// 	fmt.Println("Error marshalling JSON:", err)
	// }

	// Gửi JSON cho client
	// conn.Write(jsonData)
	fmt.Printf("list of 50 random pokemons\n")
	for i := 0; i < 50; i++ {
		fmt.Printf("ID: %s\nName: %s\nExperience: %s\nHP: %s\nAttack: %s\nDefense: %s\nSpecial Attack: %s\nSpecial Defense: %s\nSpeed: %s\nTotal EVs: %s\nCoordinate: (%d, %d)\n",
			p[i].Id, p[i].Name, p[i].Exp, p[i].HP, p[i].Attack, p[i].Defense, p[i].SpAttack, p[i].SpDefense, p[i].Speed, p[i].TotalEVs, p[i].coordinate.x, p[i].coordinate.y)
		fmt.Printf("\n")
	}
	reader := bufio.NewReader(conn)
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
			if playerX == p[i].coordinate.x && playerY == p[i].coordinate.y {
				conn.Write([]byte("RESPONSE: You have catch pokemon "))
				conn.Write([]byte(p[i].Name))
			}
		}
	}

}
