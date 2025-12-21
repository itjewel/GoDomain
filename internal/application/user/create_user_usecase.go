package user

import "github.com/jewelmia/GoDomain/internal/domain/user"

type UserService struct {
	repo user.UserRepository
}

func NewUserService(repo user.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id string) (*user.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) CreateUser(id, name, email string) (*user.User, error) {
	u := user.NewUser(id, name, email)
	if err := s.repo.Save(u); err != nil {
		return nil, err
	}
	return u, nil
}
