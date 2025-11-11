<template>
  <div class="process-approval-container app-container">
    <div class="process-approval-header">
      <el-row>
        <el-col :span="20" class="title-wrapper">
          <el-tag size="large" effect="plain">工单审批</el-tag>
        </el-col>
        <el-col :span="4" class="header-btn-wrapper">
          <el-button icon="Back" @click="handleGoBack">返回</el-button>
        </el-col>
      </el-row>
    </div>
    <div class="process-steps-wrapper">
      <ProcessSteps
        v-if="processOrderData.orderProcess.length > 0"
        :process="processOrderData.orderProcess"
        :activeIndex="processOrderData.activeIndex"
      ></ProcessSteps>
    </div>
    <div class="process-approval-main">
      <div class="form-wrapper">
        <el-form ref="formRef" label-width="100px" :model="formData" label-position="right">
          <el-form-item label="工单标题:">
            <el-tag class="graph-name-wrapper" type="info">
              {{ processOrderData.title }}
            </el-tag>
          </el-form-item>
          <el-form-item label="架构图名称:" v-if="orderType <= 3">
            <el-tag class="graph-name-wrapper" @click="handleToGraph" type="info">
              {{ processOrderData.graphName }}
            </el-tag>
          </el-form-item>
          <el-form-item label="请求名称:" v-if="orderType <= 4">
            <el-text>
              {{ processOrderData.demandName }}
            </el-text>
          </el-form-item>
          <el-form-item label="环境:" v-if="processOrderData.env.length > 0">
            <el-text>
              {{ envMapping[processOrderData.env] }}
            </el-text>
          </el-form-item>
          <el-form-item label="工单类型:">
            <el-tag type="primary">{{ processOrderTypeMapping[orderType] }}</el-tag>
          </el-form-item>
          <el-form-item label="架构图:" label-position="right" v-if="orderType <= 3">
            <div class="image-wrapper">
              <div class="image-item-wrapper">
                <el-result title="" sub-title="提交审批架构图">
                  <template #icon>
                    <el-image
                      style="width: 450px; height: 300px"
                      :src="processOrderData.imageData"
                      :zoom-rate="1.2"
                      :max-scale="7"
                      :min-scale="0.2"
                      :preview-src-list="[processOrderData.imageData]"
                      show-progress
                      :initial-index="4"
                      fit="cover"
                    />
                  </template>
                </el-result>
              </div>
              <div class="image-item-wrapper" v-if="processOrderData.enabledImageData.length > 0">
                <el-result title="" sub-title="已经生效架构图">
                  <template #icon>
                    <el-image
                      style="width: 450px; height: 300px"
                      :src="processOrderData.enabledImageData"
                      :zoom-rate="1.2"
                      :max-scale="7"
                      :min-scale="0.2"
                      :preview-src-list="[
                        processOrderData.enabledImageData,
                        processOrderData.imageData,
                      ]"
                      show-progress
                      :initial-index="4"
                      fit="cover"
                    />
                  </template>
                </el-result>
              </div>
            </div>
          </el-form-item>

          <el-form-item
            label="预检查结果:"
            v-if="['autoFirewall', 'strategyWhitelist'].includes(orderLabel)"
          >
            <pre class="task-result-item">{{ processOrderData.taskResult.checkResult }}</pre>
          </el-form-item>

          <el-form-item
            label="执行结果:"
            v-if="['autoFirewall', 'strategyWhitelist'].includes(orderLabel)"
          >
            <pre class="task-result-item">{{ processOrderData.taskResult.execResult }}</pre>
          </el-form-item>

          <div class="order-info-wrapper" v-if="orderType > 2">
            <div class="order-info-laber">
              <el-text>工单信息:</el-text>
            </div>
            <OrderForm
              :isApproval="isApproval"
              :isView="isView"
              v-model:approvalEdit="approvalEdit"
              :orderID="orderID"
              :processOrderID="processOrderID"
              :orderName="processOrderData.orderName"
              :orderType="orderType"
              v-model:layout="orderLayout"
              :demandList="demandList"
              v-model:orderFieldRets="orderFieldRets"
              v-model:processOrderInfo="processOrderData.orderInfo"
              @handleSubmit="handleSubmit"
            ></OrderForm>
          </div>
          <el-form-item label-position="right" label="架构图描述:" v-if="orderType <= 2">
            <el-text
              v-for="(text, idx) in processOrderData.description.split('\n') || []"
              :key="idx"
              style="width: 100%"
              type="info"
            >
              {{ text }}
            </el-text>
          </el-form-item>
          <el-form-item
            v-if="processOrderData.opinion.length > 0"
            label-position="right"
            label="审批记录:"
          >
            <div class="opinion-wrapper" v-for="(op, idx) in opinions" :key="idx">
              <el-text style="width: 100%" type="info">
                {{ op }}
              </el-text>
            </div>
          </el-form-item>
          <el-form-item v-if="isView !== 1" label-position="right" label="审批意见:" prop="opinion">
            <el-input
              style="width: 100%"
              :rows="10"
              type="textarea"
              v-model="formData.opinion"
              placeholder="审批意见"
            />
          </el-form-item>
        </el-form>
      </div>
      <div class="submit-wrapper">
        <div v-if="isView !== 1 && [1, 3].includes(processOrderData.status)">
          <div class="btn-wrapper">
            <el-button type="success" icon="SuccessFilled" @click="handleAssignApprover">
              分配审批人
            </el-button>
            <el-button type="success" icon="SuccessFilled" @click="handleApprove('success')">
              审批通过
            </el-button>
            <!-- <el-button type="primary" icon="Avatar" @click="handleAssign">指定审批人</el-button> -->
            <el-button type="danger" icon="CircleCloseFilled" @click="handleApprove('notpass')">
              审批不通过
            </el-button>
            <el-button type="danger" icon="CircleCloseFilled" @click="handleApprovalBack()">
              审批不通过到节点
            </el-button>
          </div>
        </div>
        <div
          v-if="
            isView === 1 &&
            processOrderData.status === 3 &&
            orderType > 2 &&
            curProcNodeType === 'procStart'
          "
        >
          <div class="btn-wrapper">
            <el-button type="success" icon="SuccessFilled" @click="handleReApply">
              重新提交
            </el-button>
          </div>
        </div>
      </div>
    </div>
    <div class="opt-process-wrapper">
      <el-dialog
        draggable
        v-model="processVisible"
        title="选择审批不通过打回节点"
        width="600px"
        @close="handleClose"
      >
        <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
          <el-form-item label="流程节点" prop="procNodeName">
            <el-select
              clearable
              filterable
              placement="bottom-end"
              v-model="procNodeName"
              placeholder="请选择状态"
            >
              <el-option
                v-for="item in backOrderProcess"
                :value="item.name"
                :label="item.name"
                :key="item.name"
              />
            </el-select>
          </el-form-item>
        </el-form>

        <template #footer>
          <div class="dialog-footer">
            <el-button type="primary" @click="handleApprove('notpass')">确 定</el-button>
            <el-button @click="handleClose">取 消</el-button>
          </div>
        </template>
      </el-dialog>
    </div>
    <div class="assign-approval-wrapper">
      <el-dialog
        draggable
        v-model="assignVisible"
        title="分配审批人"
        width="600px"
        @close="handleClose"
      >
        <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
          <el-form-item label-position="right" label="审批人" prop="approver">
            <el-select
              clearable
              filterable
              placement="bottom-end"
              v-model="approver"
              placeholder="请选择审批人"
            >
              <el-option
                v-for="item in userList"
                :value="item.userName"
                :label="item.userName"
                :key="item.id"
              />
            </el-select>
          </el-form-item>
        </el-form>

        <template #footer>
          <div class="dialog-footer">
            <el-button type="primary" @click="AssignApprover">确 定</el-button>
            <el-button @click="handleClose">取 消</el-button>
          </div>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script setup lang="ts">
