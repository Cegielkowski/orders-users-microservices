package errors

import (
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/config"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/database"
	"log"
)

// Log Print the error including the file and line, if the verbose flag is set to true.
func Log(err error){
	if config.CONFIG.Verbose {
		log.Println(err)
	}
}

// ConfigureGORMLog Set the GORM logger to false if the verbose flag is set to false.
func ConfigureGORMLog(){
	if !config.CONFIG.Verbose {
		database.DB.LogMode(true)
	}
}