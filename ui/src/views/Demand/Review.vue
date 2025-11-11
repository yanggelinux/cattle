<template>
  <div class="demand-review-container app-container">
    <div class="demand-review-header">
      <el-row>
        <el-col :span="20" class="title-wrapper">
          <el-tag size="large" effect="plain">请求详情</el-tag>
        </el-col>
        <el-col :span="4" class="header-btn-wrapper">
          <el-button icon="Back" @click="handleGoBack">返回</el-button>
        </el-col>
      </el-row>
    </div>
    <div class="demand-review-main">
      <el-tabs
        v-model="activeName"
        type="card"
        :stretch="true"
        class="demand-detail-tabs"
        @tab-change="handleTabChange"
      >
        <el-tab-pane label="请求详情" name="请求详情">
          <div class="process-steps-wrapper">
            <ProcessSteps
              v-if="demandData.reviewProcess.length > 0"
              :orderProcess="demandData.reviewProcess"
              :activeIndex="demandData.activeIndex"
            ></ProcessSteps>
          </div>
          <div class="form-wrapper">
            <el-form ref="formRef" label-width="100px" :model="formData" label-position="right">
              <el-form-item label="请求名称:">
                <el-tag type="info">{{ demandData.name }}</el-tag>
              </el-form-item>

              <el-form-item label="请求类型:">
                <el-tag type="primary">{{ demandTypeMapping[demandData.demandType] }}</el-tag>
              </el-form-item>
              <el-form-item label="业务组:">
                <el-tag type="primary">{{ demandData.biz }}</el-tag>
              </el-form-item>
              <el-form-item label="归属人:">
                <el-tag type="primary">{{ demandData.owner }}</el-tag>
              </el-form-item>
              <el-form-item label="OA工单编号:">
                <el-tag type="primary">{{ demandData.orderNo }}</el-tag>
              </el-form-item>
              <el-form-item label-position="right" label="请求描述:">
                <el-text
                  v-for="(text, idx) in demandData.description.split('\n') || []"
                  :key="idx"
                  style="width: 1030px"
                  type="info"
                >
                  {{ text }}
                </el-text>
              </el-form-item>
              <el-form-item
                v-if="demandData.opinion.length > 0"
                label-position="right"
                label="审批记录:"
              >
                <div class="opinion-wrapper" v-for="(op, idx) in opinions" :key="idx">
                  <el-text style="width: 1030px" type="info">
                    {{ op }}
                  </el-text>
                </div>
              </el-form-item>
              <el-form-item
                v-if="isView !== 1"
                label-position="right"
                label="评审意见:"
                prop="opinion"
              >
                <el-input
                  style="width: 1030px"
                  :rows="10"
                  type="textarea"
                  v-model="formData.opinion"
                  placeholder="评审意见"
                />
              </el-form-item>
              <el-form-item v-if="isView !== 1">
                <div class="btn-wrapper">
                  <el-button type="success" icon="SuccessFilled" @click="handleApprove('pass')">
                    审批通过
                  </el-button>
                  <el-button
                    type="danger"
                    icon="CircleCloseFilled"
                    @click="handleApprove('notpass')"
                  >
                    审批不通过
                  </el-button>
                </div>
              </el-form-item>
            </el-form>
          </div>
        </el-tab-pane>
        <el-tab-pane label="架构图工单" name="架构图工单">
          <div class="order-table-wrapper">
            <el-table row-key="id" :data="graphOrderList" style="width: 100%">
              <el-table-column prop="graphName" label="架构图名称">
                <template #default="scope">
                  <span>
                    {{ scope.row.graphName }}
                  </span>
                </template>
              </el-table-column>
              <el-table-column prop="orderType" label="工单类型">
                <template #default="scope">
                  <span>
                    {{ processOrderTypeMapping[scope.row.orderType] }}
                  </span>
                </template>
              </el-table-column>
              <el-table-column prop="demandName" label="请求名称">
                <template #default="">
                  <span>
                    {{ demandName }}
                  </span>
                </template>
              </el-table-column>
              <el-table-column prop="status" label="状态">
                <template #default="scope">
                  <el-tag>
                    {{ graphStatusMapping[scope.row.status] }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="createdTime" label="创建时间">
                <template #default="scope">
                  {{ scope.row.createdTime }}
                </template>
              </el-table-column>
              <!-- <el-table-column fixed="right" label="操作" width="120">
                <template #default="scope"></template>
              </el-table-column> -->
              <el-table-column fixed="right" label="操作" width="220">
                <template #default="scope">
                  <el-button
                    type="primary"
                    size="small"
                    link
                    icon="view"
                    @click="handleToApplyOrder(scope.row)"
                  >
                    申请工单
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-tab-pane>
        <el-tab-pane label="请求工单" name="请求工单">
          <div class="order-table-wrapper">
            <el-table
              class="order-table"
              row-key="id"
              :data="resourceOrderList"
              style="width: 100%"
            >
              <el-table-column prop="graphName" label="架构图名称">
                <template #default="scope">
                  <span>
                    {{ scope.row.graphName }}
                  </span>
                </template>
              </el-table-column>
              <el-table-column prop="orderType" label="工单类型">
                <template #default="scope">
                  <span>
                    {{ processOrderTypeMapping[scope.row.orderType] }}
                  </span>
                </template>
              </el-table-column>
              <el-table-column prop="demandName" label="请求名称">
                <template #default="">
                  <span>
                    {{ demandName }}
                  </span>
                </template>
              </el-table-column>
              <el-table-column prop="orderGroup" label="工单批次">
                <template #default="scope">
                  <span>
                    {{ scope.row.orderGroup }}
                  </span>
                </template>
              </el-table-column>
              <el-table-column prop="status" label="状态">
                <template #default="scope">
                  <el-tag>
                    {{ graphStatusMapping[scope.row.status] }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="createdTime" label="创建时间">
                <template #default="scope">
                  {{ scope.row.createdTime }}
                </template>
              </el-table-column>
              <!-- <el-table-column fixed="right" label="操作" width="120">
                <template #default="scope"></template>
              </el-table-column> -->
            </el-table>
          </div>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script setup lang="ts">
import DemandAPI, {
  type DemandResult,
  type DemandForm,
  type DemandDetailQuery,
  type DemandApprovalForm,
} from '@/api/demand'
import ProcessOrderAPI, {
  type ProcessOrderQuery,
  type ProcessOrderResult,
  type ProcessOrderData,
} from '@/api/process/order'
import { onMounted, reactive, ref } from 'vue'
import { type AxiosResponse } from 'axios'
import { useRoute } from 'vue-router'
import router from '@/router'
import { ElMessage } from 'element-plus'
import { graphStatusMapping, processOrderTypeMapping, demandTypeMapping } from '@/utils/constant'
import { useAuthStore } from '@/store/modules/auth.ts'
import ProcessSteps from './components/ProcessSteps.vue'
import { useThrottleAsync } from '@/hooks/useThrottleAsync'

defineOptions({
  name: 'ProcessApproval',
})

const route = useRoute()
const demandID = route.query.demandID ? Number(route.query.demandID) : 0
const demandName = route.query.demandName ? String(route.query.demandName) : ''
const isView = route.query.isView ? Number(route.query.isView) : 0
const { userName, displayName } = useAuthStore()
//ref
const activeName = ref<string>('请求详情')
const graphOrderList = ref<ProcessOrderResult[]>([])
const resourceOrderList = ref<ProcessOrderResult[]>([])
const opinions = ref<string[]>([])
const demandData = ref<DemandResult>({
  id: 0,
  name: '',
  orderNo: '',
  demandType: 0,
  biz: '',
  owner: '',
  description: '',
  opinion: '',
  status: 0,
  updatedTime: '',
  createdTime: '',
  curReviewNode: {
    name: '',
    approverGroup: '',
    approver: '',
    approverName: '',
    role: '',
    opt: '',
    status: 1,
  },
  evaluation: '',
  isEvaluate: 0,
  evaluationRes: '',
  evaluationReason: '',
  reviewProcess: [],
  activeIndex: 0,
})
const formData = reactive<DemandForm>({
  opinion: '',
})
// 生命周期
onMounted(() => {
  getDemandDetail()
})
// method

function handleTabChange(val: string) {
  activeName.value = val
  if (val === '请求详情') {
    getDemandDetail()
  } else if (val === '架构图工单' || val === '请求工单') {
    const params: ProcessOrderQuery = {
      demandName: demandName,
    }
    getProcessOrderList(params)
  }
}

// 检测是否下一节点
// 查询
async function getDemandDetail() {
  try {
    const params: DemandDetailQuery = {
      id: demandID,
      name: demandName,
    }
    const resp: AxiosResponse = await DemandAPI.getDetail(params)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: DemandResult = resp.data.data
      demandData.value = resData
      const opinion = resData.opinion
      opinions.value = opinion.split('@')
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

async function getProcessOrderList(params: ProcessOrderQuery) {
  try {
    const resp: AxiosResponse = await ProcessOrderAPI.getList(params)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      graphOrderList.value = []
      resourceOrderList.value = []
      const resData: ProcessOrderData = resp.data.data
      const retList = resData.retList
      for (const ret of retList) {
        const orderType = ret.orderType
        if (orderType > 2) {
          resourceOrderList.value.push(ret)
        } else {
          graphOrderList.value.push(ret)
        }
      }
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

// 检测是否下一节点
function isNext(): boolean {
  const reviewProcess = demandData.value.reviewProcess
  const activeIndex = demandData.value.activeIndex
  const rpl = reviewProcess.length
  if (activeIndex === rpl - 2) {
    return false
  }
  return true
}
// pass
// notpass
// success
// pass
// notpass

const handleApprove = useThrottleAsync(handleDoApprove, 1000)

async function handleDoApprove(action: string) {
  const next = isNext()
  if (action === 'pass') {
    if (!next) {
      action = 'success'
    }
  }
  await doApprove(action)
}
async function doApprove(action: string) {
  try {
    const approvalData: DemandApprovalForm = {
      id: demandID,
      action: action,
      approver: userName,
      approverName: displayName,
      opinion: formData.opinion,
    }
    const resp: AxiosResponse = await DemandAPI.approve(approvalData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('审批操作成功')
      handleGoBack()
    } else {
      console.log(msg)
      ElMessage.error('审批操作失败')
    }
  } catch (err) {
    console.error(err)
    ElMessage.error('审批操作失败')
  }
}
function handleToApplyOrder(row: ProcessOrderResult) {
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
function handleGoBack() {
  router.go(-1)
}
</script>

<style lang="scss" scoped>
.demand-review-container {
  .demand-review-header {
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
  .demand-review-main {
    margin-top: 10px;
    background-color: #fff;
    .form-wrapper {
      padding: 20px;
      .btn-wrapper {
        line-height: 100px;
        height: 100px;
        width: 1030px;
        text-align: right;
      }
      .opinion-wrapper {
        display: block;
        width: 100%;
      }
    }
    .order-table-wrapper {
      margin-top: 10px;
    }
  }
}
</style>