import ProcessOrderAPI, {
  type ProcessOrderDetailResult,
  type ProcessApprovalForm,
  type ProcessApplyForm,
  type OrderInfo,
  type AssignApproverForm,
} from '@/api/process/order'
import { type ProcessNode } from '@/api/process/process'
import { computed, onMounted, reactive, ref } from 'vue'
import ProcessSteps from '@/views/Process/List/components/ProcessSteps.vue'
import { type AxiosResponse } from 'axios'
import { processOrderTypeMapping, envMapping } from '@/utils/constant'
import { useRoute } from 'vue-router'
import router from '@/router'
import { ElMessage, ElMessageBox, type Action } from 'element-plus'
import { useAuthStore } from '@/store/modules/auth.ts'
import OrderFieldAPI, {
  type OrderFieldData,
  type OrderFieldQuery,
  type OrderFieldResult,
} from '@/api/order/orderField'
import OrderForm from '@/views/ProcessOrder/Apply/components/OrderForm.vue'

import { useThrottleAsync } from '@/hooks/useThrottleAsync'
import { useDemandStore, useUserStore } from '@/store'
import type { DemandQuery } from '@/api/demand'

defineOptions({
  name: 'ProcessApproval',
})

const route = useRoute()
const orderID = route.query.orderID ? Number(route.query.orderID) : 0
const orderType = route.query.orderType ? Number(route.query.orderType) : 0
const isView = route.query.isView ? Number(route.query.isView) : 0

