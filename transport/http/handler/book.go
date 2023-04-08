package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Manager) Get(c echo.Context) error {
	h.srv.Book.Get(c.Request().Context())
	return c.JSON(http.StatusOK, "Hello world")
}
