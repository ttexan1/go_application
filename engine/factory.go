package engine

type (
	// Factory interface allows us to provide
	// other parts of the system with a way to make
	// instances of our use-case / interactors when
	// they need to
	Factory interface {
		NewCategory() Category
	}

	factory struct {
		StorageFactory
	}
)

// NewEngine returns a Factory
func NewEngine(s StorageFactory) Factory {
	return &factory{s}
}
