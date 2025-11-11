<template>
  <div class="opt-wrapper">
    <el-dialog draggable v-model="visible" :title="title" width="800px" @close="handleCloseDialog">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="工单名" prop="name">
          <el-input v-model="formData.name" placeholder="请输入工单名称" />
        </el-form-item>
        <el-form-item label="工单组" prop="groupID">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.groupID"
            placeholder="请选择工单组"
          >
            <el-option
              v-for="item in orderGroupList"
              :value="item.id"
              :label="item.name"
              :key="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="工单流程" prop="processID">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.processID"
            placeholder="请选择工单流程"
          >
            <el-option
              v-for="item in processList"
              :value="item.id"
              :label="item.name"
              :key="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="工单类型" prop="orderType">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.orderType"
            placeholder="请选择工单类型"
          >
            <el-option
              v-for="item in newOrderTypeList"
              :value="item.value"
              :label="item.label"
              :key="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="关联节点" prop="nodeType" v-if="formData.orderType === 3">
          <el-select
            clearable
            filterable
            multiple
            placement="bottom-end"
            v-model="formData.nodeType"
            placeholder="请选择关联节点"
          >
            <el-option
              v-for="item in nodeList"
              :value="item.type"
              :label="item.text"
              :key="item.type"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="标签" prop="label">
          <el-input v-model="formData.label" placeholder="请输入标签" />
        </el-form-item>
        <el-form-item label="布局" prop="layout">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.layout"
            placeholder="请选择布局"
          >
            <el-option
              v-for="item in layoutTypeList"
              :value="item.value"
              :label="item.label"
              :key="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="是关联任务" prop="isTask">
          <el-radio-group v-model="formData.isTask">
            <el-radio :value="1" size="default" border>是</el-radio>
            <el-radio :value="0" size="default" border>否</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="任务URL" prop="taskUrl" v-if="formData.isTask === 1">
          <el-input v-model="formData.taskUrl" placeholder="请输入任务URL" />
        </el-form-item>
        <el-form-item label="任务方法" prop="taskMethod" v-if="formData.isTask === 1">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.taskMethod"
            placeholder="请选择任务方法"
          >
            <el-option v-for="item in methodTypeList" :value="item" :label="item" :key="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number
            v-model="formData.sort"
            style="width: 100px"
            controls-position="right"
            :min="0"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.status"
            placeholder="请选择状态"
          >
            <el-option
              v-for="item in processStatusList"
              :value="item.value"
              :label="item.label"
              :key="item.value"
            />
          </el-select>
        </el-form-item>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="handleSubmit">确 定</el-button>
          <el-button @click="handleCloseDialog">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import OrderAPI, { type OrderForm } from '@/api/order/order'
import { onMounted, reactive, ref, toRefs, type PropType } from 'vue'
import { ElMessage } from 'element-plus'
import {
  processStatusList,
  layoutTypeList,
  methodTypeList,
  newOrderTypeList,
} from '@/utils/constant'
import { usedNodes, otherNodes, type Node } from '@/views/Arch/GraphDraw/constant/index'
import { type AxiosResponse } from 'axios'
import OrderGroupAPI, { type OrderGroupData, type OrderGroupResult } from '@/api/order/orderGroup'
import ProcessAPI, { type ProcessData, type ProcessResult } from '@/api/process/process'
defineOptions({
  name: 'OrderOpt',
})
//props
const props = defineProps({
  formData: {
    type: Object as PropType<OrderForm>,
    required: true,
    default: () => ({}),
  },
  action: {
    type: String,
    required: true,
  },
})

const visible = defineModel('visible', {
  type: Boolean,
  required: true,
  default: true,
})

const emit = defineEmits(['submit'])
const formData = props.formData
// 小心「解构赋值」导致丢失响应式 用 toRefs 或 toRef 保持响应式。
const { action } = toRefs(props)
// ref
const formRef = ref()
const title = ref<string>(action.value === 'create' ? '创建工单' : '编辑工单')
const orderGroupList = ref<OrderGroupResult[]>([])
const processList = ref<ProcessResult[]>([])
const nodeList = ref<Node[]>([
  ...[{ type: 'line', class: 'line', text: '连线' }],
  ...usedNodes,
  ...otherNodes,
])

// reactive
const rules = reactive({
  name: [{ required: true, message: '请输入工单名称', trigger: 'blur' }],
  groupID: [{ required: true, message: '请选择工单组', trigger: 'blur' }],
  processID: [{ required: true, message: '请选择工单流程', trigger: 'blur' }],
  orderType: [{ required: true, message: '请选择工单类型', trigger: 'blur' }],
  layout: [{ required: true, message: '请选择布局', trigger: 'blur' }],
  sort: [{ required: true, message: '请输入排序', trigger: 'blur' }],
  status: [{ required: true, message: '请选择状态', trigger: 'blur' }],
})

onMounted(() => {
  getOrderGroupList()
  getProcessList()
})

//method

async function getOrderGroupList() {
  try {
    const params = {}
    const resp: AxiosResponse = await OrderGroupAPI.getList(params)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: OrderGroupData = resp.data.data
      orderGroupList.value = resData.retList
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

async function getProcessList() {
  try {
    const params = {}
    const resp: AxiosResponse = await ProcessAPI.getList(params)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: ProcessData = resp.data.data
      processList.value = resData.retList
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

function handleCloseDialog() {
  visible.value = false
  formRef.value.resetFields()
  formRef.value.clearValidate()
  formData.id = undefined
}

// 提交表单
async function handleCreate() {
  try {
    const resp: AxiosResponse = await OrderAPI.create(formData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('创建成功')
      handleCloseDialog()
      emit('submit')
    } else {
      console.log(msg)
      ElMessage.error(`创建失败${msg}`)
    }
  } catch (err) {
    console.error(err)
    ElMessage.error('创建失败')
  }
}

async function handleUpdate() {
  try {
    if (formData.isTask !== 1) {
      formData.taskUrl = ''
      formData.taskMethod = ''
    }
    const resp: AxiosResponse = await OrderAPI.update(formData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('编辑成功')
      handleCloseDialog()
      emit('submit')
    } else {
      console.log(msg)
      ElMessage.error(`编辑失败${msg}`)
    }
  } catch (err) {
    console.error(err)
    ElMessage.error('编辑失败')
  }
}

async function handleSubmit() {
  formRef.value?.validate(async (valid: boolean) => {
    if (!valid) return
    if (action.value === 'update') {
      handleUpdate()
    } else {
      handleCreate()
    }
  })
}
</script>
<style lang="scss" scoped></style>
