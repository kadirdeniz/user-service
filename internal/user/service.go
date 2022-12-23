package user

//go:generate mockgen -source=service.go -destination=../../test/mock/mock_service.go -package=mock
type IService interface{}

type Service struct{}

func NewService() IService {
	return &Service{}
}
