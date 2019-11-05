package main

import (
	"fmt"
	"github.com/davidsteinsland/ynab-go/ynab"
)

func budgetDebug() {

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

	}

	return
	/*

	   tr := &ynab.SaveTransaction{
	   AccountId:  myAccountId,
	   Date:       "2019-11-01",
	   Amount:     12345,
	   PayeeId:    "",
	   PayeeName:  "Ruben Betaler",
	   CategoryId: "",
	   Memo:       "Dette er et memofelt",
	   Cleared:    "",
	   Approved:   false,
	   FlagColor:  "",
	   ImportId:   "",
	   }

	   trd := ynab.TransactionDetail{
	   TransactionSummary: ynab.TransactionSummary{},
	   AccountName:        "",
	   PayeeName:          "",
	   CategoryName:       "",
	   SubTransactions:    nil,
	   }

	   // trd, err :=
	   trd,err := client.TransactionsService.Create(myBudgetId,tr)
	   if err != nil {
	   panic(err)
	   }


	   fmt.Println(trd.Id, " | " , trd.AccountName, " | ", trd.Memo)

	   ac,_ := client.AccountsService.Get(myBudgetId,myAccountId)
	   fmt.Println("ACNOTE " , *ac.Note)


	   	budgets, _ := client.BudgetService.List()
	   	for _, budgetSummary := range budgets {
	   		fmt.Printf("Budget %v: %v\n", budgetSummary.Id, budgetSummary.Name)

	   		budget, _ := client.BudgetService.Get(budgetSummary.Id)

	   		fmt.Printf("Accounts:\n")
	   		for _, account := range budget.Accounts {
	   			fmt.Printf("\tAccount %v: %v\n", account.Id, account.Name)
	   			fmt.Printf("\t\tBalance: %v\n", account.Balance)

	   			transactions, _ := client.TransactionsService.GetByAccount(budgetSummary.Id, account.Id)

	   			fmt.Printf("\t\tTransactions:\n")
	   			for _, transaction := range transactions {
	   				fmt.Printf("\t\t\t%v: %v\n", transaction.Date, transaction.Amount)
	   			}
	   		}
	   	}

	*/
}
