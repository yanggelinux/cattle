import request from '@/utils/request'

const BASE_URL = '/api/v1/team'

const TeamAPI = {
  /** 获取数据 */
  getList(queryParams?: TeamQuery) {
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
  create(data: TeamForm) {
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
  update(data: TeamForm) {
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

export default TeamAPI

/** 查询参数 */
export interface TeamQuery {
  /** 搜索关键字 */
  page?: number
  pageSize?: number
  name?: string
}

export interface TeamData {
  total: number
  retList: TeamResult[]
}

/** 查询结果 */
export interface TeamResult {
  /** ID */
  id: number
  /** 名称 */
  name: string
  leader: string
  director: string
  /** 修改时间 */
  updatedTime: string
  /** 创建时间 */
  createdTime: string
}

/** 表单对象 */
export interface TeamForm {
  /** ID */
  id?: number
  /** 名称 */
  name?: string
  leader?: string
  director?: string
}
