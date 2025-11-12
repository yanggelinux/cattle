package router

import (
	"github.com/yanggelinux/cattle/global"
	apiV1 "github.com/yanggelinux/cattle/internal/api/v1"
	"github.com/yanggelinux/cattle/internal/middleware"
	openapiV1 "github.com/yanggelinux/cattle/internal/openapi/v1"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

var auditIgnoreURI = []string{}

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(middleware.Recovery())
	gin.SetMode(global.ServerSetting.RunMode)

	// api v1
	test := apiV1.NewTest()
	user := apiV1.NewUser()
	role := apiV1.NewRole()
	auth := apiV1.NewAuth()
	permission := apiV1.NewPermission()
	archGroup := apiV1.NewArchGroup()
	archGraph := apiV1.NewArchGraph()
	processOrder := apiV1.NewProcessOrder()
	demand := apiV1.NewDemand()
	dashboard := apiV1.NewDashboard()
	process := apiV1.NewProcess()
	order := apiV1.NewOrder()
	orderGroup := apiV1.NewOrderGroup()
	orderField := apiV1.NewOrderField()
	team := apiV1.NewTeam()
	// 配置路由
	apiPublic := r.Group("/cattle/api/v1")
	apiGroup := r.Group("/cattle/api/v1")

	// openapi
	openAuth := openapiV1.NewAuth()
	openArchGraph := openapiV1.NewArchGraph()

	openapiGroup := r.Group("/cattle/openapi/v1")
	openapiPublicGroup := r.Group("/cattle/openapi/v1")

	// 一些中间件功能,首先jwt。然后header meta data 最后audit
	// JWT认证功能
	apiGroup.Use(middleware.JWT())

	// header 源数据检查功能
	apiPublic.Use(middleware.RequestMetadata())
	apiGroup.Use(middleware.RequestMetadata())
	openapiPublicGroup.Use(middleware.RequestMetadata())
	openapiGroup.Use(middleware.RequestMetadata())
	openapiGroup.Use(middleware.JWT())

	// api
	{

		// test
		apiPublic.GET("/test", test.DoTest)
		apiPublic.GET("/test/status", test.DoTestStatus)

		apiPublic.POST("/auth/login", auth.Login)

		apiGroup.GET("/auth/user-perm", auth.GetUserPermList)
		apiGroup.GET("/user", user.GetList)
		apiGroup.POST("/user", user.Create)
		apiGroup.PUT("/user", user.Update)
		apiGroup.DELETE("/user/:id", user.Delete)

		apiGroup.GET("/role", role.GetList)
		apiGroup.POST("/role", role.Create)
		apiGroup.PUT("/role", role.Update)
		apiGroup.DELETE("/role/:id", role.Delete)

		apiGroup.GET("/team", team.GetList)
		apiGroup.POST("/team", team.Create)
		apiGroup.PUT("/team", team.Update)
		apiGroup.DELETE("/team/:id", team.Delete)

		apiGroup.GET("/permission", permission.GetList)
		apiGroup.POST("/permission", permission.Create)
		apiGroup.PUT("/permission", permission.Update)
		apiGroup.DELETE("/permission/:id", permission.Delete)
		apiGroup.GET("/permission/role-perm", permission.GetRolePermList)
		apiGroup.PUT("/permission/role-perm", permission.UpdateRolePerm)

		apiGroup.GET("/process", process.GetList)
		apiGroup.GET("/process/:id", process.GetDetail)
		apiGroup.POST("/process", process.Create)
		apiGroup.PUT("/process", process.Update)
		apiGroup.DELETE("/process/:id", process.Delete)
		apiGroup.POST("/process/copy/:id", process.Copy)

		apiGroup.GET("/arch-group", archGroup.GetList)
		apiGroup.POST("/arch-group", archGroup.Create)
		apiGroup.PUT("/arch-group", archGroup.Update)
		apiGroup.DELETE("/arch-group/:id", archGroup.Delete)

		apiGroup.GET("/arch-graph", archGraph.GetList)
		apiGroup.POST("/arch-graph", archGraph.Create)
		apiGroup.PUT("/arch-graph", archGraph.Update)
		apiGroup.PUT("/arch-graph/save", archGraph.Save)
		apiGroup.DELETE("/arch-graph/:id", archGraph.Delete)
		apiGroup.POST("/arch-graph/copy/:id", archGraph.Copy)
		apiGroup.GET("/arch-graph/:id", archGraph.GetDetail)
		apiGroup.GET("/arch-graph/record", archGraph.GetRecordList)
		apiGroup.PUT("/arch-graph/select", archGraph.SelectRecord)
		apiGroup.GET("/arch-graph/enabled/:id", archGraph.GetEnabledRecord)

		apiGroup.GET("/arch-graph/review", archGraph.GetReviewList)
		apiGroup.POST("/arch-graph/review", archGraph.CreateReview)
		apiGroup.DELETE("/arch-graph/review/:id", archGraph.DeleteReview)

		apiGroup.GET("/process-order", processOrder.GetList)
		apiGroup.GET("/process-order/approval", processOrder.GetApprovalList)
		apiGroup.POST("/process-order", processOrder.Create)
		apiGroup.PUT("/process-order", processOrder.Update)
		apiGroup.DELETE("/process-order/:id", processOrder.Delete)
		apiGroup.GET("/process-order/:id", processOrder.GetDetail)
		apiGroup.POST("/process-order/apply", processOrder.Apply)
		apiGroup.POST("/process-order/re-apply", processOrder.ReApply)
		apiGroup.POST("/process-order/approve", processOrder.Approve)
		apiGroup.GET("/process-order/unapproved", processOrder.GetUnapprovedList)
		apiGroup.POST("/process-order/assign-approver", processOrder.AssignApprover)

		apiGroup.GET("/demand", demand.GetList)
		apiGroup.GET("/demand/detail", demand.GetDetail)
		apiGroup.POST("/demand", demand.Create)
		apiGroup.PUT("/demand", demand.Update)
		apiGroup.DELETE("/demand/:id", demand.Delete)
		apiGroup.POST("/demand/approve", demand.Approve)
		apiGroup.POST("/demand/evaluate", demand.Evaluate)

		apiGroup.GET("/dashboard/graph-info", dashboard.GetGraphInfo)
		apiGroup.GET("/dashboard/order-info", dashboard.GetOrderInfo)
		apiGroup.GET("/dashboard/demand-info", dashboard.GetDemandInfo)

		//order
		apiGroup.GET("/order", order.GetList)
		apiGroup.GET("/order/node", order.GetNodeTypeList)
		apiGroup.GET("/order/:id", order.GetDetail)
		apiGroup.POST("/order", order.Create)
		apiGroup.PUT("/order", order.Update)
		apiGroup.DELETE("/order/:id", order.Delete)
		apiGroup.POST("/order/copy/:id", order.Copy)

		apiGroup.GET("/order-group", orderGroup.GetList)
		apiGroup.POST("/order-group", orderGroup.Create)
		apiGroup.PUT("/order-group", orderGroup.Update)
		apiGroup.DELETE("/order-group/:id", orderGroup.Delete)

		apiGroup.GET("/order-field", orderField.GetList)
		apiGroup.POST("/order-field", orderField.Create)
		apiGroup.PUT("/order-field", orderField.Update)
		apiGroup.DELETE("/order-field/:id", orderField.Delete)

	}
	// open api
	{
		openapiPublicGroup.GET("/auth/token", openAuth.GetToken)
		openapiPublicGroup.GET("/process-graph/info", openArchGraph.GetGraphData)

	}
	return r
}
