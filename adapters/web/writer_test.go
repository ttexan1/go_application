package web

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"
)

type WriterSuite struct {
	webSuite
}

type testWriterEngine struct{}

func (e *testWriterEngine) Find(r *engine.FindWriterRequest) *engine.FindWriterResponse {
	if r.ID == 1 {
		return &engine.FindWriterResponse{
			Writer: &domain.Writer{ID: 1},
		}
	}
	return &engine.FindWriterResponse{
		Error: &domain.Error{Code: http.StatusNotFound},
	}
}

func (e *testWriterEngine) List(r *engine.ListWritersRequest) *engine.ListWritersResponse {
	return &engine.ListWritersResponse{
		Error: &domain.Error{Code: http.StatusNotFound},
	}
}

func (e *testWriterEngine) Create(r *engine.CreateWriterRequest) *engine.CreateWriterResponse {
	return &engine.CreateWriterResponse{
		Writer: &domain.Writer{
			ID:    1,
			Name:  "Writer1",
			Email: "writer@example.com",
		},
		Error: nil,
	}
}

func (e *testWriterEngine) Update(r *engine.UpdateWriterRequest) *engine.UpdateWriterResponse {
	if r.ID == 1 {
		return &engine.UpdateWriterResponse{
			Writer: &domain.Writer{
				ID:    1,
				Name:  "Updated Writer1",
				Email: "writer@example.com",
			},
			Error: nil,
		}
	}
	return &engine.UpdateWriterResponse{
		Writer: &domain.Writer{},
		Error:  domain.NewError(http.StatusNotFound, gorm.ErrRecordNotFound.Error()),
	}
}

func (e *testWriterEngine) Destroy(r *engine.DestroyWriterRequest) *engine.DestroyWriterResponse {
	return &engine.DestroyWriterResponse{Error: nil}
}
func (e *testWriterEngine) Login(r *engine.LoginWriterRequest) *engine.LoginWriterResponse {
	if r.Email == "writer@example.com" && r.Password == "password" {
		return &engine.LoginWriterResponse{}
	}
	return &engine.LoginWriterResponse{
		Error: domain.NewError(http.StatusBadRequest, "invalid password"),
	}
}

func TestWriterSuite(t *testing.T) {
	m := mux.NewRouter()
	server := httptest.NewServer(m)
	defer server.Close()
	initWriter(&testFactory{}, m)
	suite.Run(t, &WriterSuite{
		webSuite{
			server: server,
		},
	})
}

func (s *WriterSuite) TestListHandlerError() {
	s.doRequest(http.MethodGet, domain.PathWriters, map[string]string{}, "").
		checkStatus(http.StatusNotFound)
}

func (s *WriterSuite) TestFindHandler() {
	s.doRequest(http.MethodGet, domain.PathWriters+"/0", map[string]string{}, "").
		checkStatus(http.StatusNotFound)
	var resp engine.FindWriterResponse
	s.get(domain.PathWriters+"/1", url.Values{}, &resp.Writer, "")
}

func (s *WriterSuite) TestCreateHandler() {
	s.doRequest(http.MethodPost, domain.PathWriters, map[string]string{}, "").
		checkStatus(http.StatusCreated)
}

func (s *WriterSuite) TestLoginHandler() {
	s.doRequest(http.MethodPost, domain.PathWritersLogin, map[string]string{
		"email":    "writer@example.com",
		"password": "password",
	}, "").checkStatus(http.StatusOK)
}

func (s *WriterSuite) TestLoginHandlerError() {
	s.doRequest(http.MethodPost, domain.PathWritersLogin, map[string]string{
		"email":    "writer@example.com",
		"password": "invalid password",
	}, "").checkStatus(http.StatusBadRequest)
}
