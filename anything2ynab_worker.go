package main

import (
	"github.com/RubenOlsen/anything2ynab/BackingStore"
	prov "github.com/RubenOlsen/anything2ynab/Providers"
	log "github.com/RubenOlsen/anything2ynab/logging"
)

func main() {

	// db := BackingStore.DBController{}
	backstore := BackingStore.DBController{}
	databaseObject := backstore.Connect()
	backstore.Migrate()
	log.Info("Ready to rumble")

	YnabAccounts := FetchBudgetAccounts(databaseObject)
	for key, YAC := range YnabAccounts {
		log.Debug("Retur k:%v v.Name:%v", key, YAC.Name)
	}

	prov.Sbanken()
	// log.Info("YANB LOG ACCOUNT %", YnabLogAccount.Id)

	//budget()
}
