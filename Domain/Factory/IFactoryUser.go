package Factory

import (
	"BackandContrat/Domain/Entities"
	"BackandContrat/Utils"
)

type IFactoryUser interface {
	Create(name string, username string, password string) *Entities.User
}

type CandidatoFactory struct {
}

type RecrutadorFactory struct {
}

func (candidatoFactory *CandidatoFactory) Create(name string, username string, password string) *Entities.User {
	return &Entities.User{Id: Utils.GenerateUUID(), Name: name, Username: username, Role: "Candidato", Password: password}
}

func (recrutadorFactory *RecrutadorFactory) Create(name string, username string, password string) *Entities.User {
	return &Entities.User{Id: Utils.GenerateUUID(), Name: name, Username: username, Role: "Recrutador", Password: password}
}
