package demand

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/biz"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	demandrepo "github.com/yanggelinux/cattle/internal/repository/demand"
	rolerepo "github.com/yanggelinux/cattle/internal/repository/role"
	teamrepo "github.com/yanggelinux/cattle/internal/repository/team"
	userrepo "github.com/yanggelinux/cattle/internal/repository/user"
	"github.com/yanggelinux/cattle/internal/store/model"
	"github.com/yanggelinux/cattle/pkg/log"
	"github.com/yanggelinux/cattle/pkg/util"
	"go.uber.org/zap"
	"gorm.io/datatypes"
	"strings"
	"time"
)

type demandService struct {
	db                 *model.MDB
	demandRepo         demandrepo.DemandRepo
	demandApprovalRepo demandrepo.DemandApprovalRepo
	userRepo           userrepo.UserRepo
	roleRepo           rolerepo.RoleRepo
	teamRepo           teamrepo.TeamRepo
}

func NewDemandService() DemandService {
	return &demandService{
		db:                 model.GetDB(),
		demandRepo:         demandrepo.NewDemandRepo(),
		demandApprovalRepo: demandrepo.NewDemandApprovalRepo(),
		userRepo:           userrepo.NewUserRepo(),
		roleRepo:           rolerepo.NewRoleRepo(),
		teamRepo:           teamrepo.NewTeamRepo(),
	}
}

