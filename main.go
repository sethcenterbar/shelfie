package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type shelf struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Books []book `json:"books"`
}

// Quick helper, learning how to reference itself like self in python or this in other languages
func (s shelf) describe() {
	fmt.Println(s)
}

// Add book to shelf, using pointer so that s is mutable
func (s *shelf) addBook(b book) {
	s.Books = append(s.Books, b)
}

// Write shelf to disk
func (s shelf) save() {
	shelfbytes, _ := json.MarshalIndent(s, "", "  ")
	err := ioutil.WriteFile("shelf.json", shelfbytes, 0644)
	if err != nil {
		panic(1)
	}
}

type book struct {
	Title   string   `json:"title"`
	Authors []author `json:"authors"`
	Rating  int      `json:"rating"`
	Review  string   `json:"review"`
}

type author struct {
	Name    string `json:"name"`
	Twitter string `json:"twitter"`
}

func main() {
	fmt.Println("Welcome to shelfie!")

	myshelf := getShelf()
	jeff := author{"Jeff Bridges", "@jbridgerr"}
	samantha := author{"Samantha Vaughn", "@sammyv92"}
	authors := []author{jeff, samantha}
	newbook := book{"Getting Good at Go", authors, 5, "I learned a lot about go while making shelfie!"}
	myshelf.addBook(newbook)
	myshelf.save()
}

func getShelf() shelf {
	// Open file
	bookshelf, err := os.Open("shelf.json")
	if err != nil {
		fmt.Println("Couldn't find your bookshelf!")
	}

	var shelf shelf

	// Read File
	bookshelfbytes, err := ioutil.ReadAll(bookshelf)
	if err != nil {
		fmt.Println("Couldn't read bookshelf!")
	}

	json.Unmarshal(bookshelfbytes, &shelf)
	return shelf
}

func createShelf(name string, owner string) (err error) {
	if _, err := os.Stat("shelf.json"); err == nil {
		return errors.New("shelf exists")
	} else if os.IsNotExist(err) {
		newshelf := shelf{name, owner, nil}
		shelfbytes, _ := json.MarshalIndent(newshelf, "", "  ")
		err := ioutil.WriteFile("shelf.json", shelfbytes, 0644)
		if err != nil {
			panic(1)
		}

		return nil

	} else {
		panic("i have no idea what has happend. possibly emp?")
	}
}
