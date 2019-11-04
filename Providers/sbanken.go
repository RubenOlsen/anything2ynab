package Providers

import (
	"fmt"
	sbanken "github.com/RubenOlsen/go-sbanken"
	"regexp"
)

func sbn() {

	// fmt.Println("RUBEn", os.Getenv("RUBEN"))

	creds := sbanken.Credentials{"edbc33bf87e943b1957e85931658f609", "xhR4l=9IQjYtG-j=3SbtMV+AypbbqXWj322Ex9y9s?ixolBeAb?WEmEB6z0kneHn", "24116842358"}
	conn := sbanken.NewAPIConnection(creds)

	// fmt.Println(conn.HasToken())

	accounts := conn.GetAccounts()

	for _, account := range accounts {
		fmt.Printf("%-14s %-40s %9.2f %s\n", account.AccountNumber, account.Name, account.Balance, account.AccountID)
	}

	r := *regexp.MustCompile(`^((?P<kort>\*\d+)\s(?P<dato>\d\d\.\d\d)\s(?P<valuta>[A-Z]{3})\s(?P<belop>\d+\.\d+))*(\s*(?P<tdato>\d+\.\d+))*\s?(?P<sted>.+?)((?i)\s(kurs):\s?(?P<kurs>\d+\.?\d+))*$`)

	transactions := conn.GetTransactions("44BBE700A0F7EB22755C88C16ECA40D4")
	for _, transaction := range transactions {

		var transformation Transformation

		if transaction.CardDetailsSpecified {
			transformation.CardNumber = transaction.CardDetails.CardNumber
		} else {
			// fmt.Println("%#v", r.SubexpNames())
			res := r.FindStringSubmatch(transaction.Text)
			for i := 0; i < len(res); i++ {

				fmt.Printf("%s ||2: %i %s\n", transaction.Text, i, res[i])

			}
			fmt.Printf("\n")

		}
		/*
		   	 	fmt.Printf("%-10s | %-8s |%-10s| %s| %s | %s | %9.2f\n",
		   	 		transaction.AccountingDate,
		               // transaction.CardDetailsSpecified,
		               transaction.TransactionID,
		               transformation.CardNumber,
		   	 		// transaction.OtherAccountNumber,
		   	 		transaction.Source,
		   	 		transaction.TransactionID,
		   	 		transaction.Text,
		   	transaction.Amount)
		*/
	}

}
