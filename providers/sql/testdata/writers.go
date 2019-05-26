package testdata

import "github.com/ttexan1/golang-simple/domain"

// Writers are the seed data
var Writers = []*domain.Writer{
	&domain.Writer{
		ID:    1,
		Name:  "test1",
		Email: "writer1@example.com",
	},
	&domain.Writer{
		ID:    2,
		Name:  "test2",
		Email: "writer2@example.com",
	},
	&domain.Writer{
		ID:    3,
		Name:  "test3",
		Email: "writer3@example.com",
	},
}
