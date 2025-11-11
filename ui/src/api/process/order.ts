import request from '@/utils/request'
import { type ProcessNode } from '@/api/process/process'
import type { StringAnyListMaping } from '@/utils/constant'
import type { OrderFieldResult } from '../order/orderField'

const BASE_URL = '/api/v1/process-order'

const ProcessOrderAPI = {
  /** 获取数据 */
  getList(queryParams?: ProcessOrderQuery) {
    return request({
      url: `${BASE_URL}`,
      method: 'get',
      params: queryParams,
    })
  },

  getUnapprovedList(queryParams?: UnapprovedOrderQuery) {
    return request({
      url: `${BASE_URL}/unapproved`,
      method: 'get',
      params: queryParams,
    })
  },

  getRelatedList(queryParams?: ProcessOrderQuery) {
    return request({
      url: `${BASE_URL}/approval`,
      method: 'get',
      params: queryParams,
    })
  },

  /** 添加 */
  create(data: ProcessOrderForm) {
    return request({
      url: `${BASE_URL}`,
      method: 'post',
      data: data,
    })
  },

  /**
   * 根据id获取详情
   *
   * @param id ID
   */
  getDetail(id: number) {
    return request({
      url: `${BASE_URL}/${id}`,
      method: 'get',
    })
  },

  update(data: ProcessOrderForm) {
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
  approve(data: ProcessApprovalForm) {
    return request({
      url: `${BASE_URL}/approve`,
      method: 'post',
      data: data,
    })
  },
  assignApprover(data: AssignApproverForm) {
    return request({
      url: `${BASE_URL}/assign-approver`,
      method: 'post',
      data: data,
    })
  },
  apply(data: ProcessApplyForm) {
    return request({
      url: `${BASE_URL}/apply`,
      method: 'post',
      data: data,
    })
  },
  reApply(data: ProcessApplyForm) {
    return request({
      url: `${BASE_URL}/re-apply`,
      method: 'post',
      data: data,
    })
  },
  getFields(queryParams: OrderFieldQuery) {
    return request({
      url: `${BASE_URL}/field`,
      method: 'get',
      params: queryParams,
    })
  },
}

export default ProcessOrderAPI

/** 查询参数 */
export interface ProcessOrderQuery {
  graphName?: string
  title?: string
  demandName?: string
  approver?: string
  owner?: string
  env?: string
  orderType?: number
  startTime?: string
  endTime?: string
  label?: string
  status?: number
  page?: number
  pageSize?: number
}

export interface UnapprovedOrderQuery {
  title?: string
  graphID?: number
}

export interface ProcessOrderData {
  total: number
  retList: ProcessOrderResult[]
}

/** 查询结果 */
export interface ProcessOrderResult {
  /** ID */
  id: number
  orderID: number
  graphID: number
  title: string
  env: string
  orderName: string
  graphName: string
  demandName: string
  orderProcess: ProcessNode[]
  curOrderNode?: ProcessNode
  orderType: number
  activeIndex: number
  status: number
  hasApproval: number
  taskResult: any
  updatedTime: string
  createdTime: string
  edit: boolean
}

// 详情查询结果
export interface ProcessOrderDetailResult extends ProcessOrderResult {
  imageData: string
  enabledImageData: string
  orderInfo: OrderInfo
  orderField: OrderFieldResult[]
  opinion: string
  description: string
  demandName: string
  orderLabel: string
  orderLayout: number
}

/** 表单对象 */
export interface ProcessOrderForm {
  /** ID */
  id?: number
  graphID?: number
  orderID?: number
  /** 名称 */
  title?: string
  env?: string
  graphName?: string
  orderName?: string
  orderLabel?: string
  imageData?: string
  enabledImageData?: string
  description?: string
  demandName?: string
  owner?: string
  orderType?: number
  orderInfo?: OrderInfo
  orderField?: OrderFieldResult[]
  isApproval?: number
}

export interface ProcessApplyForm extends ProcessOrderForm {
  ids?: number[]
}

export interface ProcessApprovalForm {
  /** ID */
  id?: number
  graphID?: number
  orderLabel?: string
  /** 名称 */
  action?: string
  procNodeName?: string
  opinion?: string
  approver?: string
  approverName?: string
}

export interface AssignApproverForm {
  id?: number
  approver?: string
}

export interface OrderFieldQuery {
  orderType?: number
}

export interface OrderField {
  /** ID */
  fieldKey: string
  fieldName: string
  description: string
  placeholder: string
  component: string
  isEdit: number
}

export interface FieldData {
  field_list: OrderField[]
}

export interface OrderInfo {
  title: string
  demandName?: string
  env?: string
  baseFormData: any
  formData: any
  groupFormDataInfo: StringAnyListMaping
}
