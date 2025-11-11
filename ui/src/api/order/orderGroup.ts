import request from '@/utils/request'

const BASE_URL = '/api/v1/order-group'

const OrderGroupAPI = {
  /** 获取数据 */
  getList(queryParams?: OrderGroupQuery) {
    return request({
      url: `${BASE_URL}`,
      method: 'get',
      params: queryParams,
    })
  },
  /** 添加 */
  create(data: OrderGroupForm) {
    return request({
      url: `${BASE_URL}`,
      method: 'post',
      data: data,
    })
  },
  /**
   * 更新
   * @param data 表单数据
   */
  update(data: OrderGroupForm) {
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
}

export default OrderGroupAPI

/** 查询参数 */
export interface OrderGroupQuery {
  /** 搜索关键字 */
  page?: number
  pageSize?: number
  status?: number
  name?: string
}

export interface OrderGroupData {
  total: number
  retList: OrderGroupResult[]
}

/** 查询结果 */
export interface OrderGroupResult {
  /** ID */
  id: number
  /** 名称 */
  name: string
  sort: number
  status: number
  /** 修改时间 */
  updatedTime: string
  /** 创建时间 */
  createdTime: string
}

/** 表单对象 */
export interface OrderGroupForm {
  /** ID */
  id?: number
  /** 名称 */
  name?: string
  sort?: number
  status?: number
}
