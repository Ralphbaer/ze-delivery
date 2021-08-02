package app

import (
	"github.com/Ralphbaer/ze-delivery/common"
)

// App is the application glue where we put all top level components to be used
type App struct {
	*Server
}

// Run starts the application
// This is the only necessary code to run an app in main.go
func (app *App) Run() {
	common.NewLauncher(
		common.RunApp("service", app.Server),
	).Run()
}
