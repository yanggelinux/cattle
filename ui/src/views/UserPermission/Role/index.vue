<template>
  <div class="role-container app-container">
    <div class="role-header app-header">
      <el-row>
        <el-col :span="12">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true" @submit.prevent>
            <el-form-item prop="roleName" label="角色名称:">
              <el-input
                v-model="queryParams.roleName"
                placeholder="角色"
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
              v-hasPerm="['/role:post']"
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
    <div class="role-main">
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
          :data="roleList"
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
              <span v-if="item === 'isSuper'">
                <el-tag v-if="scope.row.isSuper === 1" type="success">是</el-tag>
                <el-tag v-else type="info">否</el-tag>
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
                icon="position"
                @click="handleOpenAssignPermDialog(scope.row)"
              >
                分配权限
              </el-button>
              <el-button
                v-hasPerm="['/role:put']"
                type="primary"
                size="small"
                link
                icon="edit"
                @click="handleOpenDialog('update', scope.row)"
              >
                编辑
              </el-button>
              <el-button
                v-hasPerm="['/role:delete']"
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
    <!-- 角色表单弹窗 -->
    <RoleOpt
      :formData="formData"
      :action="action"
      :dialog="dialog"
      @submit="handleResetQuery"
    ></RoleOpt>

    <!-- 分配权限弹窗 -->
    <AssignPerm
      v-model:permVisible="assignPermVisible"
      :checkedRole="checkedRole"
      @submit="handleResetQuery"
    ></AssignPerm>
  </div>
</template>

<script setup lang="ts">
import RoleAPI, { type RoleQuery, type RoleResult, type RoleForm } from '@/api/userPerm/role'
import { computed, onMounted, reactive, ref, watch } from 'vue'
import Pagination from '@/components/Pagination/index.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { type AxiosResponse } from 'axios'
import { useRoleStore } from '@/store/modules/role'
import RoleOpt from './components/RoleOpt.vue'
import AssignPerm, { type CheckedRole } from './components/AssignPerm.vue'
import { indexMethod } from '@/utils/view'

defineOptions({
  name: 'Role',
})

const roleStore = useRoleStore()
const { getRoleList } = roleStore
const roleList = computed(() => roleStore.roleList)
const total = computed(() => roleStore.total)

const tableItemMapping: Map<string, string> = new Map<string, string>([
  ['roleName', '角色名'],
  ['displayName', '展示名称'],
  ['isSuper', '超管'],
  ['updatedTime', '更新时间'],
  ['createdTime', '创建时间'],
])
const tableHead = ref<string[]>(['roleName', 'displayName', 'isSuper', 'createdTime'])
const tableHeadOptions = ref<string[]>([
  'roleName',
  'displayName',
  'isSuper',
  'updatedTime',
  'createdTime',
])
const checkboxVal = ref<string[]>(tableHead.value)
//ref
const queryFormRef = ref()
const loading = ref<boolean>(false)
// 编辑还是新增
const action = ref<string>('')
const assignPermVisible = ref<boolean>(false)

// reactive
const queryParams = reactive<RoleQuery>({
  page: 1,
  pageSize: 10,
})

const checkedRole = reactive<CheckedRole>({
  id: 0,
  displayName: '',
  isSuper: 0,
})

// 弹窗
const dialog = reactive({
  title: '',
  visible: false,
})
// 角色表单
const formData = reactive<RoleForm>({
  isSuper: 0,
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
  await getRoleList(queryParams)
}
//重置
// 重置查询
function handleResetQuery() {
  queryFormRef.value.resetFields()
  queryParams.page = 1
  queryParams.pageSize = 10
  queryParams.roleName = ''
  handleQuery()
}

// 打开角色弹窗
function handleOpenDialog(opt: string, rowData?: RoleResult) {
  dialog.visible = true
  action.value = opt
  if (opt === 'update') {
    dialog.title = '编辑角色'
    Object.assign(formData, rowData)
  } else {
    dialog.title = '新增角色'
  }
}

function handleDelete(rowData: RoleResult) {
  ElMessageBox.confirm(`确定要删除${rowData?.roleName}吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      loading.value = true
      try {
        const id: number = rowData.id ? rowData.id : 0
        const resp: AxiosResponse = await RoleAPI.delete(id)
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
// 分配权限
// 打开分配菜单权限弹窗
async function handleOpenAssignPermDialog(row: RoleResult) {
  const roleID = row.id
  if (roleID) {
    assignPermVisible.value = true
    checkedRole.id = roleID
    checkedRole.displayName = row.displayName
    checkedRole.isSuper = row.isSuper
  }
}
</script>

<style lang="scss" scoped>
.role-container {
  .role-header {
    .btn-wrapper {
      text-align: right;
    }
  }
  .role-main {
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
