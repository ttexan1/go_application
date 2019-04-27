package engine

import "github.com/ttexan1/golang-simple/domain"

type (
	// StorageFactory is the interface to create repos
	StorageFactory interface {
		Close()
		DropTables()
		Migrate()

		NewCategoryRepo() CategoryRepo
	}

	// CategoryRepo is the interface for the repo
	CategoryRepo interface {
		List(*ListCategoriesRequest) ([]*domain.Category, int, *domain.Error)
		Create(*domain.Category) (*domain.Category, *domain.Error)
		Find(int) (*domain.Category, *domain.Error)
		Update(*domain.Category, *domain.Category) *domain.Error
		Destroy(int) *domain.Error
	}
)