const { userName, displayName } = useAuthStore()
const demandStore = useDemandStore()
const { getDemandList } = demandStore
const demandList = computed(() => demandStore.demandList)

const userStore = useUserStore()
const { getUserList } = userStore
const userList = computed(() => userStore.userList)

//ref
const opinions = ref<string[]>([])
const processOrderID = ref<number>(0)
const orderLabel = ref<string>('')
const approvalEdit = ref<number>(0)
const curProcNodeType = ref<string>('')
const isApproval = ref<number>(isView === 1 ? 0 : 1)
const orderLayout = ref<number>(2)

const processVisible = ref<boolean>(false)
const procNodeName = ref<string>('')
const backOrderProcess = ref<ProcessNode[]>([])

const assignVisible = ref<boolean>(false)
const approver = ref<string>('')

// const iconStyle = ref<any>({ marginRight: '6px' })
const processOrderData = ref<ProcessOrderDetailResult>({
  id: 0,
  title: '',
  env: '',
  orderID: 0,
  graphID: 0,
  graphName: '',
  orderName: '',
  orderLabel: '',
  orderLayout: 2,
  demandName: '',
  orderProcess: [],
  orderType: 0,
  activeIndex: 1,
  status: 1,
  updatedTime: '',
  createdTime: '',
  imageData: '',
  enabledImageData: '',
  hasApproval: 0,
  orderInfo: {
    title: '',
    formData: {},
    baseFormData: {},
    groupFormDataInfo: {},
  },
  orderField: [],
  opinion: '',
  description: '',
  taskResult: {},
  edit: false,
})

const orderFieldRets = ref<OrderFieldResult[]>([])
const formData = reactive<ProcessApprovalForm>({
  opinion: '',
})
const rules = reactive({
  procNodeName: [{ required: true, message: '请选择流程节点', trigger: 'blur' }],
  approver: [{ required: true, message: '请选择审批人', trigger: 'blur' }],
})
// 生命周期
onMounted(async () => {
  const params: DemandQuery = {
    status: 2,
  }
  await getDemandList(params)
  await getOrderDetail()
  await getOrderFieldList()
  getUserList({})
})
// method

function handleApprovalBack() {
  processVisible.value = true
  backOrderProcess.value = processOrderData.value.orderProcess.slice(
    0,
    processOrderData.value.activeIndex
  )
}

