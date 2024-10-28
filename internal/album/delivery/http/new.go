package http

import (
	albumUsecase "github.com/thinhhuy997/go-windows/internal/album/usecase"
	"github.com/thinhhuy997/go-windows/pkg/log"
)

type handler struct {
	l      log.Logger
	albumUC albumUsecase.Usecase
}

// NewHandler returns a new instance of the HTTPHandler interface
func NewHandler(l log.Logger, albumUC albumUsecase.Usecase) handler {
	return handler{
		l:      l,
		albumUC: albumUC,
	}
}
