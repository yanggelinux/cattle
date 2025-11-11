<template>
  <div class="order-apply-list-container app-container">
    <div class="order-apply-list-header app-header">
      <el-row>
        <el-col :span="12">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true" @submit.prevent>
            <el-form-item prop="name" label="工单名:">
              <el-input
                v-model="queryParams.name"
                placeholder="工单名"
                clearable
                @keyup.enter="handleQuery"
              />
            </el-form-item>
          </el-form>
        </el-col>
        <el-col :span="12">
          <div class="btn-wrapper">
            <el-button type="primary" icon="search" @click="handleQuery">查询</el-button>
            <el-button icon="refresh" @click="handleResetQuery">重置</el-button>
          </div>
        </el-col>
      </el-row>
    </div>
    <div class="order-apply-list-main">
      <el-space wrap :size="10">
        <el-card
          class="card-body"
          shadow="hover"
          style="width: 400px"
          v-for="order in orderGroupList"
          :key="order.groupID"
        >
          <template #header>
            <div class="card-header">
              <el-space :size="5">
                <span class="icon-wrapper" :class="`i-svg:menu_order`"></span>
                <span class="name-wrapper">
                  {{ order.groupName }}
                </span>
              </el-space>
            </div>
          </template>
          <div
            class="order-list-wrapper"
            @click="handleToAppaly(ret)"
            v-for="ret in orderInfo.get(order.groupID)"
            :key="ret.id"
          >
            <el-space class="order-list-item">
              <span class="icon-wrapper" :class="`i-svg:menu_res_order`"></span>
              <el-text type="info">
                {{ ret.name }}
                <el-tag round v-if="ret.orderType === 3" type="primary" size="small">资源</el-tag>
              </el-text>
            </el-space>
          </div>
        </el-card>
      </el-space>
    </div>
  </div>
</template>

<script setup lang="ts">
import OrderAPI, { type OrderQuery, type OrderResult, type OrderData } from '@/api/order/order'

import { onMounted, reactive, ref } from 'vue'
import { type AxiosResponse } from 'axios'
import router from '@/router'
import { useThrottleAsync } from '@/hooks/useThrottleAsync'
// import { newOrderTypeMapping } from '@/utils/constant'

defineOptions({
  name: 'OrderApplyList',
})

interface OrderGroup {
  /** 搜索关键字 */
  groupID: number
  groupName: string
}

//ref
const queryFormRef = ref()
// const loading = ref<boolean>(false)
const orderGroupList = ref<OrderGroup[]>([])
const orderInfo = ref<Map<number, OrderResult[]>>(new Map<number, OrderResult[]>())

// reactive
const queryParams = reactive<OrderQuery>({
  status: 1,
})

// 生命周期
onMounted(() => {
  handleQuery()
})
// method
// 查询

const handleQuery = useThrottleAsync(doQuery, 1000)

async function doQuery() {
  try {
    const resp: AxiosResponse = await OrderAPI.getList(queryParams)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: OrderData = resp.data.data
      const retList: OrderResult[] = resData.retList
      const orderInfoMap = new Map<number, OrderResult[]>()
      const orderGroups: OrderGroup[] = []
      const has: number[] = []
      for (const ret of retList) {
        const groupID = ret.groupID
        const groupName = ret.groupName
        const orders = orderInfoMap.get(groupID)
        if (orders) {
          orders.push(ret)
          orderInfoMap.set(groupID, orders)
        } else {
          orderInfoMap.set(groupID, [ret])
        }
        if (has.includes(groupID)) {
          continue
        }
        const orderGroup: OrderGroup = {
          groupID: groupID,
          groupName: groupName,
        }
        orderGroups.push(orderGroup)
        has.push(groupID)
      }
      orderInfo.value = orderInfoMap
      orderGroupList.value = orderGroups
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

//重置
// 重置查询
function handleResetQuery() {
  queryFormRef.value.resetFields()
  queryParams.name = ''
  handleQuery()
}

// 打开弹窗
function handleToAppaly(order: OrderResult) {
  if (order.orderType === 3) {
    router.push({
      path: `/process-order/graph-list`,
      query: {
        // graphName: graphName.value,
      },
    })
  } else {
    router.push({
      path: `/process-order/apply/${order.id}`,
      query: {
        orderID: order?.id,
        orderName: order.name,
        layout: order.layout,
        orderType: order.orderType,
        orderLabel: order.label,
      },
    })
  }
}
</script>

<style lang="scss" scoped>
.order-apply-list-container {
  .order-apply-list-header {
    .btn-wrapper {
      text-align: right;
    }
  }
  .order-apply-list-main {
    .card-body {
      height: 500px;
      overflow: auto;
      :deep(.el-card__header) {
        padding: 10px 0px 10px 0px;
      }
      .card-header {
        text-align: center;
        .icon-wrapper {
          color: #67c23a;
          font-size: 26px;
        }
      }
      .order-list-wrapper {
        padding: 5px;
        cursor: pointer;
        .order-list-item {
          .icon-wrapper {
            color: #409eff;
          }
        }
        &:hover {
          background: rgb(217, 236, 255);
        }
      }
    }
  }
}
</style>
