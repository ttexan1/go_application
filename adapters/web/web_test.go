package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/stretchr/testify/suite"
	"github.com/ttexan1/golang-simple/engine"
)

type testFactory struct{}

func (f *testFactory) NewCategory() engine.Category { return &testCategoryEngine{} }
func (f *testFactory) NewArticle() engine.Article   { return nil }
func (f *testFactory) NewWriter() engine.Writer     { return nil }

type webSuite struct {
	suite.Suite
	server *httptest.Server
}

type webResp struct {
	s   suite.Suite
	res *http.Response
}

func (s *webSuite) doRequest(method, path string, body interface{}, token string) *webResp {
	posting, err := json.Marshal(&body)
	s.Require().NoError(err)
	client := &http.Client{}
	req, err := http.NewRequest(method, s.server.URL+path, bytes.NewBuffer(posting))
	s.Require().NoError(err)
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := client.Do(req)
	s.Require().NoError(err)
	return &webResp{s: s.Suite, res: res}
}

func (s *webSuite) get(path string, params url.Values, target interface{}, token string) {
	query := ""
	if len(params) > 0 {
		query = "?" + params.Encode()
	}
	s.doRequest(http.MethodGet, path+query, []byte{}, token).
		checkStatus(http.StatusOK).
		decode(target)
}

func (a *webResp) checkStatus(code int) *webResp {
	a.s.Equal(code, a.res.StatusCode)
	return a
}

func (a *webResp) decode(target interface{}) *webResp {
	defer a.res.Body.Close()
	err := json.NewDecoder(a.res.Body).Decode(target)
	a.s.Require().NoError(err)
	return a
}
