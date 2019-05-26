package engine

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/ttexan1/golang-simple/domain"
)

type WriterSuite struct {
	suite.Suite
	factory Factory
	writer  Writer
}

type testWriterRepo struct{}

func (r *testWriterRepo) List(params *ListWritersRequest) ([]*domain.Writer, int, *domain.Error) {
	if params.Limit == 1 {
		return []*domain.Writer{
			&domain.Writer{},
		}, 1, nil
	}
	return nil, 0, &domain.Error{
		Code: http.StatusBadRequest,
	}
}
func (r *testWriterRepo) Create(c *domain.Writer) (*domain.Writer, *domain.Error) {
	if c.Email == "writer@example.com" {
		c.ID = 1
		return c, nil
	}
	return nil, &domain.Error{
		Code: http.StatusBadRequest,
	}
}
func (r *testWriterRepo) Find(cond domain.Writer) (*domain.Writer, *domain.Error) {
	if cond.Email == "writer@example.com" {
		return &domain.Writer{}, nil
	}
	return nil, &domain.Error{
		Code: http.StatusBadRequest,
	}
}
func (r *testWriterRepo) Update(writer, params *domain.Writer) *domain.Error {
	if writer.ID == 1 {
		return nil
	}
	return &domain.Error{Code: http.StatusBadRequest}
}
func (r *testWriterRepo) Destroy(id int) *domain.Error {
	if id == 1 {
		return nil
	}
	return &domain.Error{Code: http.StatusBadRequest}
}

func TestWriterSuite(t *testing.T) {
	suite.Run(t, &WriterSuite{
		writer: &writer{
			repo: &testWriterRepo{},
		},
		factory: &factory{
			StorageFactory: &testStorage{},
		},
	})
}

func (s *WriterSuite) TestNewWriter() {
	s.NotNil(s.factory.NewWriter())
}

func (s *WriterSuite) TestListError() {
	resp := s.writer.List(&ListWritersRequest{})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *WriterSuite) TestList() {
	resp := s.writer.List(&ListWritersRequest{
		Limit: 1,
	})
	s.Equal(1, resp.Count)
}

func (s *WriterSuite) TestCreateError() {
	resp := s.writer.Create(&CreateWriterRequest{})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *WriterSuite) TestCreate() {
	resp := s.writer.Create(&CreateWriterRequest{
		Name:  "writer",
		Email: "writer@example.com",
	})
	s.Nil(resp.Error)
	s.Equal(1, resp.Writer.ID)
}

func (s *WriterSuite) TestFindError() {
	resp := s.writer.Find(&FindWriterRequest{})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *WriterSuite) TestUpdateError() {
	resp := s.writer.Update(&UpdateWriterRequest{ID: 1})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *WriterSuite) TestDestroyError() {
	resp := s.writer.Destroy(&DestroyWriterRequest{})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *WriterSuite) TestDestroy() {
	resp := s.writer.Destroy(&DestroyWriterRequest{
		ID: 1,
	})
	s.Nil(resp.Error)
}
