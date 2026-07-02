package repo

import (
        "os"
        "path/filepath"
        "strings"

        "github.com/lulugodev/gosrc/internal/parser"
        "github.com/lulugodev/gosrc/internal/pkg"
)

func (r *Repository) Search(name string) ([]*pkg.Package, error) {

        var packages []*pkg.Package

        err := filepath.Walk(r.Path, func(path string, info os.FileInfo, err error) error {

                if err != nil {
                        return err
                }

                if info.IsDir() {
                        return nil
                }

                if filepath.Ext(path) != ".toml" {
    return nil
    }

                p, err := parser.Parse(path)
                if err != nil {
                        return err
                }

                if strings.Contains(p.Name, name) {
                        packages = append(packages, p)
                }

                return nil
        })

        return packages, err
}
