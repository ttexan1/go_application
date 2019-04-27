package psql

import (
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"
)

var tables = []interface{}{
	&domain.Category{},
}

var tableNames = []string{
	tblCategories,
}

type factory struct {
	db *string
}

// NewStorage creates a storage factory
func NewStorage(driver, url string) (f engine.StorageFactory, err error) {
	db := "db"
	f = &factory{&db}
	return
}

func (f *factory) Close()      {}
func (f *factory) DropTables() {}
func (f *factory) Migrate()    {}
func (f *factory) NewCategoryRepo() engine.CategoryRepo {
	return newCategoryRepo(f.db)
}
