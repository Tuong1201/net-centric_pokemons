package main

import (
	"bufio"

	// "encoding/json"
	"fmt"
	"net"
	"net-centric_pokemons/Db"
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
	// Create players
	var Playerarr []Db.Player
	Player1, Player2, Player3 := Db.PlayerDb()
	Playerarr = append(Playerarr, Player1, Player2, Player3)
	//
	reader := bufio.NewReader(conn)
	//clasify username and password
	//username
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username:", err)
		conn.Write([]byte("ERROR: Failed to read username.\n"))
		return
	}
	username = strings.TrimSpace(username)
	//password
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username:", err)
		conn.Write([]byte("ERROR: Failed to read username.\n"))
		return
	}
	password = strings.TrimSpace(password)
	var Player *Db.Player
	for _, player := range Playerarr {
		if username == player.Username {
			if password == player.Password {
				fmt.Printf("You are connected to the server as %s.\n", player.Username)
				conn.Write([]byte("AUTH_SUCCED: valid password. Connection as " + player.Username + "\n"))
				Player = &player
				break
			} else {
				fmt.Printf("Invalid password. Connection closed.\n")
				conn.Write([]byte("AUTH_FAIL: Invalid password. Connection closed.\n"))
				return
			}
		}
	}
	conn.Write([]byte(strconv.Itoa(Player.PlayerCoordinate.PlayerX) + "\n"))
	conn.Write([]byte(strconv.Itoa(Player.PlayerCoordinate.PlayerY) + "\n"))
	conn.Write([]byte("Welcome to pokemons! \n"))
	fmt.Println("User authenticated successfully.")

	fmt.Println(Player)
	savePlayersToJSON(Player, "Player.json")

	for {
		conn.Write([]byte("INPUT: Input your coordinate \n"))
		playerXStr, _ := reader.ReadString('\n')   //1
		playerXStr = strings.TrimSpace(playerXStr) //1
		//
		playerYStr, _ := reader.ReadString('\n')   //2
		playerYStr = strings.TrimSpace(playerYStr) //2
		//stop condition
		if playerXStr == "stop" && playerYStr == "stop" {
			conn.Write([]byte("Exit: Exit program\n"))
			fmt.Printf("Server disconnect \n")
			break
		}
		ReadJSONFile("Player.json", &Player)
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
				conn.Write([]byte("RESPONSE_POKEMON_CATCH: You have catch pokemon " + p[i].Name + "\n"))
				Player.PlayerCoordinate.PlayerX = playerX
				Player.PlayerCoordinate.PlayerY = playerY
				Player.CapturedPokemons = append(Player.CapturedPokemons, p[i].Name)
				break
			}
		}

		Player.PlayerCoordinate.PlayerX = playerX
		Player.PlayerCoordinate.PlayerY = playerY
		savePlayersToJSON(Player, "Player.json")
		fmt.Println(Player)
	}

}
