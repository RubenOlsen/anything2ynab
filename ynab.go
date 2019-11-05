package main

import (
	"github.com/RubenOlsen/anything2ynab/BackingStore"
	log "github.com/RubenOlsen/anything2ynab/logging"
	"github.com/davidsteinsland/ynab-go/ynab"
	"regexp"
)

type YnabAccount struct {
	Name     string
	Id       string
	SyncData string
}

type YnabLogAccount struct {
	Id string
}

func FetchBudgetAccounts(dc BackingStore.DBController) map[string]YnabAccount {

	BudgetData := BackingStore.BudgetProviders{}
	BudgetData = dc.FetchBudgetProviders()

	log.Info("Budget name: %s | API key %s", BudgetData.BudgetName, BudgetData.Apikey)

	client := ynab.NewDefaultClient(BudgetData.Apikey)

	rex := *regexp.MustCompile(`^sync\s+(\w+)`)
	returnMap := make(map[string]YnabAccount)

	accounts, _ := client.AccountsService.List(BudgetData.BudgetId)
	for _, account := range accounts {
		if !account.Closed {

			Note := "" //
			if account.Note != nil {
				Note = *account.Note
				res := rex.FindStringSubmatch(Note)
				if len(res) > 0 {

					YAC := YnabAccount{
						Name:     account.Name,
						Id:       account.Id,
						SyncData: res[1],
					}

					returnMap[res[1]] = YAC
					// YnabAccounts = append(YnabAccounts, YAC)
					log.Debug("Adding %s SD:%s ID:%s", account.Name, res[1], account.Id)
				}
			}

		}

	}
	return returnMap
}
