package service

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) SaveData() error {
	return nil
}
