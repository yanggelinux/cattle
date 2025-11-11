import request from '@/utils/request'

const BASE_URL = '/api/v1/arch-group'

const ArchGroupAPI = {
  /** 获取数据 */
  getList(queryParams?: ArchGroupQuery) {
    return request({
      url: `${BASE_URL}`,
      method: 'get',
      params: queryParams,
    })
  },

  /** 添加 */
  create(data: ArchGroupForm) {
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
  update(data: ArchGroupForm) {
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

export default ArchGroupAPI

/** 查询参数 */
export interface ArchGroupQuery {
  /** 搜索关键字 */
  parentID?: number
  groupName?: string
}

export interface ArchGroupData {
  retList: ArchGroupResult[]
}

/** 查询结果 */
export interface ArchGroupResult {
  /** ID */
  id: number
  parentID: number
  /** 名称 */
  groupName: string
  itemCount: number
  /** 修改时间 */
  updatedTime: string
  /** 创建时间 */
  createdTime: string
}

/** 表单对象 */
export interface ArchGroupForm {
  /** ID */
  id?: number
  /** 名称 */
  parentID?: number
  groupName?: string
}

export interface GroupTreeNode {
  label: string
  value: number
  children?: GroupTreeNode[]
}
