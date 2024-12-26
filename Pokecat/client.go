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
	//Valid username and password
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	conn.Write([]byte(strings.TrimSpace(username) + "\n"))
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	conn.Write([]byte(strings.TrimSpace(password) + "\n"))
	//
	authResponse, _ := serverReader.ReadString('\n')
	fmt.Print("Server: " + authResponse)
	//

	fmt.Printf("Your current position: \n")
	Xresponse, _ := serverReader.ReadString('\n')
	fmt.Print("X= ", Xresponse)
	Yresponse, _ := serverReader.ReadString('\n')
	fmt.Print("Y= ", Yresponse)

	AuthMess, _ := serverReader.ReadString('\n')
	if strings.HasPrefix(AuthMess, "AUTH_SUCCED:") {
		fmt.Print(AuthMess + "\n")
	} else if strings.HasPrefix(AuthMess, "AUTH_FAIL:") {
		fmt.Print(AuthMess + "\n")
	}
	//
	for {
		fmt.Print("Enter player coordinate x moves: ")
		X, _ := reader.ReadString('\n')                 //1
		conn.Write([]byte(strings.TrimSpace(X) + "\n")) //1
		//
		fmt.Print("Enter player coordinate y moves: ")
		Y, _ := reader.ReadString('\n')                 //2
		conn.Write([]byte(strings.TrimSpace(Y) + "\n")) //2
		//
		ServerMessage1, _ := serverReader.ReadString('\n')
		fmt.Print(ServerMessage1 + "\n")
		ServerMessage, _ := serverReader.ReadString('\n')
		if strings.HasPrefix(ServerMessage, "RESPONSE_POKEMON_CATCH:") {
			fmt.Print(ServerMessage + "\n")
		} else if strings.HasPrefix(ServerMessage, "ERROR:") {
			fmt.Print(ServerMessage + "\n")
		} else if strings.HasPrefix(ServerMessage, "Exit:") {
			fmt.Print(ServerMessage + "\n")
			break
		} else if strings.HasPrefix(ServerMessage, "INPUT:") {
			continue
		}
	}

}
