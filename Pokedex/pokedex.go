package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Newpokemons struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Exp       string `json:"exp"`
	HP        string `json:"hp"`
	Attack    string `json:"attack"`
	Defense   string `json:"defense"`
	SpAttack  string `json:"sp_attack"`
	SpDefense string `json:"sp_defense"`
	Speed     string `json:"speed"`
	TotalEVs  string `json:"total_evs"`
}

func main() {
	// Make a GET request to the WEBTOON homepage
	resp, err := http.Get("https://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_effort_value_yield_in_Generation_IX")
	if err != nil {
		fmt.Println("Error fetching pokedex homepage:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Parse the HTML content
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}
	var pokemons []Newpokemons
	var newValues []string
	values := findParagraphs(doc)
	for i := 0; i < (len(values) - 30); i += 11 {
		newValues = append(newValues, values[i], values[i+2], values[i+3], values[i+4], values[i+5], values[i+6], values[i+7], values[i+8], values[i+9], values[i+10])
	}

	for i := 0; i < len(newValues); i += 10 {
		pokemons = append(pokemons, Newpokemons{
			Id:        newValues[i],
			Name:      newValues[i+1],
			Exp:       newValues[i+2],
			HP:        newValues[i+3],
			Attack:    newValues[i+4],
			Defense:   newValues[i+5],
			SpAttack:  newValues[i+6],
			SpDefense: newValues[i+7],
			Speed:     newValues[i+8],
			TotalEVs:  newValues[i+9],
		})
	}
	fmt.Println("len=", len(newValues))
	// Save links to a JSON file
	err = saveLinksToJSON(pokemons, "pokedex.json")
	if err != nil {
		fmt.Println("Error saving links to JSON:", err)
	} else {
		fmt.Println("Links successfully saved to pokedex.json")
	}
}

func findParagraphs(n *html.Node) []string {
	var values []string
	var walk func(n *html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "td" {
			content := getTextContent(n)
			if content != "" || content != "\n" {
				values = append(values, content)
			}
		}
		// Recursively walk through child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	var result []string
	result = values[1:]
	return result
}

// Function to get the text content of a node
func getTextContent(n *html.Node) string {
	if n.Type == html.TextNode {
		return strings.TrimSpace(n.Data)
	}
	var content string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		content += getTextContent(c)
	}
	return content
}

// Function to save links to a JSON file
func saveLinksToJSON(links []Newpokemons, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode links as JSON and write to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Add indentation for better readability
	return encoder.Encode(links)
}
