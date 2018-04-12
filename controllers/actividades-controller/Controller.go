package actividades_controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"github.com/fredysosa/ejemploGORM/models/actividad"
	"github.com/fredysosa/ejemploGORM/modules/general"
)

func GetActividades(ctx *gin.Context) {
	if ctx.GetHeader("Accept") != "application/json" && ctx.GetHeader("Accept") != "application/xml" {
		ctx.Status(http.StatusNotAcceptable)
		return
	}

	pagina, err := strconv.ParseInt(ctx.DefaultQuery("pagina", "1"), 10, 64)
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

	actividades, err := actividad.GetActividades(pagina, limite+1, name)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	if general.EstaVacio(actividades) {
		ctx.Status(http.StatusNotFound)
		return
	}

	respuesta := make(map[string]interface{}, 0)
	respuesta["actividad"] = actividades

	if len(*actividades) > int(limite) {
		if name == "%%" {
			respuesta["links"] = map[string]string{"siguiente": "/actividades?pagina=" + strconv.FormatInt(pagina+2, 10) + "&limite=" + strconv.FormatInt(limite, 10)}
		} else {
			respuesta["links"] = map[string]string{"siguiente": "/actividades?pagina=" + strconv.FormatInt(pagina+2, 10) + "&limite=" + strconv.FormatInt(limite, 10) + "&nombre=" + name}
		}
	}
	if pagina > 0 {
		if name == "%%" {
			respuesta["links"] = map[string]string{"siguiente": "/actividades?pagina=" + strconv.FormatInt(pagina, 10) + "&limite=" + strconv.FormatInt(limite, 10)}
		} else {
			respuesta["links"] = map[string]string{"siguiente": "/actividades?pagina=" + strconv.FormatInt(pagina, 10) + "&limite=" + strconv.FormatInt(limite, 10) + "&nombre=" + name}
		}
	}

	ctx.JSON(http.StatusOK, respuesta)
	return
}
