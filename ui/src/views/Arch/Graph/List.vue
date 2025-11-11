<template>
  <div class="graph-list-container app-container">
    <div class="graph-list-header app-header">
      <el-row>
        <el-col :span="16">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true" @submit.prevent>
            <el-form-item prop="graphName" label="架构图名:">
              <el-input
                style="width: 200px"
                v-model="queryParams.graphName"
                placeholder="架构图名"
                clearable
                @keyup.enter="handleQuery"
              />
            </el-form-item>
            <el-form-item label-position="right" label="架构图组" prop="groupID">
              <el-cascader
                :props="cascaderProps"
                v-model="groupIDs"
                :options="archGroupTree"
                @change="handleQuery"
                style="width: 400px"
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
    <div class="graph-list-main">
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
          :data="archGraphList"
          highlight-current-row
          bgraph:true
        >
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
                  {{ graphStatusMapping[scope.row[item]] }}
                </el-tag>
                <el-tag v-else-if="scope.row[item] === 3" type="danger">
                  {{ graphStatusMapping[scope.row[item]] }}
                </el-tag>
                <el-tag v-else type="primary">{{ graphStatusMapping[scope.row[item]] }}</el-tag>
              </span>
              <span v-else-if="item === 'groupName'" class="click">
                <el-tag @click="handleToGraph(scope.row)" type="primary" effect="plain">
                  {{ scope.row[item] }}
                </el-tag>
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
                @click="handleToView(scope.row, 1)"
              >
                申请工单
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
import ArchGraphAPI, {
  type ArchGraphQuery,
  type ArchGraphResult,
  type ArchGraphData,
} from '@/api/arch/graph'
import { computed, onMounted, reactive, ref, watch } from 'vue'
import Pagination from '@/components/Pagination/index.vue'
import { type AxiosResponse } from 'axios'
import { indexMethod } from '@/utils/view'
import { graphStatusMapping } from '@/utils/constant'
import router from '@/router'
import { useArchGroupStore } from '@/store'

defineOptions({
  name: 'ArchGraphList',
})

const archGroupStore = useArchGroupStore()
const { getArchGroupTree, getNodePathMap } = archGroupStore
const archGroupTree = computed(() => archGroupStore.archGroupTree)

const tableItemMapping: Map<string, string> = new Map<string, string>([
  ['graphName', '架构图名称'],
  ['groupName', '架构图组'],
  ['groupPath', '架构组路径'],
  ['owner', '归属人'],
  ['status', '状态'],
  ['updatedTime', '更新时间'],
  ['createdTime', '创建时间'],
])
const tableHead = ref<string[]>([
  'graphName',
  'groupName',
  'groupPath',
  'owner',
  'status',
  'createdTime',
])
const tableHeadOptions = ref<string[]>([
  'graphName',
  'groupName',
  'groupPath',
  'owner',
  'status',
  'updatedTime',
  'createdTime',
])
const cascaderProps = {
  showPrefix: false,
  checkStrictly: true,
  checkOnClickNode: true,
}
const checkboxVal = ref<string[]>(tableHead.value)
//ref
const queryFormRef = ref()
const loading = ref<boolean>(false)
const total = ref<number>(0)
const archGraphList = ref<ArchGraphResult[]>([])
const groupIDs = ref<number[]>([])
const groupPath = ref<Map<number, string>>(new Map<number, string>())

// reactive
const queryParams = reactive<ArchGraphQuery>({
  page: 1,
  pageSize: 10,
  status: 2,
})

// 生命周期
onMounted(async () => {
  await getArchGroupTree({})
  genGroupPath()
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
function genGroupPath() {
  const results = getNodePathMap(archGroupTree.value)
  results.forEach((path, value) => {
    groupPath.value.set(value, '/' + path.join('/'))
  })
}

async function handleQuery() {
  try {
    if (groupIDs.value.length > 0) {
      const groupID = groupIDs.value[groupIDs.value.length - 1]
      queryParams.groupID = groupID
    }
    const resp: AxiosResponse = await ArchGraphAPI.getList(queryParams)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      archGraphList.value = []
      const resData: ArchGraphData = resp.data.data
      const retList = resData.retList
      for (const ret of retList) {
        ret.groupPath = groupPath.value.get(ret.groupID) || ''
        archGraphList.value.push(ret)
      }
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
  groupIDs.value = []
  queryParams.page = 1
  queryParams.pageSize = 10
  queryParams.status = 2
  handleQuery()
}

function handleToGraph(graph: ArchGraphResult) {
  router.push({
    path: `/arch/graph/${graph.groupID}`,
    query: {
      groupID: graph.groupID,
      groupName: graph.groupName,
    },
  })
}
function handleToView(graph: ArchGraphResult, silentMode?: number) {
  router.push({
    path: `/arch/graph-draw/${graph.id}`,
    query: {
      graphID: graph.id,
      graphName: graph.graphName,
      owner: graph.owner,
      silentMode: silentMode,
      status: graph.status,
      hasOrder: 1,
    },
  })
}
</script>

<style lang="scss" scoped>
.graph-list-container {
  .graph-list-header {
    .btn-wrapper {
      text-align: right;
    }
  }
  .graph-list-main {
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
    .click {
      cursor: pointer;
    }
  }
}
</style>
