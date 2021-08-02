package common

import (
	"fmt"
	"sync"
)

// App represents an application that will run as a deployable component.
// It's an entrypoint at main.go
type App interface {
	Run(launcher *Launcher) error
}

//LauncherOption defines a function option for Launcher
type LauncherOption func(l *Launcher)

// RunApp start all process registered before to the launcher
func RunApp(name string, app App) LauncherOption {
	return func(l *Launcher) {
		l.Add(name, app)
	}
}

// Launcher manages apps
type Launcher struct {
	apps    map[string]App
	wg      *sync.WaitGroup
	Verbose bool
}

// Add runs an applicaton in a goroutine.
func (l *Launcher) Add(appName string, a App) *Launcher {
	l.apps[appName] = a
	return l
}

// Run run every application registered before with Run method.
func (l *Launcher) Run() {

	count := len(l.apps)
	l.wg.Add(count)

	fmt.Println("Launcher Run")

	print(fmt.Sprintf("Starting %d app(s)\n", count))

	for name, app := range l.apps {

		go func(name string, app App) {
			print("--")
			print(fmt.Sprintf("Launcher: App \u001b[33m(%s)\u001b[0m starting\n", name))

			if err := app.Run(l); err != nil {
				print(fmt.Sprintf("Launcher: App (%s) error:", name))
				print(fmt.Sprintln("\u001b[31m", err, "\u001b[0m"))
			}

			l.wg.Done()

			print(fmt.Sprintf("Launcher: App (%s) finished\n", name))
		}(name, app)
	}

	l.wg.Wait()

	print("Launcher: Terminated")
}

// NewLauncher create an instance of Launch
func NewLauncher(opts ...LauncherOption) *Launcher {

	l := &Launcher{
		apps:    make(map[string]App),
		wg:      new(sync.WaitGroup),
		Verbose: true,
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}
