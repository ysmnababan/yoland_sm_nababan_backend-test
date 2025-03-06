package murid

import (
	"net/http"
	"soal2/murid/api/service"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service.MuridService
}

func (s *Controller) GetAllMurid(c echo.Context) error {
	data:= s.MuridService.GetAll();
	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"message": "Login success",
			"token":   data,
		})
}