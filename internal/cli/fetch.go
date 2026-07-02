package cli

import (
        "github.com/lulugodev/gosrc/internal/config"
        "github.com/lulugodev/gosrc/internal/fetch"
        "github.com/lulugodev/gosrc/internal/repo"
)

func Fetch(atom string) error {

        cfg, err := config.Load()
        if err != nil {
                return err
        }

        r := repo.New(cfg.RepositoryDir)

        p, err := r.Open(atom)
        if err != nil {
                return err
        }

        f := fetch.New(cfg.DistfilesDir)

        return f.Fetch(p)
}
