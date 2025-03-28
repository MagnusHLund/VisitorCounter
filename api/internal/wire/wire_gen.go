// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/MagnusHLund/VisitorCounter/internal/config"
	"github.com/MagnusHLund/VisitorCounter/internal/database"
	"github.com/MagnusHLund/VisitorCounter/internal/handlers"
	"github.com/MagnusHLund/VisitorCounter/internal/services"
)

// Injectors from wire.go:

// CreateApp initializes all dependencies
func CreateApp() (*App, error) {
	configConfig := config.NewConfig()
	db, err := database.NewDatabase(configConfig)
	if err != nil {
		return nil, err
	}
	pageHandler := handlers.NewPageHandler(db)
	visitorHandler := handlers.NewVisitorHandler(db)
	pageService := services.NewPageService(db)
	visitorService := services.NewVisitorService(db)
	app := NewApp(pageHandler, visitorHandler, pageService, visitorService)
	return app, nil
}
