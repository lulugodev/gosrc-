package fetch

import (
        "fmt"
        "io"
        "net/http"
        "os"
        "path/filepath"

        "github.com/lulugodev/gosrc/internal/pkg"
)

type Fetcher struct {
        DistfilesDir string
}

func New(dir string) *Fetcher {
        return &Fetcher{
                DistfilesDir: dir,
        }
}

func (f *Fetcher) Fetch(p *pkg.Package) error {

        fmt.Println("==> Fetching", p.Name)
        fmt.Println("==>", p.Source.URL)

        if err := os.MkdirAll(f.DistfilesDir, 0755); err != nil {
                return err
        }

        filename := filepath.Base(p.Source.URL)
        target := filepath.Join(f.DistfilesDir, filename)

        if _, err := os.Stat(target); err == nil {
                fmt.Println("==> Already downloaded")
                return nil
        }

        resp, err := http.Get(p.Source.URL)
if err != nil {
    return err
}

fmt.Println("Status:", resp.Status)

if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("download failed: %s", resp.Status)
}
        defer resp.Body.Close()

        out, err := os.Create(target)
        if err != nil {
                return err
        }
        defer out.Close()

        _, err = io.Copy(out, resp.Body)
        if err != nil {
                return err
        }

        fmt.Println("==> Saved to")
        fmt.Println(target)

        return nil
}
