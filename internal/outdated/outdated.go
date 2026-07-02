package outdated

import (
        "github.com/lulugodev/gosrc/internal/database"
        "github.com/lulugodev/gosrc/internal/repo"
)

type PackageUpdate struct {
        Name      string
        Installed string
        Latest    string
}

func Check(r *repo.Repository) ([]PackageUpdate, error) {

        installed, err := database.ListPackages()
        if err != nil {
                return nil, err
        }

        var updates []PackageUpdate

        for _, pkg := range installed {

                p, err := r.Open(pkg.Category + "/" + pkg.Name)
                if err != nil {
                        continue
                }

                if p.Version != pkg.Version {

                        updates = append(updates, PackageUpdate{
                                Name:      pkg.Name,
                                Installed: pkg.Version,
                                Latest:    p.Version,
                        })
                }
        }

        return updates, nil
}
