package cli

import (
        "path/filepath"

        "github.com/lulugodev/gosrc/internal/config"
        "github.com/lulugodev/gosrc/internal/repo"
        "github.com/lulugodev/gosrc/internal/unpack"
)

func Unpack(atom string) error {

        cfg, err := config.Load()
        if err != nil {
                return err
        }

        r := repo.New(cfg.RepositoryDir)

        p, err := r.Open(atom)
        if err != nil {
                return err
        }

        archive := filepath.Join(
                cfg.DistfilesDir,
                filepath.Base(p.Source.URL),
        )

        buildDir := filepath.Join(
                cfg.BuildDir,
                filepath.Dir(atom),
                p.Name+"-"+p.Version,
        )

        u := unpack.New()

        _, err = u.Unpack(archive, buildDir)
        if err != nil {
                return err
        }

        return nil
}
