package main

import (
	"DBDock/services"
	"embed"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func initConfig() *services.ThemeService {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return services.NewThemeService("")
	}
	appConfigDir := filepath.Join(configDir, "dbdock")
	return services.NewThemeService(appConfigDir)
}

func main() {
	config := initConfig()

	app := NewApp(config)

	err := wails.Run(&options.App{
		Title:  "DBDock",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
