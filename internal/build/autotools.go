package build

import (
        "fmt"
        "os"
        "os/exec"
        "path/filepath"

        "github.com/lulugodev/gosrc/internal/compiler"
)

type Autotools struct{}

func fileExists(path string) bool {
        _, err := os.Stat(path)
        return err == nil
}

func run(dir string, name string, args ...string) error {

        cmd := exec.Command(name, args...)
        cmd.Dir = dir
        cmd.Env = compiler.Env()
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        return cmd.Run()
}

func (a *Autotools) Build(dir string) error {

        fmt.Println("==> Configuring")
        fmt.Println("Build directory:", dir)

        configure := filepath.Join(dir, "configure")

        // Bootstrap if configure is missing
        if !fileExists(configure) {

                switch {

                case fileExists(filepath.Join(dir, "autogen.sh")):

                        fmt.Println("==> Running autogen.sh")

                        if err := run(dir, "sh", "autogen.sh"); err != nil {
                                return err
                        }

                case fileExists(filepath.Join(dir, "bootstrap.sh")):

                        fmt.Println("==> Running bootstrap.sh")

                        if err := run(dir, "sh", "bootstrap.sh"); err != nil {
                                return err
                        }

                case fileExists(filepath.Join(dir, "bootstrap")):

                        fmt.Println("==> Running bootstrap")

                        if err := run(dir, "sh", "bootstrap"); err != nil {
                                return err
                        }

                case fileExists(filepath.Join(dir, "configure.ac")),
                        fileExists(filepath.Join(dir, "configure.in")):

                        fmt.Println("==> Running autoreconf")

                        if err := run(dir, "autoreconf", "-fi"); err != nil {
                                return err
                        }

                default:
                        return fmt.Errorf("configure not found")
                }

                if !fileExists(configure) {
                        return fmt.Errorf("failed to generate configure")
                }
        }

        if err := run(
                dir,
                configure,
                "--prefix=/usr",
        ); err != nil {
                return err
        }

        fmt.Println("==> Building")

        if err := run(dir, "make"); err != nil {
                return err
        }

        fmt.Println("==> Build complete")

        return nil
}

func (a *Autotools) Install(dir, destdir string) error {

        fmt.Println("==> Installing")

        if err := run(
                dir,
                "make",
                "DESTDIR="+destdir,
                "PREFIX=/usr",
                "prefix=/usr",
                "install",
        ); err != nil {
                return err
        }

        fmt.Println("==> Install complete")

        return nil
}
