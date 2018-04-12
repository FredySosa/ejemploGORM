package actividad

import "github.com/fredysosa/ejemploGORM/models"

var db = models.GetDB()

func init() {
	if !db.HasTable(&Actividad{}) {
		db.AutoMigrate(&Actividad{})
	}
}

func GetActividades(pagina, limite int64, nombre string) ([]Actividad, error) {
	db := models.GetDB()

	var actividades []Actividad

	db.Where("nomact LIKE ?", "%"+nombre+"%").Limit(limite).Offset(pagina * limite).Find(&actividades)

	return actividades, nil
}
