<template>
  <div class="order-form-container app-container">
    <div class="order-form-main">
      <div class="order-base-info">
        <el-card class="card-body" shadow="always" style="width: 100%">
          <template #header>
            <div class="card-header">
              <el-space :size="5">
                <span class="icon-wrapper" :class="`i-svg:menu_res_order`"></span>
                <span class="name-wrapper">基础信息</span>
              </el-space>
            </div>
          </template>
          <div class="form-wrapper">
            <el-form ref="baseFormRef" :model="baseFormData" :rules="baseRules" label-width="150px">
              <el-form-item label="架构图节点:" v-if="graphNodeInfo">
                <el-text type="primary" truncated>{{ graphNodeInfo }}</el-text>
              </el-form-item>
              <el-form-item label="工单标题" prop="title">
                <el-input
                  v-if="!isBaseFieldDisabled(isView, isApproval)"
                  v-model="baseFormData.title"
                  placeholder="不输入工单标题默认生成标题"
                />
                <el-text v-else>{{ baseFormData.title }}</el-text>
              </el-form-item>
              <el-form-item label="环境" prop="env">
                <el-select
                  v-if="!isBaseFieldDisabled(isView, isApproval)"
                  clearable
                  filterable
                  placement="bottom-end"
                  v-model="baseFormData.env"
                  placeholder="请选择环境"
                >
                  <el-option
                    v-for="item in envList"
                    :value="item.value"
                    :label="item.label"
                    :key="item.value"
                  />
                </el-select>
                <el-text v-else>{{ baseFormData.env }}</el-text>
              </el-form-item>
              <el-form-item v-if="orderType !== 5" label="请求名称" prop="demandName">
                <el-select
                  v-if="!isBaseFieldDisabled(isView, isApproval)"
                  v-model="baseFormData.demandName"
                  placeholder="请选择请求名称"
                  clearable
                  filterable
                  style="width: 100%"
                >
                  <el-option
                    v-for="item in demandList"
                    :key="item.id"
                    :label="item.name"
                    :value="item.name"
                  />
                </el-select>
                <el-text v-else>{{ baseFormData.demandName }}</el-text>
              </el-form-item>
            </el-form>
          </div>
        </el-card>
      </div>
      <div class="order-form-info">
        <el-card class="card-body" shadow="always" style="width: 100%">
          <template #header>
            <div class="card-header">
              <el-space :size="5">
                <span class="icon-wrapper" :class="`i-svg:menu_res_order`"></span>
                <span class="name-wrapper">表单信息</span>
              </el-space>
            </div>
          </template>
          <div class="form-wrapper">
            <el-form ref="formRef" :model="formData" :rules="rules" label-width="160px">
              <el-row :gutter="10">
                <template v-for="field in orderFieldList" :key="field.key">
                  <el-col
                    :span="layoutSpan"
                    v-if="
                      (field.displayVal.length > 0 &&
                        formData[field.displayField] === field.displayVal) ||
                      field.displayVal.length === 0
                    "
                  >
                    <el-form-item
                      :label="`${field.name}:`"
                      :prop="field.key"
                      v-if="!isFieldDisabled(isView, isApproval, approvalEdit, field)"
                    >
                      <template #label v-if="field.description.length > 0">
                        <el-tooltip :content="field.description" placement="top" effect="dark">
                          <el-text>
                            {{ field.name }}
                            <el-icon><InfoFilled /></el-icon>
                            :
                          </el-text>
                        </el-tooltip>
                      </template>
                      <!-- select -->
                      <el-select
                        v-if="field.component === 'select'"
                        v-model="formData[field.key]"
                        clearable
                        filterable
                        :placeholder="field.placeholder"
                        style="width: 100%"
                        placement="bottom-end"
                      >
                        <el-option
                          v-for="(val, key) in field.enum.split('\n')"
                          :key="key"
                          :label="val"
                          :value="val"
                        />
                      </el-select>

                      <!-- multipleSelect -->
                      <el-select
                        v-else-if="field.component === 'multipleSelect'"
                        v-model="formData[field.key]"
                        clearable
                        filterable
                        multiple
                        :placeholder="field.placeholder"
                        style="width: 100%"
                        placement="bottom-end"
                      >
                        <el-option
                          v-for="(val, key) in field.enum.split('\n')"
                          :key="key"
                          :label="val"
                          :value="val"
                        />
                      </el-select>

                      <!-- inputNumber -->
                      <el-input-number
                        v-else-if="field.component === 'inputNumber'"
                        v-model="formData[field.key]"
                        :min="0"
                        controls-position="right"
                        :placeholder="field.placeholder"
                        style="width: 100%"
                      />

                      <!-- textarea 类型 -->
                      <el-input
                        v-else-if="
                          ['textarea', 'smallTextarea', 'largeTextarea'].includes(field.component)
                        "
                        type="textarea"
                        v-model="formData[field.key]"
                        :placeholder="field.placeholder"
                        :autosize="getTextareaAutosize(field.component)"
                        style="width: 100%"
                      />

                      <el-date-picker
                        v-model="formData[field.key]"
                        v-else-if="field.component === 'dateTimePicker'"
                        type="datetime"
                        placeholder="请选择时间"
                        format="YYYY-MM-DD hh:mm:ss"
                        value-format="YYYY-MM-DD hh:mm:ss"
                        style="width: 100%"
                      />
                      <el-date-picker
                        v-model="formData[field.key]"
                        v-else-if="field.component === 'datePicker'"
                        type="date"
                        placeholder="请选择日期"
                        format="YYYY-MM-DD"
                        value-format="YYYY-MM-DD"
                        style="width: 100%"
                      />

                      <UploadFile
                        :fieldKey="field.key"
                        :fieldVal="formData[field.key]"
                        @handleUpload="handleUploadFile"
                        v-else-if="field.component === 'uploadFile'"
                      ></UploadFile>

                      <!-- 普通 input -->
                      <el-input
                        v-else
                        v-model="formData[field.key]"
                        :placeholder="field.placeholder"
                        style="width: 100%"
                      />
                    </el-form-item>

                    <el-form-item
                      :label="`${field.name}:`"
                      :prop="field.key"
                      v-if="isFieldDisabled(isView, isApproval, approvalEdit, field)"
                    >
                      <template #label v-if="field.description.length > 0">
                        <el-tooltip :content="field.description" placement="top" effect="dark">
                          <el-text>
                            {{ field.name }}
                            <el-icon><InfoFilled /></el-icon>
                            :
                          </el-text>
                        </el-tooltip>
                      </template>

                      <pre
                        class="pre-item"
                        v-if="
                          ['textarea', 'smallTextarea', 'largeTextarea'].includes(field.component)
                        "
                        >{{ formData[field.key] }}</pre
                      >
                      <el-link
                        @click="downloadFile(formData[field.key])"
                        type="primary"
                        v-else-if="field.component === 'uploadFile'"
                      >
                        {{ formData[field.key] }}
                      </el-link>
                      <el-text v-else>{{ formData[field.key] }}</el-text>
                    </el-form-item>
                  </el-col>
                </template>
              </el-row>
            </el-form>
          </div>
        </el-card>
      </div>
      <div class="order-form-group" v-if="orderGroupFieldInfo.size > 0">
        <el-card class="card-body" shadow="always" style="width: 100%">
          <template #header>
            <div class="card-header">
              <el-space :size="5">
                <span class="icon-wrapper" :class="`i-svg:menu_res_order`"></span>
                <span class="name-wrapper">表单组信息</span>
              </el-space>
            </div>
          </template>
          <div
            class="order-group-info"
            v-for="[groupName, orderGroupFieldList] in Array.from(orderGroupFieldInfo.entries())"
            :key="groupName"
          >
            <div class="order-group-info-title">
              <el-row>
                <el-col :span="12">
                  <el-tag>{{ groupName }}</el-tag>
                </el-col>
                <el-col :span="12" class="add-btn-wrapper">
                  <el-button
                    v-if="isView === 0"
                    size="small"
                    type="primary"
                    icon="plus"
                    @click="handleAddGroupItem(groupName)"
                  >
                    新增
                  </el-button>
                </el-col>
              </el-row>
            </div>
            <div class="order-group-wrapper">
              <div v-for="(orderGroupFields, idx) in orderGroupFieldList" :key="idx">
                <div class="delete-btn-wrapper">
                  <el-button
                    v-if="isView === 0"
                    size="small"
                    type="danger"
                    icon="delete"
                    circle
                    @click="handleDeleteGroupItem(groupName, idx)"
                  ></el-button>
                </div>
                <el-form
                  ref="groupFormRef"
                  class="order-group-form"
                  :model="groupFormDataInfo.get(groupName)![idx]"
                  :rules="groupFormRulesInfo.get(groupName)"
                  label-width="160px"
                >
                  <el-row :gutter="10">
                    <template v-for="field in orderGroupFields" :key="field.key">
                      <el-col
                        :span="layoutSpan"
                        v-if="
                          (field.displayVal.length > 0 &&
                            groupFormDataInfo.get(groupName)![idx][field.displayField] ===
                              field.displayVal) ||
                          field.displayVal.length === 0
                        "
                      >
                        <el-form-item
                          :label="`${field.name}:`"
                          :prop="field.key"
                          v-if="!isFieldDisabled(isView, isApproval, approvalEdit, field)"
                        >
                          <template #label v-if="field.description.length > 0">
                            <el-tooltip :content="field.description" placement="top" effect="dark">
                              <el-text>
                                {{ field.name }}
                                <el-icon><InfoFilled /></el-icon>
                                :
                              </el-text>
                            </el-tooltip>
                          </template>
                          <el-select
                            clearable
                            filterable
                            v-if="field.component === 'select'"
                            v-model="groupFormDataInfo.get(groupName)![idx][field.key]"
                            :placeholder="`${field.placeholder}`"
                            style="width: 100%"
                            placement="bottom-end"
                          >
                            <el-option
                              v-for="(val, key) in field.enum.split('\n')"
                              :key="key"
                              :label="val"
                              :value="val"
                            />
                          </el-select>

                          <el-select
                            clearable
                            filterable
                            multiple
                            v-else-if="field.component === 'multipleSelect'"
                            v-model="groupFormDataInfo.get(groupName)![idx][field.key]"
                            :placeholder="`${field.placeholder}`"
                            style="width: 100%"
                            placement="bottom-end"
                          >
                            <el-option
                              v-for="(val, key) in field.enum.split('\n')"
                              :key="key"
                              :label="val"
                              :value="val"
                            />
                          </el-select>
                          <el-input-number
                            v-else-if="field.component === 'inputNumber'"
                            v-model="groupFormDataInfo.get(groupName)![idx][field.key]"
                            :placeholder="`${field.placeholder}`"
                            controls-position="right"
                            :min="0"
                            style="width: 100%"
                          />

                          <el-input
                            v-else-if="
                              ['textarea', 'smallTextarea', 'largeTextarea'].includes(
                                field.component
                              )
                            "
                            v-model="groupFormDataInfo.get(groupName)![idx][field.key]"
                            type="textarea"
                            :placeholder="field.placeholder"
                            :autosize="getTextareaAutosize(field.component)"
                            style="width: 100%"
                          />

                          <el-date-picker
                            v-model="groupFormDataInfo.get(groupName)![idx][field.key]"
                            v-else-if="field.component === 'dateTimePicker'"
                            type="datetime"
                            placeholder="请选择时间"
                            format="YYYY-MM-DD hh:mm:ss"
                            value-format="YYYY-MM-DD hh:mm:ss"
                            style="width: 100%"
                          />
                          <el-date-picker
                            v-model="groupFormDataInfo.get(groupName)![idx][field.key]"
                            v-else-if="field.component === 'datePicker'"
                            type="date"
                            placeholder="请选择日期"
                            format="YYYY-MM-DD"
                            value-format="YYYY-MM-DD"
                            style="width: 100%"
                          />

                          <UploadFile
                            :fieldKey="field.key"
                            :fieldVal="groupFormDataInfo.get(groupName)![idx][field.key]"
                            :groupName="groupName"
                            :idx="idx"
                            @handleUpload="handleUploadFile"
                            v-else-if="field.component === 'uploadFile'"
                          ></UploadFile>

                          <el-input
                            v-else
                            v-model="groupFormDataInfo.get(groupName)![idx][field.key]"
                            :placeholder="`${field.placeholder}`"
                            style="width: 100%"
                          />
                        </el-form-item>

                        <el-form-item
                          :label="`${field.name}:`"
                          :prop="field.key"
                          v-if="isFieldDisabled(isView, isApproval, approvalEdit, field)"
                        >
                          <template #label v-if="field.description.length > 0">
                            <el-tooltip :content="field.description" placement="top" effect="dark">
                              <el-text>
                                {{ field.name }}
                                <el-icon><InfoFilled /></el-icon>
                                :
                              </el-text>
                            </el-tooltip>
                          </template>

                          <pre
                            class="pre-item"
                            v-if="
                              ['textarea', 'smallTextarea', 'largeTextarea'].includes(
                                field.component
                              )
                            "
                            >{{ groupFormDataInfo.get(groupName)![idx][field.key] }}</pre
                          >
                          <el-link
                            @click="downloadFile(groupFormDataInfo.get(groupName)![idx][field.key])"
                            type="primary"
                            v-else-if="field.component === 'uploadFile'"
                          >
                            {{ groupFormDataInfo.get(groupName)![idx][field.key] }}
                          </el-link>
                          <el-text v-else>
                            {{ groupFormDataInfo.get(groupName)![idx][field.key] }}
                          </el-text>
                        </el-form-item>
                      </el-col>
                    </template>
                  </el-row>
                </el-form>
              </div>
            </div>
          </div>
        </el-card>
      </div>
      <div class="form-submit-wrapper" v-if="isView === 0">
        <div class="submit-btn-wrapper">
          <el-button size="large" v-if="isApproval === 0" @click="handleCancel()">取消</el-button>
          <el-button
            size="large"
            v-if="isApproval === 0 || approvalEdit === 1"
            type="primary"
            @click="handleSubmit"
          >
            {{ submitText }}
          </el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { type OrderFieldResult } from '@/api/order/orderField'
