package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"repotest/model"
)

// CreateUser godoc
// @Summary      Создание пользователя
// @Description  Создание пользователя
// @ID           CreateOrder
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        rq   body      model.UserCreateReq  true  "Входящие данные"
// @Success	     200  {object}  model.CreateResp
// @Router       /api/v1/user [post]
func (h Manager) CreateUser(c echo.Context) error {
	var req model.User
	if err := c.Bind(&req); err != nil {
		return err
	}
	resp, err := h.srv.User.Create(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, resp)
}
