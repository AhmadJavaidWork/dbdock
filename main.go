package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func initConfig() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", nil
	}

	appPath := filepath.Join(configDir, "dbdock")

	err = os.MkdirAll(appPath, 0755)
	return appPath, err

}

func main() {
	appConfigDir, err := initConfig()
	if err != nil {
		fmt.Println("error", err)
	}

	app := NewApp(appConfigDir)

	err = wails.Run(&options.App{
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
