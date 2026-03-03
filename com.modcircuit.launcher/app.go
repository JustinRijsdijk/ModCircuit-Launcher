package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"os/exec"
	goRuntime "runtime"
	"time"

	"com.modcircuit.launcher/internal/config"
	"com.modcircuit.launcher/internal/fs"
	"com.modcircuit.launcher/internal/repository"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx           context.Context
	isMaximized   bool
	defaultWidth  int
	defaultHeight int
	paths         *fs.Paths
	settings      *config.Settings
	repoManager   *repository.Manager
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	paths, err := fs.EnsurePaths()
	if err != nil {
		return
	}
	a.paths = paths

	// Load settings — fall back to defaults on any error
	s, err := config.Load(paths.SettingsPath())
	if err != nil {
		s = config.Default()
	}
	a.settings = s

	// Create repository manager and start background ping
	mgr, err := repository.NewManager(ctx)
	if err == nil {
		a.repoManager = mgr
		interval := time.Duration(s.PingIntervalMinutes) * time.Minute
		mgr.StartBackgroundPing(interval)
	}

	screens, err := runtime.ScreenGetAll(ctx)
	if err != nil || len(screens) == 0 {
		return
	}

	primary := screens[0]

	a.defaultWidth = 1200
	a.defaultHeight = 800

	runtime.WindowSetSize(a.ctx, a.defaultWidth, a.defaultHeight)
	runtime.WindowCenter(a.ctx)

	runtime.WindowSetMaxSize(
		ctx,
		primary.Size.Width,
		primary.Size.Height,
	)
}

// domReady is called after front-end resources have been loaded
func (a *App) domReady(ctx context.Context) {
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	if a.repoManager != nil {
		a.repoManager.StopBackgroundPing()
	}
}

/*
Custom functions for window controls
*/

func (a *App) Minimize() {
	runtime.WindowMinimise(a.ctx)
}

// Toggle Maximize / Restore
func (a *App) Maximize() {
	if a.isMaximized {
		// Restore to default size
		runtime.WindowSetSize(a.ctx, a.defaultWidth, a.defaultHeight)
		runtime.WindowCenter(a.ctx)
		a.isMaximized = false
	} else {
		// Maximize manually to primary screen
		screens, err := runtime.ScreenGetAll(a.ctx)
		if err != nil || len(screens) == 0 {
			return
		}

		primary := screens[0]

		runtime.WindowSetSize(
			a.ctx,
			primary.Size.Width,
			primary.Size.Height,
		)
		runtime.WindowSetPosition(a.ctx, 0, 0)
		a.isMaximized = true
	}
}

func (a *App) Restore() {
	runtime.WindowUnmaximise(a.ctx)
}

func (a *App) Close() {
	runtime.Quit(a.ctx)
}

func (a *App) GetAppDir() (string, error) {
	if a.paths == nil {
		return "", fmt.Errorf("paths not initialized")
	}

	return a.paths.AppRoot(), nil
}

func (a *App) GetInstancesDir() (string, error) {
	if a.paths == nil {
		return "", fmt.Errorf("paths not initialized")
	}

	return a.paths.InstancesPath(), nil
}

func (a *App) OpenInstancesDir() error {
	if a.paths == nil {
		return fmt.Errorf("paths not initialized")
	}

	dir := a.paths.InstancesPath()

	var cmd *exec.Cmd

	switch goRuntime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", dir)
	case "darwin":
		cmd = exec.Command("open", dir)
	default:
		cmd = exec.Command("xdg-open", dir)
	}

	return cmd.Start()
}

// -----------------------------------------------------------------------------
// Settings
// -----------------------------------------------------------------------------

func (a *App) GetSettings() *config.Settings {
	if a.settings == nil {
		return config.Default()
	}
	return a.settings
}

func (a *App) SaveSettings(s config.Settings) error {
	if a.paths == nil {
		return fmt.Errorf("paths not initialized")
	}
	if err := config.Save(a.paths.SettingsPath(), &s); err != nil {
		return err
	}
	a.settings = &s
	if a.repoManager != nil {
		interval := time.Duration(s.PingIntervalMinutes) * time.Minute
		a.repoManager.RestartBackgroundPing(interval)
	}
	return nil
}

// -----------------------------------------------------------------------------
// Repositories
// -----------------------------------------------------------------------------

func (a *App) ListRepositories() []repository.Repository {
	if a.repoManager == nil {
		return nil
	}
	return a.repoManager.List()
}

func (a *App) AddRepository(name, url, description string, enabled bool, priority int) (repository.Repository, error) {
	if a.repoManager == nil {
		return repository.Repository{}, fmt.Errorf("manager not initialized")
	}
	repo := repository.Repository{
		ID:          newID(),
		Name:        name,
		URL:         url,
		Description: description,
		Enabled:     enabled,
		Priority:    priority,
	}
	return repo, a.repoManager.Add(repo)
}

func (a *App) UpdateRepository(id, name, url, description string, enabled bool, priority int) error {
	if a.repoManager == nil {
		return fmt.Errorf("manager not initialized")
	}
	return a.repoManager.Update(repository.Repository{
		ID:          id,
		Name:        name,
		URL:         url,
		Description: description,
		Enabled:     enabled,
		Priority:    priority,
	})
}

func (a *App) RemoveRepository(id string) error {
	if a.repoManager == nil {
		return fmt.Errorf("manager not initialized")
	}
	return a.repoManager.Remove(id)
}

func (a *App) PingAll() {
	if a.repoManager != nil {
		a.repoManager.PingAll()
	}
}

func (a *App) PingRepository(id string) {
	if a.repoManager != nil {
		a.repoManager.Ping(id)
	}
}

// -----------------------------------------------------------------------------
// Helpers
// -----------------------------------------------------------------------------

// newID generates a random UUID v4 string.
func newID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}