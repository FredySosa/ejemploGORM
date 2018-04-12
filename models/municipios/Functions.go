package municipios

import "github.com/fredysosa/ejemploGORM/models"

var db = models.GetDB()

func init() {
	if !db.HasTable(&Municipio{}) {
		db.AutoMigrate(&Municipio{})
	}
}
