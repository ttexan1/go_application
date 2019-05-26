package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/scoville/validations"

	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"

	// to use postgresql
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var tables = []interface{}{
	&domain.Writer{},
	&domain.Category{},
	&domain.Article{},
}

var tableNames = []string{
	tblArticles,
	tblCategories,
	tblWriters,
}

type (
	factory struct {
		db *gorm.DB
	}
)

// NewStorage creates a storage factory
func NewStorage(driver, url string) (f engine.StorageFactory, err error) {
	db, err := gorm.Open(driver, url)
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
	// Don't use AutoMigrate in production
	f.db.AutoMigrate(tables...)
}

func (f *factory) NewArticleRepo() engine.ArticleRepo {
	return newArticleRepo(f.db)
}
func (f *factory) NewCategoryRepo() engine.CategoryRepo {
	return newCategoryRepo(f.db)
}
func (f *factory) NewWriterRepo() engine.WriterRepo {
	return newWriterRepo(f.db)
}
