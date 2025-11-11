<!--  -->
<template>
  <div class="graph-apply-wrapper">
    <div class="apply-content">
      <el-form ref="formRef" label-position="left" style="max-width: 500px">
        <el-form-item label="架构图名:">
          <el-tag type="info">{{ graphName }}</el-tag>
        </el-form-item>
      </el-form>
      <el-divider content-position="left">工单列表</el-divider>
      <div class="search-wrapper">
        <el-input
          v-model="title"
          placeholder="工单标题"
          clearable
          style="width: 200px"
          @keyup.enter="handleQuery"
          @clear="handleQuery"
        />
        <el-button icon="search" type="primary" @click="handleQuery"></el-button>
      </div>
      <el-table
        row-key="id"
        v-loading="loading"
        :data="orderList"
        @selection-change="handleSelectionChange"
        style="width: 100%"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="title" label="工单标题">
          <template #default="scope">
            <span>
              {{ scope.row.title }}
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="80px">
          <template #default="scope">
            <el-tag>
              {{ graphStatusMapping[scope.row.status] }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="120">
          <template #default="scope">
            <el-button
              type="danger"
              size="small"
              icon="delete"
              @click="handleDelete(scope.row)"
              circle
            ></el-button>
            <el-button
              type="warning"
              size="small"
              icon="edit"
              @click="handleUpdate(scope.row)"
              circle
            ></el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <div class="apply-footer">
      <el-button v-if="orderList.length > 0" type="primary" @click="handleSubmit">
        提交审批
      </el-button>
    </div>
    <OrderPanel
      v-if="orderVisible"
      :graphID="graphID"
      :graphName="graphName"
      :orderID="curOrderID"
      :orderName="curOrderName"
      v-model:processOrderID="processOrderID"
      v-model:graphNodeInfo="graphNodeInfo"
      v-model:orderVisible="orderVisible"
      @handleSubmit="handleQuery"
    ></OrderPanel>
  </div>
</template>

<script setup lang="ts">
import ProcessOrderAPI, {
  type ProcessApplyForm,
  type ProcessOrderDetailResult,
  type ProcessOrderData,
  type ProcessOrderResult,
  type UnapprovedOrderQuery,
} from '@/api/process/order'
import { onMounted, ref, toRefs } from 'vue'
import OrderPanel from './OrderPanel.vue'
import type { AxiosResponse } from 'axios'
import { graphStatusMapping, processOrderTypeMapping } from '@/utils/constant'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAuthStore } from '@/store'

defineOptions({
  name: 'ApplyPanel',
})

const props = defineProps({
  graphID: {
    required: true,
    type: Number,
  },

  graphName: {
    required: true,
    type: String,
  },
})
const orderVisible = defineModel('orderVisible', {
  type: Boolean,
  required: true,
  default: () => false,
})
const graphNodeInfo = defineModel('graphNodeInfo', {
  required: true,
  type: String,
  default: () => '',
})
const curOrderID = defineModel('curOrderID', {
  required: true,
  type: Number,
  default: () => 0,
})
const curOrderName = defineModel('curOrderName', {
  required: true,
  type: String,
  default: () => '',
})

const authStore = useAuthStore()
const { userName } = authStore
const { graphID, graphName } = toRefs(props)
const formRef = ref()
const title = ref<string>('')
const processOrderID = ref<number>(0)
const orderType = ref<number>(3)
const orderList = ref<ProcessOrderResult[]>([])
const loading = ref<boolean>(false)
const selectedIDs = ref<number[]>([])
const selectedOrderGroups = ref<string[]>([])
const demandName = ref<string>('')

onMounted(() => {
  handleQuery()
})

const emit = defineEmits<{
  (e: 'handleApprove', orderType: number): void
}>()

async function handleQuery() {
  try {
    const params: UnapprovedOrderQuery = {
      graphID: graphID.value,
      title: title.value,
    }
    const resp: AxiosResponse = await ProcessOrderAPI.getUnapprovedList(params)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: ProcessOrderData = resp.data.data
      orderList.value = resData.retList
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

const handleSelectionChange = (vals: ProcessOrderResult[]) => {
  selectedIDs.value = []
  for (const val of vals) {
    selectedIDs.value.push(val.id)
    demandName.value = val.demandName
  }
}

async function handleUpdate(rowData: ProcessOrderDetailResult) {
  processOrderID.value = rowData.id
  curOrderID.value = rowData.orderID
  curOrderName.value = rowData.orderName
  orderVisible.value = true
}

function handleDelete(rowData: ProcessOrderResult) {
  ElMessageBox.confirm(`确定要删除${processOrderTypeMapping[rowData.orderType]}吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      loading.value = true
      try {
        const id: number = rowData.id ? rowData.id : 0
        const resp: AxiosResponse = await ProcessOrderAPI.delete(id)
        const status = resp.data.status
        const msg = resp.data.msg
        if (status === 200) {
          ElMessage.success('删除成功')
          handleQuery()
        } else {
          console.log(msg)
          ElMessage.error('删除失败')
        }
      } catch (err) {
        console.error(err)
        ElMessage.error('删除失败')
      } finally {
        loading.value = false
      }
    })
    .catch(() => {
      // 用户点击“取消” 或 关闭弹窗，什么都不做即可避免报错
    })
}

function handleSubmit() {
  ElMessageBox.confirm(`确定要提交审批工单吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      try {
        if (selectedIDs.value.length == 0) {
          selectedOrderGroups.value = []
          orderList.value.forEach((item: ProcessOrderResult) => {
            selectedIDs.value.push(item.id)
          })
        }
        const orderData: ProcessApplyForm = {
          ids: selectedIDs.value,
          orderType: 3,
          graphName: graphName.value,
          owner: userName,
          graphID: graphID.value,
        }
        const resp: AxiosResponse = await ProcessOrderAPI.apply(orderData)
        const status = resp.data.status
        const msg = resp.data.msg
        if (status === 200) {
          ElMessage.success('提交成功')
          // 选择成功后重新获取后端数据，render
          emit('handleApprove', orderType.value)
        } else {
          console.log(msg)
          ElMessage.error('提交失败')
        }
      } catch (err) {
        ElMessage.error('提交失败')
        console.log(err)
      }
    })
    .catch(() => {
      // 用户点击“取消” 或 关闭弹窗，什么都不做即可避免报错
    })
}
</script>

<style lang="scss" scoped>
.graph-apply-wrapper {
  height: 100%;
  display: flex;
  flex-direction: column;
  .apply-content {
    flex: 1;
    height: 100%;
    .search-wrapper {
      text-align: left;
      margin-bottom: 10px;
      // margin-left: 210px;
    }
    .btn-wrapper {
      margin-top: 20px;
      text-align: right;
      // margin-right: 20px;
    }
  }
  .apply-footer {
    text-align: right;
  }
}
</style>
