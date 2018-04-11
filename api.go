package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"fmt"
)

type Actividad struct {
	Idact int64 `gorm:"column:idact;type:integer;AUTO_INCREMENT;PRIMARY_KEY"`
	Nombre string `gorm:"column:nomact;type:varchar(250)"`
}

func (a Actividad) TableName() string {
	return "actividad"
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=hpw password=Frodo7751@")
	defer db.Close()
	if err != nil {
		log.Fatal("Adios", err)
	}

	if !db.HasTable(&Actividad{}){
		db.AutoMigrate(&Actividad{})
	}else{
		fmt.Println("Existe vato")
	}

	var act []Actividad

	db.Find(&act)
	fmt.Println("Actividad",act)

}
