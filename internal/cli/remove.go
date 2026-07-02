package cli

import "github.com/lulugodev/gosrc/internal/install"

func Remove(atom string) error {
        return install.Remove(atom)
}
