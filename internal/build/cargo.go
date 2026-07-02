package build

import (
        "fmt"
        "os"
        "os/exec"

        "github.com/lulugodev/gosrc/internal/compiler"
)

type Cargo struct{}

func (c *Cargo) Build(dir string) error {

        fmt.Println("==> Building")

        cmd := exec.Command(
                "cargo",
                "build",
                "--release",
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

func (c *Cargo) Install(dir, destdir string) error {

        fmt.Println("==> Installing")

        cmd := exec.Command(
                "cargo",
                "install",
                "--path",
                ".",
                "--root",
                destdir+"/usr",
                "--locked",
        )

        cmd.Dir = dir
        cmd.Env = compiler.Env()
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        if err := cmd.Run(); err != nil {
                return err
        }

        fmt.Println("==> Install complete")

        return nil
}
