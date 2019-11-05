package main

import (
	"github.com/RubenOlsen/anything2ynab/BackingStore"
	log "github.com/RubenOlsen/anything2ynab/logging"
)

func main() {

	// db := BackingStore.DBController{}
	backstore := BackingStore.DBController{}
	databaseObject := backstore.Connect()
	backstore.Migrate()
	log.Info("Ready to rumble")

	// budgetinfo = BackingStore.BudgetProviders{}
	// budgetinfo = backstore.FetchBudgetProviders()
	// YnabAccounts := []YnabAccount
	YnabAccounts := FetchBudgetAccounts(databaseObject)
	for key, YAC := range YnabAccounts {
		log.Info("Retur k:%v v.Name:%v", key, YAC.Name)
	}

	// log.Info("YANB LOG ACCOUNT %", YnabLogAccount.Id)

	//budget()
}
