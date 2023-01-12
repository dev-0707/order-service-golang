package service

import (
	models "order-service/internal/pkg/models/users"
	"order-service/internal/pkg/persistence"
)

type UserService struct {
	repository persistence.UserRepository
}

//you can get a pointer using the & operator (e.g. &c)
//you can dereference a pointer to get the value using *
func NewUserService(repository persistence.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) Get(id string) (*models.User, error) {
	return persistence.GetUserRepository().Get(id)
}

func GetByUsername(username string) (*models.User, error) {
	return persistence.GetUserRepository().GetByUsername(username)
}

func (s *UserService) All() (*[]models.User, error) {
	return persistence.GetUserRepository().All()
}

func (s *UserService) Query(q *models.User) (*[]models.User, error) {
	return persistence.GetUserRepository().Query(q)
}

func (s *UserService) Add(user *models.User) error {
	return persistence.GetUserRepository().Add(user)
}

func (s *UserService) Update(user *models.User) error {
	return persistence.GetUserRepository().Update(user)
}

func (s *UserService) Delete(user *models.User) error {
	return persistence.GetUserRepository().Delete(user)
}
