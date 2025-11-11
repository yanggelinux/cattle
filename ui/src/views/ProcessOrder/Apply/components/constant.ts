export function isFieldDisabled(
  isView: number,
  isApproval: number,
  approvalEdit: number,
  field: any
) {
  // 查看详情的情况
  if (isView === 1) {
    return true
  }
  // 工单申请的情况
  if (isView === 0 && isApproval === 0) {
    return false
  }
  // 工单审批的情况
  if (isView === 0 && isApproval === 1) {
    // 节点没有审批权限都disabled
    if (approvalEdit === 0) {
      return true
    }
    // 节点有审批权限 但是字段没编辑权限 disabled
    if (approvalEdit === 1 && field.isEdit === 0) {
      return true
    }
  }
  return false
}

// 根据 textarea 类型返回 autosize 配置
export function getTextareaAutosize(type: string) {
  switch (type) {
    case 'smallTextarea':
      return { minRows: 1, maxRows: 5 }
    case 'largeTextarea':
      return { minRows: 4, maxRows: 20 }
    default: // 'textarea'
      return { minRows: 2, maxRows: 10 }
  }
}
