package actividad

import (
	"github.com/fredysosa/ejemploGORM/models"
	"strconv"
)

var db = models.GetDB()

func init() {
	if !db.HasTable(&Actividad{}) {
		db.AutoMigrate(&Actividad{})
	}
}

func GetActividades(pagina, limite int64, nombre string) (*[]ActividadRespuesta, error) {
	db := models.GetDB()

	var actividades []Actividad

	db.Where("nomact LIKE ?", "%"+nombre+"%").Limit(limite).Offset(pagina * limite).Find(&actividades)

	actividadesResp := make([]ActividadRespuesta, 0, len(actividades))
	for _, actividad := range actividades {
		actividadesResp = append(actividadesResp, ActividadRespuesta{actividad.Idact, actividad.Nombre, "/actividades/" + strconv.FormatInt(actividad.Idact, 10)})
	}

	return &actividadesResp, nil
}
