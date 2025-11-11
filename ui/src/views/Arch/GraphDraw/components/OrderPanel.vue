<!--  -->
<template>
  <div class="order-pannel-wrapper">
    <el-dialog
      v-model="orderVisible"
      draggable
      title="请求资源工单"
      width="70%"
      @close="handleClose"
    >
      <OrderForm
        v-if="orderVisible"
        :isApproval="0"
        :isView="isView"
        v-model:approvalEdit="approvalEdit"
        :orderID="orderID"
        :processOrderID="processOrderID"
        :orderName="orderName"
        :orderType="orderType"
        v-model:layout="layout"
        v-model:orderFieldRets="orderFieldRets"
        :demandList="demandList"
        v-model:processOrderInfo="processOrderInfo"
        :cancelFunc="handleClose"
        :graphNodeInfo="graphNodeInfo"
        @handleSubmit="handleSubmit"
      ></OrderForm>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import OrderFieldAPI, {
  type OrderFieldQuery,
  type OrderFieldResult,
  type OrderFieldData,
} from '@/api/order/orderField'
import { computed, onMounted, reactive, ref, toRefs } from 'vue'
import { type AxiosResponse } from 'axios'
import type { DemandQuery } from '@/api/demand'
import { useAuthStore, useDemandStore } from '@/store'
import OrderForm from '@/views/ProcessOrder/Apply/components/OrderForm.vue'
import { ElMessage } from 'element-plus'
import ProcessOrderAPI, {
  type OrderInfo,
  type ProcessApplyForm,
  type ProcessOrderDetailResult,
} from '@/api/process/order'
import OrderAPI, { type OrderResult } from '@/api/order/order'

defineOptions({
  name: 'OrderPanel',
})

const orderType = 3
const isView = 0
const approvalEdit = 1

const props = defineProps({
  orderID: {
    required: true,
    type: Number,
    default: 0,
  },
  orderName: {
    required: true,
    type: String,
    default: '',
  },
  graphID: {
    required: true,
    type: Number,
    default: 0,
  },
  graphName: {
    required: true,
    type: String,
    default: '',
  },
})
const orderVisible = defineModel('orderVisible', {
  type: Boolean,
  required: true,
  default: false,
})

const graphNodeInfo = defineModel('graphNodeInfo', {
  required: true,
  type: String,
  default: '',
})

const processOrderID = defineModel('processOrderID', {
  required: true,
  type: Number,
  default: 0,
})

const emit = defineEmits<{
  (e: 'handleSubmit'): void
}>()

const { userName } = useAuthStore()
const demandStore = useDemandStore()
const { getDemandList } = demandStore
const demandList = computed(() => demandStore.demandList)

const { orderID, orderName, graphName, graphID } = toRefs(props)

const orderFieldRets = ref<OrderFieldResult[]>([])
const processOrderInfo = ref<any>({})
const processOrderField = ref<OrderFieldResult[]>([])

const orderLabel = ref<string>('')
const layout = ref<number>(2)

const queryParams = reactive<OrderFieldQuery>({
  orderID: orderID.value,
  status: 1,
})

// 生命周期
onMounted(async () => {
  const params: DemandQuery = {
    status: 2,
  }
  await getDemandList(params)
  if (processOrderID.value > 0) {
    await getProcessOrderDetail()
  }
  await getOrderDetail()
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

async function getOrderDetail() {
  try {
    const resp: AxiosResponse = await OrderAPI.getDetail(orderID.value)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: OrderResult = resp.data.data
      orderLabel.value = resData.label
      layout.value = resData.layout
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

async function getProcessOrderDetail() {
  try {
    const resp: AxiosResponse = await ProcessOrderAPI.getDetail(processOrderID.value)
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
      orderID: orderID.value,
      orderName: orderName.value,
      graphID: graphID.value,
      orderLabel: orderLabel.value,
      graphName: graphName.value,
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
      emit('handleSubmit')
      handleClose()
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
      id: processOrderID.value,
      demandName: orderInfo.demandName,
      orderLabel: orderLabel.value,
      orderType: orderType,
      owner: userName,
      orderInfo: orderInfo,
    }
    const resp: AxiosResponse = await ProcessOrderAPI.update(orderData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('提交成功')
      emit('handleSubmit')
      handleClose()
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
  if (processOrderID.value === 0) {
    await handleCreate(orderInfo)
  } else {
    await handleUpdate(orderInfo)
  }
}
function handleClose() {
  processOrderID.value = 0
  orderVisible.value = false
}
</script>

<style lang="scss" scoped>
.order-pannel-wrapper {
  .btn-wrapper {
    text-align: right;
    margin-right: 50px;
  }
  .opt-form-item {
    height: 50px;
    line-height: 50px;
  }
  .fire-add-wrapper {
    display: flex;
    justify-content: center;
  }
  .waring-text-wrapper {
    color: red;
    font-size: 12px;
    margin-left: 10px;
    display: inline-block;
    vertical-align: top;
    line-height: 20px;
  }
}
</style>
