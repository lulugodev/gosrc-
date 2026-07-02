package build

import (
        "fmt"
        "os"
        "os/exec"

        "github.com/lulugodev/gosrc/internal/compiler"
)

type Python struct{}

func (p *Python) Build(dir string) error {

        fmt.Println("==> Building")

        var cmd *exec.Cmd

        switch {
        case exists(dir, "pyproject.toml"):
                cmd = exec.Command(
                        "python3",
                        "-m",
                        "build",
                )

        default:
                cmd = exec.Command(
                        "python3",
                        "setup.py",
                        "build",
                )
        }

        cmd.Dir = dir
        cmd.Env = compiler.Env()
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        return cmd.Run()
}

func (p *Python) Install(dir, destdir string) error {

        fmt.Println("==> Installing")

        var cmd *exec.Cmd

        switch {
        case exists(dir, "pyproject.toml"):
                cmd = exec.Command(
                        "python3",
                        "-m",
                        "installer",
                        "--destdir",
                        destdir,
                        "dist/*.whl",
                )

        default:
                cmd = exec.Command(
                        "python3",
                        "setup.py",
                        "install",
                        "--root="+destdir,
                        "--prefix=/usr",
                )
        }

        cmd.Dir = dir
        cmd.Env = compiler.Env()
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        return cmd.Run()
}

func exists(dir, file string) bool {
        _, err := os.Stat(dir + "/" + file)
        return err == nil
}
