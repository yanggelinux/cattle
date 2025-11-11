import request from '@/utils/request'

const BASE_URL = '/api/v1/arch-graph'

const ArchGraphAPI = {
  /** 获取数据 */
  getList(queryParams?: ArchGraphQuery) {
    return request({
      url: `${BASE_URL}`,
      method: 'get',
      params: queryParams,
    })
  },

  getRecordList(queryParams?: ArchGraphRecordQuery) {
    return request({
      url: `${BASE_URL}/record`,
      method: 'get',
      params: queryParams,
    })
  },

  /** 添加 */
  create(data: ArchGraphForm) {
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
  update(data: ArchGraphForm) {
    return request({
      url: `${BASE_URL}`,
      method: 'put',
      data: data,
    })
  },

  selectRecord(data: ArchGraphSelectForm) {
    return request({
      url: `${BASE_URL}/select`,
      method: 'put',
      data: data,
    })
  },

  save(data: ArchGraphForm) {
    return request({
      url: `${BASE_URL}/save`,
      method: 'put',
      data: data,
    })
  },

  /**
   * 根据id获取详情
   *
   * @param id ID
   */
  getDetail(id: number) {
    return request({
      url: `${BASE_URL}/${id}`,
      method: 'get',
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

  sync(id: number) {
    return request({
      url: `${BASE_URL}/sync/${id}`,
      method: 'post',
    })
  },
  getEnabled(id: number) {
    return request({
      url: `${BASE_URL}/enabled/${id}`,
      method: 'get',
    })
  },

  getReviewList(queryParams?: ArchGraphReviewQuery) {
    return request({
      url: `${BASE_URL}/review`,
      method: 'get',
      params: queryParams,
    })
  },
  createReview(data: ArchGraphForm) {
    return request({
      url: `${BASE_URL}/review`,
      method: 'post',
      data: data,
    })
  },
  deleteReview(id: number) {
    return request({
      url: `${BASE_URL}/review/${id}`,
      method: 'delete',
    })
  },
}

export default ArchGraphAPI

/** 查询参数 */
export interface ArchGraphQuery {
  page?: number
  pageSize?: number
  status?: number
  groupID?: number
  graphName?: string
}
export interface ArchGraphReviewQuery {
  page?: number
  pageSize?: number
  graphID?: number
  graphKey?: string
}

export interface ArchGraphRecordQuery {
  graphID: number
}

export interface ArchGraphData {
  total: number
  retList: ArchGraphResult[]
}

/** 查询结果 */
export interface ArchGraphResult {
  /** ID */
  id: number
  groupID: number
  /** 名称 */
  graphName: string
  groupName: string
  groupPath: string
  graphKey: string
  graphLabel: string
  imageData: string
  imageHash: string
  owner: string
  status: number
}

export interface ArchGraphRecordResult {
  /** ID */
  id: number
  graphID: number
  imageData: string
  imageHash: string
  createdTime: string
}

export interface ArchGraphReviewResult {
  /** ID */
  id: number
  graphID: number
  graphKey: string
  content: string
  reviewer: string
  createdTime: string
}

export interface ArchGraphRecordData {
  retList: ArchGraphRecordResult[]
}

/** 创建结果 */
export interface ArchGraphOptResult {
  /** ID */
  id: number
}
// 详情查询结果
export interface ArchGraphDetailResult extends ArchGraphResult {
  nodeData: any
  edgeData: any
  owner: string
  isSync: number
}

export interface ArchGraphSelectForm {
  /** ID */
  id: number
  graphID: number
}

/** 表单对象 */
export interface ArchGraphForm {
  /** ID */
  id?: number
  groupID?: number
  /** 名称 */
  graphName?: string
  graphLabel?: string
  nodeData?: any[]
  edgeData?: any[]
  imageData?: string
  owner?: string
  status?: number
  action?: string
  isShared?: number
}

export interface ArchGraphReviewForm {
  /** ID */
  id?: number
  graphID?: number
  graphName?: string
  graphKey?: string
  content?: string
  notifyParty?: string[]
}
