package verify

import (
        "crypto/sha256"
        "encoding/hex"
        "fmt"
        "io"
        "os"
)

type Verifier struct{}

func New() *Verifier {
        return &Verifier{}
}

func (v *Verifier) Verify(path string, expected string) error {

        fmt.Println("==> Verifying")
        fmt.Println(path)

        file, err := os.Open(path)
        if err != nil {
                return err
        }
        defer file.Close()

        hash := sha256.New()

        if _, err := io.Copy(hash, file); err != nil {
                return err
        }

        actual := hex.EncodeToString(hash.Sum(nil))

        fmt.Println("Expected:", expected)
        fmt.Println("Actual:  ", actual)

        if actual != expected {
                return fmt.Errorf("sha256 verification failed")
        }

        fmt.Println("==> Verification successful")

        return nil
}
