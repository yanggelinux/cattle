<template>
  <div class="process-order-container app-container">
    <div class="process-order-header app-header">
      <el-row>
        <el-col :span="16">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true" @submit.prevent>
            <el-form-item prop="title" label="工单标题:">
              <el-input
                style="width: 180px"
                v-model="queryParams.title"
                placeholder="工单标题"
                clearable
                @keyup.enter="handleQuery"
              />
            </el-form-item>
            <el-form-item label="工单类型:" prop="orderType">
              <el-select
                style="width: 180px"
                clearable
                filterable
                placement="bottom-end"
                v-model="queryParams.orderType"
                placeholder="请选工单类型"
              >
                <el-option
                  v-for="item in processOrderTypeList"
                  :value="item.value"
                  :label="item.label"
                  :key="item.value"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="工单状态:" prop="status" width="300px">
              <el-select
                style="width: 180px"
                clearable
                filterable
                placement="bottom-end"
                v-model="queryParams.status"
                placeholder="请选工单状态"
              >
                <el-option
                  v-for="item in graphStatusList"
                  :value="item.value"
                  :label="item.label"
                  :key="item.value"
                />
              </el-select>
            </el-form-item>
            <el-form-item prop="demandName" label="请求名称:">
              <el-input
                style="width: 180px"
                v-model="queryParams.demandName"
                placeholder="请求名称"
                clearable
                @keyup.enter="handleQuery"
              />
            </el-form-item>
          </el-form>
        </el-col>
        <el-col :span="8">
          <div class="btn-wrapper">
            <el-button type="primary" icon="search" @click="handleQuery">查询</el-button>
            <el-button icon="refresh" @click="handleResetQuery">重置</el-button>
          </div>
        </el-col>
      </el-row>
    </div>
    <div class="process-order-main">
      <div class="filter-wrapper">
        <el-collapse accordion>
          <el-collapse-item>
            <template #title>
              <span class="filter-text">
                表格显示字段选择
                <el-icon><CaretBottom /></el-icon>
              </span>
            </template>
            <el-checkbox-group class="filter-group" v-model="checkboxVal" :min="1">
              <el-checkbox v-for="item in tableHeadOptions" :key="item" :value="item">
                <span>{{ tableItemMapping.get(item) }}</span>
              </el-checkbox>
            </el-checkbox-group>
          </el-collapse-item>
        </el-collapse>
      </div>
      <div class="table-wrapper">
        <el-table
          ref="dataTableRef"
          v-loading="loading"
          :data="processOrderList"
          highlight-current-row
          border:true
        >
          <el-table-column type="expand">
            <template #default="props">
              <ProcessSteps
                :process="props.row.orderProcess"
                :activeIndex="props.row.activeIndex"
              ></ProcessSteps>
            </template>
          </el-table-column>
          <el-table-column label="序号" type="index" :index="indexMethod" width="60" />
          <el-table-column
            v-for="item in tableHead"
            :key="item"
            :label="tableItemMapping.get(item)"
            :prop="item"
          >
            <template #default="scope">
              <span v-if="item === 'orderType'">
                {{ processOrderTypeMapping[scope.row[item]] }}
              </span>
              <span v-else-if="item === 'status'">
                <el-tag v-if="scope.row[item] === 2" type="success">
                  {{ graphStatusMapping[scope.row[item]] }}
                </el-tag>
                <el-tag v-else-if="scope.row[item] === 3" type="danger">
                  {{ graphStatusMapping[scope.row[item]] }}
                </el-tag>
                <el-tag v-else type="primary">{{ graphStatusMapping[scope.row[item]] }}</el-tag>
              </span>
              <span v-else-if="item === 'curOrderNode'">
                <span type="primary" effect="plain">{{ scope.row[item].name }}</span>
              </span>

              <span v-else>{{ scope.row[item] }}</span>
            </template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" width="220">
            <template #default="scope">
              <el-button
                type="primary"
                size="small"
                link
                icon="view"
                @click="handleToApproval(scope.row, 1)"
              >
                详情
              </el-button>
            </template>
          </el-table-column>
        </el-table>

        <Pagination
          v-if="total > 0"
          v-model:total="total"
          v-model:page="queryParams.page"
          v-model:limit="queryParams.pageSize"
          @pagination="handleQuery"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import ProcessOrderAPI, {
  type ProcessOrderQuery,
  type ProcessOrderResult,
  type ProcessOrderData,
} from '@/api/process/order'
import { onMounted, reactive, ref, watch } from 'vue'
import Pagination from '@/components/Pagination/index.vue'
import ProcessSteps from '@/views/Process/List/components/ProcessSteps.vue'
import { type AxiosResponse } from 'axios'
import { indexMethod } from '@/utils/view'
import {
  graphStatusMapping,
  graphStatusList,
  processOrderTypeMapping,
  processOrderTypeList,
} from '@/utils/constant'
import router from '@/router'
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/store'

