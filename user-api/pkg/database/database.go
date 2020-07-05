package database

import (
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/config"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/database/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

// CreateDB Creates the database connection with GORM, this is meant to be an database agnostic API, it only requires.
// changes here, to change the database technology.
func CreateDB() error {
	var err error
	c := config.CONFIG.Databases
	switch c.Use {
	case "mysql":
		DB, err = mysql.NewMysql()
	}

	if err != nil{
		log.Println(err)
	}
	return err
}

var DB *gorm.DB