import { computed, reactive, ref, toRefs, watch, type PropType } from 'vue'
import { envList } from '@/utils/constant'
import type { DemandResult } from '@/api/demand'
import { useAuthStore } from '@/store'
import { getCurrentMinutesString, removeAtImmutable } from '@/utils'
import type { OrderInfo } from '@/api/process/order'
import { validateFormsSequentially } from '@/utils/validator'
import { useThrottleAsync } from '@/hooks/useThrottleAsync'
import { ElMessage } from 'element-plus'
import { isFieldDisabled, getTextareaAutosize } from './constant'
import UploadFile from './UploadFile.vue'
import type { UploadFileInfo } from '@/api/upload'

defineOptions({
  name: 'OrderForm',
})

const { userName } = useAuthStore()

const props = defineProps({
  demandList: {
    type: Array as PropType<DemandResult[]>,
    required: false,
    default: () => [],
  },
  isView: {
    type: Number,
    required: false,
    default: () => 0,
  },
  isApproval: {
    type: Number,
    required: false,
    default: () => 0,
  },
  orderID: {
    type: Number,
    required: true,
  },
  processOrderID: {
    type: Number,
    required: true,
  },
  orderName: {
    type: String,
    required: true,
  },
  graphNodeInfo: {
    type: String,
    required: false,
  },
  orderType: {
    type: Number,
    required: true,
  },
  cancelFunc: {
    type: Function as PropType<() => void>,
    required: false,
    default: () => {},
  },
})

