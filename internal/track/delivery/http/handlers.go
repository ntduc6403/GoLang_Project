package http

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thinhhuy997/go-windows/pkg/response"
)

// func (h handler) list(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	albums, err := h.albumUC.List(ctx)
// 	if err != nil {
// 		h.l.Error(ctx, "album.handler.list.albumUC.List: %s", err)
// 		response.ErrorWithMap(c, err, errMap)
// 		return
// 	}

// 	response.OK(c, newListResponse(albums))
// }


func (h handler) list(c *gin.Context) {
	ctx := c.Request.Context()
	tracks, err := h.trackUC.All(ctx)
	if err != nil {
		h.l.Error(ctx, "track.handler.list.trackUC.All: %s", err)
		response.ErrorWithMap(c, err, errMap)
		return
	}

	response.OK(c, newListResponse(tracks))
}

func (h handler) add (c *gin.Context) {
	ctx := c.Request.Context()
	var req addRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errInvalidRequestBody)
		return
	}

	fmt.Printf("req: %+v\n", req)

	input, err := req.toInput()

	// print input
	

	if err != nil {
		response.Error(c, err)
		return
	}

	if err := h.trackUC.Create(ctx, input); err != nil {
		h.l.Error(ctx, "track.handler.add.trackUC.Create: %s", err)
		response.ErrorWithMap(c, err, errMap)
		return
	}

	response.OK(c, nil)
}

func (h handler) detail(c *gin.Context) {
	ctx := c.Request.Context()

	rawID := c.Param("id")

	if strings.TrimSpace(rawID) == "" {
		response.Error(c, errIDIsRequired)
		return
	}

	// id, err := strconv.Atoi(rawID)
	id, err := strconv.Atoi(rawID)

	track, err := h.trackUC.Detail(ctx, id)

	if err != nil {
		h.l.Error(ctx, "track.handler.detail.trackUC.detail")
		return
	}

	response.OK(c, newDetailResponse(track))
 }