package unpack

import (
        "archive/tar"
        "compress/gzip"
        "fmt"
        "io"
        "os"
        "path/filepath"
)

type Unpacker struct{}

func New() *Unpacker {
        return &Unpacker{}
}

func (u *Unpacker) Unpack(archive, destination string) (string, error) {

        fmt.Println("==> Unpacking")
        fmt.Println("Archive:     ", archive)
        fmt.Println("Destination: ", destination)

        _ = os.RemoveAll(destination)

        if err := os.MkdirAll(destination, 0755); err != nil {
                return "", err
        }

        file, err := os.Open(archive)
        if err != nil {
                return "", err
        }
        defer file.Close()

        gzr, err := gzip.NewReader(file)
        if err != nil {
                return "", err
        }
        defer gzr.Close()

        tr := tar.NewReader(gzr)

        for {

                header, err := tr.Next()

                if err == io.EOF {
                        break
                }

                if err != nil {
                        return "", err
                }

                target := filepath.Join(destination, header.Name)

                switch header.Typeflag {

                case tar.TypeDir:

                        if err := os.MkdirAll(target, os.FileMode(header.Mode)); err != nil {
                                return "", err
                        }

                        // Preserve directory timestamp
                        _ = os.Chtimes(target, header.ModTime, header.ModTime)

                case tar.TypeReg:

                        if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
                                return "", err
                        }

                        out, err := os.OpenFile(
                                target,
                                os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
                                os.FileMode(header.Mode),
                        )
                        if err != nil {
                                return "", err
                        }

                        if _, err := io.Copy(out, tr); err != nil {
                                out.Close()
                                return "", err
                        }

                        if err := out.Close(); err != nil {
                                return "", err
                        }

                        // Preserve permissions
                        if err := os.Chmod(target, os.FileMode(header.Mode)); err != nil {
                                return "", err
                        }

                        // Preserve original timestamp
                        if err := os.Chtimes(target, header.ModTime, header.ModTime); err != nil {
                                return "", err
                        }

                case tar.TypeSymlink:

                        if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
                                return "", err
                        }

                        if err := os.Symlink(header.Linkname, target); err != nil && !os.IsExist(err) {
                                return "", err
                        }

                case tar.TypeLink:

                        link := filepath.Join(destination, header.Linkname)

                        if err := os.Link(link, target); err != nil {
                                return "", err
                        }
                }
        }

        entries, err := os.ReadDir(destination)
        if err != nil {
                return "", err
        }

        if len(entries) == 1 && entries[0].IsDir() {
                source := filepath.Join(destination, entries[0].Name())
                fmt.Println("==> Source:", source)
                return source, nil
        }

        fmt.Println("==> Source:", destination)

        return destination, nil
}
