package sql

import (
	"github.com/jinzhu/gorm"

	"github.com/scoville/validations"
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"

	// to use postgresql
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/mattn/go-sqlite3"
)

var tables = []interface{}{
	&domain.Category{},
}

var tableNames = []string{
	tblCategories,
}

type (
	factory struct {
		db *gorm.DB
	}
)

// NewStorage creates a storage factory
func NewStorage(driver, url string) (f engine.StorageFactory, err error) {
	db, err := gorm.Open(driver, url)
	// db, err := gorm.Open(driver, url)
	if err != nil {
		return
	}
	validations.RegisterCallbacks(db)
	f = &factory{db}
	return
}

func (f *factory) Close() {
	f.db.Close()
}

func (f *factory) DropTables() {
	for i := len(tables) - 1; i >= 0; i-- {
		f.db.DropTableIfExists(tables[i])
	}
}

func (f *factory) Migrate() {
	// We shouldn't use AutoMigrate in production
	f.db.AutoMigrate(tables...)
}

func (f *factory) NewCategoryRepo() engine.CategoryRepo {
	return newCategoryRepo(f.db)
}
