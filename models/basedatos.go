package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"github.com/fredysosa/ejemploGORM/modules/constantes"
)

func init() {

	config := constantes.GetConfig().Postgres

	var err error
	db, err = gorm.Open("postgres", "host="+config.Servidor+" port="+config.Puerto+" user="+config.Usuario+" dbname="+config.NombreBase+" password="+config.Pass)

	if err != nil {
		log.Fatal("Existió un error en la conexión a la base de datos.", err)
	}
}

var (
	db *gorm.DB
)

func GetDB() *gorm.DB {
	return db
}
