package config

import (
        "os"
        "path/filepath"
)

type Config struct {
        RepositoryDir string
        CacheDir      string
        DistfilesDir  string
        BuildDir      string
}

func Load() (*Config, error) {

        home, err := os.UserHomeDir()
        if err != nil {
                return nil, err
        }

        cache := filepath.Join(home, ".cache", "gosrc")

        return &Config{
                RepositoryDir: "./repos/core",
                CacheDir:      cache,
                DistfilesDir:  filepath.Join(cache, "distfiles"),
                BuildDir:      filepath.Join(cache, "build"),
        }, nil
}
