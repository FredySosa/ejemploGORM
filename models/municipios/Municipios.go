package municipios

type Municipio struct {
	Idmunicipio int64 `gorm:"column:idmunicipio;type:integer;AUTO_INCREMENT;PRIMARY_KEY"`
	Nombre string `gorm:"column:nombre;type:varchar(250)"`
}

func (m Municipio) TableName() string {
	return "municipio"
}