package cli

import "fmt"
import "errors"

func Run(args []string) error {

        if len(args) == 0 {
                return fmt.Errorf("no command specified")
        }

        switch args[0] {

        case "version":
                fmt.Println("gosrc 0.1.0")

        case "search":
                if len(args) != 2 {
                        return fmt.Errorf("usage: gosrc search <package>")
                }
                return Search(args[1])

        case "info":
                if len(args) != 2 {
                        return fmt.Errorf("usage: gosrc info <category/package>")
                }
                return Info(args[1])

        case "update":
        return Update()

        case "outdated":
        return Outdated()

        case "fetch":
                if len(args) != 2 {
                        return fmt.Errorf("usage: gosrc fetch <category/package>")
                }
                return Fetch(args[1])

        case "verify":
                if len(args) != 2 {
                        return fmt.Errorf("usage: gosrc verify <category/package>")
                }
                return Verify(args[1])

        case "unpack":
                if len(args) != 2 {
                        return fmt.Errorf("usage: gosrc unpack <category/package>")
                }
                return Unpack(args[1])

        case "build":
                if len(args) != 2 {
                        return fmt.Errorf("usage: gosrc build <category/package>")
                }
                return Build(args[1])

        case "remove":
                if len(args) != 2 {
                        return fmt.Errorf("usage: gosrc remove <package>")
                }
                return Remove(args[1])

        case "install":
        if len(args) != 2 {
                return fmt.Errorf("usage: gosrc install <category/package>")
        }
        return Install(args[1])

        case "new":
    if len(args) < 2 {
        return errors.New("usage: gosrc new <category/package>")
    }
    return New(args[1])

        default:
                return fmt.Errorf("unknown command: %s", args[0])
        }

        return nil
}
