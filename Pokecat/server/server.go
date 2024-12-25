package main

import (
	"encoding/json"
	"fmt"
	"net"
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
	pokemons = Rand50Pokemon(pokemons) //50 random pokemon
	for i := 0; i <= 50; i++ {
		pokeworld[pokemons[i].coordinate.x][pokemons[i].coordinate.y] = pokemons[i].Id
	}
	username := conn.RemoteAddr().String() //Lấy địa chỉ IP của client và hiển thị
	fmt.Println("New client connected:", username)
	jsonData, err := json.Marshal(pokemons)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}

	// Gửi JSON cho client
	conn.Write(jsonData)
	fmt.Println(pokemons)

}
