package cli

import (
        "fmt"

        "github.com/lulugodev/gosrc/internal/repo"
)

func Search(name string) error {

        r := repo.New("./repos/core")

        packages, err := r.Search(name)
        if err != nil {
                return err
        }

        for _, p := range packages {
                fmt.Printf("%s %s\n", p.Name, p.Version)
        }

        return nil
}
