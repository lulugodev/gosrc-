package cli

import (
        "fmt"
        "os"
        "path/filepath"

        "github.com/lulugodev/gosrc/internal/build"
        "github.com/lulugodev/gosrc/internal/config"
        "github.com/lulugodev/gosrc/internal/fetch"
        "github.com/lulugodev/gosrc/internal/install"
        "github.com/lulugodev/gosrc/internal/repo"
        "github.com/lulugodev/gosrc/internal/unpack"
        "github.com/lulugodev/gosrc/internal/verify"
)

func Install(atom string) error {

        cfg, err := config.Load()
        if err != nil {
                return err
        }

        r := repo.New(cfg.RepositoryDir)

        p, err := r.Open(atom)
        if err != nil {
                return err
        }

        // Fetch
        f := fetch.New(cfg.DistfilesDir)
        if err := f.Fetch(p); err != nil {
                return err
        }

        archive := filepath.Join(
                cfg.DistfilesDir,
                filepath.Base(p.Source.URL),
        )

        // Verify
        verifier := verify.New()
        if err := verifier.Verify(
                archive,
                p.Source.SHA256,
        ); err != nil {
                return err
        }

        // Unpack
        buildDir := filepath.Join(
                cfg.BuildDir,
                filepath.Dir(atom),
                p.Name+"-"+p.Version,
        )

        u := unpack.New()

        sourceDir, err := u.Unpack(archive, buildDir)
        if err != nil {
                return err
        }

        // Detect build system if omitted
        system := p.Build.System

        if system == "" {
                system, err = build.DetectBuildSystem(sourceDir)
                if err != nil {
                        return err
                }

                fmt.Println("==> Detected build system:", system)
        }

        // Build
        builder, err := build.New(system)
        if err != nil {
                return err
        }

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
        return install.Install(
                filepath.Dir(atom),
                p.Name,
                p.Version,
                files,
        )
}
