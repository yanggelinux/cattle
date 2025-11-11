<template>
  <div class="team-container app-container">
    <div class="team-header app-header">
      <el-row>
        <el-col :span="16">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true" @submit.prevent>
            <el-form-item prop="name" label="团队名:">
              <el-input
                style="width: 200px"
                v-model="queryParams.name"
                placeholder="团队名"
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
            <el-button
              v-hasPerm="['/team:post']"
              type="success"
              icon="plus"
              @click="handleTeamOpt('create')"
            >
              创建
            </el-button>
          </div>
        </el-col>
      </el-row>
    </div>
    <div class="team-main">
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
          :data="teamList"
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
              <span>{{ scope.row[item] }}</span>
            </template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" width="350">
            <template #default="scope">
              <el-button
                v-hasPerm="['/team:put']"
                type="primary"
                size="small"
                link
                icon="edit"
                @click="handleTeamOpt('update', scope.row)"
              >
                编辑
              </el-button>
              <el-button
                type="danger"
                size="small"
                link
                icon="delete"
                v-hasPerm="['/team:delete']"
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
    <TeamOpt
      v-if="teamVisible"
      v-model:visible="teamVisible"
      :action="action"
      :formData="formData"
      @submit="handleResetQuery"
    ></TeamOpt>
  </div>
</template>

<script setup lang="ts">
import TeamAPI, {
  type TeamQuery,
  type TeamResult,
  type TeamData,
  type TeamForm,
} from '@/api/process/team'
import { onMounted, reactive, ref, watch } from 'vue'
import Pagination from '@/components/Pagination/index.vue'
import { type AxiosResponse } from 'axios'
import { indexMethod } from '@/utils/view'
import { ElMessage, ElMessageBox } from 'element-plus'
import TeamOpt from './components/TeamOpt.vue'

defineOptions({
  name: 'TeamList',
})

const tableItemMapping: Map<string, string> = new Map<string, string>([
  ['name', '组名称'],
  ['leader', '组长'],
  ['director', '总监'],
  ['updatedTime', '更新时间'],
  ['createdTime', '创建时间'],
])
const tableHead = ref<string[]>(['name', 'leader', 'director', 'createdTime'])
const tableHeadOptions = ref<string[]>(['name', 'leader', 'director', 'createdTime'])
const checkboxVal = ref<string[]>(tableHead.value)
//ref
const queryFormRef = ref()
const loading = ref<boolean>(false)
const total = ref<number>(0)
const teamList = ref<TeamResult[]>([])
const action = ref<string>('')
const teamVisible = ref<boolean>(false)

// reactive
const queryParams = reactive<TeamQuery>({
  page: 1,
  pageSize: 10,
})

const formData = reactive<TeamForm>({})

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
    const resp: AxiosResponse = await TeamAPI.getList(queryParams)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: TeamData = resp.data.data
      teamList.value = resData.retList
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

function handleDelete(row: TeamResult) {
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
        const resp: AxiosResponse = await TeamAPI.delete(id)
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

function handleTeamOpt(opt: string, team?: TeamResult) {
  action.value = opt
  if (opt === 'update') {
    Object.assign(formData, team)
  }
  teamVisible.value = true
}
</script>

<style lang="scss" scoped>
.team-container {
  .team-header {
    .btn-wrapper {
      text-align: right;
    }
  }
  .team-main {
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
