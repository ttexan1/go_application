package testseed

import "github.com/ttexan1/golang-simple/domain"

// Categories are the seed data
var Categories = []*domain.Category{
	&domain.Category{
		ID:   1,
		Name: "test1",
	},
	&domain.Category{
		ID:   2,
		Name: "test2",
	},
	&domain.Category{
		ID:   3,
		Name: "test3",
	},
}
