package httpserver

import (
	"github.com/thinhhuy997/go-windows/internal/appconfig/dbmanger"
	pkgLog "github.com/thinhhuy997/go-windows/pkg/log"

	"github.com/gin-gonic/gin"
)

const (
	DBDriverPG = "postgres"
)

type HTTPServer struct {
	gin      *gin.Engine
	l        pkgLog.Logger
	port     int
	dbConfig dbmanger.DBConfig
}

type Config struct {
	Port     int
	DBConfig dbmanger.DBConfig
}

func New(l pkgLog.Logger, cfg Config) *HTTPServer {
	return &HTTPServer{
		l:        l,
		gin:      gin.Default(),
		port:     cfg.Port,
		dbConfig: cfg.DBConfig,
	}
}
