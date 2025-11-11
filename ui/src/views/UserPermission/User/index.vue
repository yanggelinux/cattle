<template>
  <div class="user-container app-container">
    <div class="user-header app-header">
      <el-row>
        <el-col :span="12">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true" @submit.prevent>
            <el-form-item prop="userName" label="用户名称:">
              <el-input
                v-model="queryParams.userName"
                placeholder="用户"
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
              v-hasPerm="['/user:post']"
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
    <div class="user-main">
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
          :data="userList"
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
              <span v-if="item === 'origin'">{{ userOriginMapping[scope.row[item]] }}</span>
              <span v-else>{{ scope.row[item] }}</span>
            </template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" width="220">
            <template #default="scope">
              <el-button
                v-hasPerm="['/user:put']"
                type="primary"
                size="small"
                link
                icon="edit"
                @click="handleOpenDialog('update', scope.row)"
              >
                编辑
              </el-button>
              <el-button
                v-hasPerm="['/user:delete']"
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
    <!-- 用户表单弹窗 -->
    <UserOpt
      :formData="formData"
      :action="action"
      :dialog="dialog"
      @submit="handleResetQuery"
    ></UserOpt>
  </div>
</template>

<script setup lang="ts">
import UserAPI, { type UserQuery, type UserResult, type UserForm } from '@/api/userPerm/user'
import { computed, onMounted, reactive, ref, watch } from 'vue'
import Pagination from '@/components/Pagination/index.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { type AxiosResponse } from 'axios'
import { useUserStore } from '@/store/modules/user'
import UserOpt from './components/UserOpt.vue'
import { userOriginMapping } from '@/utils/constant'

import { indexMethod } from '@/utils/view'

defineOptions({
  name: 'User',
})

const userStore = useUserStore()
const { getUserList } = userStore
const userList = computed(() => userStore.userList)
const total = computed(() => userStore.total)

const tableItemMapping: Map<string, string> = new Map<string, string>([
  ['userName', '用户名'],
  ['displayName', '展示名称'],
  ['email', '邮箱'],
  ['deptName', '部门'],
  ['roleNames', '角色'],
  ['origin', '来源'],
  ['lastLoginTime', '最后登录时间'],
  ['updatedTime', '更新时间'],
  ['createdTime', '创建时间'],
])
const tableHead = ref<string[]>([
  'userName',
  'displayName',
  'email',
  'deptName',
  'roleNames',
  'origin',
  'lastLoginTime',
  'createdTime',
])
const tableHeadOptions = ref<string[]>([
  'userName',
  'displayName',
  'email',
  'deptName',
  'roleNames',
  'origin',
  'lastLoginTime',
  'updatedTime',
  'createdTime',
])
const checkboxVal = ref<string[]>(tableHead.value)

//ref
const queryFormRef = ref()
const loading = ref<boolean>(false)
// 编辑还是新增
const action = ref<string>('')

// reactive
const queryParams = reactive<UserQuery>({
  page: 1,
  pageSize: 10,
})

// 弹窗
const dialog = reactive({
  title: '',
  visible: false,
})
// 用户表单
const formData = reactive<UserForm>({})
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
  await getUserList(queryParams)
}
//重置
// 重置查询
function handleResetQuery() {
  queryFormRef.value.resetFields()
  queryParams.page = 1
  queryParams.pageSize = 10
  queryParams.userName = ''
  handleQuery()
}

// 打开用户弹窗
async function handleOpenDialog(opt: string, rowData?: UserResult) {
  dialog.visible = true
  action.value = opt
  if (opt === 'update') {
    dialog.title = '编辑用户'
    Object.assign(formData, rowData)
    // const ePassword = rowData?.password
    // if (ePassword) {
    //   const ePasswords = ePassword.split('@')
    //   let password = ''
    //   if (ePasswords.length == 2) {
    //     const iv = ePasswords[0]
    //     const cipher = ePasswords[1]
    //     password = await decryptPassword(iv, cipher)
    //   }
    //   formData.password = password
    // }
  } else {
    dialog.title = '新增用户'
  }
}

function handleDelete(rowData: UserResult) {
  ElMessageBox.confirm(`确定要删除${rowData?.userName}吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      loading.value = true
      try {
        const id: number = rowData.id ? rowData.id : 0
        const resp: AxiosResponse = await UserAPI.delete(id)
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
</script>

<style lang="scss" scoped>
.user-container {
  .user-header {
    .btn-wrapper {
      text-align: right;
    }
  }
  .user-main {
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
