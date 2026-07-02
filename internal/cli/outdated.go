package cli

import (
        "fmt"

        "github.com/lulugodev/gosrc/internal/config"
        "github.com/lulugodev/gosrc/internal/outdated"
        "github.com/lulugodev/gosrc/internal/repo"
)

func Outdated() error {

        cfg, err := config.Load()
        if err != nil {
                return err
        }

        r := repo.New(cfg.RepositoryDir)

        updates, err := outdated.Check(r)
        if err != nil {
                return err
        }

        if len(updates) == 0 {
                fmt.Println("All packages are up to date.")
                return nil
        }

        fmt.Println()

        for _, p := range updates {
                fmt.Printf(
                        "%-20s %s -> %s\n",
                        p.Name,
                        p.Installed,
                        p.Latest,
                )
        }

        return nil
}