const { graphNodeInfo, demandList, processOrderID, orderName, orderType, isView, isApproval } =
  toRefs(props)

const orderFieldRets = defineModel<OrderFieldResult[]>('orderFieldRets', {
  required: true,
  default: [],
})

const approvalEdit = defineModel('approvalEdit', {
  type: Number,
  required: true,
  default: () => 0,
})
const layout = defineModel('layout', {
  type: Number,
  required: true,
  default: () => 2,
})

const processOrderInfo = defineModel('processOrderInfo', {
  type: Object as PropType<any>,
  required: true,
})

const emit = defineEmits<{
  (e: 'handleSubmit', orderInfo: any): void
}>()
// 计算想要保证响应式 使用computed
const layoutSpan = computed(() => Math.round(24 / layout.value))

// 编辑还是新增
const orderFieldList = ref<OrderFieldResult[]>([])

// reactive
const orderGroupFieldList = ref<Map<string, OrderFieldResult[]>>(
  new Map<string, OrderFieldResult[]>()
)
const orderGroupFieldInfo = reactive<Map<string, OrderFieldResult[][]>>(
  new Map<string, OrderFieldResult[][]>()
)
const formRef = ref()
const baseFormRef = ref()
const groupFormRef = ref()
const submitText = ref<string>(orderType.value != 3 ? '提交工单' : '确认')
const formData = reactive<any>({})
const baseFormData = reactive<any>({})
const groupFormDataInfo = reactive<Map<string, any[]>>(new Map<string, any[]>())
const groupFormRulesInfo = reactive<Map<string, any>>(new Map<string, any>())

