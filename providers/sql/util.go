package sql

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/inflection"
)

func pk(tbl string) string {
	return tbl + ".id"
}

func fk(tbl, to string) string {
	return tbl + "." + inflection.Singular(to) + "_id"
}

func limit(db *gorm.DB, num, dft int) *gorm.DB {
	if num < 0 {
		return db
	}
	if num > 0 {
		return db.Limit(num)
	}
	if dft > 0 {
		return db.Limit(dft)
	}
	return db
}

func offset(db *gorm.DB, num int) *gorm.DB {
	if num > 0 {
		return db.Offset(num)
	}
	return db
}

type orBuilder struct {
	state  string
	params []interface{}
	cond   bool
}

func newOrBuilder(s string, params ...interface{}) orBuilder {
	return orBuilder{s, params, true}
}

func newOrBuilderIf(cond bool, s string, params ...interface{}) orBuilder {
	return orBuilder{s, params, cond}
}

func or(db *gorm.DB, builders []orBuilder) *gorm.DB {
	states := make([]string, 0, len(builders))
	params := make([]interface{}, 0, len(builders))
	for _, b := range builders {
		if b.cond {
			states = append(states, b.state)
			params = append(params, b.params...)
		}
	}
	if len(states) == 0 {
		return db
	}
	return db.Where(strings.Join(states, " OR "), params...)
}

func preload(db *gorm.DB, columns map[string]bool) *gorm.DB {
	for col := range columns {
		db = db.Preload(col)
	}
	return db
}

func sort(db *gorm.DB, val string, whitelist []string, dft string) *gorm.DB {
	if val != "" {
		val = strings.Replace(val, "+", " ", -1)
		column := strings.Split(val, " ")[0]
		for _, e := range whitelist {
			if column == e {
				return db.Order(val)
			}
		}
	}
	return db.Order(dft)
}
