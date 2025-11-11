import request from '@/utils/request'

const BASE_URL = '/api/v1/process'

const ProcessAPI = {
  /** 获取数据 */
  getList(queryParams?: ProcessQuery) {
    return request({
      url: `${BASE_URL}`,
      method: 'get',
      params: queryParams,
    })
  },
  getDetail(id: number) {
    return request({
      url: `${BASE_URL}/${id}`,
      method: 'get',
    })
  },
  /** 添加 */
  create(data: ProcessForm) {
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
  update(data: ProcessForm) {
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
  copy(id: number) {
    return request({
      url: `${BASE_URL}/copy/${id}`,
      method: 'post',
    })
  },
}

export default ProcessAPI

/** 查询参数 */
export interface ProcessQuery {
  /** 搜索关键字 */
  page?: number
  pageSize?: number
  status?: number
  name?: string
}

export interface ProcessData {
  total: number
  retList: ProcessResult[]
}

/** 查询结果 */
export interface ProcessResult {
  /** ID */
  id: number
  /** 名称 */
  name: string
  procInfo: ProcessNode[]
  nodeData: any[]
  edgeData: any[]
  status: number
  /** 修改时间 */
  updatedTime: string
  /** 创建时间 */
  createdTime: string
}

/** 表单对象 */
export interface ProcessForm {
  /** ID */
  id?: number
  /** 名称 */
  name?: string
  nodeData?: any[]
  edgeData?: any[]
  status?: number
}

export interface ProcessNodeForm {
  /** 名称 */
  name: string
  role: string
  roleName: string
  approvalType: number
  approvalEdit: number
}

export interface ProcessOptResult {
  id: number
}

export interface ProcessNode {
  name: string
  type: string
  deptName: string
  approvalType: number
  approvalEdit: number
  approver: string
  status: number
  approvalInfo: ApprovalInfo[]
}

export interface ApprovalInfo {
  approver: string
  approverName: string
  role: string
  roleName: string
  status: number
}
