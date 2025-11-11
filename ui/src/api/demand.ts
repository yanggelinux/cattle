import request from '@/utils/request'

const BASE_URL = '/api/v1/demand'

const DemandAPI = {
  /** 获取数据 */
  getList(queryParams?: DemandQuery) {
    return request({
      url: `${BASE_URL}`,
      method: 'get',
      params: queryParams,
    })
  },
  getDetail(queryParams?: DemandDetailQuery) {
    return request({
      url: `${BASE_URL}/detail`,
      method: 'get',
      params: queryParams,
    })
  },

  /** 添加 */
  create(data: DemandForm) {
    return request({
      url: `${BASE_URL}`,
      method: 'post',
      data: data,
    })
  },

  /**
   * 更新
   *
   * @param id ID
   * @param data 表单数据
   */
  update(data: DemandForm) {
    return request({
      url: `${BASE_URL}`,
      method: 'put',
      data: data,
    })
  },

  /**
   * 删除
   *
   * @param id ID
   */
  delete(id: number) {
    return request({
      url: `${BASE_URL}/${id}`,
      method: 'delete',
    })
  },

  approve(data: DemandApprovalForm) {
    return request({
      url: `${BASE_URL}/approve`,
      method: 'post',
      data: data,
    })
  },
  evaluate(data: EvaluateDemandForm) {
    return request({
      url: `${BASE_URL}/evaluate`,
      method: 'post',
      data: data,
    })
  },
}

export default DemandAPI

/** 查询参数 */
export interface DemandQuery {
  /** 搜索关键字 */
  page?: number
  pageSize?: number
  status?: number
  name?: string
}

export interface DemandDetailQuery {
  /** 搜索关键字 */
  id?: number
  name?: string
}

export interface DemandData {
  total: number
  retList: DemandResult[]
}

/** 查询结果 */
export interface DemandResult {
  /** ID */
  id: number
  /** 名称 */
  name: string
  demandType: number
  orderNo: string
  biz: string
  owner: string
  description: string
  opinion: string
  reviewProcess: ProcessNode[]
  curReviewNode: ProcessNode
  activeIndex: number
  evaluation: any
  isEvaluate: number
  evaluationRes: string
  evaluationReason: string
  hasReview: number
  status: number
  /** 修改时间 */
  updatedTime: string
  /** 创建时间 */
  createdTime: string
}

/** 表单对象 */
export interface DemandForm {
  /** ID */
  id?: number
  /** 名称 */
  name?: string
  demandType?: number
  orderNo?: string
  biz?: string
  owner?: string
  description?: string
  opinion?: string
  status?: number
  action?: string
}

export interface DemandApprovalForm {
  /** ID */
  id?: number
  /** 名称 */
  action?: string
  opinion?: string
  approver?: string
  approverName?: string
}

export interface EvaluateDemandForm {
  id?: number
  opsEvaluation?: string
  opsReason?: string
  resEvaluation?: string
  resReason?: string
  netEvaluation?: string
  netReason?: string
}

export interface ProcessNode {
  name: string
  approverGroup: string
  approver: string
  approverName: string
  role: string
  opt: string
  status: number
}
