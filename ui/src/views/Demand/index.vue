<template>
  <div class="demand-container app-container">
    <div class="demand-header app-header">
      <el-row>
        <el-col :span="16">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true" @submit.prevent>
            <el-form-item prop="name" label="请求名:">
              <el-input
                style="width: 200px"
                v-model="queryParams.name"
                placeholder="请求名"
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
                  v-for="item in demandStatusList"
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
              v-hasPerm="['/demand:post']"
              type="success"
              icon="plus"
              @click="handleDemandOpt('create')"
            >
              创建
            </el-button>
          </div>
        </el-col>
      </el-row>
    </div>
    <div class="demand-main">
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
          :data="demandList"
          highlight-current-row
          border:true
        >
          <el-table-column type="expand">
            <template #default="props">
              <ProcessSteps
                :orderProcess="props.row.reviewProcess"
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
              <span v-if="item === 'status'">
                <el-tag v-if="scope.row[item] === 2" type="success">
                  {{ demandStatusMapping[scope.row[item]] }}
                </el-tag>
                <el-tag v-else-if="scope.row[item] === 3" type="danger">
                  {{ demandStatusMapping[scope.row[item]] }}
                </el-tag>
                <el-tag v-else type="primary">{{ demandStatusMapping[scope.row[item]] }}</el-tag>
              </span>
              <span v-else-if="item === 'biz'">
                <el-tag type="primary" effect="plain">{{ scope.row[item] }}</el-tag>
              </span>
              <span v-else-if="item === 'curReviewNode'">
                <el-tag type="primary" effect="plain">{{ scope.row[item].name }}</el-tag>
              </span>
              <span v-else-if="item === 'demandType'">
                <span>{{ demandTypeMapping[scope.row[item]] }}</span>
              </span>
              <span v-else>{{ scope.row[item] }}</span>
            </template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" width="320">
            <template #default="scope">
              <el-button
                type="primary"
                size="small"
                link
                icon="view"
                @click="handleToReview(scope.row, 1)"
              >
                详情
              </el-button>
              <el-button
                type="primary"
                size="small"
                link
                icon="Comment"
                :disabled="!hasEvaluate(scope.row)"
                @click="handleEvaluate(scope.row)"
              >
                评价
              </el-button>
              <el-button
                v-hasPerm="['/demand:put']"
                type="primary"
                size="small"
                link
                icon="edit"
                :disabled="hasEdit(scope.row)"
                @click="handleDemandOpt('update', scope.row)"
              >
                编辑
              </el-button>
              <el-button
                v-hasPerm="['/demand/review:put']"
                type="primary"
                size="small"
                link
                icon="stamp"
                :disabled="scope.row.hasReview === 0"
                @click="handleToReview(scope.row, 0)"
              >
                评审
              </el-button>
              <el-button
                type="danger"
                size="small"
                link
                icon="delete"
                v-hasPerm="['/demand:delete']"
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
    <DemandOpt
      v-if="demandVisible"
      v-model:visible="demandVisible"
      :action="action"
      :formData="formData"
      @submit="handleResetQuery"
    ></DemandOpt>
    <Evaluation
      v-if="evaluateVisible"
      v-model:visible="evaluateVisible"
      :formData="evaluationFormData"
      @submit="handleResetQuery"
    ></Evaluation>
  </div>
</template>

<script setup lang="ts">
import DemandAPI, {
  type DemandQuery,
  type DemandResult,
  type DemandForm,
  type EvaluateDemandForm,
} from '@/api/demand'
import { computed, onMounted, reactive, ref, watch } from 'vue'
import Pagination from '@/components/Pagination/index.vue'
import { type AxiosResponse } from 'axios'
import { indexMethod } from '@/utils/view'
import { demandStatusMapping, demandStatusList, demandTypeMapping } from '@/utils/constant'
import router from '@/router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAuthStore, useDemandStore } from '@/store'
import DemandOpt from './components/DemandOpt.vue'
import Evaluation from './components/Evaluation.vue'
import ProcessSteps from './components/ProcessSteps.vue'
import { useRoute } from 'vue-router'

defineOptions({
  name: 'DemandList',
})

const route = useRoute()
let demandStatus = route.query.status ? Number(route.query.status) : -1

