package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/service/dashboard"
)

type Dashboard struct {
}

func NewDashboard() *Dashboard {
	return &Dashboard{}
}

func (a *Dashboard) GetGraphInfo(c *gin.Context) {
	svc := dashboard.NewDashboardService()
	resultData, err := svc.GetGraphInfo(c)
	app.Response(c, resultData, err)
}
func (a *Dashboard) GetOrderInfo(c *gin.Context) {
	svc := dashboard.NewDashboardService()
	resultData, err := svc.GetOrderInfo(c)
	app.Response(c, resultData, err)
}
func (a *Dashboard) GetDemandInfo(c *gin.Context) {
	svc := dashboard.NewDashboardService()
	resultData, err := svc.GetDemandInfo(c)
	app.Response(c, resultData, err)
}
