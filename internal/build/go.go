package build

import (
        "fmt"
        "os"
        "os/exec"

        "github.com/lulugodev/gosrc/internal/compiler"
)

type Go struct{}

func (g *Go) Build(dir string) error {

        fmt.Println("==> Building")

        cmd := exec.Command(
                "go",
                "build",
                "-v",
                "./...",
        )

        cmd.Dir = dir
        cmd.Env = compiler.Env()
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        if err := cmd.Run(); err != nil {
                return err
        }

        fmt.Println("==> Build complete")

        return nil
}

func (g *Go) Install(dir, destdir string) error {

        fmt.Println("==> Installing")

        binDir := destdir + "/usr/bin"

        if err := os.MkdirAll(binDir, 0755); err != nil {
                return err
        }

        cmd := exec.Command(
                "go",
                "install",
        )

        cmd.Dir = dir
        cmd.Env = append(
                compiler.Env(),
                "GOBIN="+binDir,
        )

        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        if err := cmd.Run(); err != nil {
                return err
        }

        fmt.Println("==> Install complete")

        return nil
}
