package main

import (
	"fmt"
	poker "lgwt/httpserver"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, closeStore, err := poker.FileSystemPlayStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer closeStore()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record win")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
