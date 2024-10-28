package middleware

import (
	"github.com/thinhhuy997/go-windows/pkg/log"

	"github.com/thinhhuy997/go-windows/pkg/jwt"
)

type Middleware struct {
	l          log.Logger
	jwtManager jwt.Manager
}

func New(l log.Logger, jwtManager jwt.Manager) Middleware {
	return Middleware{
		l:          l,
		jwtManager: jwtManager,
	}
}
