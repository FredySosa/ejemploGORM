package actividades_controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"github.com/fredysosa/ejemploGORM/models/actividad"
	"github.com/fredysosa/ejemploGORM/modules/general"
)

func GetActividades(ctx *gin.Context) {

	header := ctx.GetHeader("Accept")

	if header != "application/json" && header != "application/xml" {
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

	actividades, limit, err := actividad.GetActividades(pagina, limite, name)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	if general.EstaVacio(actividades) {
		ctx.Status(http.StatusNotFound)
		return
	}

	var respuesta actividad.ResponseActividad
	respuesta.Actividad = actividades

	links := make([]actividad.Links, 0, 1)

	if limit {
		if name == "%%" {
			links = append(links, actividad.Links{"siguiente", "/actividades?pagina=" + strconv.FormatInt(pagina+2, 10) + "&limite=" + strconv.FormatInt(limite, 10)})
		} else {
			links = append(links, actividad.Links{"siguiente", "/actividades?pagina=" + strconv.FormatInt(pagina+2, 10) + "&limite=" + strconv.FormatInt(limite, 10) + "&nombre=" + name})
		}
	}

	if pagina > 0 {
		if name == "%%" {
			links = append(links, actividad.Links{"anterior", "/actividades?pagina=" + strconv.FormatInt(pagina, 10) + "&limite=" + strconv.FormatInt(limite, 10)})
		} else {
			links = append(links, actividad.Links{"anterior", "/actividades?pagina=" + strconv.FormatInt(pagina, 10) + "&limite=" + strconv.FormatInt(limite, 10) + "&nombre=" + name})
		}
	}

	respuesta.Links = links
	if header == "application/json" {
		ctx.JSON(http.StatusOK, respuesta)
	} else {
		ctx.XML(http.StatusOK, respuesta)
	}

	return
}
