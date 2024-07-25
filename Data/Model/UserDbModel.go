package Model

import (
	"BackandContrat/Domain/Entities"
	"gorm.io/gorm"
)

type UserDbModel struct {
	gorm.Model
	Id       string `gorm:"primaryKey;type:uuid"`
	Name     string `gorm:"index"`
	Username string `gorm:"unique"`
	Role     string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}

type IDtoToModel interface {
	Convert(userEntity *Entities.User) *UserDbModel
}

type DtoToModel struct {
}

func (dtoToModel *DtoToModel) Convert(userEntity *Entities.User) *UserDbModel {
	newUserDb := UserDbModel{}
	newUserDb.Id = userEntity.Id
	newUserDb.Name = userEntity.Name
	newUserDb.Username = userEntity.Username
	newUserDb.Role = userEntity.Role
	newUserDb.Password = userEntity.Password
	return &newUserDb
}
