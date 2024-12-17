package service

type Service interface {
}

type service struct {
}

var _ Service = (*service)(nil)

func NewService() Service {
	return &service{}
}
