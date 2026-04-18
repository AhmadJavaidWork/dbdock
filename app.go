package main

import (
	"DBDock/db"
	"DBDock/models"
	"DBDock/services"
	"context"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	appConfigDir      string
	ctx               context.Context
	connectionManager *services.ConnectionManager
	connectionService *services.ConnectionService
	databaseService   *services.DatabaseService
	themeService      *services.ThemeService
}

// NewApp creates a new App application struct
func NewApp(appConfigDir string) *App {
	db.Init(appConfigDir)
	return &App{
		appConfigDir:      appConfigDir,
		connectionManager: services.NewConnectionManager(),
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

func (a *App) shutdown(ctx context.Context) {
	a.connectionManager.DisconnectAll()
}

func (a *App) Reload() {
	runtime.WindowReload(a.ctx)
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

func (a *App) TestConnection(conn models.CreateDBConnection, driver models.DatabaseDriver) (string, error) {
	err := a.connectionService.Test(conn, driver)
	if err != nil {
		return "", err
	}
	return "Connection successfull", nil
}

func (a *App) GetThemeConfig() models.ThemeConfig {
	return a.themeService.LoadTheme()
}

func (a *App) CreateConnection(conn models.CreateDBConnection) (models.SaveDBConnectionResponse, error) {
	connection, err := a.connectionService.Create(conn)
	if err != nil {
		return models.SaveDBConnectionResponse{}, err
	}
	return models.SaveDBConnectionResponse{
		Connection: models.ConnectionWithStatus{
			DBConnection: connection,
			IsConnected:  false,
		},
		Message: "Connection added successfully",
	}, nil
}

func (a *App) GetConnections() ([]models.ConnectionWithStatus, error) {
	connections, err := a.connectionService.GetAll()
	if err != nil {
		return nil, err
	}

	var result []models.ConnectionWithStatus
	for _, conn := range connections {
		result = append(result, models.ConnectionWithStatus{
			DBConnection: conn,
			IsConnected:  a.connectionManager.IsConnected(conn.ID),
		})
	}

	return result, nil
}

func (a *App) UpdateConnection(conn models.DBConnection) (models.SaveDBConnectionResponse, error) {
	connection, err := a.connectionService.Update(conn)
	if err != nil {
		return models.SaveDBConnectionResponse{}, err
	}

	err = a.connectionManager.Disconnect(conn.ID)
	if err != nil {
		return models.SaveDBConnectionResponse{}, err
	}

	return models.SaveDBConnectionResponse{
		Connection: models.ConnectionWithStatus{
			DBConnection: connection,
			IsConnected:  false,
		},
		Message: "Connection updated successfully",
	}, nil
}

func (a *App) DeleteConnection(connectionID int) error {
	err := a.connectionManager.Disconnect(connectionID)
	if err != nil {
		return err
	}

	return a.connectionService.Delete(connectionID)
}

func (a *App) ReadDroppedFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (a *App) ConnectToDatabase(conn models.DBConnection) (string, error) {
	if err := a.connectionManager.Connect(conn); err != nil {
		return "", err
	}
	return "Connected successfully", nil
}

func (a *App) DisconnectFromDatabase(id int) (string, error) {
	if err := a.connectionManager.Disconnect(id); err != nil {
		return "", err
	}

	return "Disconnected successfully", nil
}

func (a *App) GetActiveConnections() ([]models.ConnectionWithStatus, error) {
	connections, err := a.connectionService.GetAll()
	if err != nil {
		return nil, err
	}

	var result []models.ConnectionWithStatus
	for _, conn := range connections {
		if a.connectionManager.IsConnected(conn.ID) {
			result = append(result, models.ConnectionWithStatus{
				DBConnection: conn,
				IsConnected:  true,
			})
		}
	}

	return result, nil
}

func (a *App) RunQuery(connectionID int, query string) ([]map[string]interface{}, error) {
	return a.connectionManager.Query(connectionID, query)
}

func (a *App) ListTables(connectionID int, driver string) ([]models.Table, error) {
	return a.connectionManager.ListTables(connectionID, driver)
}

func (a *App) IsConnected(connectionID int) bool {
	return a.connectionManager.IsConnected(connectionID)
}

func (a *App) GetTableData(connectionID int, driver string, table string, limit, offset int, orderBy string, order models.Order) ([]models.Column, error) {
	return a.connectionManager.GetTableData(connectionID, driver, table, limit, offset, orderBy, order)
}
