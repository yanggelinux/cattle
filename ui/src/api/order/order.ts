import request from '@/utils/request'

const BASE_URL = '/api/v1/order'

const OrderAPI = {
  /** 获取数据 */
  getList(queryParams?: OrderQuery) {
    return request({
      url: `${BASE_URL}`,
      method: 'get',
      params: queryParams,
    })
  },
  getOrderNodeList(queryParams?: OrderQuery) {
    return request({
      url: `${BASE_URL}/node`,
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
  create(data: OrderForm) {
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
  update(data: OrderForm) {
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

export default OrderAPI

/** 查询参数 */
export interface OrderQuery {
  /** 搜索关键字 */
  page?: number
  pageSize?: number
  status?: number
  orderType?: number
  name?: string
}

export interface OrderData {
  total: number
  retList: OrderResult[]
}

/** 查询结果 */
export interface OrderResult {
  /** ID */
  id: number
  /** 名称 */
  name: string
  groupID: number
  processID: number
  processNmae: string
  groupName: string
  orderType: number
  nodeType: string
  label: string
  layout: number
  isTask: number
  taskUrl: string
  taskMethod: string
  sort: number
  status: number
  /** 修改时间 */
  updatedTime: string
  /** 创建时间 */
  createdTime: string
}

/** 表单对象 */
export interface OrderForm {
  /** ID */
  id?: number
  /** 名称 */
  name?: string
  groupID?: number
  processID?: number
  orderType?: number
  nodeType?: string
  label?: string
  layout?: number
  isTask?: number
  taskUrl?: string
  taskMethod?: string
  sort?: number
  status?: number
}
