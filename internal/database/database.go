package database

import (
        "os"
        "path/filepath"
        "strings"

        "github.com/BurntSushi/toml"
)

const DatabaseDir = "/home/whoisluca/.local/share/gosrc/packages"

type Package struct {
        Category string   `toml:"category"`
        Name     string   `toml:"name"`
        Version  string   `toml:"version"`
        Files    []string `toml:"files"`
}

func CreatePackage(category, name, version string, files []string) error {

        if err := os.MkdirAll(DatabaseDir, 0755); err != nil {
                return err
        }

        pkg := Package{
                Category: category,
                Name:     name,
                Version:  version,
                Files:    files,
        }

        file, err := os.Create(filepath.Join(DatabaseDir, name+".toml"))
        if err != nil {
                return err
        }
        defer file.Close()

        return toml.NewEncoder(file).Encode(pkg)
}

func LoadPackage(name string) (*Package, error) {

        var pkg Package

        _, err := toml.DecodeFile(
                filepath.Join(DatabaseDir, name+".toml"),
                &pkg,
        )

        if err != nil {
                return nil, err
        }

        return &pkg, nil
}

func ListPackages() ([]Package, error) {

        var packages []Package

        entries, err := os.ReadDir(DatabaseDir)
        if err != nil {
                if os.IsNotExist(err) {
                        return packages, nil
                }
                return nil, err
        }

        for _, entry := range entries {

                if entry.IsDir() {
                        continue
                }

                if !strings.HasSuffix(entry.Name(), ".toml") {
                        continue
                }

                name := strings.TrimSuffix(entry.Name(), ".toml")

                pkg, err := LoadPackage(name)
                if err != nil {
                        continue
                }

                packages = append(packages, *pkg)
        }

        return packages, nil
}

func RemovePackage(name string) error {
        return os.Remove(filepath.Join(DatabaseDir, name+".toml"))
}

func IsInstalled(name string) bool {

        _, err := os.Stat(filepath.Join(DatabaseDir, name+".toml"))

        return err == nil
}

func Exists(name string) bool {
        return IsInstalled(name)
}

func Version(name string) (string, error) {

        pkg, err := LoadPackage(name)
        if err != nil {
                return "", err
        }

        return pkg.Version, nil
}
