import request from '@/utils/request'

const AUTH_BASE_URL = '/api/v1/auth'

const AuthAPI = {
  /** 登录接口*/
  login(data: LoginFormData) {
    return request({
      url: `${AUTH_BASE_URL}/login`,
      method: 'post',
      data: data,
    })
  },
  getRolePerm(params: GetRolePermParams) {
    return request({
      url: `${AUTH_BASE_URL}/user-perm`,
      method: 'get',
      params: params,
    })
  },
}

export default AuthAPI

/** 登录请求参数 */
export interface LoginFormData {
  /** 用户名 */
  userName: string
  /** 密码 */
  password?: string
}

/** 登录响应 */
export interface LoginResult extends RolePermResult {
  userID: number
  userName: string
  displayName: string
  email: string
  deptName: string
  /** 访问令牌 */
  token: string
  authorization: string
  project: string
}

export interface GetRolePermParams {
  userID: number
}

export interface RolePermResult {
  isSuper: number
  roleNames: string
  roleDisplayNames: string
  menus: string[]
  uris: string[]
}
