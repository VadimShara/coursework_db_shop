package app

import (
	"github.com/VadimShara/coursework_db_shop/internal/service"
	"github.com/VadimShara/coursework_DB_shop/internal/repo"
	"github.com/VadimShara/coursework_DB_shop/internal/config"

)

type App struct {
	service service.Service
}

func NewApp() *App{
	config := config.Load()

	db := repo.NewDB(config.PGConfig)

	service := service.NewService(db)

	return &App{
		service: service,
	}
}

func (a *App) Run() {
	a.service.Auth()
}