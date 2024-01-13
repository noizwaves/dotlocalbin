package pkg

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

const (
	localBinPath = ".local/bin"
	configPath   = ".garb.yml"
)

type Context struct {
	Binaries     []Binary
	BinPath      string
	ConfigPath   string
	Config       *configRoot
	Platform     string
	Architecture string
}

func NewContext() (Context, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Context{}, fmt.Errorf("error determining home directory: %w", err)
	}

	configPath := path.Join(homeDir, configPath)
	config, err := loadConfig(configPath)
	if err != nil {
		return Context{}, fmt.Errorf("error loading config: %w", err)
	}

	binaries := make([]Binary, 0)
	for _, cb := range config.Binaries {
		binary, err := NewBinary(cb)
		if err != nil {
			return Context{}, fmt.Errorf("error constructing binary %q: %w", cb.Name, err)
		}

		binaries = append(binaries, binary)
	}

	return Context{
		Binaries:     binaries,
		BinPath:      path.Join(homeDir, localBinPath),
		ConfigPath:   configPath,
		Config:       &config,
		Platform:     runtime.GOOS,
		Architecture: runtime.GOARCH,
	}, err
}