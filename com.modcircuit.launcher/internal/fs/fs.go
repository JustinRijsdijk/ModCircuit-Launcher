package fs

import (
	"os"
	"path/filepath"
)

type Paths struct {
	appRoot       string
	settingsPath  string
	instancesPath string
}

// EnsurePaths initializes the app directory, settings file, and instances directory
func EnsurePaths() (*Paths, error) {
	// Get user home directory
	base, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// App root
	appRoot := filepath.Join(base, ".ModCircuitLauncher")
	if err := os.MkdirAll(appRoot, 0o755); err != nil {
		return nil, err
	}

	// Settings file
	settingsPath := filepath.Join(appRoot, "settings.json")
	if info, err := os.Stat(settingsPath); err != nil {
		if os.IsNotExist(err) {
			f, err := os.OpenFile(settingsPath, os.O_CREATE|os.O_WRONLY, 0o644)
			if err != nil {
				return nil, err
			}
			_, _ = f.Write([]byte("{}"))
			_ = f.Close()
		} else {
			return nil, err
		}
	} else if info.IsDir() {
		return nil, os.ErrInvalid
	}

	// Instances directory
	instancesPath := filepath.Join(appRoot, "instances")
	if err := os.MkdirAll(instancesPath, 0o755); err != nil {
		return nil, err
	}

	return &Paths{
		appRoot:       appRoot,
		settingsPath:  settingsPath,
		instancesPath: instancesPath,
	}, nil
}

// Returns app root path
func (p *Paths) AppRoot() string {
	return p.appRoot
}

// Returns settings file path
func (p *Paths) SettingsPath() string {
	return p.settingsPath
}

// Returns instances directory path
func (p *Paths) InstancesPath() string {
	return p.instancesPath
}
