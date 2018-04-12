package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
	"github.com/fredysosa/ejemploGORM/models"
	"github.com/fredysosa/ejemploGORM/modules/constantes"
	"github.com/fredysosa/ejemploGORM/controllers/actividades-controller"
)

func main() {

	db := models.GetDB()
	defer db.Close()

	app := gin.Default()
	app.Use(cors.New(constantes.Maincors))

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Hola mundo"})
		return
	})

	rutasActividades := app.Group("/actividades")
	{
		rutasActividades.GET("", actividades_controller.GetActividades)
	}

	app.Run(":" + constantes.GetConfig().Default.Puerto)
	return

}
