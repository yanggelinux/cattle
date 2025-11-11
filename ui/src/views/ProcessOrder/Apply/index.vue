<template>
  <div class="order-apply-container app-container">
    <div class="order-apply-header app-header">
      <el-row>
        <el-col :span="12">
          <div class="info-wrapper">
            <span class="icon-wrapper" :class="`i-svg:menu_order`"></span>
            <el-text class="info-item">新建工单：</el-text>
            <el-tag>{{ orderName }}</el-tag>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="btn-wrapper">
            <el-button icon="Back" @click="handleGoBack">返回</el-button>
          </div>
        </el-col>
      </el-row>
    </div>
    <OrderForm
      :isApproval="0"
      :isView="isView"
      :orderID="orderID"
      :processOrderID="processOrderID"
      :orderName="orderName"
      :orderType="orderType"
      v-model:layout="layout"
      v-model:approvalEdit="approvalEdit"
      v-model:orderFieldRets="orderFieldRets"
      :demandList="demandList"
      v-model:processOrderInfo="processOrderInfo"
      :cancelFunc="handleGoBack"
      @handleSubmit="handleSubmit"
    ></OrderForm>
  </div>
</template>

<script setup lang="ts">
import OrderFieldAPI, {
  type OrderFieldQuery,
  type OrderFieldResult,
  type OrderFieldData,
} from '@/api/order/orderField'
import { computed, onMounted, reactive, ref } from 'vue'
import { type AxiosResponse } from 'axios'
import { useRoute } from 'vue-router'
import router from '@/router'
import type { DemandQuery } from '@/api/demand'
import { useAuthStore, useDemandStore } from '@/store'
import OrderForm from './components/OrderForm.vue'
import { ElMessage } from 'element-plus'
import ProcessOrderAPI, {
  type OrderInfo,
  type ProcessApplyForm,
  type ProcessOrderDetailResult,
} from '@/api/process/order'

defineOptions({
  name: 'OrderApply',
})

const route = useRoute()
const isView = 0
const approvalEdit = 1
const orderID = route.query.orderID ? Number(route.query.orderID) : 0
const orderName = route.query.orderName ? String(route.query.orderName) : ''
const orderLabel = route.query.orderLabel ? String(route.query.orderLabel) : ''
const layout = route.query.layout ? Number(route.query.layout) : 2
const orderType = route.query.orderType ? Number(route.query.orderType) : 0
const processOrderID = route.query.processOrderID ? Number(route.query.processOrderID) : 0

const { userName } = useAuthStore()
const demandStore = useDemandStore()
const { getDemandList } = demandStore
const demandList = computed(() => demandStore.demandList)
//ref
// const loading = ref<boolean>(false)
// 编辑还是新增
const orderFieldRets = ref<OrderFieldResult[]>([])
const processOrderInfo = ref<any>({})
const processOrderField = ref<OrderFieldResult[]>([])

const queryParams = reactive<OrderFieldQuery>({
  orderID: orderID,
  status: 1,
})

// 生命周期
onMounted(async () => {
  const params: DemandQuery = {
    status: 2,
  }
  await getDemandList(params)
  if (processOrderID > 0) {
    await getProcessOrderDetail()
  }
  await getOrderField()
})
// method
// 查询
async function getOrderField() {
  try {
    const resp: AxiosResponse = await OrderFieldAPI.getList(queryParams)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: OrderFieldData = resp.data.data
      const retList = resData.retList
      if (processOrderField.value.length === 0) {
        orderFieldRets.value = retList
      } else {
        orderFieldRets.value = processOrderField.value
      }
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

async function getProcessOrderDetail() {
  try {
    const resp: AxiosResponse = await ProcessOrderAPI.getDetail(processOrderID)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: ProcessOrderDetailResult = resp.data.data
      processOrderInfo.value = resData.orderInfo
      processOrderField.value = resData.orderField
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

async function handleCreate(orderInfo: OrderInfo) {
  try {
    const orderData: ProcessApplyForm = {
      title: orderInfo.title,
      orderID: orderID,
      orderName: orderName,
      orderLabel: orderLabel,
      orderType: orderType, // process order 的order_type 和 order 的order_type 不一样 3 请求类资源工单 4 请求类非资源工单 5 非请求工单
      demandName: orderInfo.demandName,
      env: orderInfo.env,
      owner: userName,
      orderInfo: orderInfo,
    }
    const resp: AxiosResponse = await ProcessOrderAPI.create(orderData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('提交成功')
      handleToApprove()
    } else {
      console.log(msg)
      ElMessage.error(`提交失败:${msg}`)
    }
  } catch (err) {
    ElMessage.error('提交失败')
    console.log(err)
  }
}

async function handleUpdate(orderInfo: OrderInfo) {
  try {
    const orderData: ProcessApplyForm = {
      id: processOrderID,
      demandName: orderInfo.demandName,
      orderLabel: orderLabel,
      orderType: orderType,
      owner: userName,
      orderInfo: orderInfo,
    }
    const resp: AxiosResponse = await ProcessOrderAPI.update(orderData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('提交成功')
      handleToApprove()
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
  if (processOrderID === 0) {
    await handleCreate(orderInfo)
  } else {
    await handleUpdate(orderInfo)
  }
}

async function handleToApprove() {
  // 再执行一次保存操作
  router.push({
    path: `/process-order/order`,
    query: {
      // graphName: graphName.value,
    },
  })
}

function handleGoBack() {
  router.go(-1)
}
</script>

<style lang="scss" scoped>
.order-apply-container {
  .order-apply-header {
    background-color: #fff;
    height: 60px;
    line-height: 60px;
    padding: 0px 20px 0px 20px;
    margin-bottom: 10px;
    border: 1px solid var(--el-border-color-light);
    border-radius: 4px;
    box-shadow: var(--el-box-shadow-light);
    .info-wrapper {
      display: flex;
      align-items: center;
      .icon-wrapper {
        color: #67c23a;
        font-size: 30px;
      }
      .info-item {
        // margin-right: 10px;
        margin-left: 10px;
      }
    }

    .btn-wrapper {
      text-align: right;
      .btn-space-wrapper {
        margin-top: 10px !important;
      }
    }
  }
}
</style>
