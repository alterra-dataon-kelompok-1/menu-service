package http

import (
	"github.com/alterra-kelompok-1/menu-service/internal/app/menu"
	"github.com/alterra-kelompok-1/menu-service/internal/factory"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	menu.NewHandler(f).Route(e.Group("/v1/menus"))
}
