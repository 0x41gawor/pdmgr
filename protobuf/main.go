package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gawor.com/protobuf/protos"
	"google.golang.org/protobuf/proto"
)

func main() {
	fname := "players.bin"

	book := &protos.Players{}

	// Read the existing book.
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	print("Existing players in book:\n")
	print("Name:     , Country:     \n")
	for _, x := range book.List {
		out := fmt.Sprintf("{%s, %s}", x.Name, x.Country)
		println(out)
	}

	print("------------------------\nCreation of a new player !!!\n")
	reader := bufio.NewReader(os.Stdin)
	print("Player's fullname: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	print("Player's country: ")
	country, _ := reader.ReadString('\n')
	country = strings.TrimSpace(country)

	// create new player
	a := protos.FootballPlayer{
		Name:    name,
		Country: country,
	}

	book.List = append(book.List, &a)

	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := os.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
