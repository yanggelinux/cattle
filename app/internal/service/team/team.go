package team

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type TeamService interface {
	GetList(context.Context, *request.GetTeamReq) (*result.TeamResult, error)
	Create(context.Context, *request.CreateTeamReq) error
	Update(context.Context, *request.UpdateTeamReq) error
	Delete(context.Context, int64) error
}
