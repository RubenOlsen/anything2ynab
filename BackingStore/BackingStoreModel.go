package BackingStore

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BankProviders struct {
	gorm.Model
	id            int `gorm:"AUTO_INGREMENT"`
	Memonic       string
	Providername  string
	Username      string
	Apikey        string
	CustomerId    string
	Password      string
	Status        int
	LastSync      string // date of last sync
	LastMonthSync int    // YYYYMM
	Param1        string
	Param2        string
}

type BudgetProviders struct {
	gorm.Model
	id         int `gorm:"AUTO_INGREMENT"`
	BudgetName string
	BudgetId   string
	Apikey     string
}

type TransactionCache struct {
	id      int    `gorm:"AUTO_INCREMENT"`
	Hash    string `gorm:"UNIQUE_INDEX"`
	Account string
	TrxDate string
	Type    string
	Ref     string
	Debet   int
	Kredit  int
	Rawdata string
}

type LogStatements struct {
	id      int `gorm:"AUTO_INCREMENT"`
	Date    *time.Time
	Level   string
	Source  string
	Message string
}
