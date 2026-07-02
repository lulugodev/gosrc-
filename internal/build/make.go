package build

import (
        "fmt"
        "os"
        "os/exec"

        "github.com/lulugodev/gosrc/internal/compiler"
)

type Make struct{}

func (m *Make) Build(dir string) error {

        fmt.Println("==> Building")

        cmd := exec.Command("make")
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

func (m *Make) Install(dir, destdir string) error {

        fmt.Println("==> Installing")

        cmd := exec.Command(
                "make",
                "DESTDIR="+destdir,
                "PREFIX=/usr",
                "prefix=/usr",
                "install",
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
