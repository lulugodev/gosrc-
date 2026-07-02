package cli

import (
        "os"
        "path/filepath"

        "github.com/lulugodev/gosrc/internal/build"
        "github.com/lulugodev/gosrc/internal/config"
        "github.com/lulugodev/gosrc/internal/dependency"
        "github.com/lulugodev/gosrc/internal/install"
        "github.com/lulugodev/gosrc/internal/repo"
)

func Build(atom string) error {

        cfg, err := config.Load()
        if err != nil {
                return err
        }

        r := repo.New(cfg.RepositoryDir)

        p, err := r.Open(atom)
        if err != nil {
                return err
        }

        if err := dependency.Resolve(p); err != nil {
                return err
        }

        sourceDir := filepath.Join(
                cfg.BuildDir,
                filepath.Dir(atom),
                p.Name+"-"+p.Version,
        )

        builder, err := build.New(p.Build.System)
        if err != nil {
                return err
        }

        // Build
        if err := builder.Build(sourceDir); err != nil {
                return err
        }

        // Clean staging directory
        _ = os.RemoveAll(install.StageDir)

        if err := os.MkdirAll(install.StageDir, 0755); err != nil {
                return err
        }

        // Install into staging
        if err := builder.Install(sourceDir, install.StageDir); err != nil {
                return err
        }

        // Copy staged files
        files, err := install.CopyToRoot()
        if err != nil {
                return err
        }

        // Register package
        if err := install.Install(
                filepath.Dir(atom),
                p.Name,
                p.Version,
                files,
        ); err != nil {
                return err
        }

        return nil
}
