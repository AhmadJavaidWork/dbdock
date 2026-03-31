package main

import (
	"DBDock/models"
	"DBDock/services"
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx               context.Context
	connectionService *services.ConnectionService
	databaseService   *services.DatabaseService
	themeService      *services.ThemeService
}

// NewApp creates a new App application struct
func NewApp(themeService *services.ThemeService) *App {
	return &App{
		connectionService: services.NewConnectionService(),
		databaseService:   services.NewDatabaseService(),
		themeService:      themeService,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) SetTheme(dark bool) {
	if dark {
		runtime.WindowSetDarkTheme(a.ctx)
	} else {
		runtime.WindowSetLightTheme(a.ctx)
	}
}

func (a *App) GetSupportedDatabases() []models.DatabaseDriver {
	return a.databaseService.ListSupported()
}

func (a *App) TestConnection(conn models.DBConnection) (string, error) {
	err := a.connectionService.Test(conn)
	if err != nil {
		return "", err
	}
	return "Connection successfull", nil
}

func (a *App) GetThemeConfig() models.ThemeConfig {
	return a.themeService.LoadTheme()
}
