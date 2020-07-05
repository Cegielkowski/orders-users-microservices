package mysql

import (
	"fmt"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// user/pass/host/port/dbname.
const mysqlConnString = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true"

// NewMysql Creates a connection in mysql.
func NewMysql() (*gorm.DB, error) {
	c := config.CONFIG.Databases.Mysql

	db, err := gorm.Open("mysql", fmt.Sprintf(mysqlConnString, c.User, c.Password, c.Host, c.Port, c.Database))
	if err != nil{
		log.Println(err)
		return nil, err
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.Set("gorm:table_options", "collation_connection=utf8_general_ci")

	return db, nil
}

