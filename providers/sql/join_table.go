package sql

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type joinTable map[string]bool

func newJoinTable() joinTable {
	return joinTable{}
}

func (jt joinTable) InnerJoin(db *gorm.DB, tbl, pkTbl, fkTbl string) *gorm.DB {
	return jt.join("INNER JOIN", db, tbl, pkTbl, fkTbl)
}

func (jt joinTable) LeftOuterJoin(db *gorm.DB, tbl, pkTbl, fkTbl string) *gorm.DB {
	return jt.join("LEFT OUTER JOIN", db, tbl, pkTbl, fkTbl)
}

func (jt joinTable) join(state string, db *gorm.DB, tbl, pkTbl, fkTbl string) *gorm.DB {
	if _, ok := jt[tbl]; !ok {
		jt[tbl] = true
		return db.Joins(fmt.Sprintf("%s %s ON %s = %s", state, tbl, pk(pkTbl), fk(fkTbl, pkTbl)))
	}
	return db
}
