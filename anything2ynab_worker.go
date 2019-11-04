package main

import (
	"fmt"
	"github.com/RubenOlsen/anything2ynab/BackingStore"
)

func main() {

	BackintStore := BackingStore.DBController{}
	BackintStore.Connect()
	BackintStore.Migrate()
	fmt.Println("Ready to rumble")

	budget()
}
