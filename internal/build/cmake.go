package build

import (
        "fmt"
        "os"
        "os/exec"
        "path/filepath"

        "github.com/lulugodev/gosrc/internal/compiler"
)

type CMake struct{}

func findSourceDir(dir string) (string, error) {

        // Does the current directory contain CMakeLists.txt?
        if _, err := os.Stat(filepath.Join(dir, "CMakeLists.txt")); err == nil {
                return dir, nil
        }

        // Otherwise look one level down.
        entries, err := os.ReadDir(dir)
        if err != nil {
                return "", err
        }

        for _, entry := range entries {
                if !entry.IsDir() {
                        continue
                }

                subdir := filepath.Join(dir, entry.Name())

                if _, err := os.Stat(filepath.Join(subdir, "CMakeLists.txt")); err == nil {
                        return subdir, nil
                }
        }

        return "", fmt.Errorf("could not locate CMakeLists.txt")
}

func (c *CMake) Build(dir string) error {

        sourceDir, err := findSourceDir(dir)
        if err != nil {
                return err
        }

        buildDir := filepath.Join(sourceDir, "build")

        if err := os.MkdirAll(buildDir, 0755); err != nil {
                return err
        }

        env := []string{
    "CC=clang",
    "CXX=clang++",
    "CFLAGS=-O2 -pipe",
    "CXXFLAGS=-O2 -pipe",
    "LDFLAGS=",
}

        fmt.Println("==> Configuring (CMake)")
        fmt.Println("Source:", sourceDir)

  configure := exec.Command(
    "cmake",
    "..",
    "-DCMAKE_BUILD_TYPE=Release",
    "-DCMAKE_INSTALL_PREFIX=/usr",
    "-DBTOP_BUILD_TESTS=OFF",
)

        configure.Dir = buildDir
        configure.Stdout = os.Stdout
        configure.Stderr = os.Stderr
        cc, err := exec.LookPath("clang")
if err != nil {
    return err
}

cxx, err := exec.LookPath("clang++")
if err != nil {
    return err
}

configure.Env = append(os.Environ(),
    "CC="+cc,
    "CXX="+cxx,
)
        configure.Env = append(os.Environ(), env...)

        if err := configure.Run(); err != nil {
                return err
        }

        fmt.Println("==> Building")

        build := exec.Command("cmake", "--build", ".")

        build.Dir = buildDir
        build.Stdout = os.Stdout
        build.Stderr = os.Stderr

        if err := build.Run(); err != nil {
                return err
        }

        fmt.Println("==> Build complete")

        return nil
}

func (c *CMake) Install(dir, destdir string) error {

        fmt.Println("==> Installing")

        cmd := exec.Command(
                "cmake",
                "--install",
                "build",
        )

        cmd.Dir = dir
        cmd.Env = compiler.Env()
        cmd.Env = append(cmd.Env, "DESTDIR="+destdir)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        return cmd.Run()
}
