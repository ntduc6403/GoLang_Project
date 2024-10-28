package http

import (
	trackUsecase "github.com/thinhhuy997/go-windows/internal/track/usecase"
	"github.com/thinhhuy997/go-windows/pkg/log"
)

type handler struct {
	l      log.Logger
	trackUC trackUsecase.Usecase
}

// // NewHandler returns a new instance of the HTTPHandler interface
// func NewHandler(l log.Logger, albumUC albumUsecase.Usecase) handler {
// 	return handler{
// 		l:      l,
// 		albumUC: albumUC,
// 	}
// }

func NewHandler(l log.Logger, trackUC trackUsecase.Usecase) handler {
	return handler{
		l: l,
		trackUC: trackUC,
	}
}