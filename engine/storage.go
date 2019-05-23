package engine

import "github.com/ttexan1/golang-simple/domain"

type (
	// StorageFactory is the interface to create repos
	StorageFactory interface {
		Close()
		DropTables()
		Migrate()

		NewArticleRepo() ArticleRepo
		NewCategoryRepo() CategoryRepo
		NewWriterRepo() WriterRepo
	}

	// ArticleRepo is the interface for the repo
	ArticleRepo interface {
		List(*ListArticlesRequest) ([]*domain.Article, int, *domain.Error)
		Create(*domain.Article) (*domain.Article, *domain.Error)
		Find(int) (*domain.Article, *domain.Error)
		Update(*domain.Article, *domain.Article) *domain.Error
		Destroy(int) *domain.Error
	}

	// CategoryRepo is the interface for the repo
	CategoryRepo interface {
		List(*ListCategoriesRequest) ([]*domain.Category, int, *domain.Error)
		Create(*domain.Category) (*domain.Category, *domain.Error)
		Find(int) (*domain.Category, *domain.Error)
		Update(*domain.Category, *domain.Category) *domain.Error
		Destroy(int) *domain.Error
	}

	// WriterRepo is the interface for the repo
	WriterRepo interface {
		List(*ListWritersRequest) ([]*domain.Writer, int, *domain.Error)
		Create(*domain.Writer) (*domain.Writer, *domain.Error)
		Find(domain.Writer) (*domain.Writer, *domain.Error)
		Update(*domain.Writer, *domain.Writer) *domain.Error
		Destroy(int) *domain.Error
	}
)
