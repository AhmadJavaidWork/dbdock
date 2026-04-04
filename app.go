package main

import (
	"DBDock/db"
	"DBDock/models"
	"DBDock/services"
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	appConfigDir      string
	ctx               context.Context
	connectionService *services.ConnectionService
	databaseService   *services.DatabaseService
	themeService      *services.ThemeService
}

// NewApp creates a new App application struct
func NewApp(appConfigDir string) *App {
	db.Init(appConfigDir)
	return &App{
		appConfigDir:      appConfigDir,
		connectionService: services.NewConnectionService(),
		databaseService:   services.NewDatabaseService(),
		themeService:      services.NewThemeService(appConfigDir),
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

func (a *App) GetSupportedDatabases() ([]models.DatabaseDriver, error) {
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

func (a *App) CreateConnection(conn models.DBConnection) (string, error) {
	err := a.connectionService.Create(conn)
	if err != nil {
		return "", err
	}
	return "Connection saved successfully", nil
}
