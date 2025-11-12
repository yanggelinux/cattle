package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	"github.com/yanggelinux/cattle/internal/service/archgraph"
)

type ArchGraph struct {
}

func NewArchGraph() *ArchGraph {
	return &ArchGraph{}
}

func (a *ArchGraph) GetGraphData(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetArchGraphDataReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgraph.NewArchGraphService()
	resultData, err := svc.GetDataByLabel(c, req)
	app.Response(c, resultData, err)
}
