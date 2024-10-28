package main

import (
	"github.com/thinhhuy997/go-windows/config"
	"github.com/thinhhuy997/go-windows/internal/appconfig/dbmanger"
	"github.com/thinhhuy997/go-windows/internal/httpserver"
	pkgLog "github.com/thinhhuy997/go-windows/pkg/log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	dbConfig, err := dbmanger.NewDBConfig(cfg)
	if err != nil {
		panic(err)
	}

	l := pkgLog.InitializeZapLogger(pkgLog.ZapConfig{
		Level:    cfg.Logger.Level,
		Mode:     cfg.Logger.Mode,
		Encoding: cfg.Logger.Encoding,
	})

	srv := httpserver.New(l, httpserver.Config{
		Port:     cfg.HTTPServer.Port,
		DBConfig: dbConfig,
	})
	srv.Run()
}
