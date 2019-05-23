package sql

import (
	"strings"

	"github.com/jinzhu/gorm"
)

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
