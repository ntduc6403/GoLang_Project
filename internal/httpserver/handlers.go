package httpserver

import (
	albumHttp "github.com/thinhhuy997/go-windows/internal/album/delivery/http"
	albumRepository "github.com/thinhhuy997/go-windows/internal/album/repository"
	albumPG "github.com/thinhhuy997/go-windows/internal/album/repository/pg"
	albumUsecase "github.com/thinhhuy997/go-windows/internal/album/usecase"
	trackHttp "github.com/thinhhuy997/go-windows/internal/track/delivery/http"
	trackRepository "github.com/thinhhuy997/go-windows/internal/track/repository"
	trackPG "github.com/thinhhuy997/go-windows/internal/track/repository/pg"
	trackUseCase "github.com/thinhhuy997/go-windows/internal/track/usecase"
)

func (srv HTTPServer) mapHandlers() {
	// Repositories
	// var todoRepo todoRepository.Repository
	var albumRepo albumRepository.Repository
	var trackRepo trackRepository.Repository
	// todoRepo = todoPG.NewRepository(srv.l, srv.dbConfig.GetPG())
	albumRepo = albumPG.NewRepository(srv.l, srv.dbConfig.GetPG())
	trackRepo = trackPG.NewRepository(srv.l, srv.dbConfig.GetPG())

	// Usecases
	albumUC := albumUsecase.New(srv.l, albumRepo)
	trackUC := trackUseCase.New(srv.l, trackRepo)

	// Handlers
	album := albumHttp.NewHandler(srv.l, albumUC)
	track := trackHttp.NewHandler(srv.l, trackUC)

	api := srv.gin.Group("/api")

	album.MapRoutes(api.Group("/albums"))
	track.MapRoutes(api.Group("/tracks"))
}
