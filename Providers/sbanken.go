package Providers

import (
	"fmt"
	log "github.com/RubenOlsen/anything2ynab/logging"
	sbanken "github.com/RubenOlsen/go-sbanken"
	"os"
)

func Sbanken() {

	// fmt.Println("RUBEn", os.Getenv("RUBEN"))
	log.Debug("Inside Sbn")

	creds := sbanken.Credentials{"edbc33bf87e943b1957e85931658f609",
		"cwf?SE3aL3Ddqw=XtZRuj?jKTb?4dkoYT3EICcnp0EbMA-g+0z+HI8y26eBjqzZR",
		"24116842358"}
	conn := sbanken.NewAPIConnection(creds)

	log.Debug("%v", conn)

	accounts := conn.GetAccounts()
	if len(accounts) < 1 {
		log.Fatal("No accounts found. Exiting.")
		os.Exit(1)
	}
	fmt.Printf("\n\n\n\n")
	fmt.Println(accounts)
	fmt.Printf("\n\n\n\n")

	log.Debug("Got accounts")

	for _, account := range accounts {
		fmt.Printf("%-14s %-40s %9.2f %s\n", account.AccountNumber, account.Name, account.Balance, account.AccountID)
	}

}
