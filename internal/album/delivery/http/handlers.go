package http

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thinhhuy997/go-windows/pkg/response"
)

// func (h handler) add(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	var req addRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		response.Error(c, errInvalidRequestBody)
// 		return
// 	}

// 	input, err := req.toInput()
// 	if err != nil {
// 		response.Error(c, err)
// 		return
// 	}

// 	if err := h.todoUC.Create(ctx, input); err != nil {
// 		h.l.Error(ctx, "todo.handler.add.todoUC.Create: %s", err)
// 		response.ErrorWithMap(c, err, errMap)
// 		return
// 	}

// 	response.OK(c, nil)
// }

// func (h handler) detail(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	rawID := c.Param("id")

// 	if strings.TrimSpace(rawID) == "" {
// 		response.Error(c, errIDIsRequired)
// 		return
// 	}

// 	id, err := strconv.Atoi(rawID)

// 	todo, err := h.todoUC.Detail(ctx, id)
// 	if err != nil {
// 		h.l.Error(ctx, "todo.handler.detail.todoUC.detail")
// 		return
// 	}
// 	response.OK(c, newDetailResponse(todo))
// }

func (h handler) add(c *gin.Context) {
	ctx := c.Request.Context()
	var req addRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errInvalidRequestBody)
		return
	}

	input, err := req.toInput()

	if err != nil {
		response.Error(c, err)
		return
	}

	if err := h.albumUC.Create(ctx, input); err != nil {
		h.l.Error(ctx, "album.handler.add.albumUC.Create: %s", err)
		response.ErrorWithMap(c, err, errMap)
		return
	}
}

func (h handler) list(c *gin.Context) {
	ctx := c.Request.Context()

	albums, err := h.albumUC.List(ctx)
	if err != nil {
		h.l.Error(ctx, "album.handler.list.albumUC.List: %s", err)
		response.ErrorWithMap(c, err, errMap)
		return
	}

	response.OK(c, newListResponse(albums))
}

func (h handler) detail(c *gin.Context) {
	ctx := c.Request.Context()

	rawID := c.Param("id")

	if strings.TrimSpace(rawID) == "" {
		response.Error(c, errIDIsRequired)
		return
	}

	id, err := strconv.Atoi(rawID)

	album, err := h.albumUC.Detail(ctx, id)

	if err != nil {
		h.l.Error(ctx, "album.handler.detail.albumUC.detail")
		return
	}

	response.OK(c, newDetailResponse(album))
}
// func (h handler) delete(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	rawID := c.Param("id")
// 	if strings.TrimSpace(rawID) == "" {
// 		response.Error(c, errIDIsRequired)
// 		return
// 	}

// 	id, err := strconv.Atoi(rawID)
// 	if err != nil {
// 		h.l.Warnf(ctx, "todo.handler.delete.strconv.Atoi(%s): %s", rawID, err)
// 		response.Error(c, errInvalidID)
// 		return
// 	}

// 	if err := h.todoUC.Delete(ctx, id); err != nil {
// 		h.l.Errorf(ctx, "todo.handler.delete.todoUC.Delete(ctx, %d): %s", id, err)
// 		response.ErrorWithMap(c, err, errMap)
// 		return
// 	}

// 	response.OK(c, nil)
// }
