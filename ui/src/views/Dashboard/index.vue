<template>
  <div class="dashboard-container app-container">
    <div class="stat-count-wrapper">
      <el-row :gutter="10" class="row-card-wrapper">
        <el-col :span="8" class="col-card-wrapper">
          <el-card shadow="hover">
            <div class="card-header-wrapper">
              <span card-header-text>架构图信息:</span>
            </div>
            <div class="card-content-wrapper">
              <div class="content-item" @click="handleToView('graph')">
                <span class="icon-wrapper primary-color" :class="`i-svg:menu_graph`"></span>
                <span class="text-wrapper">总数：</span>
                <span class="num-wrapper primary-color">{{ graphInfo.totalCount }}</span>
              </div>
              <div class="content-item" @click="handleToView('graph')">
                <span class="icon-wrapper default-color" :class="`i-svg:unapproval`"></span>
                <span class="text-wrapper">未审批：</span>
                <span class="num-wrapper default-color">{{ graphInfo.unapprovedCount }}</span>
              </div>
              <div class="content-item" @click="handleToView('graph')">
                <span class="icon-wrapper warn-color" :class="`i-svg:approving`"></span>
                <span class="text-wrapper">审批中：</span>
                <span class="num-wrapper warn-color">{{ graphInfo.approvingCount }}</span>
              </div>
              <div class="content-item" @click="handleToView('graph')">
                <span class="icon-wrapper success-color" :class="`i-svg:approval_success`"></span>
                <span class="text-wrapper">审批成功：</span>
                <span class="num-wrapper success-color">{{ graphInfo.successCount }}</span>
              </div>
              <div class="content-item" @click="handleToView('graph')">
                <span class="icon-wrapper failed-color" :class="`i-svg:approval_failed`"></span>
                <span class="text-wrapper">审批失败：</span>
                <span class="num-wrapper failed-color">{{ graphInfo.failedCount }}</span>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="8" class="col-card-wrapper">
          <el-card shadow="hover">
            <div class="card-header-wrapper">
              <span card-header-text>工单信息:</span>
            </div>
            <div class="card-content-wrapper">
              <div class="content-item" @click="handleToView('order')">
                <span class="icon-wrapper primary-color" :class="`i-svg:menu_order_list`"></span>
                <span class="text-wrapper">总数：</span>
                <span class="num-wrapper primary-color">{{ orderInfo.totalCount }}</span>
              </div>
              <div class="content-item" @click="handleToView('order', 0)">
                <span class="icon-wrapper default-color" :class="`i-svg:unapproval`"></span>
                <span class="text-wrapper">未审批：</span>
                <span class="num-wrapper default-color">{{ orderInfo.unapprovedCount }}</span>
              </div>
              <div class="content-item" @click="handleToView('order', 1)">
                <span class="icon-wrapper warn-color" :class="`i-svg:approving`"></span>
                <span class="text-wrapper">审批中：</span>
                <span class="num-wrapper warn-color">{{ orderInfo.approvingCount }}</span>
              </div>
              <div class="content-item" @click="handleToView('order', 2)">
                <span class="icon-wrapper success-color" :class="`i-svg:approval_success`"></span>
                <span class="text-wrapper">审批成功：</span>
                <span class="num-wrapper success-color">{{ orderInfo.successCount }}</span>
              </div>
              <div class="content-item" @click="handleToView('order', 3)">
                <span class="icon-wrapper failed-color" :class="`i-svg:approval_failed`"></span>
                <span class="text-wrapper">审批失败：</span>
                <span class="num-wrapper failed-color">{{ orderInfo.failedCount }}</span>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="8" class="col-card-wrapper">
          <el-card shadow="hover">
            <div class="card-header-wrapper">
              <span card-header-text>请求信息:</span>
            </div>
            <div class="card-content-wrapper">
              <div class="content-item" @click="handleToView('demand')">
                <span class="icon-wrapper primary-color" :class="`i-svg:menu_demand`"></span>
                <span class="text-wrapper">总数：</span>
                <span class="num-wrapper primary-color">{{ demandInfo.totalCount }}</span>
              </div>
              <div class="content-item" @click="handleToView('demand', 0)">
                <span class="icon-wrapper default-color" :class="`i-svg:unapproval`"></span>
                <span class="text-wrapper">未审批：</span>
                <span class="num-wrapper default-color">{{ demandInfo.unapprovedCount }}</span>
              </div>
              <div class="content-item" @click="handleToView('demand', 1)">
                <span class="icon-wrapper warn-color" :class="`i-svg:approving`"></span>
                <span class="text-wrapper">审批中：</span>
                <span class="num-wrapper warn-color">{{ demandInfo.approvingCount }}</span>
              </div>
              <div class="content-item" @click="handleToView('demand', 2)">
                <span class="icon-wrapper success-color" :class="`i-svg:approval_success`"></span>
                <span class="text-wrapper">审批成功：</span>
                <span class="num-wrapper success-color">{{ demandInfo.successCount }}</span>
              </div>
              <div class="content-item" @click="handleToView('demand', 3)">
                <span class="icon-wrapper failed-color" :class="`i-svg:approval_failed`"></span>
                <span class="text-wrapper">审批失败：</span>
                <span class="num-wrapper failed-color">{{ demandInfo.failedCount }}</span>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <div class="stat-rate-wrapper">
      <el-row>
        <el-col :span="8" class="pie-chart-item">
          <EChart :width="'100%'" :height="'400px'" :options="graphRateOptions"></EChart>
        </el-col>
        <el-col :span="8" class="pie-chart-item">
          <EChart :width="'100%'" :height="'400px'" :options="orderRateOptions"></EChart>
        </el-col>
        <el-col :span="8" class="pie-chart-item">
          <EChart :width="'100%'" :height="'400px'" :options="demandRateOptions"></EChart>
        </el-col>
      </el-row>
    </div>
    <div class="stat-dist-wrapper">
      <EChart :width="'100%'" :height="'400px'" :options="orderTypeLineOptions"></EChart>
    </div>
  </div>
