package BackingStore

import (
	log "github.com/RubenOlsen/anything2ynab/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DBController struct {
	DB *gorm.DB
}

func (dc *DBController) Connect() (ret DBController) {
	var err error
	dc.DB, err = gorm.Open("sqlite3", "/Users/ruben/go/src/test.sqlite")
	if err != nil {
		log.Fatal(err.Error())
		panic("Could not open database")
		dc.DB.Close()
	}

	return *dc
}

func (dc *DBController) Migrate() {
	log.Info("Migrating database tabels")
	dc.DB.AutoMigrate(&BudgetProviders{})
	dc.DB.AutoMigrate(&BankProviders{})
	dc.DB.AutoMigrate(&TransactionCache{})
	dc.DB.AutoMigrate(&LogStatements{})
}

func (dc *DBController) FetchBudgetProviders() (BudgetProviders BudgetProviders) {
	dc.DB.First(&BudgetProviders)
	return BudgetProviders
}
