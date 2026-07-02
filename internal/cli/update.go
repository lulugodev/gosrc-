package cli

import (
        "fmt"
        "os"
        "os/exec"

        "github.com/lulugodev/gosrc/internal/config"
)

func Update() error {

        cfg, err := config.Load()
        if err != nil {
                return err
        }

        fmt.Println("==> Updating repository")

        cmd := exec.Command("git", "pull", "--ff-only")
        cmd.Dir = cfg.RepositoryDir
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        if err := cmd.Run(); err != nil {
                return err
        }

        fmt.Println("==> Repository is up to date")

        return nil
}
