package repository

type Repository struct {
	baseUrl string
}

func New(baseUrl string) (*Repository, error) {
	return &Repository{
		baseUrl: baseUrl,
	}, nil
}