const baseRules = reactive({
  title: [{ required: false, message: '请输入工单标题', trigger: 'blur' }],
  env: [{ required: true, message: '请选择环境', trigger: 'blur' }],
  demandName: [{ required: orderType.value !== 5, message: '请选择请求', trigger: 'blur' }],
})

const rules = reactive<any>({})

watch(
  orderFieldRets,
  () => {
    initOrder()
  },
  { immediate: true, deep: true }
)

// method
// 查询

function isBaseFieldDisabled(isView: number, isApproval: number) {
  return isView === 1 || (isView === 0 && isApproval === 1)
}

function initOrder() {
  const groupKeyNum = new Map<string, number>()
  if (processOrderID.value > 0) {
    // 初始化baseForm
    if (processOrderInfo?.value.baseFormData) {
      Object.entries(processOrderInfo?.value.baseFormData).forEach(([key, value]) => {
        baseFormData[key] = value
      })
    }
    if (processOrderInfo?.value.formData) {
      // 初始化form
      Object.entries(processOrderInfo?.value.formData).forEach(([key, value]) => {
        formData[key] = value
      })
    }

    //初始化 groupFormDataInfo
    const groupFormDatas = processOrderInfo?.value.groupFormDataInfo
    if (groupFormDatas) {
      // const groupFormDataMap = Object.entries(groupFormDatas)
      for (const [key, value] of Object.entries(groupFormDatas) as [string, any[]][]) {
        groupFormDataInfo.set(key, value)
        groupKeyNum.set(key, value.length)
      }
    }
  }

  const orderFields: OrderFieldResult[] = []
  const orderGroupFieldMap = new Map<string, OrderFieldResult[]>()
  for (const ret of orderFieldRets.value) {
    // 生成表单信息
    const groupName = ret.groupName
    if (groupName.length === 0) {
      orderFields.push(ret)
      // 处理默认值的情况
      const defaultVal = ret.defaultVal
      const key = ret.key
      if (defaultVal.trim().length > 0 && !formData[key]) {
        formData[key] = defaultVal
      }
    } else {
      // 生成表单组信息
      const orderGroupFields = orderGroupFieldMap.get(groupName)
      if (orderGroupFields) {
        orderGroupFields.push(ret)
        orderGroupFieldMap.set(groupName, orderGroupFields)
      } else {
        orderGroupFieldMap.set(groupName, [ret])
      }
    }
  }
  orderFieldList.value = orderFields
  orderGroupFieldList.value = orderGroupFieldMap
  for (const [key, orderFiled] of orderGroupFieldMap) {
    if (processOrderID.value == 0) {
      // 初始化动态表单组信息
      orderGroupFieldInfo.set(key, [orderFiled])
      // 在这里初始化表单信息
      const obj = createFormData(orderFiled)
      groupFormDataInfo.set(key, [obj])
    } else {
      const orderFileds = []
      const num = groupKeyNum.get(key) || 1
      for (let i = 0; i < num; i++) {
        orderFileds.push(orderFiled)
      }
      orderGroupFieldInfo.set(key, orderFileds)
    }
  }
  // 生成表单校验规则
  genFormRules()
  genGroupFormRules()
}

