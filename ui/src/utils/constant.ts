export interface NumberStringMaping {
  [key: number]: string
}

export interface StringAnyMaping {
  [key: string]: any
}

export interface StringAnyListMaping {
  [key: string]: any[]
}
export interface StringStringMaping {
  [key: string]: string
}

export const userOriginMapping: NumberStringMaping = {
  1: '系统用户',
}

export const envList: LabelStringValue[] = [
  { label: '生产环境', value: 'PROD' },
  { label: '测试环境', value: 'TEST' },
  { label: '开发环境', value: 'DEV' },
]

export const envMapping: StringStringMaping = {
  PROD: '生产环境',
  TEST: '测试环境',
  DEV: '开发环境',
}

export interface LabelValue {
  label: string
  value: number
}

export interface LabelStringValue {
  label: string
  value: string
}

export interface StringLabelValuesMaping {
  [key: string]: LabelValue[]
}

export const graphStatusList: LabelValue[] = [
  { label: '未审批', value: 0 },
  { label: '审批中', value: 1 },
  { label: '审批成功', value: 2 },
  { label: '审批不通过', value: 3 },
]
export const graphStatusMapping: NumberStringMaping = {
  0: '未审批',
  1: '审批中',
  2: '审批成功',
  3: '审批不通过',
}

export const demandStatusList: LabelValue[] = [
  { label: '未评审', value: 0 },
  { label: '评审中', value: 1 },
  { label: '评审成功', value: 2 },
  { label: '评审不通过', value: 3 },
]
export const demandStatusMapping: NumberStringMaping = {
  0: '未评审',
  1: '评审中',
  2: '评审成功',
  3: '评审不通过',
}

export const processStatusList: LabelValue[] = [
  { label: '生效', value: 1 },
  { label: '失效', value: 0 },
]
export const processStatusMapping: NumberStringMaping = {
  0: '失效',
  1: '生效',
}

export const approvalInfoList: LabelValue[] = [
  { label: '一个审批节点审批就通过', value: 0 },
  { label: '所有审批节点审批才通过', value: 1 },
]
export const approvalInfoMapping: NumberStringMaping = {
  0: '一个节点通过',
  1: '所有节点都通过',
}

export const demandTypeList: LabelValue[] = [
  { label: '常规', value: 1 },
  { label: '紧急', value: 2 },
]
export const demandTypeMapping: NumberStringMaping = {
  1: '常规',
  2: '紧急',
}

export const layoutTypeList: LabelValue[] = [
  { label: '单列分布', value: 1 },
  { label: '双列分布', value: 2 },
  { label: '三列分布', value: 3 },
]
export const layoutTypeMapping: NumberStringMaping = {
  1: '单列分布',
  2: '双列分布',
  3: '三列分布',
}

export const methodTypeList: string[] = ['GET', 'POST', 'PUT', 'DELETE']

export const bizGroupList: string[] = []

export const evaluationList: string[] = ['满意', '不满意']

export const processOrderTypeList: LabelValue[] = [
  { label: '架构图申请', value: 1 },
  { label: '架构图变更', value: 2 },
  { label: '请求资源工单', value: 3 },
  { label: '请求非资源工单', value: 4 },
  { label: '非请求工单', value: 5 },
]

export const processOrderTypeMapping: NumberStringMaping = {
  1: '架构图申请',
  2: '架构图变更',
  3: '请求资源工单',
  4: '请求非资源工单',
  5: '非请求工单',
}

export const newOrderTypeList: LabelValue[] = [
  { label: '请求资源工单', value: 3 },
  { label: '请求非资源工单', value: 4 },
  { label: '非请求工单', value: 5 },
]

export const newOrderTypeMapping: NumberStringMaping = {
  3: '请求资源工单',
  4: '请求非资源工单',
  5: '非请求工单',
}

export const yesOrNoList: LabelValue[] = [
  { label: '是', value: 1 },
  { label: '否', value: 0 },
]

export const yesOrNoMapping: NumberStringMaping = {
  1: '是',
  0: '否',
}

export const verRuleMapping: NumberStringMaping = {
  1: '无规则',
  2: '仅数字',
  3: '仅英文字母',
  4: '仅小写英文字母',
  5: '仅大写英文字母',
  6: '仅英文字母和数字',
  7: '仅英文字母、数字和下划线',
  8: '仅英文字母、数字和中横线',
  9: '仅中文',
  10: '邮件',
  11: '手机号',
  12: 'IP地址',
  13: '时间字符串,YYYY-mm-dd HH:MM:SS',
  14: '日期字符串,YYYY-mm-dd',
}

export const verRuleList: LabelValue[] = [
  { label: '无规则', value: 1 },
  { label: '仅数字', value: 2 },
  { label: '仅英文字母', value: 3 },
  { label: '仅小写英文字母', value: 4 },
  { label: '仅大写英文字母', value: 5 },
  { label: '仅英文字母和数字', value: 6 },
  { label: '仅英文字母、数字和下划线', value: 7 },
  { label: '仅英文字母、数字和中横线', value: 8 },
  { label: '仅中文', value: 9 },
  { label: '邮件', value: 10 },
  { label: '手机号', value: 11 },
  { label: 'IP地址', value: 12 },
  { label: '时间字符串,YYYY-mm-dd HH:MM:SS', value: 13 },
  { label: '日期字符串,YYYY-mm-dd', value: 14 },
]

export const componentList: LabelStringValue[] = [
  { label: '输入框', value: 'input' },
  { label: '数字输入框', value: 'inputNumber' },
  { label: '文本域', value: 'textarea' },
  { label: '小文本域', value: 'smallTextarea' },
  { label: '大文本域', value: 'largeTextarea' },
  { label: '单选选择器', value: 'select' },
  { label: '多选选择器', value: 'multipleSelect' },
  { label: '日期选择器', value: 'datePicker' },
  { label: '日期时间选择器', value: 'dateTimePicker' },
  { label: '文件上传', value: 'uploadFile' },
]

export const componentMapping: StringStringMaping = {
  input: '输入框',
  inputNumber: '数字输入框',
  textarea: '文本域',
  smallTextarea: '小文本域',
  largeTextarea: '大文本域',
  select: '单选选择器',
  multipleSelect: '多选选择器',
  dateTimePicker: '日期时间选择器',
  datePicker: '日期选择器',
  uploadFile: '文件上传',
}
