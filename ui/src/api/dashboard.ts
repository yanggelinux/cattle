import request from '@/utils/request'

const BASE_URL = '/api/v1/dashboard'

const DashboardAPI = {
  /** 获取数据 */
  getGraphInfo(queryParams?: DashboardGraphQuery) {
    return request({
      url: `${BASE_URL}/graph-info`,
      method: 'get',
      params: queryParams,
    })
  },
  getOrderInfo(queryParams?: DashboardOrderQuery) {
    return request({
      url: `${BASE_URL}/order-info`,
      method: 'get',
      params: queryParams,
    })
  },
  getDemandInfo(queryParams?: DashboardDemandQuery) {
    return request({
      url: `${BASE_URL}/demand-info`,
      method: 'get',
      params: queryParams,
    })
  },
}

export default DashboardAPI

export interface DashboardGraphQuery {}
export interface DashboardOrderQuery {}
export interface DashboardDemandQuery {}

export interface DashboardGraphResult {
  totalCount: number
  unapprovedCount: number
  approvingCount: number
  successCount: number
  failedCount: number
}
export interface DashboardOrderResult {
  totalCount: number
  unapprovedCount: number
  approvingCount: number
  successCount: number
  failedCount: number
  graphApplyDist: Dist[]
  graphChangeDist: Dist[]
  resApplyChangeDist: Dist[]
}
export interface DashboardDemandResult {
  totalCount: number
  unapprovedCount: number
  approvingCount: number
  successCount: number
  failedCount: number
}

export interface Dist {
  count: number
  dt: string
}
