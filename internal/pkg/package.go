package pkg

type Package struct {
        Name        string `toml:"name"`
        Version     string `toml:"version"`

        Description string `toml:"description"`
        Homepage    string `toml:"homepage"`
        License     string `toml:"license"`

        Source Source `toml:"source"`
        Build  Build  `toml:"build"`

        Dependencies Dependencies `toml:"dependencies"`
}

type Source struct {
        URL    string `toml:"url"`
        SHA256 string `toml:"sha256"`
}

type Build struct {
        System string `toml:"system"`
}

type Dependencies struct {
        Build   []string `toml:"build"`
        Runtime []string `toml:"runtime"`
}
