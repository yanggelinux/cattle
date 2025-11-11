package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/service/test"
)

type Test struct {
}

func NewTest() *Test {
	return &Test{}
}

func (a *Test) DoTest(c *gin.Context) {

	data := make(map[string]interface{})

	app.Response(c, data, nil)
}

func (a *Test) DoTestStatus(c *gin.Context) {
	data := make(map[string]interface{})

	s := test.NewTestService()
	data, err := s.DoTestStatus(c)
	app.Response(c, data, err)

}
