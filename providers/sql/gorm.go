package sql

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/ttexan1/golang-simple/domain"
)

func newErrorByGormError(gormErr error) *domain.Error {
	switch gormErr {
	case nil:
		return nil
	case gorm.ErrRecordNotFound:
		return domain.NewError(http.StatusNotFound, gormErr.Error())
	case gorm.ErrInvalidSQL, gorm.ErrInvalidTransaction, gorm.ErrCantStartTransaction, gorm.ErrUnaddressable:
		return domain.NewError(http.StatusInternalServerError, gormErr.Error())
	default:
		return domain.NewError(http.StatusUnprocessableEntity, gormErr.Error())
	}
}
