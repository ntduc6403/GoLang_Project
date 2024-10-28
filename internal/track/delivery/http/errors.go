package http

import (
	"github.com/thinhhuy997/go-windows/internal/album/usecase"
	pkgErrors "github.com/thinhhuy997/go-windows/pkg/errors"
	"github.com/thinhhuy997/go-windows/pkg/response"
)

var errMap = response.ErrorMapping{
	usecase.ErrNotFound: pkgErrors.NewHTTPError(404, "album not found"),
}

var (
	errIDIsRequired       = pkgErrors.NewHTTPError(402, "id is required")
	errInvalidID          = pkgErrors.NewHTTPError(403, "id is invalid")
	errInvalidRequestBody = pkgErrors.NewHTTPError(501, "invalid request body")
)

const (
	errMsgNameIsRequired = "name is required"
)
