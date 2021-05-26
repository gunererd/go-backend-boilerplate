package repositories

import (
	"fmt"

	"gitlab.com/gunererd/dummy-challange/src/models"
	"gitlab.com/gunererd/dummy-challange/src/utils"
)

type UserStore struct {
	users map[string]models.User
}

type userStore interface {
	Get() (models.User, error)
	List() []models.User
	Set(username string, password string)
}

func NewUserStore() *UserStore {
	us := &UserStore{}

	users := make(map[string]models.User)

	us.users = users
	return us
}

func (us *UserStore) Set(username string, password string) {

	u := models.User{
		Username: username,
		Password: password,
	}

	us.users[u.Username] = u
}

func (us *UserStore) Get(username string) (models.User, error) {

	u, ok := us.users[username]
	if ok {
		return u, nil
	} else {
		err := &utils.FError{
			Code:      404,
			ErrorCode: "errors.userNotFound",
			Message:   fmt.Sprintf("user with username=%s not found", username),
		}
		return models.User{}, err
	}
}

func (us *UserStore) List() []models.User {

	users := us.users

	userList := make([]models.User, 0)

	for _, value := range users {
		userList = append(userList, value)
	}

	return userList
}
