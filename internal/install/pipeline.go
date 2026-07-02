package install

import (
        "fmt"

        "github.com/lulugodev/gosrc/internal/pkg"
)

type Pipeline struct {
        Package *pkg.Package
}

func NewPipeline(p *pkg.Package) *Pipeline {
        return &Pipeline{
                Package: p,
        }
}

func (p *Pipeline) Run() error {

        fmt.Println("==> Starting install pipeline")

        // Dependency resolution
        // Fetch
        // Verify
        // Unpack
        // Build
        // Install
        // Register

        fmt.Println("==> Pipeline finished")

        return nil
}