function genFormRules() {
  orderFieldList.value.forEach((field) => {
    if (field.isRequired === 1 && field.displayVal.length === 0) {
      rules[field.key] = [{ required: true, message: `${field.name}不能为空`, trigger: 'change' }]
    }
  })
}

function genGroupFormRules() {
  orderGroupFieldList.value.forEach((fields, groupName) => {
    const groupRules: any = {}
    fields.forEach((field) => {
      if (field.isRequired === 1 && field.displayVal.length === 0) {
        groupRules[field.key] = [
          { required: true, message: `${field.name}不能为空`, trigger: 'change' },
        ]
      }
    })
    groupFormRulesInfo.set(groupName, groupRules)
  })
}

function handleAddGroupItem(groupName: string) {
  // 先新增 groupFormDataInfo 表单里面数据
  const groupFormDatas = groupFormDataInfo.get(groupName)
  const orderGroupFields = orderGroupFieldList.value.get(groupName)
  if (groupFormDatas && orderGroupFields) {
    // const clonedFields = JSON.parse(JSON.stringify(orderGroupFields))
    const obj = createFormData(orderGroupFields)
    groupFormDatas?.push(obj)
    groupFormDataInfo.set(groupName, [...groupFormDatas])
    // 在新增 orderGroupFieldInfo
    const fields = orderGroupFieldInfo.get(groupName) || []
    fields.push(orderGroupFields)
    orderGroupFieldInfo.set(groupName, [...fields])
  }
}

