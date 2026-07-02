package build

import "fmt"

type Builder interface {
    Build(dir string) error
    Install(dir, destdir string) error
}

func New(system string) (Builder, error) {
        switch system {

case "autotools":
        return &Autotools{}, nil

case "cmake":
        return &CMake{}, nil

case "make":
        return &Make{}, nil

case "cargo":
        return &Cargo{}, nil

case "go":
        return &Go{}, nil

case "python":
        return &Python{}, nil

case "meson":
        return &Meson{}, nil

default:
        return nil, fmt.Errorf("unknown build system: %s", system)
}
}
