package main

import (
	"fmt"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/config"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/database"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/database/redis"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/errors"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/models"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/utils/general"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	var a app

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = general.ReadConfigJson(&config.CONFIG, filepath.Join(dir, "/config/config.json"))
	if err != nil {
		log.Fatal(err)
	}

	err = database.CreateDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := database.DB.Close()
		if err != nil {
			log.Print(err.Error())
		}
	}()

	redisConn := redis.Conn
	defer func() {
		err := redisConn.Close()
		if err != nil {
			log.Print(err.Error())
		}
	}()

	models.Migrate()

	errors.ConfigureGORMLog()

	a.Initialize()
	fmt.Println("Waiting for connections...")
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}