defineOptions({
  name: 'MyCreateOrder',
})

const route = useRoute()
let orderStatus = route.query.status ? Number(route.query.status) : -1
const { displayName } = useAuthStore()
const tableItemMapping: Map<string, string> = new Map<string, string>([
  ['title', '工单标题'],
  ['graphName', '架构图名称'],
  ['demandName', '请求名称'],
  ['orderType', '工单类型'],
  ['orderName', '工单名称'],
  ['owner', '创建人'],
  ['env', '环境'],
  ['status', '状态'],
  ['curOrderNode', '当前工单节点'],
  ['updatedTime', '更新时间'],
  ['createdTime', '创建时间'],
])
const tableHead = ref<string[]>([
  'title',
  'graphName',
  'demandName',
  'orderType',
  'orderName',
  'owner',
  'env',
  'status',
  'curOrderNode',
  'createdTime',
])
const tableHeadOptions = ref<string[]>([
  'title',
  'graphName',
  'demandName',
  'orderType',
  'orderName',
  'owner',
  'env',
  'status',
  'curOrderNode',
  'updatedTime',
  'createdTime',
])
const checkboxVal = ref<string[]>(tableHead.value)
//ref
const queryFormRef = ref()
const loading = ref<boolean>(false)
const total = ref<number>(0)
const processOrderList = ref<ProcessOrderResult[]>([])

// reactive
const queryParams = reactive<ProcessOrderQuery>({
  approver: displayName,
  page: 1,
  pageSize: 10,
})

// 生命周期
onMounted(() => {
  handleQuery()
})
watch(
  () => checkboxVal.value,
  (valArr) => {
    tableHead.value = tableHeadOptions.value.filter((i) => valArr.indexOf(i) >= 0)
  },
  { immediate: true }
)
// method
// 查询
async function handleQuery() {
  try {
    if (orderStatus !== -1) {
      queryParams.status = orderStatus
    }
    const resp: AxiosResponse = await ProcessOrderAPI.getRelatedList(queryParams)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: ProcessOrderData = resp.data.data
      processOrderList.value = resData.retList
      total.value = resData.total
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
  queryParams.page = 1
  queryParams.pageSize = 10
  queryParams.title = ''
  queryParams.demandName = ''
  orderStatus = -1
  handleQuery()
}
function handleToApproval(rowData: ProcessOrderResult, isView?: number) {
  router.push({
    path: `/process-order/approval`,
    query: {
      orderID: rowData.id,
      orderType: rowData.orderType,
      isView: isView,
    },
  })
}
</script>

<style lang="scss" scoped>
.process-order-container {
  .process-order-header {
    .btn-wrapper {
      text-align: right;
    }
  }
  .process-order-main {
    .filter-wrapper {
      text-align: left;
      .filter-text {
        margin-left: 10px;
        font-size: 14px;
        color: #409eff;
      }
      .filter-group {
        margin-left: 10px;
      }
    }
    .table-wrapper {
      margin-top: 3px;
    }
  }
}
</style>
