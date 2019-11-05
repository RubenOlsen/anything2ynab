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
	FetchBudgetAccounts(databaseObject)

	//budget()
}