function handleAssignApprover() {
  assignVisible.value = true
}

function handleClose() {
  processVisible.value = false
  assignVisible.value = false
}

// 检测是否下一节点
function isNext(): boolean {
  const orderProcess = processOrderData.value.orderProcess
  const activeIndex = processOrderData.value.activeIndex
  const rpl = orderProcess.length
  if (activeIndex === rpl - 2) {
    return false
  }
  return true
}
function genNextNode(): ProcessNode {
  const orderProcess = processOrderData.value.orderProcess
  const activeIndex = processOrderData.value.activeIndex
  const rpl = orderProcess.length
  if (activeIndex <= rpl - 2) {
    return orderProcess[activeIndex + 1]
  }
  return orderProcess[activeIndex]
}
// 查询
async function getOrderDetail() {
  try {
    const resp: AxiosResponse = await ProcessOrderAPI.getDetail(orderID)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: ProcessOrderDetailResult = resp.data.data
      const curProcessNode = resData.curOrderNode
      if (curProcessNode) {
        approvalEdit.value = curProcessNode?.approvalEdit
        curProcNodeType.value = curProcessNode.type
      }
      orderLayout.value = resData.orderLayout
      processOrderData.value = resData
      processOrderID.value = resData.id
      orderLabel.value = resData.orderLabel
      const opinion = resData.opinion
      opinions.value = opinion.split('@')
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

async function getOrderFieldList() {
  try {
    const params: OrderFieldQuery = {
      orderID: processOrderData.value.orderID,
      status: 1,
    }
    const resp: AxiosResponse = await OrderFieldAPI.getList(params)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: OrderFieldData = resp.data.data
      const retList = resData.retList
      if (processOrderData.value.orderField.length === 0) {
        orderFieldRets.value = retList
      } else {
        orderFieldRets.value = processOrderData.value.orderField
      }
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

async function AssignApprover() {
  try {
    const approvalData: AssignApproverForm = {
      id: orderID,
      approver: userName,
    }
    const resp: AxiosResponse = await ProcessOrderAPI.assignApprover(approvalData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('分配操作成功')
      handleGoBack()
      procNodeName.value = ''
    } else {
      console.log(msg)
      ElMessage.error('分配操作失败')
    }
  } catch (err) {
    console.error(err)
    ElMessage.error('分配操作失败')
  }
}

// success
// pass
// notpass
const handleApprove = useThrottleAsync(handleDoApprove, 1000)

async function handleDoApprove(action: string) {
  const next = isNext()
  const nextNode = genNextNode()
  if (next && action !== 'notpass' && orderType <= 2) {
    ElMessageBox.confirm(`直接审批成功还是，审批流转到${nextNode.name}节点`, '提示', {
      confirmButtonText: '直接审批成功',
      cancelButtonText: `${nextNode.name}`,
      distinguishCancelAndClose: true,
      type: 'success',
      lockScroll: false,
    })
      .then(async () => {
        action = 'success'
        await doApprove(action)
      })
      .catch(async (act: Action) => {
        if (act === 'cancel') {
          action = 'pass'
          await doApprove(action)
        }
      })
  } else {
    if (orderType > 2 && action === 'success') {
      action = 'pass'
    }
    await doApprove(action)
  }
}
async function doApprove(action: string) {
  try {
    const approvalData: ProcessApprovalForm = {
      id: orderID,
      graphID: processOrderData.value.graphID,
      action: action,
      procNodeName: procNodeName.value,
      approver: userName,
      orderLabel: orderLabel.value,
      approverName: displayName,
      opinion: formData.opinion,
    }
    const resp: AxiosResponse = await ProcessOrderAPI.approve(approvalData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('审批操作成功')
      handleGoBack()
      procNodeName.value = ''
    } else {
      console.log(msg)
      ElMessage.error('审批操作失败')
    }
  } catch (err) {
    console.error(err)
    ElMessage.error('审批操作失败')
  }
}

async function handleUpdate(orderInfo: OrderInfo) {
  try {
    const orderData: ProcessApplyForm = {
      id: processOrderID.value,
      demandName: orderInfo.demandName,
      orderLabel: orderLabel.value,
      isApproval: isApproval.value,
      orderType: orderType,
      owner: userName,
      orderInfo: orderInfo,
    }
    const resp: AxiosResponse = await ProcessOrderAPI.update(orderData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('提交成功')
    } else {
      console.log(msg)
      ElMessage.error(`提交失败:${msg}`)
    }
  } catch (err) {
    ElMessage.error('提交失败')
    console.log(err)
  }
}

async function handleSubmit(orderInfo: OrderInfo) {
  await handleUpdate(orderInfo)
}

async function handleReApply() {
  if (processOrderData.value.orderType != 3) {
    router.push({
      path: `/process-order/apply/${processOrderData.value.orderID}`,
      query: {
        orderID: processOrderData.value.orderID,
        orderName: processOrderData.value.orderName,
        layout: processOrderData.value.orderLayout,
        orderType: processOrderData.value.orderType,
        orderLabel: processOrderData.value.orderLabel,
        processOrderID: processOrderData.value.id,
      },
    })
    return
  }
  try {
    const reApplyData: ProcessApplyForm = {
      id: processOrderData.value.id,
    }
    const resp: AxiosResponse = await ProcessOrderAPI.reApply(reApplyData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('重新提交操作成功')
      handleToApplyOrder(processOrderData.value)
    } else {
      console.log(msg)
      ElMessage.error('重新提交操作失败')
    }
  } catch (err) {
    console.error(err)
    ElMessage.error('重新提交操作失败')
  }
}

function handleToApplyOrder(row: ProcessOrderDetailResult) {
  router.push({
    path: `/arch/graph-draw/${row.graphID}`,
    query: {
      graphID: row.graphID,
      graphName: row.graphName,
      owner: '',
      silentMode: 1,
      status: 2,
      hasOrder: 1,
    },
  })
}

function handleToGraph() {
  const row: ProcessOrderDetailResult = processOrderData.value
  router.push({
    path: `/arch/graph-draw/${row.graphID}`,
    query: {
      graphID: row.graphID,
      graphName: row.graphName,
      owner: '',
      silentMode: 1,
      status: row.status,
      hasOrder: 0,
    },
  })
}

function handleGoBack() {
  router.go(-1)
}
</script>

<style lang="scss" scoped>
.process-approval-container {
  .process-approval-header {
    height: 50px;
    line-height: 50px;
    background-color: #fff;
    .title-wrapper {
      padding-left: 30px;
    }
    .header-btn-wrapper {
      padding-right: 30px;
      text-align: right;
    }
  }
  .process-steps-wrapper {
    margin-top: 10px;
    padding-top: 60px;
    padding-bottom: 10px;
    background-color: #fff;
  }
  .process-approval-main {
    margin-top: 10px;
    .form-wrapper {
      background-color: #fff;
      padding: 20px;
      .graph-name-wrapper {
        cursor: pointer;
      }
      .image-wrapper {
        display: flex;
        justify-content: space-around;
        .image-item-wrapper {
          border-right: solid 1px var(--el-border-color);
          border-left: solid 1px var(--el-border-color);
        }
      }
      .task-result-item {
        margin: 0px;
        color: rgb(96, 98, 102);
        font-family:
          'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei',
          微软雅黑, Arial, sans-serif;
      }
      .order-info-wrapper {
        width: 100%;
        margin-left: 20px;
        display: flex;
        .order-info-laber {
          width: 70px;
        }
      }

      .opinion-wrapper {
        display: block;
        width: 100%;
      }
    }
    .submit-wrapper {
      margin-top: 10px;
      background-color: #fff;
      .btn-wrapper {
        line-height: 80px;
        height: 80px;
        margin-right: 20px;
        text-align: right;
      }
    }
  }
}
</style>