func (s *demandService) GetList(ctx context.Context, req *request.GetDemandReq) (*result.DemandResult, error) {

	filter := &demandrepo.DemandFilter{
		Status:   req.Status,
		Name:     req.Name,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	records, total, err := s.demandRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	userName := util.GetUserName(ctx)

	roleNames, err := s.getRolesByUser(ctx)
	if err != nil {
		return nil, errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
	}
	teamNameInfo, _ := s.genTeamInfo(ctx)

	retList := make([]*result.DemandRet, 0, len(records))
	for _, record := range records {

		var reviewProcess []*result.PorcessNode
		err := json.Unmarshal(record.ReviewProcess, &reviewProcess)
		if err != nil {
			reviewProcess = ReviewProcess
		}
		activeIndex := s.genActiveIndex(reviewProcess)
		curNodeIndex := activeIndex
		if activeIndex == int8(len(reviewProcess)) {
			curNodeIndex = activeIndex - 1
		}
		var curReviewNode *result.PorcessNode
		if len(reviewProcess) > 0 {
			curReviewNode = reviewProcess[curNodeIndex]
		}

		hasReview := s.isReview(userName, roleNames, curReviewNode, teamNameInfo)

		evaluation := &result.Evaluation{}
		var evaluationRes string
		var evaluationReason string
		err = json.Unmarshal(record.Evaluation, &evaluation)
		if err == nil {
			evaluationRes, evaluationReason = s.dealEvaluationInfo(evaluation)
		}
		ret := &result.DemandRet{}
		ret.ID = record.ID
		ret.Name = record.Name
		ret.DemandType = record.DemandType
		ret.OrderNo = record.OrderNo
		ret.Biz = record.Biz
		ret.Owner = record.Owner
		ret.Description = record.Description
		ret.Opinion = record.Opinion
		ret.ReviewProcess = reviewProcess
		ret.CurReviewNode = curReviewNode
		ret.Status = record.Status
		ret.ActiveIndex = activeIndex
		ret.IsEvaluate = record.IsEvaluate
		ret.EvaluationRes = evaluationRes
		ret.EvaluationReason = evaluationReason
		ret.HasReview = hasReview
		ret.Evaluation = evaluation
		ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.DemandResult{}
	resultData.Total = total
	resultData.RetList = retList
	return resultData, nil
}

func (s *demandService) dealEvaluationInfo(evaluation *result.Evaluation) (string, string) {
	evaluationRes := "满意"
	evaluationReason := make([]string, 0, 0)
	if evaluation.OpsEvaluation == "不满意" {
		evaluationRes = "不满意"
		evaluationReason = append(evaluationReason, fmt.Sprintf("应用运维组:%s", evaluation.OpsReason))
	}
	if evaluation.ResEvaluation == "不满意" {
		evaluationRes = "不满意"
		evaluationReason = append(evaluationReason, fmt.Sprintf("资源管理组:%s", evaluation.ResReason))
	}
	if evaluation.NetEvaluation == "不满意" {
		evaluationRes = "不满意"
		evaluationReason = append(evaluationReason, fmt.Sprintf("安全支撑组:%s", evaluation.NetReason))
	}
	return evaluationRes, strings.Join(evaluationReason, ",")
}

func (s *demandService) genActiveIndex(process []*result.PorcessNode) int8 {
	i := 0
	for _, v := range process {
		if v.Status == 0 {
			break
		}
		i += 1
	}
	return int8(i)
}

func (s *demandService) isReview(approver string, roleNames []string, curNode *result.PorcessNode, teamNameInfo map[string]*biz.TeamInfo) int8 {
	var leaders, directors []string
	deptName := curNode.DeptName
	teamInfo, ok := teamNameInfo[deptName]
	if ok {
		leaders = strings.Split(teamInfo.Leader, ",")
		directors = strings.Split(teamInfo.Director, ",")
	}
	leaderRoles := []string{
		"leader",
		"teamLeader",
	}
	directorRoles := []string{
		"director",
	}
	if lo.Contains(leaderRoles, curNode.Role) {
		if len(leaders) > 0 {
			if lo.Contains(leaders, approver) {
				return 1
			} else {
				return 0
			}
		}
	}
	if lo.Contains(directorRoles, curNode.Role) {
		if len(directors) > 0 {
			if lo.Contains(directors, approver) {
				return 1
			} else {
				return 0
			}
		}
	}
	if curNode.Status == 0 && lo.Contains(roleNames, curNode.Role) {
		return 1
	}
	return 0
}

func (s *demandService) genTeamInfo(ctx context.Context) (map[string]*biz.TeamInfo, error) {
	teamNameInfo := make(map[string]*biz.TeamInfo)
	teams, err := s.teamRepo.GetAll(ctx)
	if err != nil {
		return teamNameInfo, errors.Wrap(err, "get team info error")
	}
	for _, team := range teams {
		teamNameInfo[team.Name] = &biz.TeamInfo{
			Name:     team.Name,
			Leader:   team.Leader,
			Director: team.Director,
		}
	}
	return teamNameInfo, nil
}

func (s *demandService) getRolesByUser(ctx context.Context) ([]string, error) {
	roleNames := make([]string, 0, 0)
	userName := util.GetUserName(ctx)
	user, err := s.userRepo.GetByName(ctx, userName)
	if err != nil {
		return roleNames, errors.Wrap(err, "获取角色信息时，查询用户信息失败")
	}
	userID := user.ID
	roles, err := s.roleRepo.GetRoles(ctx, userID)
	if err != nil {
		return roleNames, errors.Wrap(err, "获取角色信息时，查询角色信息失败")
	}
	for _, role := range roles {
		roleNames = append(roleNames, role.RoleName)
	}
	return roleNames, nil
}

func (s *demandService) GetDetail(ctx context.Context, req *request.GetDemandDetailReq) (*result.DemandRet, error) {

	var (
		err    error
		record *model.Demand
	)

	if req.ID != nil && *req.ID != 0 {
		id := *req.ID
		record, err = s.demandRepo.GetByID(ctx, id)
		if err != nil {
			err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
			return nil, err
		}
	} else if req.Name != nil && len(*req.Name) > 0 {
		name := *req.Name
		record, err = s.demandRepo.GetByName(ctx, name)
		if err != nil {
			err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
			return nil, err
		}
	} else {
		return nil, errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
	}

	var reviewProcess []*result.PorcessNode
	err = json.Unmarshal(record.ReviewProcess, &reviewProcess)
	if err != nil {
		reviewProcess = ReviewProcess
	}
	activeIndex := s.genActiveIndex(reviewProcess)
	curNodeIndex := activeIndex
	if activeIndex == int8(len(reviewProcess)) {
		curNodeIndex = activeIndex - 1
	}
	var curReviewNode *result.PorcessNode
	if len(reviewProcess) > 0 {
		curReviewNode = reviewProcess[curNodeIndex]
	}

	resultData := &result.DemandRet{}
	resultData.ID = record.ID
	resultData.Name = record.Name
	resultData.DemandType = record.DemandType
	resultData.OrderNo = record.OrderNo
	resultData.Biz = record.Biz
	resultData.Owner = record.Owner
	resultData.Description = record.Description
	resultData.Opinion = s.getApprovalOpinion(ctx, record.ID)
	resultData.ActiveIndex = activeIndex
	resultData.ReviewProcess = reviewProcess
	resultData.CurReviewNode = curReviewNode
	resultData.Status = record.Status
	resultData.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
	resultData.CreatedTime = util.FormatTimeToString(record.CreatedTime)
	return resultData, nil
}

func (s *demandService) Create(ctx context.Context, req *request.CreateDemandReq) error {
	demand := model.NewDemand()
	demand.Name = *req.Name
	demand.Biz = *req.Biz
	demand.DemandType = *req.DemandType
	demand.Owner = *req.Owner
	demand.Description = *req.Description
	if req.Status != nil {
		demand.Status = *req.Status
	}
	userName := util.GetUserName(ctx)
	deptName := util.GetDeptName(ctx)
	//工单必须长度大于10
	if req.OrderNo != nil && len(*req.OrderNo) > 10 {
		demand.OrderNo = *req.OrderNo
		ReviewProcess[0].ApproverName = userName
		ReviewProcess[0].Approver = userName
		ReviewProcess[1].Status = 1
	} else {
		ReviewProcess[0].ApproverName = userName
		ReviewProcess[0].Approver = userName
	}
	for _, node := range ReviewProcess {
		node.DeptName = deptName
	}
	// 通过判断oa编号 来确定是否走组长审批
	reviewProcess, err := json.Marshal(ReviewProcess)
	if err != nil {
		reviewProcess = datatypes.JSON("")
	}
	demand.ReviewProcess = reviewProcess

	checked := s.demandRepo.CheckDuplicateEntry(ctx, *req.Name)
	if checked {
		return errors.WithCodeError(ce.ErrorDBDuplicateEntry.Code(), errors.New("Duplicate Entry error"))
	}
	err = s.demandRepo.Create(ctx, demand)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return err
}

func (s *demandService) getApprovalOpinion(ctx context.Context, id int64) string {

	records, err := s.demandApprovalRepo.GetByDemand(ctx, id)
	if err != nil {
		return ""
	}
	opinions := make([]string, 0, len(records))
	for _, record := range records {
		approver := record.Approver
		opinion := record.Opinion
		action := record.Action
		createTime := util.FormatTimeToString(record.CreatedTime)
		txt := fmt.Sprintf("审批时间：%s，审批人：%s，审批操作：%s，审批意见：%s", createTime, approver, action, opinion)
		opinions = append(opinions, txt)
	}
	return strings.Join(opinions, "@")
}

func (s *demandService) Update(ctx context.Context, req *request.UpdateDemandReq) error {
	id := *req.ID
	demand, err := s.demandRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	updateField := make(map[string]interface{})
	if req.Name != nil {
		updateField[demand.NameField()] = *req.Name
	}
	if req.OrderNo != nil {
		updateField[demand.OrderNoField()] = *req.OrderNo
	}
	if req.Biz != nil {
		updateField[demand.BizField()] = *req.Biz
	}
	if req.Owner != nil {
		updateField[demand.OwnerField()] = *req.Owner
	}
	if req.Description != nil {
		updateField[demand.DescriptionField()] = *req.Description
	}
	if req.Status != nil {
		updateField[demand.StatusField()] = *req.Status
	}

	// 评审失败重新提交工单的情况
	if req.Status != nil && *req.Status == 1 && demand.Status == 3 {
		userName := util.GetUserName(ctx)
		reviewProcess := demand.ReviewProcess
		var reviewProcessData []*result.PorcessNode
		err = json.Unmarshal(reviewProcess, &reviewProcessData)
		if err != nil {
			return errors.WithCodeError(ce.Error.Code(), err)
		}
		reviewProcessData[0].Status = 1
		reviewProcessData[0].Approver = userName
		reviewProcessData[0].ApproverName = userName
		if len(demand.OrderNo) > 10 || (req.OrderNo != nil && len(*req.OrderNo) > 10) {
			reviewProcessData[1].Status = 1
		}
		reviewProcessJSON, err := json.Marshal(reviewProcessData)
		if err != nil {
			return errors.WithCodeError(ce.Error.Code(), err)
		}
		updateField[demand.ReviewProcessField()] = reviewProcessJSON
	}
	err = s.demandRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *demandService) Delete(ctx context.Context, id int64) error {
	// 删除组前先检查是否有
	_, err := s.demandRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	err = s.demandRepo.DeleteByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *demandService) Approve(ctx context.Context, req *request.ApproveDemandReq) error {
	id := *req.ID
	action := *req.Action
	// 审批人
	approver := *req.Approver
	approverName := *req.ApproverName
	demand, err := s.demandRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}

	demandApproval := &model.DemandApproval{}
	demandApproval.Approver = approverName
	demandApproval.DemandID = id
	demandApproval.Opinion = *req.Opinion
	// 图状态 0 未审批 1 审批中 2 审批成功 3审批失败 4 审批成功失效
	status := int8(1)
	actionText := "通过"
	//审批成功
	if action == "success" {
		status = 2
		// 审批不通过
	} else if action == "notpass" {
		status = 3
		actionText = "不通过"
		// 下一节点
	} else {
		status = 1
	}
	demandApproval.Status = status
	demandApproval.Action = actionText
	reviewProcess := demand.ReviewProcess
	var reviewProcessData []*result.PorcessNode
	err = json.Unmarshal(reviewProcess, &reviewProcessData)
	if err != nil {
		return errors.WithCodeError(ce.Error.Code(), err)
	}
	activeIndex := s.genActiveIndex(reviewProcessData)
	// 审批成功
	bizGroupInfo := approvalMap[demand.Biz]
	for i, v := range reviewProcessData {
		// 在这里确认个组长的审批权限
		if v.Name == "业务责人评审" {
			v.Approver = bizGroupInfo.Approver
			v.ApproverName = bizGroupInfo.ApproverName
		}

		if int8(i) == activeIndex {
			v.Approver = approver
			v.ApproverName = approverName
		}
		if status == 2 {
			v.Status = 1
		} else if status == 3 {
			//if i == 0 {
			//	continue
			//}
			v.Status = 0
		} else {
			v.Status = 1
			if int8(i) == activeIndex {
				break
			}
		}
	}
	reviewProcessJSON, err := json.Marshal(reviewProcessData)
	if err != nil {
		return errors.WithCodeError(ce.Error.Code(), err)
	}
	// 事务相关
	tx := s.db.WithContext(ctx).Begin()
	defer model.RecoverRollback(tx)

	demandTx := s.demandRepo.WithTx(tx)
	demandApprovalTx := s.demandApprovalRepo.WithTx(tx)

	// 创建一条审批记录
	err = demandApprovalTx.Create(ctx, demandApproval)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	// 更新订单，状态，节点信息
	updateField := make(map[string]interface{})
	updateField[demand.StatusField()] = status
	updateField[demand.ReviewProcessField()] = reviewProcessJSON
	err = demandTx.Update(ctx, id, updateField)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	tx.Commit()
	return nil
}

func (s *demandService) Evaluate(ctx context.Context, req *request.EvaluateDemandReq) error {
	id := *req.ID
	demand, err := s.demandRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	evaluationJson, err := json.Marshal(req)
	if err != nil {
		return errors.WithCodeError(ce.Error.Code(), err)
	}
	updateField := make(map[string]interface{})
	updateField[demand.EvaluationField()] = evaluationJson
	updateField[demand.IsEvaluateField()] = 1
	err = s.demandRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *demandService) CheckEvaluation(ctx context.Context) error {
	status := int8(2)
	demands, err := s.demandRepo.GetByStatus(ctx, status)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
	}
	for _, demand := range demands {
		id := demand.ID
		updateTime := demand.UpdatedTime
		nowTime := time.Now()
		diff := nowTime.Sub(updateTime)
		//三天时间
		if diff <= 72*time.Hour {
			continue
		}
		evaluationData := &result.Evaluation{}
		evaluationData.OpsEvaluation = "满意"
		evaluationData.OpsReason = ""
		evaluationData.ResEvaluation = "满意"
		evaluationData.ResReason = ""
		evaluationData.NetEvaluation = "满意"
		evaluationData.NetReason = ""
		evaluationJson, err := json.Marshal(evaluationData)
		if err != nil {
			return errors.WithCodeError(ce.Error.Code(), err)
		}
		updateField := make(map[string]interface{})
		updateField[demand.EvaluationField()] = evaluationJson
		updateField[demand.IsEvaluateField()] = 1
		err = s.demandRepo.Update(ctx, id, updateField)
		if err != nil {
			log.Logger.Error("check update evaluation error", zap.Error(err))
			continue
		}
	}
	return nil
}
