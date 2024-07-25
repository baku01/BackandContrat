package Service

import (
	"BackandContrat/Data"
	"gorm.io/gorm"
)

type GormRepository struct {
	Connect *gorm.DB
}

func NewGormRepository() *GormRepository {
	return &GormRepository{Connect: Data.GormDB}
}
