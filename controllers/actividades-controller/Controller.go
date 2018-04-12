package actividades_controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"github.com/fredysosa/ejemploGORM/models/actividad"
)

func GetActividades(ctx *gin.Context) {
	if ctx.GetHeader("Accept") != "application/json" && ctx.GetHeader("Accept") != "application/xml" {
		ctx.Status(http.StatusNotAcceptable)
		return
	}

	pagina, err := strconv.ParseInt(ctx.DefaultQuery("pagina", "0"), 10, 64)
	if err != nil {
		ctx.Status(http.StatusNotAcceptable)
		return
	}
	if pagina <= 0 {
		ctx.Status(http.StatusNotAcceptable)
		return
	}
	pagina--
	limite, err := strconv.ParseInt(ctx.DefaultQuery("limite", "50"), 10, 64)
	if err != nil {
		ctx.Status(http.StatusNotAcceptable)
		return
	}
	name := ctx.DefaultQuery("nombre", "%%")

	actividades, err := actividad.GetActividades(pagina, limite, name)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, actividades)
	return
}
