package test

import (
	"context"
)

type ITestService interface {
	DoTestStatus(context.Context) (map[string]interface{}, error)
}
