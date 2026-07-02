package cli

import (
        "fmt"
        "strings"

        "github.com/lulugodev/gosrc/internal/repo"
)

func Info(atom string) error {

        r := repo.New("./repos/core")

        p, err := r.Open(atom)
        if err != nil {
                return err
        }

        fmt.Println("Name:       ", p.Name)
        fmt.Println("Version:    ", p.Version)
        fmt.Println("Description:", p.Description)
        fmt.Println("Homepage:   ", p.Homepage)
        fmt.Println("License:    ", p.License)
        fmt.Println("Build:      ", p.Build.System)

        fmt.Println()

        fmt.Println("Source")
        fmt.Println("------")
        fmt.Println(p.Source.URL)

        fmt.Println()

        fmt.Println("Build Dependencies")
        fmt.Println("------------------")

        if len(p.Dependencies.Build) == 0 {
                fmt.Println("None")
        } else {
                fmt.Println(strings.Join(p.Dependencies.Build, "\n"))
        }

        fmt.Println()

        fmt.Println("Runtime Dependencies")
        fmt.Println("--------------------")

        if len(p.Dependencies.Runtime) == 0 {
                fmt.Println("None")
        } else {
                fmt.Println(strings.Join(p.Dependencies.Runtime, "\n"))
        }

        return nil
}