const { userName } = useAuthStore()

const demandStore = useDemandStore()
const { getDemandList } = demandStore
const demandList = computed(() => demandStore.demandList)
const total = computed(() => demandStore.total)

const tableItemMapping: Map<string, string> = new Map<string, string>([
  ['name', '请求名称'],
  ['demandType', '请求类型'],
  ['biz', '业务组'],
  ['owner', '归属人'],
  ['evaluationRes', '评价'],
  ['evaluationReason', '不满意原因'],
  ['status', '状态'],
  ['curReviewNode', '当前评审节点'],
  ['updatedTime', '更新时间'],
  ['createdTime', '创建时间'],
])
const tableHead = ref<string[]>([
  'name',
  'demandType',
  'biz',
  'owner',
  'evaluationRes',
  'status',
  'curReviewNode',
  'createdTime',
])
const tableHeadOptions = ref<string[]>([
  'name',
  'demandType',
  'biz',
  'owner',
  'evaluationRes',
  'status',
  'evaluationReason',
  'curReviewNode',
  'updatedTime',
  'createdTime',
])
const checkboxVal = ref<string[]>(tableHead.value)
//ref
const queryFormRef = ref()
const loading = ref<boolean>(false)
const action = ref<string>('')
const demandVisible = ref<boolean>(false)
const evaluateVisible = ref<boolean>(false)

// reactive
const queryParams = reactive<DemandQuery>({
  page: 1,
  pageSize: 10,
})

const formData = reactive<DemandForm>({
  owner: userName,
})

const evaluationFormData = reactive<EvaluateDemandForm>({})

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
    if (demandStatus !== -1) {
      queryParams.status = demandStatus
    }
    getDemandList(queryParams)
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
  demandStatus = -1
  handleQuery()
}

function handleDelete(row: DemandResult) {
  ElMessageBox.confirm(`确定要删除工单${row?.name}吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      loading.value = true
      try {
        const id: number = row.id ? row.id : 0
        const resp: AxiosResponse = await DemandAPI.delete(id)
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

// function hasReview(demand: DemandResult): boolean {
//   const reviewStatus: number[] = [0, 1]
//   if (!reviewStatus.includes(demand.status)) {
//     return true
//   }
//   const nodeApproveRoles: Map<string, string[]> = new Map<string, string[]>([
//     ['业务组负责人审批', ['admin', 'teamLeader']],
//     ['业务组责人审批', ['admin', 'teamLeader']],
//     ['架构负责人审批', ['admin', 'architect']],
//     ['运维负责人审批', ['admin', 'ops']],
//   ])
//   const curNodeName = demand.curReviewNode.name
//   const approveRoles: string[] = nodeApproveRoles.get(curNodeName) || []
//   for (const role of approveRoles) {
//     if (roleList.value.includes(role)) {
//       return false
//     }
//   }
//   return true
// }

function hasEdit(demand: DemandResult): boolean {
  if (demand.owner !== userName) {
    return true
  }
  const editStatus: number[] = [0, 3]
  if (!editStatus.includes(demand.status)) {
    return true
  }
  return false
}

function hasEvaluate(demand: DemandResult): boolean {
  let flag = true
  if (demand.status !== 2) {
    flag = false
  }
  if (demand.isEvaluate === 1) {
    flag = false
  }

  if (demand.owner !== userName) {
    flag = false
  }
  return flag
}

function handleDemandOpt(opt: string, demand?: DemandResult) {
  action.value = opt
  if (opt === 'update') {
    Object.assign(formData, demand)
  }
  demandVisible.value = true
}

function handleEvaluate(demand: DemandResult) {
  const evaluation = demand.evaluation
  evaluationFormData.id = demand.id
  Object.assign(evaluationFormData, evaluation)
  evaluateVisible.value = true
}

function handleToReview(demand: DemandResult, isView?: number) {
  router.push({
    path: `/demand/review`,
    query: {
      demandID: demand?.id,
      demandName: demand.name,
      isView: isView,
    },
  })
}
</script>

<style lang="scss" scoped>
.demand-container {
  .demand-header {
    .btn-wrapper {
      text-align: right;
    }
  }
  .demand-main {
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
