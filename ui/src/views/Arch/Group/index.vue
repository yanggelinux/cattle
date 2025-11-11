<template>
  <div class="arch-group-container app-container">
    <div class="arch-group-header app-header">
      <el-row>
        <el-col :span="12">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true" @submit.prevent>
            <el-form-item prop="roleName" label="组名称:">
              <el-input
                v-model="queryParams.groupName"
                placeholder="组名称"
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
              v-hasPerm="['/arch-group:post']"
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
    <div class="arch-group-main">
      <GroupList
        :archGroupList="archGroupList"
        :groupID="0"
        @handle-query="handleResetQuery"
      ></GroupList>
    </div>
    <GroupOpt
      :formData="formData"
      :action="action"
      :dialog="dialog"
      :parentID="0"
      @submit="handleResetQuery"
    ></GroupOpt>
  </div>
</template>

<script setup lang="ts">
import { type ArchGroupQuery, type ArchGroupForm, type ArchGroupResult } from '@/api/arch/group'
import { computed, onMounted, reactive, ref } from 'vue'
import GroupOpt from './components/GroupOpt.vue'
import GroupList from './components/GroupList.vue'
import { useArchGroupStore } from '@/store/modules/archGroup'

defineOptions({
  name: 'ArchGroup',
})

const archGroupStore = useArchGroupStore()
const { getArchGroupList } = archGroupStore
const archGroupList = computed(() => archGroupStore.archGroupList)

const queryFormRef = ref()
const action = ref<string>('')
// reactive
const queryParams = reactive<ArchGroupQuery>({
  parentID: 0,
  groupName: '',
})
// 弹窗
const dialog = reactive({
  title: '',
  visible: false,
})
// 架构组表单
const formData = reactive<ArchGroupForm>({})
// 生命周期
onMounted(() => {
  handleQuery()
})
// 查询
async function handleQuery() {
  await getArchGroupList(queryParams)
}
//重置
// 重置查询
function handleResetQuery() {
  queryFormRef.value.resetFields()
  queryParams.groupName = ''
  handleQuery()
}

// 打开弹窗
function handleOpenDialog(opt: string, rowData?: ArchGroupResult) {
  dialog.visible = true
  action.value = opt
  if (opt === 'update') {
    dialog.title = '编辑架构组'
    Object.assign(formData, rowData)
  } else {
    dialog.title = '新增架构组'
  }
}
</script>

<style lang="scss" scoped>
.arch-group-container {
  .arch-group-header {
    .btn-wrapper {
      text-align: right;
    }
  }
}
</style>
