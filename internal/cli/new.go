package cli

import (
        "fmt"
        "os"
        "path/filepath"
        "strings"

        "github.com/lulugodev/gosrc/internal/config"
)

func New(atom string) error {

        cfg, err := config.Load()
        if err != nil {
                return err
        }

        dir := filepath.Join(cfg.RepositoryDir, atom)

        if err := os.MkdirAll(dir, 0755); err != nil {
                return err
        }

        name := filepath.Base(atom)

        template := fmt.Sprintf(`name = "%s"
version = ""

description = ""
homepage = ""
license = ""

[source]
url = ""
sha256 = ""

[build]
system = ""

[dependencies]
build = []
runtime = []
`, name)

        return os.WriteFile(
                filepath.Join(dir, "package.toml"),
                []byte(strings.TrimSpace(template)+"\n"),
                0644,
        )
}
