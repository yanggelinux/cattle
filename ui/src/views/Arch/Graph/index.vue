<!--  -->
<template>
  <div class="arch-graph-container app-container">
    <div class="arch-graph-header app-header">
      <el-row>
        <el-col :span="12">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true" @submit.prevent>
            <el-form-item prop="roleName" label="图名称:">
              <el-input
                v-model="queryParams.graphName"
                placeholder="图名称"
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
              v-hasPerm="['/arch-graph:post']"
              type="success"
              icon="PictureFilled"
              @click="handleOpenDialog('create')"
            >
              创建图
            </el-button>
            <el-button
              v-hasPerm="['/arch-group:post']"
              type="primary"
              icon="Management"
              @click="handleOpenGroupDialog('create')"
            >
              创建组
            </el-button>
            <el-button icon="back" @click="handleGoBack">返回</el-button>
          </div>
        </el-col>
      </el-row>
    </div>
    <div class="arch-graph-main">
      <div class="group-wrapper" v-if="archGroupList.length > 0">
        <GroupList
          :archGroupList="archGroupList"
          :groupID="groupID"
          @handle-query="handleGroupQuery"
        ></GroupList>
      </div>
      <el-divider
        content-position="left"
        v-if="archGroupList.length > 0 && archGraphList.length > 0"
      >
        架构图列表
      </el-divider>

      <div class="graph-wrapper"></div>
      <el-space wrap :size="10">
        <el-card
          class="card-body"
          shadow="always"
          style="width: 400px"
          v-for="graph in archGraphList"
          :key="graph.id"
        >
          <template #header>
            <div class="card-header">
              <el-row :gutter="2">
                <el-col :span="14">
                  <span class="name-wrapper" @click="handleToGraphDetail(graph, 0)">
                    {{ graph.graphName }}
                  </span>
                </el-col>
                <el-col :span="10" class="tag-wrapper">
                  <span v-if="graph.graphLabel.length > 0">
                    {{ graph.graphLabel }}
                  </span>
                </el-col>
              </el-row>
            </div>
          </template>
          <el-image
            v-if="graph.imageData.length > 0"
            class="image-wrapper"
            @click="handleToGraphDetail(graph, 0)"
            :src="graph.imageData"
          />
          <el-image
            v-if="graph.imageData.length === 0"
            class="image-wrapper"
            @click="handleToGraphDetail(graph, 0)"
            :src="defaultImgSrc"
          ></el-image>

          <template #footer>
            <el-row class="card-footer">
              <el-col :span="14">
                <el-space>
                  <span>{{ graph.owner }}</span>
                  <el-tag v-if="graph.status === 2" type="success">
                    {{ graphStatusMapping[graph.status] }}
                  </el-tag>
                  <el-tag v-else-if="graph.status === 3" type="danger">
                    {{ graphStatusMapping[graph.status] }}
                  </el-tag>
                  <el-tag v-else>
                    {{ graphStatusMapping[graph.status] }}
                  </el-tag>
                </el-space>
              </el-col>
              <el-col :span="10" class="icon-wrapper">
                <el-button size="small" @click="handleToGraphDetail(graph, 1)" icon="view" circle />
                <el-button
                  v-hasPerm="['/arch-graph/copy:post']"
                  size="small"
                  @click="handleCopy(graph)"
                  icon="CopyDocument"
                  circle
                />
                <el-button
                  v-hasPerm="['/arch-graph:put']"
                  size="small"
                  @click="handleOpenDialog('update', graph)"
                  icon="edit"
                  circle
                />
                <el-button
                  v-hasPerm="['/arch-graph:delete']"
                  size="small"
                  v-if="hasDelete(graph)"
                  @click="handleDelete(graph)"
                  icon="delete"
                  circle
                />
              </el-col>
            </el-row>
          </template>
        </el-card>
        <div>
          <el-icon
            v-if="archGroupList.length === 0 || archGraphList.length > 0"
            v-hasPerm="['/arch-graph:post']"
            @click="handleOpenDialog('create')"
            class="plus-wrapper"
          >
            <Plus />
          </el-icon>
        </div>
      </el-space>
    </div>
    <GraphOpt
      v-if="dialog.visible"
      :groupID="groupID"
      :formData="formData"
      :action="action"
      :dialog="dialog"
      @submit="handleResetQuery"
    ></GraphOpt>
    <GroupOpt
      :formData="groupFormData"
      :action="groupAction"
      :dialog="groupDialog"
      :parentID="groupID"
      @submit="handleGroupQuery"
    ></GroupOpt>
  </div>
</template>

