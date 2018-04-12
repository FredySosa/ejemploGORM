package actividad

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
