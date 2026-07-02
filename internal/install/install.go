package install

import (
        "fmt"

        "github.com/lulugodev/gosrc/internal/database"
)

func Install(category, name, version string, files []string) error {

        fmt.Println("==> Registering package")

        if err := database.CreatePackage(
                category,
                name,
                version,
                files,
        ); err != nil {
                return err
        }

        fmt.Printf(
                "==> Installed %s-%s (%d files)\n",
                name,
                version,
                len(files),
        )

        return nil
}