</template>

<script setup lang="ts">
import DashboardAPI, {
  type DashboardGraphResult,
  type DashboardOrderResult,
  type DashboardDemandResult,
} from '@/api/dashboard'
import router from '@/router'
import type { AxiosResponse } from 'axios'
import { onMounted, ref } from 'vue'
import EChart from '@/components/EChart/index.vue'
import type { EChartsCoreOption } from 'echarts'
import { pieOption, lineOption } from './options'

const graphInfo = ref<DashboardGraphResult>({
  totalCount: 0,
  unapprovedCount: 0,
  approvingCount: 0,
  successCount: 0,
  failedCount: 0,
})
const orderInfo = ref<DashboardOrderResult>({
  totalCount: 0,
  unapprovedCount: 0,
  approvingCount: 0,
  successCount: 0,
  failedCount: 0,
  graphApplyDist: [],
  graphChangeDist: [],
  resApplyChangeDist: [],
})
const demandInfo = ref<DashboardDemandResult>({
  totalCount: 0,
  unapprovedCount: 0,
  approvingCount: 0,
  successCount: 0,
  failedCount: 0,
})
const graphRateOptions = ref<EChartsCoreOption>({})
const orderRateOptions = ref<EChartsCoreOption>({})
const demandRateOptions = ref<EChartsCoreOption>({})
const orderTypeLineOptions = ref<EChartsCoreOption>({})

onMounted(async () => {
  await getGraphInfo()
  await getOrderInfo()
  await getDemandInfo()
  genGraphRateOptions()
  genOrderRateOptions()
  genDemandRateOptions()
  genOrderTypeOptions()
})