<script setup lang="ts">
import ArchGraphAPI, {
  type ArchGraphQuery,
  type ArchGraphData,
  type ArchGraphForm,
  type ArchGraphResult,
} from '@/api/arch/graph'
import ArchGroupAPI, {
  type ArchGroupQuery,
  type ArchGroupResult,
  type ArchGroupForm,
  type ArchGroupData,
} from '@/api/arch/group'
import { onMounted, reactive, ref } from 'vue'
import { type AxiosResponse } from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useRoute } from 'vue-router'
import router from '@/router'
import GraphOpt from './components/GraphOpt.vue'
import GroupList from '@/views/Arch/Group/components/GroupList.vue'
import GroupOpt from '@/views/Arch/Group/components/GroupOpt.vue'
import { graphStatusMapping } from '@/utils/constant'
import { defaultImgSrc } from './components/constant'
import { useAuthStore } from '@/store/modules/auth'
defineOptions({
  name: 'ArchGraph',
})

const route = useRoute()
const groupID = route.query.groupID ? Number(route.query.groupID) : 0

const { isSuper, userName } = useAuthStore()

const archGroupList = ref<ArchGroupResult[]>([])
const archGraphList = ref<ArchGraphResult[]>([])
const queryFormRef = ref()
const action = ref<string>('')
const groupAction = ref<string>('')
// reactive
const queryParams = reactive<ArchGraphQuery>({
  groupID: groupID,
  graphName: '',
})
// 弹窗
const dialog = reactive({
  title: '',
  visible: false,
})
const groupDialog = reactive({
  title: '',
  visible: false,
})
// 架构组表单
const formData = reactive<ArchGraphForm>({
  isShared: 0,
  status: 0,
})
const groupFormData = reactive<ArchGroupForm>({})
// 生命周期
onMounted(() => {
  handleQuery()
  handleGroupQuery()
})
// 查询
async function handleGroupQuery() {
  const parmas: ArchGroupQuery = {
    parentID: groupID,
  }
  try {
    const resp: AxiosResponse = await ArchGroupAPI.getList(parmas)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const roleData: ArchGroupData = resp.data.data
      archGroupList.value = roleData.retList
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}
async function handleQuery() {
  try {
    const resp: AxiosResponse = await ArchGraphAPI.getList(queryParams)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const archGraphData: ArchGraphData = resp.data.data
      archGraphList.value = archGraphData.retList
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
  queryParams.graphName = ''
  handleQuery()
}

function hasDelete(rowData: ArchGraphResult) {
  const status = rowData.status
  const owner = rowData.owner
  if (isSuper == 1) {
    return true
  }
  if (status === 2) {
    return false
  }
  if (owner !== userName) {
    return false
  }
  return true
}

function handleDelete(rowData: ArchGraphResult) {
  ElMessageBox.confirm(`确定要删除${rowData?.graphName}吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      try {
        const id: number = rowData.id ? rowData.id : 0
        const resp: AxiosResponse = await ArchGraphAPI.delete(id)
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
      }
    })
    .catch(() => {
      // 用户点击“取消” 或 关闭弹窗，什么都不做即可避免报错
    })
}

function handleCopy(rowData: ArchGraphResult) {
  ElMessageBox.confirm(`确定要复制${rowData?.graphName}吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      try {
        const id: number = rowData.id ? rowData.id : 0
        const resp: AxiosResponse = await ArchGraphAPI.copy(id)
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
      }
    })
    .catch(() => {
      // 用户点击“取消” 或 关闭弹窗，什么都不做即可避免报错
    })
}
// 打开弹窗
function handleOpenDialog(opt: string, rowData?: ArchGraphResult) {
  dialog.visible = true
  action.value = opt
  if (opt === 'update') {
    dialog.title = '编辑架构图'
    Object.assign(formData, rowData)
  } else {
    dialog.title = '新增架构图'
  }
}

function handleOpenGroupDialog(opt: string, rowData?: ArchGroupResult) {
  groupDialog.visible = true
  groupAction.value = opt
  if (opt === 'update') {
    groupDialog.title = '编辑架构组'
    Object.assign(formData, rowData)
  } else {
    groupDialog.title = '新增架构组'
  }
}

function handleToGraphDetail(graph: ArchGraphResult, silentMode: number) {
  router.push({
    path: `/arch/graph-draw/${graph.id}`,
    query: {
      graphID: graph.id,
      graphName: graph.graphName,
      owner: graph.owner,
      silentMode: silentMode,
      status: graph.status,
    },
  })
}
function handleGoBack() {
  router.go(-1)
}
</script>

<style lang="scss" scoped>
.arch-graph-container {
  .arch-graph-header {
    .btn-wrapper {
      text-align: right;
    }
  }
  .arch-graph-main {
    .image-wrapper {
      height: 160px;
      width: 100%;
      cursor: pointer;
    }
    .card-header {
      height: 35px;

      color: rgb(96, 98, 102);
      .name-wrapper {
        cursor: pointer;
        font-size: 15px;
      }
      .tag-wrapper {
        // line-height: 35px;
        word-break: break-word;
        font-size: 12px;
      }
    }
    .card-footer {
      height: 20px;
      // line-height: 25px;
      .icon-wrapper {
        text-align: right;
      }
    }

    .plus-wrapper {
      margin-left: 10px;
      font-size: 55px;
      cursor: pointer;
    }
  }
}
</style>
