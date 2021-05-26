package services

import (
	"gitlab.com/gunererd/dummy-challange/src/models"
	"gitlab.com/gunererd/dummy-challange/src/repositories"
)

type UserService interface {
	Create(username string, password string)
	ChangePassword(username string, password string) error
	Get(username string) (models.User, error)
	List() []models.User
}

type userService struct {
	userStore *repositories.UserStore
}

// New initialize a UserService
func NewUserService(userStore *repositories.UserStore) *userService {
	us := &userService{
		userStore: userStore,
	}
	return us
}

// Create creates a new user in the UserStore. Since this is not a serious project
// password is stored as a plain text. In more serious cases pasword should be hashed before storing.
func (us *userService) Create(username string, password string) {
	us.userStore.Set(username, password)
}

// ChangePassword changes password of User.
func (us *userService) ChangePassword(username string, password string) error {

	_, err := us.userStore.Get(username)

	if err != nil {
		return err
	}

	us.userStore.Set(username, password)

	return nil
}

func (us *userService) Get(username string) (models.User, error) {
	u, err := us.userStore.Get(username)

	if err != nil {
		return u, err
	}

	return u, nil
}

func (us *userService) List() []models.User {

	return us.userStore.List()
}
