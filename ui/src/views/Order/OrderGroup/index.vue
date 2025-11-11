<template>
  <div class="order-group-container app-container">
    <div class="order-group-header app-header">
      <el-row>
        <el-col :span="12">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true" @submit.prevent>
            <el-form-item prop="name" label="工单组:">
              <el-input
                v-model="queryParams.name"
                placeholder="工单组"
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
            <el-button
              v-hasPerm="['/order-group:post']"
              type="success"
              icon="plus"
              @click="handleOpenDialog('create')"
            >
              创建
            </el-button>
          </div>
        </el-col>
      </el-row>
    </div>
    <div class="order-group-main">
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
          :data="orderGroupList"
          highlight-current-row
          border:true
        >
          <el-table-column label="序号" type="index" :index="indexMethod" width="60" />
          <el-table-column
            v-for="item in tableHead"
            :key="item"
            :label="tableItemMapping.get(item)"
            :prop="item"
          >
            <template #default="scope">
              <span v-if="item === 'status'">{{ processStatusMapping[scope.row[item]] }}</span>
              <span v-else>{{ scope.row[item] }}</span>
            </template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" width="220">
            <template #default="scope">
              <el-button
                v-hasPerm="['/order-group:put']"
                type="primary"
                size="small"
                link
                icon="edit"
                @click="handleOpenDialog('update', scope.row)"
              >
                编辑
              </el-button>
              <el-button
                v-hasPerm="['/order-group:delete']"
                type="danger"
                size="small"
                link
                icon="delete"
                @click="handleDelete(scope.row)"
              >
                删除
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
    <!-- 表单弹窗 -->
    <OrderGroupOpt
      :formData="formData"
      :action="action"
      v-model:visible="visible"
      @submit="handleResetQuery"
    ></OrderGroupOpt>
  </div>
</template>

<script setup lang="ts">
import OrderGroupAPI, {
  type OrderGroupQuery,
  type OrderGroupResult,
  type OrderGroupData,
  type OrderGroupForm,
} from '@/api/order/orderGroup'
import { onMounted, reactive, ref, watch } from 'vue'
import Pagination from '@/components/Pagination/index.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { type AxiosResponse } from 'axios'
import OrderGroupOpt from './components/OrderGroupOpt.vue'
import { processStatusMapping } from '@/utils/constant'
import { indexMethod } from '@/utils/view'

defineOptions({
  name: 'OrderGroup',
})

const tableItemMapping: Map<string, string> = new Map<string, string>([
  ['name', '工单组名'],
  ['sort', '排序'],
  ['status', '状态'],
  ['updatedTime', '更新时间'],
  ['createdTime', '创建时间'],
])
const tableHead = ref<string[]>(['name', 'sort', 'status', 'createdTime'])
const tableHeadOptions = ref<string[]>(['name', 'sort', 'status', 'updatedTime', 'createdTime'])
const checkboxVal = ref<string[]>(tableHead.value)

//ref
const queryFormRef = ref()
const loading = ref<boolean>(false)
const visible = ref<boolean>(false)
// 编辑还是新增
const action = ref<string>('')
const orderGroupList = ref<OrderGroupResult[]>([])
const total = ref<number>(0)

// reactive
const queryParams = reactive<OrderGroupQuery>({
  page: 1,
  pageSize: 10,
})

// 表单
const formData = reactive<OrderGroupForm>({})
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
    const resp: AxiosResponse = await OrderGroupAPI.getList(queryParams)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: OrderGroupData = resp.data.data
      orderGroupList.value = resData.retList
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
  queryParams.name = ''
  handleQuery()
}

// 打开弹窗
async function handleOpenDialog(opt: string, rowData?: OrderGroupResult) {
  visible.value = true
  action.value = opt
  if (opt === 'update') {
    Object.assign(formData, rowData)
  }
}

function handleDelete(rowData: OrderGroupResult) {
  ElMessageBox.confirm(`确定要删除${rowData?.name}吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      loading.value = true
      try {
        const id: number = rowData.id ? rowData.id : 0
        const resp: AxiosResponse = await OrderGroupAPI.delete(id)
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
      // 点击“取消” 或 关闭弹窗，什么都不做即可避免报错
    })
}
</script>

<style lang="scss" scoped>
.order-group-container {
  .order-group-header {
    .btn-wrapper {
      text-align: right;
    }
  }
  .order-group-main {
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
