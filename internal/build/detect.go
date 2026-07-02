package build

import (
        "fmt"
        "os"
        "path/filepath"
)

func DetectBuildSystem(dir string) (string, error) {

        check := func(name string) bool {
                _, err := os.Stat(filepath.Join(dir, name))
                return err == nil
        }

        switch {

        case check("meson.build"):
                return "meson", nil

        case check("CMakeLists.txt"):
                return "cmake", nil

        case check("configure"),
                check("configure.ac"),
                check("configure.in"),
                check("autogen.sh"),
                check("bootstrap"),
                check("bootstrap.sh"):
                return "autotools", nil

        case check("pyproject.toml"),
        check("setup.py"),
        check("setup.cfg"):
        return "python", nil

        case check("Cargo.toml"):
                return "cargo", nil

        case check("go.mod"):
                return "go", nil

        case check("Makefile"),
                check("makefile"),
                check("GNUmakefile"):
                return "make", nil

        default:
                return "", fmt.Errorf("unable to detect build system")
        }
}
