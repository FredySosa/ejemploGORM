package constantes

import (
	"github.com/gin-contrib/cors"
	"time"
	"io/ioutil"
	"log"
	"encoding/json"
)

type DataCfg struct {
	Servidor   string `json:"servidor"`
	Puerto     string `json:"puerto"`
	Usuario    string `json:"usuario"`
	Pass       string `json:"pass"`
	Protocolo  string `json:"protocolo"`
	NombreBase string `json:"nombre_base"`
}

type Config struct {
	Default  *DataCfg `json:"default"`
	Postgres *DataCfg `json:"postgres"`
}

func init() {
	bytes, err := ioutil.ReadFile(FileConfigName)
	if err != nil {
		log.Fatal("No se puede iniciar sin configuracion")
	}

	err = json.Unmarshal(bytes, &Globalconfig)
	if err != nil {
		log.Fatal("No se puede iniciar sin configuracion")
	}
}

func GetConfig() *Config {
	return &Globalconfig
}

var (
	Maincors = cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Accept", "Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	Globalconfig Config
)

const (
	FileConfigName = "config.json"
)
