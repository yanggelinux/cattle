package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	"github.com/yanggelinux/cattle/internal/service/archgraph"
	"strconv"
)

type ArchGraph struct {
}

func NewArchGraph() *ArchGraph {
	return &ArchGraph{}
}

func (a *ArchGraph) GetList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetArchGraphReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgraph.NewArchGraphService()
	resultData, err := svc.GetList(c, req)
	app.Response(c, resultData, err)
}

func (a *ArchGraph) Create(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.CreateArchGraphReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgraph.NewArchGraphService()
	resultData, err := svc.Create(c, req)
	app.Response(c, resultData, err)
}
func (a *ArchGraph) Update(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.UpdateArchGraphReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgraph.NewArchGraphService()
	err = svc.Update(c, req)
	app.Response(c, data, err)
}

func (a *ArchGraph) Save(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.UpdateArchGraphReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgraph.NewArchGraphService()
	err = svc.Save(c, req)
	app.Response(c, data, err)
}

func (a *ArchGraph) Delete(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgraph.NewArchGraphService()
	err = svc.Delete(c, id)
	app.Response(c, data, err)
}

func (a *ArchGraph) Copy(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgraph.NewArchGraphService()
	err = svc.Copy(c, id)
	app.Response(c, data, err)
}

func (a *ArchGraph) GetDetail(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgraph.NewArchGraphService()
	resultData, err := svc.GetByID(c, id)
	app.Response(c, resultData, err)
}
func (a *ArchGraph) GetRecordList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetArchGraphRecordReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgraph.NewArchGraphService()
	resultData, err := svc.GetRecordList(c, req)
	app.Response(c, resultData, err)
}

func (a *ArchGraph) SelectRecord(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.SelectArchGraphReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgraph.NewArchGraphService()
	err = svc.SelectRecord(c, req)
	app.Response(c, data, err)
}

func (a *ArchGraph) GetEnabledRecord(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgraph.NewArchGraphService()
	resultData, err := svc.GetEnabledRecord(c, id)
	app.Response(c, resultData, err)
}

// review
func (a *ArchGraph) GetReviewList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetArchGraphReviewReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgraph.NewArchGraphService()
	resultData, err := svc.GetReviewList(c, req)
	app.Response(c, resultData, err)
}

func (a *ArchGraph) CreateReview(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.CreateArchGraphReviewReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgraph.NewArchGraphService()
	err = svc.CreateReview(c, req)
	app.Response(c, data, err)
}

func (a *ArchGraph) DeleteReview(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgraph.NewArchGraphService()
	err = svc.DeleteReview(c, id)
	app.Response(c, data, err)
}
