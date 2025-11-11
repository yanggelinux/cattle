import request from '@/utils/request'

const BASE_URL = '/api/v1/user'

const UserAPI = {
  /** 获取数据 */
  getList(queryParams?: UserQuery) {
    return request({
      url: `${BASE_URL}`,
      method: 'get',
      params: queryParams,
    })
  },

  /** 添加 */
  create(data: UserForm) {
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
  update(data: UserForm) {
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

export default UserAPI

/** 分页查询参数 */
export interface UserQuery {
  /** 搜索关键字 */
  page?: number
  pageSize?: number
  userName?: string
}

export interface UserData {
  total: number
  retList: UserResult[]
}

/** 查询结果 */
export interface UserResult {
  /** ID */
  id: number
  /** 名称 */
  userName: string
  password: string
  email: string
  /**展示名称  */
  displayName: string
  deptName: string
  roleNames: string
  roleIDs: number[]
  origin: number
  lastLoginTime: string
  /** 修改时间 */
  updatedTime: string
  /** 创建时间 */
  createdTime: string
  value?: string
}

/** 表单对象 */
export interface UserForm {
  /** ID */
  id?: number
  /** 名称 */
  userName?: string
  password?: string
  iv?: string
  cipher?: string
  email?: string
  /** 展示名称 */
  displayName?: string
  deptName?: string
  roleIDs?: number[]
  origin?: number
}
