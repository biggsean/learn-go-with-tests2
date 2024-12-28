package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/biggsean/learn-go-with-tests2/app"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("let's play poker")
	fmt.Println(`Type "{Name} wine" to record a win`)
	poker.NewCLI(store, os.Stdin)
}
