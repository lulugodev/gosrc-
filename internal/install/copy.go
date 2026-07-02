package install

import (
        "io"
        "os"
        "path/filepath"
        "fmt"
        "strings"
)

func CopyToRoot() ([]string, error) {
        var installedFiles []string

        err := filepath.Walk(StageDir, func(path string, info os.FileInfo, err error) error {
                if err != nil {
                        return err
                }

                // Skip the staging directory itself
                if path == StageDir {
                        return nil
                }

    rel, err := filepath.Rel(StageDir, path)
    if err != nil {
        return err
}

     home, err := os.UserHomeDir()
     if err != nil {
         return err
}

    if strings.HasPrefix(rel, "usr/") {
    rel = rel[4:] // remove "usr/"
}

    dst := filepath.Join(home, ".local", rel)

                fmt.Printf("COPY: %s -> %s\n", path, dst)

                // Create directories
                if info.IsDir() {
                        return os.MkdirAll(dst, info.Mode())
                }

                // Skip symlinks for now
                if info.Mode()&os.ModeSymlink != 0 {
                        return nil
                }

                // Ensure parent directory exists
                if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
                        return err
                }

                // Open source file
                src, err := os.Open(path)
                if err != nil {
                        return err
                }
                defer src.Close()

                // Create destination file
                dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
                if err != nil {
                        return err
                }
                defer dstFile.Close()

                // Copy file contents
                if _, err := io.Copy(dstFile, src); err != nil {
                        return err
                }

                // Preserve permissions
                if err := os.Chmod(dst, info.Mode()); err != nil {
                        return err
                }

                // Record installed file
                installedFiles = append(installedFiles, dst)

                return nil
        })

        if err != nil {
                return nil, err
        }

        return installedFiles, nil
}
