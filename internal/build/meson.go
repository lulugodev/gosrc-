package build

import (
        "fmt"
        "os"
        "os/exec"

        "github.com/lulugodev/gosrc/internal/compiler"
)

type Meson struct{}

func (m *Meson) Build(dir string) error {

        fmt.Println("==> Configuring (Meson)")

        cmd := exec.Command(
                "meson",
                "setup",
                "build",
                "--prefix=/usr",
                "--buildtype=release",
                "-Dwerror=false",
        )

        cmd.Dir = dir
        cmd.Env = compiler.Env()
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        if err := cmd.Run(); err != nil {
                return err
        }

        fmt.Println("==> Building")

        cmd = exec.Command(
                "meson",
                "compile",
                "-C",
                "build",
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

func (m *Meson) Install(dir, destdir string) error {

        fmt.Println("==> Installing")

        cmd := exec.Command(
                "meson",
                "install",
                "-C",
                "build",
                "--destdir",
                destdir,
        )

        cmd.Dir = dir
        cmd.Env = compiler.Env()
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        return cmd.Run()
}
