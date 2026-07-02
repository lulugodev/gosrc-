package install

import (
        "fmt"
        "os"
        "path/filepath"

        "github.com/lulugodev/gosrc/internal/database"
)

func Remove(name string) error {
        pkg, err := database.LoadPackage(name)
        if err != nil {
                return err
        }

        fmt.Println("==> Removing", name)

        // Remove files
        for _, file := range pkg.Files {
                fmt.Println("DELETE:", file)

                if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
                        return err
                }

                removeEmptyParents(file)
        }

        return database.RemovePackage(name)
}

func removeEmptyParents(path string) {
        dir := filepath.Dir(path)

        home, _ := os.UserHomeDir()
        stop := filepath.Join(home, ".local")

        for dir != stop && dir != "/" {
                err := os.Remove(dir)
                if err != nil {
                        break
                }

                dir = filepath.Dir(dir)
        }
}
