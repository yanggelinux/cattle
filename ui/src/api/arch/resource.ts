import request from '@/utils/request'

const BASE_URL = '/api/v1/resource'

const ResourceAPI = {
  /** 获取数据 */
  getList(queryParams?: ResourceQuery) {
    return request({
      url: `${BASE_URL}`,
      method: 'get',
      params: queryParams,
    })
  },
}

export default ResourceAPI

/** 查询参数 */
export interface ResourceQuery {
  resourceType?: string
  relKeys?: string
}

export interface ResourceData {
  retList: ResourceResult[]
  relRetList: ResourceResult[]
}

/** 查询结果 */
export interface ResourceResult {
  /** ID */
  id: number
  resourceID: string
  resourceName: string
  ip: string
  resourceKey: string
  resourceType: string
  port: number
  spec: string
  cpu: number
  mem: string
}
