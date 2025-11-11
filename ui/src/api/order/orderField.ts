import request from '@/utils/request'

const BASE_URL = '/api/v1/order-field'

const OrderFieldAPI = {
  /** 获取数据 */
  getList(queryParams?: OrderFieldQuery) {
    return request({
      url: `${BASE_URL}`,
      method: 'get',
      params: queryParams,
    })
  },
  /** 添加 */
  create(data: OrderFieldForm) {
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
  update(data: OrderFieldForm) {
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

export default OrderFieldAPI

/** 查询参数 */
export interface OrderFieldQuery {
  /** 搜索关键字 */
  page?: number
  pageSize?: number
  orderID?: number
  status?: number
  name?: string
}

export interface OrderFieldData {
  total: number
  retList: OrderFieldResult[]
}

/** 查询结果 */
export interface OrderFieldResult {
  /** ID */
  id: number
  /** 名称 */
  name: string
  key: string
  orderID: number
  orderName: string
  component: string
  placeholder: string
  verRule: number
  defaultVal: string
  isRequired: number
  isTitle: number
  isEdit: number
  isClear: number
  displayField: string
  displayVal: string
  description: string
  enum: string
  groupName: string
  sort: number
  status: number
  value: string
  /** 修改时间 */
  updatedTime: string
  /** 创建时间 */
  createdTime: string
}

/** 表单对象 */
export interface OrderFieldForm {
  /** ID */
  id?: number
  /** 名称 */
  name?: string
  key?: string
  orderID?: number
  component?: string
  placeholder?: string
  verRule?: number
  defaultVal?: string
  isRequired?: number
  isTitle?: number
  isEdit?: number
  isClear?: number
  displayField?: string
  displayVal?: string
  description?: string
  enum?: string
  groupName?: string
  sort?: number
  status?: number
}
