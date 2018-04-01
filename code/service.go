package code

type Repository interface {
	FetchData(id string) (string, error)
}

type Service struct {
	repo Repository
}

func (s Service) Do() error {
	_, err := s.repo.FetchData("42")
	if err != nil {
		return err
	}
	// do smth. with data
	return nil
}
