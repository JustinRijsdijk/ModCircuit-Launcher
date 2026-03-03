package config

import (
	"encoding/json"
	"os"
)

// JavaSettings holds Java runtime configuration.
type JavaSettings struct {
	PreferredPath *string  `json:"preferredPath"`
	MaxMemoryMb   int      `json:"maxMemoryMb"`
	MinMemoryMb   int      `json:"minMemoryMb"`
	ExtraArgs     []string `json:"extraArgs"`
}

// Settings holds all persisted launcher settings.
type Settings struct {
	SchemaVersion       int          `json:"schemaVersion"`
	PingIntervalMinutes int          `json:"pingIntervalMinutes"`
	Java                JavaSettings `json:"java"`
}

// Default returns a Settings struct populated with sensible defaults.
func Default() *Settings {
	return &Settings{
		SchemaVersion:       1,
		PingIntervalMinutes: 5,
		Java: JavaSettings{
			MaxMemoryMb: 4096,
			MinMemoryMb: 512,
			ExtraArgs:   []string{},
		},
	}
}

// Load reads and parses the settings file at path.
// Missing fields keep their default values.
func Load(path string) (*Settings, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	s := Default()
	if err := json.Unmarshal(data, s); err != nil {
		return nil, err
	}
	return s, nil
}

// Save writes settings to disk as indented JSON.
func Save(path string, s *Settings) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}