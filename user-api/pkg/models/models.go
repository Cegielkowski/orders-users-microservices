package models

import "github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/database"

func Migrate(){
	db := database.DB
	db.AutoMigrate(&User{})
}