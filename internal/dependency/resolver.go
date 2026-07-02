package dependency

import (
        "fmt"

        "github.com/lulugodev/gosrc/internal/pkg"
)

func Resolve(p *pkg.Package) error {

        fmt.Println("==> Resolving dependencies")

        if len(p.Dependencies.Build) == 0 &&
                len(p.Dependencies.Runtime) == 0 {

                fmt.Println("No dependencies.")

                return nil
        }

        if len(p.Dependencies.Build) > 0 {
                fmt.Println("Build dependencies:")

                for _, dep := range p.Dependencies.Build {
                        fmt.Println("  -", dep)
                }
        }

        if len(p.Dependencies.Runtime) > 0 {
                fmt.Println("Runtime dependencies:")

                for _, dep := range p.Dependencies.Runtime {
                        fmt.Println("  -", dep)
                }
        }

        return nil
}
