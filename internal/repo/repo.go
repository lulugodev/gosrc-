package repo

type Repository struct {
        Path string
}

func New(path string) *Repository {
        return &Repository{
                Path: path,
        }
}
