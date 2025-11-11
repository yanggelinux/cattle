<template>
  <div class="process-container app-container">
    <div class="process-header app-header">
      <el-row>
        <el-col :span="16">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true" @submit.prevent>
            <el-form-item prop="name" label="流程名:">
              <el-input
                style="width: 200px"
                v-model="queryParams.name"
                placeholder="流程名"
                clearable
                @keyup.enter="handleQuery"
              />
            </el-form-item>
            <el-form-item label-position="right" label="状态" prop="status">
              <el-select
                clearable
                filterable
                placement="bottom-end"
                v-model="queryParams.status"
                placeholder="状态"
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
        </el-col>
        <el-col :span="8">
          <div class="btn-wrapper">
            <el-button type="primary" icon="search" @click="handleQuery">查询</el-button>
            <el-button icon="refresh" @click="handleResetQuery">重置</el-button>
            <el-button
              v-hasPerm="['/process:post']"
              type="success"
              icon="plus"
              @click="handleProcessOpt('create')"
            >
              创建
            </el-button>
          </div>
        </el-col>
      </el-row>
    </div>
    <div class="process-main">
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
          :data="processList"
          highlight-current-row
          border:true
        >
          <el-table-column type="expand">
            <template #default="props">
              <ProcessSteps
                v-if="props.row.procInfo"
                :process="props.row.procInfo"
                :activeIndex="0"
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
              <span v-if="item === 'status'">
                <el-tag v-if="scope.row[item] === 1" type="success">
                  {{ processStatusMapping[scope.row[item]] }}
                </el-tag>
                <el-tag v-else-if="scope.row[item] === 0" type="danger">
                  {{ processStatusMapping[scope.row[item]] }}
                </el-tag>
                <el-tag v-else type="primary">{{ processStatusMapping[scope.row[item]] }}</el-tag>
              </span>
              <span v-else>{{ scope.row[item] }}</span>
            </template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" width="350">
            <template #default="scope">
              <el-button
                type="primary"
                size="small"
                link
                icon="view"
                @click="handleToProcess(scope.row, 1)"
              >
                查看
              </el-button>
              <el-button
                v-hasPerm="['/process:put']"
                type="primary"
                size="small"
                link
                icon="edit"
                @click="handleProcessOpt('update', scope.row)"
              >
                编辑
              </el-button>
              <el-button
                v-hasPerm="['/process/draw:put']"
                type="primary"
                size="small"
                link
                icon="edit"
                @click="handleToProcess(scope.row, 0)"
              >
                编辑流程
              </el-button>
              <el-button
                type="primary"
                size="small"
                link
                icon="CopyDocument"
                v-hasPerm="['/process/copy:post']"
                @click="handleCopy(scope.row)"
              >
                复制
              </el-button>
              <el-button
                type="danger"
                size="small"
                link
                icon="delete"
                v-hasPerm="['/process:delete']"
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
    <ProcessOpt
      v-if="processVisible"
      v-model:visible="processVisible"
      :action="action"
      :formData="formData"
      @submit="handleResetQuery"
    ></ProcessOpt>
  </div>
</template>

<script setup lang="ts">
import ProcessAPI, {
  type ProcessQuery,
  type ProcessResult,
  type ProcessData,
  type ProcessForm,
} from '@/api/process/process'
import { onMounted, reactive, ref, watch } from 'vue'
import Pagination from '@/components/Pagination/index.vue'
import { type AxiosResponse } from 'axios'
import { indexMethod } from '@/utils/view'
import { processStatusMapping, processStatusList } from '@/utils/constant'
import router from '@/router'
import { ElMessage, ElMessageBox } from 'element-plus'
import ProcessOpt from './components/ProcessOpt.vue'
import ProcessSteps from './components/ProcessSteps.vue'

defineOptions({
  name: 'ProcessList',
})

const tableItemMapping: Map<string, string> = new Map<string, string>([
  ['name', '流程名称'],
  ['status', '状态'],
  ['updatedTime', '更新时间'],
  ['createdTime', '创建时间'],
])
const tableHead = ref<string[]>(['name', 'status', 'createdTime'])
const tableHeadOptions = ref<string[]>(['name', 'status', 'updatedTime', 'createdTime'])
const checkboxVal = ref<string[]>(tableHead.value)
//ref
const queryFormRef = ref()
const loading = ref<boolean>(false)
const total = ref<number>(0)
const processList = ref<ProcessResult[]>([])
const action = ref<string>('')
const processVisible = ref<boolean>(false)

// reactive
const queryParams = reactive<ProcessQuery>({
  page: 1,
  pageSize: 10,
})

const formData = reactive<ProcessForm>({})

// 生命周期
onMounted(async () => {
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
    const resp: AxiosResponse = await ProcessAPI.getList(queryParams)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: ProcessData = resp.data.data
      processList.value = resData.retList
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

function handleDelete(row: ProcessResult) {
  ElMessageBox.confirm(`确定要删除流程${row?.name}吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      loading.value = true
      try {
        const id: number = row.id ? row.id : 0
        const resp: AxiosResponse = await ProcessAPI.delete(id)
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

function handleCopy(row: ProcessResult) {
  ElMessageBox.confirm(`确定要复制流程${row?.name}吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      loading.value = true
      try {
        const id: number = row.id ? row.id : 0
        const resp: AxiosResponse = await ProcessAPI.copy(id)
        const status = resp.data.status
        const msg = resp.data.msg
        if (status === 200) {
          ElMessage.success('复制成功')
          handleQuery()
        } else {
          console.log(msg)
          ElMessage.error('复制失败')
        }
      } catch (err) {
        console.error(err)
        ElMessage.error('复制失败')
      } finally {
        loading.value = false
      }
    })
    .catch(() => {
      // 用户点击“取消” 或 关闭弹窗，什么都不做即可避免报错
    })
}

function handleProcessOpt(opt: string, process?: ProcessResult) {
  action.value = opt
  if (opt === 'update') {
    Object.assign(formData, process)
  }
  processVisible.value = true
}

function handleToProcess(process: ProcessResult, silentMode?: number) {
  router.push({
    path: `/process/draw/${process.id}`,
    query: {
      processID: process?.id,
      processName: process.name,
      silentMode: silentMode,
    },
  })
}
</script>

<style lang="scss" scoped>
.process-container {
  .process-header {
    .btn-wrapper {
      text-align: right;
    }
  }
  .process-main {
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
