package Service

import (
	"BackandContrat/Domain/Entities"
	"BackandContrat/Domain/Factory"
)

type IUserRepository interface {
	CreateUser(factory Factory.IFactoryUser, Name string, Username string, Password string)
	ListAllUsers() []Entities.User
}

type UserRepository struct {
}

func (userRepository *UserRepository) CreateUser(factory Factory.IFactoryUser, Name string, Username string, Password string) {
	user := factory.Create(Name, Username, Password)
	gormRepository := NewGormRepository()
	gormRepository.Connect.Create(user)
}

func (userRepository *UserRepository) ListAllUsers() []Entities.User {
	gormRepository := NewGormRepository()

	var users []Entities.User
	gormRepository.Connect.Find(&users)

	return users
}
