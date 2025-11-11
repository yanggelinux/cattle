import request from '@/utils/request'

const BASE_URL = '/api/v1/role'

const RoleAPI = {
  /** 获取角色数据 */
  getList(queryParams?: RoleQuery) {
    return request({
      url: `${BASE_URL}`,
      method: 'get',
      params: queryParams,
    })
  },

  /** 添加角色 */
  create(data: RoleForm) {
    return request({
      url: `${BASE_URL}`,
      method: 'post',
      data: data,
    })
  },

  /**
   * 更新角色
   *
   * @param id 角色ID
   * @param data 角色表单数据
   */
  update(data: RoleForm) {
    return request({
      url: `${BASE_URL}`,
      method: 'put',
      data: data,
    })
  },

  /**
   * 删除角色
   *
   * @param id 角色ID
   */
  delete(id: number) {
    return request({
      url: `${BASE_URL}/${id}`,
      method: 'delete',
    })
  },
}

export default RoleAPI

/** 角色分页查询参数 */
export interface RoleQuery {
  /** 搜索关键字 */
  page?: number
  pageSize?: number
  roleName?: string
}

export interface RoleData {
  total: number
  retList: RoleResult[]
}

/** 角色查询结果 */
export interface RoleResult {
  /** 角色ID */
  id: number
  /** 角色名称 */
  roleName: string
  /**展示名称  */
  displayName: string
  /** 是否超管 */
  isSuper: number
  /** 修改时间 */
  updatedTime: string
  /** 创建时间 */
  createdTime: string
}

/** 角色表单对象 */
export interface RoleForm {
  /** 角色ID */
  id?: number
  /** 角色名称 */
  roleName?: string
  /** 展示名称 */
  displayName?: string
  /** 是否超管 */
  isSuper?: number
}
