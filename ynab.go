package main

import (
	"fmt"
	"github.com/RubenOlsen/anything2ynab/BackingStore"
)

type BudgetData struct {
}

func FetchBudgetAccounts(dc BackingStore.DBController) {

	BudgetData := BackingStore.BudgetProviders{}
	BudgetData = dc.FetchBudgetProviders()
	fmt.Println(BudgetData)

	/*
			const accessToken = "559e9289d528bfa34c45a6a68c0c5006f9d0cca04eb749033239591fdd9df937"
			client := ynab.NewDefaultClient(accessToken)

			// myBudgetId := "3acbb288-829f-47a5-8f0b-63ef13f4c6c1"
			myBudgetId := "3acbb288-829f-47a5-8f0b-63ef13f4c6c1"
			// myAccountId := "b1ebfb88-dee2-486b-8c4d-310a8623ad76"
			//

			accounts, _ := client.AccountsService.List(myBudgetId)
			for _, account := range accounts {
			Note := "" //
			if account.Note != nil {
			Note = *account.Note
			}
			fmt.Println("\nName: ", account.Name,
			"\nNote: ", Note,
			"\nBal:  ", account.Balance,
			"\nCB:   ", account.ClearedBalance,
			"\nId:   ", account.Id,
			"\nOB:   ", account.OnBudget,
			"\nTy:   ", account.Type,
			"\nUC:   ", account.UnclearedBalance)

	}*/
}
