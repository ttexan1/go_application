package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"
)

const tblWriters = "writers"

type writerRepo struct {
	db *gorm.DB
}

func newWriterRepo(db *gorm.DB) *writerRepo {
	return &writerRepo{db}
}

func (r *writerRepo) Destroy(id int) (err *domain.Error) {
	gormErr := r.db.Table(tblWriters).
		Delete(domain.Writer{}, "id = ?", id).Error

	err = newErrorByGormError(gormErr)
	return
}

func (r *writerRepo) List(params *engine.ListWritersRequest) (writers []*domain.Writer, count int, err *domain.Error) {
	db := r.db.Table(tblWriters)
	db = db.Count(&count)
	db = sort(db, params.Sort, []string{"id", "display_order"}, "id DESC")
	db = limit(db, params.Limit, 20)
	db = offset(db, params.Offset)
	gormErr := db.Find(&writers).Error
	err = newErrorByGormError(gormErr)
	return
}

func (r *writerRepo) Find(id int) (writer *domain.Writer, err *domain.Error) {
	writer = &domain.Writer{}
	gormErr := r.db.Table(tblWriters).
		Where(domain.Writer{ID: id}).
		First(&writer).Error

	err = newErrorByGormError(gormErr)
	return
}

func (r *writerRepo) Create(params *domain.Writer) (*domain.Writer, *domain.Error) {
	if err := newErrorByGormError(
		r.db.Table(tblWriters).
			Create(&params).Error); err != nil {
		return nil, err
	}
	return r.Find(params.ID)
}

func (r *writerRepo) Update(writer, params *domain.Writer) *domain.Error {
	return newErrorByGormError(
		r.db.Model(writer).
			Updates(params).Error)
}
