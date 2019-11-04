package BackingStore

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"log"
)

type DBController struct {
	DB *gorm.DB
}

func (dc *DBController) Connect() {
	var err error
	dc.DB, err = gorm.Open("sqlite3", "/Users/ruben/go/src/test.sqlite")
	if err != nil {
		log.Fatal(err)
		panic("Could not open database")
		dc.DB.Close()
	}

}

func (dc *DBController) Migrate() {
	fmt.Println("INFO migrating tabels")
	dc.DB.AutoMigrate(&BudgetProviders{})
	dc.DB.AutoMigrate(&BankProviders{})
	dc.DB.AutoMigrate(&TransactionCache{})
	dc.DB.AutoMigrate(&LogStatements{})
}