function createFormData(orderFields: OrderFieldResult[]) {
  const obj: Record<string, any> = {}
  for (const field of orderFields) {
    const defaultVal = field.defaultVal
    const key = field.key
    // 处理默认值的情况
    if (defaultVal.trim().length > 0 && !obj[key]) {
      obj[key] = defaultVal
    }
  }
  return obj
}

function handleDeleteGroupItem(groupName: string, idx: number) {
  // 先新增 groupFormDataInfo 表单里面数据
  const groupFormDatas = groupFormDataInfo.get(groupName)
  if (groupFormDatas) {
    const newGroupFormDatas = removeAtImmutable(groupFormDatas, idx)
    groupFormDataInfo.set(groupName, newGroupFormDatas)
  }
  // 在新增 orderGroupFieldInfo
  const orderGroupFields = orderGroupFieldList.value.get(groupName)
  if (orderGroupFields) {
    const fields = orderGroupFieldInfo.get(groupName) || []
    const newFields = removeAtImmutable(fields, idx)
    orderGroupFieldInfo.set(groupName, newFields)
  }
}

function handleUploadFile(fileInfo: UploadFileInfo) {
  const fieldKey = fileInfo.fieldKey
  const fieldVal = fileInfo.fieldVal
  const groupName = fileInfo.groupName
  const idx = fileInfo.idx
  // form 的情况
  if (groupName.length === 0) {
    formData[fieldKey] = fieldVal
  } else {
    const groupForms = groupFormDataInfo.get(groupName) || []
    const groupForm = groupForms[idx]
    groupForm[fieldKey] = fieldVal
  }
}

function downloadFile(fileUrl: string) {
  const link = document.createElement('a')
  link.href = fileUrl
  link.download = '' // 如果不指定文件名，浏览器会用原始文件名
  link.target = '_blank' // 可选：新窗口打开
  link.click()
}

const handleSubmit = useThrottleAsync(doSubmit, 2000)

async function doSubmit() {
  const formValid = await validateFormsSequentially([baseFormRef.value, formRef.value])
  if (!formValid) {
    return
  }

  // 在这里判断子组件的校验
  const validate = formData._validate
  if (validate === 0) {
    const msg = formData._msg
    ElMessage.error(`${msg}`)
    return
  }
  const timestr = getCurrentMinutesString()
  let newName = ''
  if (baseFormData.title) {
    // 如果包含姓名不自动生成
    if (!baseFormData.title.includes(userName)) {
      newName = `${userName}-${baseFormData.title}-${timestr}`
      baseFormData.title = newName
    }
  } else {
    newName = `${userName}-${orderName.value}-${timestr}`
    baseFormData.title = newName
  }
  if (graphNodeInfo?.value) {
    baseFormData.graphNodeInfo = graphNodeInfo.value
  }
  const orderInfo: OrderInfo = {
    title: newName,
    demandName: baseFormData.demandName || '',
    env: baseFormData.env || '',
    formData: { ...formData },
    baseFormData: { ...baseFormData },
    groupFormDataInfo: Object.fromEntries(groupFormDataInfo),
  }
  // 调用 创建流程工单的接口
  emit('handleSubmit', orderInfo)
}

function handleCancel() {
  props.cancelFunc()
}
</script>

<style lang="scss" scoped>
.order-form-container {
  width: 100%;
  .order-form-main {
    .order-base-info {
      margin-top: 10px;
      .edit-btn-wrapper {
        text-align: right;
      }
    }
    .order-form-info {
      margin: 10px 0px 10px 0px;
    }
    .order-form-group {
      margin-top: 10px;
      .order-group-info {
        margin-top: 10px;
        .order-group-info-title {
          margin-bottom: 10px;
          .add-btn-wrapper {
            text-align: right;
          }
        }

        .order-group-wrapper {
          border: 1px solid #e6e8eb;
          .delete-btn-wrapper {
            text-align: right;
            margin: 3px 10px 3px 0px;
          }
          .order-group-form {
            padding: 10px 10px 5px 10px;
            border-bottom: 1px solid #e6e8eb;
          }
        }
      }
    }
    .form-submit-wrapper {
      background-color: #fff;
      padding: 10px 10px 10px 10px;
      margin-top: 10px;
      text-align: right;
      border: 1px solid var(--el-border-color-light);
      border-radius: 4px;
      box-shadow: var(--el-box-shadow-light);
    }
    .pre-item {
      overflow: auto;
      margin: 0px;
      color: rgb(96, 98, 102);
      max-height: 500px;
      font-family:
        'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', 微软雅黑,
        Arial, sans-serif;
    }
  }
}
</style>
