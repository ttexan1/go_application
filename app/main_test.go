package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ttexan1/golang-simple/adapters/web"
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"
	"github.com/ttexan1/golang-simple/providers/sql"
)

var normalEmail = "tetsuji@example.com"
var normalPassword = "tetsuji"

func TestMain(t *testing.T) {
	config, err := domain.NewConfig("")
	require.NoError(t, err)
	domain.JWTSigningKey = "test"
	config.Port = ":8081"
	config.PostgresURL = "host=localhost dbname=golang_practice_test sslmode=disable"
	if dbURL := os.Getenv("TEST_DB_URL"); dbURL != "" {
		config.PostgresURL = dbURL
	}

	s, err := sql.NewStorage(config.Driver, config.PostgresURL)
	require.NoError(t, err)
	defer s.Close()

	s.DropTables()
	s.Migrate()

	w := &domain.Writer{
		Email:  normalEmail,
		Status: domain.WriterStatusValid,
		Name:   "てつじ",
	}
	w.SetPassword(normalPassword)
	_, derr := s.NewWriterRepo().Create(w)
	if derr != nil {
		t.Fatal(derr)
	}

	e := engine.NewEngine(s)

	go func() {
		log.Printf("Listening port %s ...\n", config.Port)
		http.ListenAndServe(config.Port, web.NewAdapter(e, config))
	}()

	time.Sleep(100 * time.Millisecond)

	et := &e2etest{t, "http://localhost" + config.Port}
	et.run()
}

type e2etest struct {
	*testing.T
	baseURL string
}

func (t *e2etest) run() {
	// First log them in.
	adminToken := t.login(domain.PathWritersLogin, map[string]interface{}{
		"email":    normalEmail,
		"password": normalPassword,
	})

	// create a category
	category := domain.Category{}
	t.doRequest(http.MethodPost, domain.PathCategories, map[string]string{
		"name": "category",
	}, adminToken).checkStatus(http.StatusCreated).decode(&category)
	for i := 1; i < 21; i++ {
		t.doRequest(http.MethodPost, domain.PathCategories, map[string]string{
			"name": fmt.Sprintf("category%d", i),
		}, adminToken).checkStatus(http.StatusCreated)
	}

	// list categories
	var categories []*domain.Category
	(func() {
		res, err := http.Get(t.baseURL + "/categories")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode)
		defer res.Body.Close()
		err = json.NewDecoder(res.Body).Decode(&categories)
		require.NoError(t, err)

		assert.Equal(t, 20, len(categories))
		assert.Equal(t, "http://localhost:8081/categories?offset=20", res.Header.Get("X-Next-Page"))
		assert.Equal(t, "", res.Header.Get("X-Prev-Page"))
		assert.Equal(t, "21", res.Header.Get("X-Total-Count"))
	})()

	// create an article
	article := domain.Article{}
	t.doRequest(http.MethodPost, domain.PathArticles, map[string]interface{}{
		"category_id": category.ID,
		"status":      domain.ArticleStatusDraft,
		"title":       "this article is awesome",
		"writer_id":   1,
	}, adminToken).checkStatus(http.StatusCreated).decode(&article)
	assert.Equal(t, domain.ArticleStatusDraft, article.Status)
}

func (t *e2etest) get(path string, params url.Values, target interface{}, token string) {
	query := ""
	if len(params) > 0 {
		query = "?" + params.Encode()
	}
	t.doRequest(http.MethodGet, path+query, []byte{}, token).
		checkStatus(http.StatusOK).
		decode(target)
}

func (t *e2etest) login(path string, target interface{}) string {
	var m map[string]string
	t.doRequest(http.MethodPost, path, target, "").
		checkStatus(http.StatusOK).
		decode(&m)
	token, ok := m["token"]
	require.True(t, ok)
	return token
}

func (t *e2etest) doRequest(method, path string, body interface{}, token string) *appResp {
	posting, err := json.Marshal(&body)
	require.NoError(t, err)
	client := &http.Client{}
	req, err := http.NewRequest(method, t.baseURL+path, bytes.NewBuffer(posting))
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := client.Do(req)
	require.NoError(t, err)
	return &appResp{t: t, res: res}
}

type appResp struct {
	t   *e2etest
	res *http.Response
}

func (a *appResp) checkStatus(code int) *appResp {
	require.Equal(a.t, code, a.res.StatusCode)
	return a
}

func (a *appResp) decode(target interface{}) *appResp {
	defer a.res.Body.Close()
	err := json.NewDecoder(a.res.Body).Decode(target)
	require.NoError(a.t, err)
	return a
}
