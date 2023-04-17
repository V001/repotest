package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"repotest/config"
	"repotest/service"
	"repotest/storage"
	"strings"
	"testing"
)

var (
	createUserJSON = `{
  "email": "r1test",
    "password": "test",
  "username": "7018077771"
}`
)

func TestUserHandler_Create(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(createUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := createUserHandler(t)

	if assert.NoError(t, h.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, createUserJSON, rec.Body.String())
	}
}

func createUserHandler(t *testing.T) *Manager {
	cfg, _ := config.New()
	repo, err := storage.New(context.Background(), cfg)
	if err != nil {
		t.Errorf(err.Error())
	}
	srvManager, _ := service.NewManager(repo)
	return NewManager(cfg, srvManager)
}
