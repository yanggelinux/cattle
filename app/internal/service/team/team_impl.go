package team

import (
	"context"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	teamrepo "github.com/yanggelinux/cattle/internal/repository/team"
	"github.com/yanggelinux/cattle/internal/store/model"
	"github.com/yanggelinux/cattle/pkg/util"
	"strings"
)

type teamService struct {
	teamRepo teamrepo.TeamRepo
}

func NewTeamService() TeamService {
	return &teamService{teamRepo: teamrepo.NewTeamRepo()}
}

func (s *teamService) GetList(ctx context.Context, req *request.GetTeamReq) (*result.TeamResult, error) {

	//进行一个copy
	filter := &teamrepo.TeamFilter{
		Name:     req.Name,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	records, total, err := s.teamRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	retList := make([]*result.TeamRet, 0, len(records))
	for _, record := range records {
		ret := &result.TeamRet{}
		ret.ID = record.ID
		ret.Name = record.Name
		ret.Leader = record.Leader
		ret.Director = record.Director
		ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.TeamResult{}
	resultData.Total = total
	resultData.RetList = retList
	return resultData, nil
}

func (s *teamService) Create(ctx context.Context, req *request.CreateTeamReq) error {
	team := model.NewTeam()
	team.Name = *req.Name
	team.Leader = strings.Replace(*req.Leader, "，", ",", -1)
	team.Director = strings.Replace(*req.Director, "，", ",", -1)
	checked := s.teamRepo.CheckDuplicateEntry(ctx, *req.Name)
	if checked {
		return errors.WithCodeError(ce.ErrorDBDuplicateEntry.Code(), errors.New("Duplicate Entry error"))
	}
	err := s.teamRepo.Create(ctx, team)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return err
}

func (s *teamService) Update(ctx context.Context, req *request.UpdateTeamReq) error {
	id := *req.ID
	team, err := s.teamRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	updateField := make(map[string]interface{})

	if req.Name != nil {
		updateField[team.NameField()] = *req.Name
	}
	if req.Leader != nil {
		updateField[team.LeaderField()] = strings.Replace(*req.Leader, "，", ",", -1)
	}
	if req.Director != nil {
		updateField[team.DirectorField()] = strings.Replace(*req.Director, "，", ",", -1)
	}
	err = s.teamRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *teamService) Delete(ctx context.Context, id int64) error {
	_, err := s.teamRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	err = s.teamRepo.DeleteByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}
