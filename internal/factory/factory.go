package factory

import (
	"github.com/alterra-kelompok-1/menu-service/database"
	"github.com/alterra-kelompok-1/menu-service/internal/repository"
)

type Factory struct {
	MenuRepository repository.Menu
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		MenuRepository: repository.NewMenu(db),
	}
}
