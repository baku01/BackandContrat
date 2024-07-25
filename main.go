package main

import (
	"BackandContrat/Application"
	"BackandContrat/Data"
	"BackandContrat/Data/Model"
)

func main() {
	Data.InitDBConnection()
	err := Data.GormDB.AutoMigrate(&Model.UserDbModel{})
	if err != nil {
		return
	}
	Application.InitWebServer()
}
