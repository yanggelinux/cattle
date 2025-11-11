import request from '@/utils/request'

const BASE_URL = '/api/v1/permission'

const PermssionAPI = {
  /** 获取权限数据 */
  getList(queryParams?: PermQuery) {
    return request({
      url: `${BASE_URL}`,
      method: 'get',
      params: queryParams,
    })
  },
  /**
   * 获取权限的菜单ID集合
   *
   * @param roleId 权限ID
   * @returns 权限的菜单ID集合
   */
  getRolePermList(queryPermParams: RolePermQuery) {
    return request({
      url: `${BASE_URL}/role-perm`,
      method: 'get',
      params: queryPermParams,
    })
  },

  /**
   * 分配菜单权限
   *
   * @param roleId 权限ID
   * @param data 菜单ID集合
   */
  updateRolePerm(RolePermdata: RolePermData) {
    return request({
      url: `${BASE_URL}/role-perm`,
      method: 'put',
      data: RolePermdata,
    })
  },

  /** 添加权限 */
  create(data: PermForm) {
    return request({
      url: `${BASE_URL}`,
      method: 'post',
      data: data,
    })
  },

  /**
   * 更新权限
   *
   * @param id 权限ID
   * @param data 权限表单数据
   */
  update(data: PermForm) {
    return request({
      url: `${BASE_URL}`,
      method: 'put',
      data: data,
    })
  },

  /**
   * 删除权限
   *
   * @param id 权限ID字符串
   */
  delete(id: number) {
    return request({
      url: `${BASE_URL}/${id}`,
      method: 'delete',
    })
  },
}

export default PermssionAPI

/** 权限分页查询参数 */
export interface PermQuery {
  page?: number
  pageSize?: number
  name?: string
  code?: string
  project?: string
}

export interface RolePermQuery {
  roleID: number
  project: string
  isSuper: number
}

export interface RolePermData {
  roleID: number
  project: string
  permIDList: number[]
}

export interface PermData {
  total: number
  retList: PermResult[]
}

/** 权限查询结果 */
export interface PermResult {
  /** 权限ID */
  id: number
  /** 权限父ID */
  parentID: number
  /** 权限名称 */
  name: string
  /**权限编码  */
  code: string
  /**权限uri  */
  uri: string
  /**请求方法  */
  method: string
  project: string
  /** 权限类型 */
  permType: number
  /** 是否有效 */
  isEnabled: number
  sort: number
  /** 修改时间 */
  updatedTime: string
  /** 创建时间 */
  createdTime: string
}

/** 权限表单对象 */
export interface PermForm {
  id?: number
  /** 权限父ID */
  parentID: number
  /** 权限名称 */
  name: string
  /**权限编码  */
  code: string
  /**权限uri  */
  uri?: string
  /**请求方法  */
  method?: string
  project?: string
  /** 权限类型 */
  permType: number
  /** 是否有效 */
  isEnabled?: number
  sort?: number
}

export interface RolePermdata {
  allPermTreeData: PermTreeResult[]
  rolePermIDList: number[]
}

export interface PermTreeResult extends PermResult {
  children?: PermTreeResult[]
}

export interface PermOptionResult {
  /** 值 */
  id: string
  /** 文本 */
  name: string
  /** 子列表  */
  children?: PermOptionResult[]
}