async function getGraphInfo() {
  try {
    const resp: AxiosResponse = await DashboardAPI.getGraphInfo()
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      graphInfo.value = resp.data.data
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

async function getOrderInfo() {
  try {
    const resp: AxiosResponse = await DashboardAPI.getOrderInfo()
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      orderInfo.value = resp.data.data
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}
async function getDemandInfo() {
  try {
    const resp: AxiosResponse = await DashboardAPI.getDemandInfo()
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      demandInfo.value = resp.data.data
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}
function handleToView(t: string, status?: number) {
  let path = ''
  let query = {}
  if (t === 'graph') {
    path = `/arch/group`
  } else if (t === 'order') {
    path = `/process-order/order`
    query = { status: status }
  } else if (t === 'demand') {
    path = `/demand/list`
    query = { status: status }
  }
  router.push({
    path: path,
    query: query,
  })
}
function genGraphRateOptions() {
  const graphPieOptions = structuredClone(pieOption)
  graphPieOptions.title.text = '架构图状态分布'
  const legendData: string[] = ['未审批', '审批中', '审批成功', '审批失败']
  const data = [
    { value: graphInfo.value.unapprovedCount, name: '未审批' },
    { value: graphInfo.value.approvingCount, name: '审批中' },
    { value: graphInfo.value.successCount, name: '审批成功' },
    { value: graphInfo.value.failedCount, name: '审批失败' },
  ]
  graphPieOptions.legend.data = legendData
  graphPieOptions.series[0].data = data
  graphRateOptions.value = graphPieOptions
}

function genOrderRateOptions() {
  // 进行深copy
  const orderPieOptions = structuredClone(pieOption)
  orderPieOptions.title.text = '工单状态分布'
  const legendData: string[] = ['未审批', '审批中', '审批成功', '审批失败']

  const data = [
    { value: orderInfo.value.unapprovedCount, name: '未审批' },
    { value: orderInfo.value.approvingCount, name: '审批中' },
    { value: orderInfo.value.successCount, name: '审批成功' },
    { value: orderInfo.value.failedCount, name: '审批失败' },
  ]
  orderPieOptions.legend.data = legendData
  orderPieOptions.series[0].data = data
  orderRateOptions.value = orderPieOptions
}

function genDemandRateOptions() {
  // 进行深copy
  const demandPieOptions = structuredClone(pieOption)
  demandPieOptions.title.text = '请求状态分布'
  const legendData: string[] = ['未审批', '审批中', '审批成功', '审批失败']

  const data = [
    { value: demandInfo.value.unapprovedCount, name: '未审批' },
    { value: demandInfo.value.approvingCount, name: '审批中' },
    { value: demandInfo.value.successCount, name: '审批成功' },
    { value: demandInfo.value.failedCount, name: '审批失败' },
  ]
  demandPieOptions.legend.data = legendData
  demandPieOptions.series[0].data = data
  demandRateOptions.value = demandPieOptions
}

function genOrderTypeOptions() {
  const orderTypeOptions = structuredClone(lineOption)
  orderTypeOptions.title.text = '工单类型趋势分布'
  const legendData: string[] = ['架构图申请', '架构图变更', '资源申请变更']
  orderTypeOptions.legend.data = legendData
  const xAxisData: string[] = []
  const applyData: number[] = []
  const changeData: number[] = []
  const resApplyChangeData: number[] = []
  for (const dist of orderInfo.value.graphApplyDist) {
    xAxisData.push(dist.dt)
    applyData.push(dist.count)
  }
  for (const dist of orderInfo.value.graphChangeDist) {
    changeData.push(dist.count)
  }
  for (const dist of orderInfo.value.resApplyChangeDist) {
    resApplyChangeData.push(dist.count)
  }
  const series = [
    {
      name: '架构图申请',
      type: 'line',
      data: applyData,
    },
    {
      name: '架构图变更',
      type: 'line',
      data: changeData,
    },
    {
      name: '资源申请变更',
      type: 'line',
      data: resApplyChangeData,
    },
  ]
  orderTypeOptions.xAxis.data = xAxisData
  orderTypeOptions.series = series
  orderTypeLineOptions.value = orderTypeOptions
}
</script>

<style lang="scss" scoped>
.dashboard-container {
  .primary-color {
    color: #409eff;
  }
  .success-color {
    color: #67c23a;
  }
  .failed-color {
    color: #f56c6c;
  }
  .warn-color {
    color: #e6a23c;
  }
  .default-color {
    color: #3fc9c6;
  }
  .stat-count-wrapper {
    margin: 10px;
    .row-card-wrapper {
      height: 300px;
      .col-card-wrapper {
        height: 300px;
        .el-card {
          height: 300px;
          :deep(.el-card__body) {
            height: 300px;
            padding: 10px;
            .card-header-wrapper {
              text-align: left;
              border-bottom: rgb(220, 223, 230) 1px solid;
              height: 25px;
              line-height: 25px;
              color: rgba(0, 0, 0, 0.45);
              font-size: 16px;
            }
            .card-content-wrapper {
              display: flex;
              height: 265px;
              flex-direction: column;
              .content-item {
                height: 50px;
                line-height: 50px;
                border-bottom: rgb(220, 223, 230) 1px dashed;
                text-align: left;
                cursor: pointer;
                .icon-wrapper {
                  font-size: 28px;
                  margin-right: 10px;
                  margin-left: 20px;
                }
                .text-wrapper {
                  color: rgba(0, 0, 0, 0.45);
                  font-size: 15px;
                  font-weight: 500;
                }
                .num-wrapper {
                  font-size: 30px;
                  font-weight: 500;
                }
              }
            }
          }
        }
      }
    }
  }
  .stat-rate-wrapper {
    margin: 10px;
    .pie-chart-item {
      padding: 10px;
      background-color: #fff;
    }
  }
  .stat-dist-wrapper {
    margin: 10px;
    padding: 10px;
    background-color: #fff;
  }
}
</style>
