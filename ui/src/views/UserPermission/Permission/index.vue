<template>
  <div class="permission-container app-container">
    <div class="permission-header app-header">
      <el-row>
        <el-col :span="12">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true" @submit.prevent>
            <el-form-item label="" prop="name">
              <el-tag>权限管理</el-tag>
            </el-form-item>
          </el-form>
        </el-col>
        <el-col :span="12">
          <div class="btn-wrapper">
            <!-- <el-button type="primary" icon="search" @click="handleQuery">查询</el-button> -->
            <!-- <el-button icon="refresh" @click="handleResetQuery">重置</el-button> -->
            <el-button
              v-hasPerm="['/permission:post']"
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

    <el-card shadow="never">
      <el-table
        v-loading="loading"
        :data="PermissionList"
        highlight-current-row
        row-key="id"
        :tree-props="{
          children: 'children',
          hasChildren: 'hasChildren',
        }"
      >
        <el-table-column label="权限名称">
          <template #default="scope">
            <span :class="`i-svg:${scope.row.permType === 1 ? 'menu_menu' : 'api'}`" />
            {{ scope.row.name }}
          </template>
        </el-table-column>

        <el-table-column label="类型" align="center" width="80">
          <template #default="scope">
            <el-tag v-if="scope.row.permType === 1" type="warning">菜单</el-tag>
            <el-tag v-if="scope.row.permType === 2" type="success">API</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="权限编码" align="left" width="150" prop="code" />
        <el-table-column label="权限URI" align="left" width="150" prop="uri" />
        <el-table-column label="权限方法" align="left" width="250" prop="method" />
        <el-table-column label="状态" align="center" width="80">
          <template #default="scope">
            <el-tag v-if="scope.row.isEnabled === 1" type="success">启用</el-tag>
            <el-tag v-else type="info">禁用</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="排序" align="left" width="100" prop="sort" />
        <el-table-column fixed="right" align="center" label="操作" width="220">
          <template #default="scope">
            <el-button
              v-if="scope.row.permType == 1"
              v-hasPerm="['/permission:post']"
              type="primary"
              link
              size="small"
              icon="plus"
              @click.stop="handleOpenDialog('create', scope.row)"
            >
              新增
            </el-button>

            <el-button
              v-hasPerm="['/permission:put']"
              type="primary"
              link
              size="small"
              icon="edit"
              @click.stop="handleOpenDialog('update', scope.row)"
            >
              编辑
            </el-button>
            <el-button
              v-hasPerm="['/permission:delete']"
              type="danger"
              link
              size="small"
              icon="delete"
              @click.stop="handleDelete(scope.row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <PermissionOpt
      v-model:curPermType="curPermType"
      :menuOptions="menuOptions"
      :formData="formData"
      :action="action"
      :dialog="dialog"
      @submit="handleResetQuery"
    ></PermissionOpt>
  </div>
</template>

<script setup lang="ts">
defineOptions({
  name: 'Permission',
  inheritAttrs: false,
})

import PermissionAPI, {
  type PermQuery,
  type PermForm,
  type PermResult,
  type PermData,
} from '@/api/userPerm/permission'
import type { AxiosResponse } from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'
import { onMounted, reactive, ref } from 'vue'
import PermissionOpt from './components/PermissionOpt.vue'
import { type OptionType } from '@/types/global.d'

const queryFormRef = ref()
const loading = ref(false)
// 菜单表格数据
const PermissionList = ref<PermResult[]>([])
// 顶级菜单下拉选项
const menuOptions = ref<OptionType[]>([])
const curPermType = ref<number | undefined>(1)

const action = ref<string>('')

const dialog = reactive({
  title: '新增菜单',
  visible: false,
})
// 查询参数
const queryParams = reactive<PermQuery>({
  project: 'cattle',
})
// 菜单表单数据
const formData = reactive<PermForm>({
  parentID: 0,
  name: '',
  code: '',
  permType: 1,
  isEnabled: 1,
})
// 表单验证规则

onMounted(() => {
  handleQuery()
})

// 查询菜单
async function handleQuery() {
  try {
    loading.value = true
    const resp: AxiosResponse = await PermissionAPI.getList(queryParams)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const respData: PermData = resp.data.data
      PermissionList.value = respData.retList
      menuOptions.value = [{ value: 0, label: '顶级菜单', children: filterTree(respData.retList) }]
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  } finally {
    loading.value = false
  }
}
function filterTree(data: any) {
  return data
    .filter((item: any) => item.permType !== 2)
    .map((item: any) => {
      const newItem: OptionType = {
        value: item.id,
        label: item.name,
      }

      if (item.children) {
        const filteredChildren = filterTree(item.children)
        if (filteredChildren.length > 0) {
          newItem.children = filteredChildren
        }
      }
      return newItem
    })
}

// 重置查询
function handleResetQuery() {
  handleQuery()
}
/**
 * 打开表单弹窗
 *
 * @param parentID 父菜单ID
 * @param permID 菜单ID
 */
function handleOpenDialog(opt: string, rowData?: PermResult) {
  dialog.visible = true
  action.value = opt
  if (opt === 'update') {
    dialog.title = '编辑权限'
    const ptype = rowData?.permType
    curPermType.value = ptype
    Object.assign(formData, rowData)
  } else {
    dialog.title = '新增权限'
    curPermType.value = 1
    formData.permType = 1
    if (rowData) {
      formData.parentID = rowData.id
    }
  }
}

// 删除菜单
function handleDelete(rowData: PermForm) {
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
        const resp: AxiosResponse = await PermissionAPI.delete(id)
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
.permission-container {
  .permission-header {
    .btn-wrapper {
      text-align: right;
    }
  }
}
</style>
