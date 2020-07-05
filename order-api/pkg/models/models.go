package models

import "github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/database"

func Migrate(){
	db := database.DB
	db.AutoMigrate(&Order{})
}