package actividad

import "encoding/xml"

type Actividad struct {
	Idact  int64  `gorm:"column:idact;type:integer;AUTO_INCREMENT;PRIMARY_KEY"`
	Nombre string `gorm:"column:nomact;type:varchar(250)"`
}

func (a Actividad) TableName() string {
	return "actividad"
}

type ActividadRespuesta struct {
	Id     int64  `json:"id" xml:"id"`
	Nombre string `json:"nombre" xml:"nombre"`
	URL    string `json:"url" xml:"url"`
}

type ResponseActividad struct {
	XMLName   xml.Name              `json:"-" xml:"actividades"`
	Actividad *[]ActividadRespuesta `json:"actividad" xml:"actividad"`
	Links     []Links               `json:"links" xml:"Links>link"`
}

type Links struct {
	Clave string `json:"link" xml:"link,attr"`
	Valor string `json:"url" xml:"url,attr"`
}
