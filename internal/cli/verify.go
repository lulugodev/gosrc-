package cli

import (
        "path/filepath"

        "github.com/lulugodev/gosrc/internal/config"
        "github.com/lulugodev/gosrc/internal/repo"
        "github.com/lulugodev/gosrc/internal/verify"
)

func Verify(atom string) error {

        cfg, err := config.Load()
        if err != nil {
                return err
        }

        r := repo.New(cfg.RepositoryDir)

        p, err := r.Open(atom)
        if err != nil {
                return err
        }

        filename := filepath.Base(p.Source.URL)
        archive := filepath.Join(cfg.DistfilesDir, filename)

        v := verify.New()

        return v.Verify(archive, p.Source.SHA256)
}
