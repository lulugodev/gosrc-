package parser

import (
        "github.com/BurntSushi/toml"

        "github.com/lulugodev/gosrc/internal/pkg"
)

func Parse(path string) (*pkg.Package, error) {

        var p pkg.Package

        _, err := toml.DecodeFile(path, &p)
        if err != nil {
                return nil, err
        }

        return &p, nil
}
