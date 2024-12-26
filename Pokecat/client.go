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
	authResponse, _ := serverReader.ReadString('\n')
	fmt.Print("Server: " + authResponse)
	//
	for {
		fmt.Printf("Your current position: \n")
		Xresponse, _ := serverReader.ReadString('\n')
		fmt.Print("X= ", Xresponse)

		Yresponse, _ := serverReader.ReadString('\n')
		fmt.Print("Y= ", Yresponse)

		//Enter Player X and Y
		XMessage, _ := serverReader.ReadString('\n')
		fmt.Print("Server: " + XMessage)
		X, _ := reader.ReadString('\n')

		if strings.TrimSpace(X) == "stop" {
			fmt.Println("Game over. Exiting.")
			break
		}

		conn.Write([]byte(strings.TrimSpace(X) + "\n"))
		//
		YMessage, _ := serverReader.ReadString('\n')
		fmt.Print("Server: " + YMessage)
		Y, _ := reader.ReadString('\n')

		if strings.TrimSpace(Y) == "stop" {
			fmt.Println("Game over. Exiting.")
			break
		}

		conn.Write([]byte(strings.TrimSpace(Y) + "\n"))
		Response, _ := serverReader.ReadString('\n')
		fmt.Print("Server Response: " + Response + "\n")

		CapResponse, _ := serverReader.ReadString('\n')
		fmt.Print("Server Response: pokemons Capture " + CapResponse)

	}

}
