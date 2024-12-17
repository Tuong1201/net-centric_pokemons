package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// Make a GET request to the WEBTOON homepage
	resp, err := http.Get("https://pokedex.org/")
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

	values := findParagraphs(doc)
	// Save links to a JSON file
	err = saveLinksToJSON(values, "pokemons.json")
	if err != nil {
		fmt.Println("Error saving links to JSON:", err)
	} else {
		fmt.Println("Links successfully saved to pokemons.json")
	}
}

func findParagraphs(n *html.Node) []string {
	var values []string
	var walk func(n *html.Node)
	walk = func(n *html.Node) {
		// Check if the node is a <p> tag
		if n.Type == html.ElementNode && n.Data == "ul" {
			// Check for class="subj"
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == "monsters-list" {
					// Get the content inside the <p> tag
					content := getTextContent(n)
					fmt.Println(" ", content)
					values = append(values, content)
					break
				}
			}
		}
		// Recursively walk through child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return values
}

// Function to get the text content of a node
func getTextContent(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var content string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		content += getTextContent(c)
		content += " "
	}
	return content
}

// Function to save links to a JSON file
func saveLinksToJSON(links []string, filename string) error {
	// Open or create the JSON file
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
