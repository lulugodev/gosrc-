package repo

import (
        "path/filepath"
        "strings"

        "github.com/lulugodev/gosrc/internal/parser"
        "github.com/lulugodev/gosrc/internal/pkg"
)

func (r *Repository) Open(atom string) (*pkg.Package, error) {

        parts := strings.Split(atom, "/")
        if len(parts) != 2 {
                return nil, ErrInvalidAtom
        }

        packagePath := filepath.Join(
                r.Path,
                parts[0],
                parts[1],
                "package.toml",
        )

        return parser.Parse(packagePath)
}
