<template>
  <div class="group-list-wrapper">
    <el-space wrap>
      <el-card
        class="card-body"
        shadow="always"
        style="width: 300px"
        v-for="group in archGroupList"
        :key="group.id"
      >
        <el-row :gutter="5">
          <el-col @click="handleToGraph(group)" :span="16" class="name-wrapper">
            <el-badge :value="group.itemCount" color="rgb(149, 212, 117)" :offset="[15, 10]">
              <span>{{ group.groupName }}</span>
            </el-badge>
          </el-col>
          <el-col :span="8" class="icon-wrapper">
            <el-button
              v-hasPerm="['/arch-group:put']"
              size="small"
              @click="handleOpenDialog('update', group)"
              icon="edit"
              circle
            />
            <el-button
              v-hasPerm="['/arch-group:delete']"
              size="small"
              @click="handleDelete(group)"
              icon="delete"
              circle
            />
          </el-col>
        </el-row>
      </el-card>
      <div>
        <el-icon
          v-hasPerm="['/arch-group:post']"
          @click="handleOpenDialog('create')"
          class="plus-wrapper"
        >
          <Plus />
        </el-icon>
      </div>
    </el-space>
    <GroupOpt
      :formData="formData"
      :action="action"
      :dialog="dialog"
      :parentID="groupID"
      @submit="handleQuery"
    ></GroupOpt>
  </div>
</template>

<script setup lang="ts">
import ArchGroupAPI, { type ArchGroupForm, type ArchGroupResult } from '@/api/arch/group'
import { reactive, ref, toRefs, type PropType } from 'vue'
import { type AxiosResponse } from 'axios'
import GroupOpt from './GroupOpt.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import router from '@/router'

defineOptions({
  name: 'GroupList',
})

const props = defineProps({
  archGroupList: {
    type: Object as PropType<ArchGroupResult[]>,
    required: true,
    default: () => ({
      isSper: 1,
    }),
  },
  groupID: {
    type: Number,
    required: true,
  },
})
const { archGroupList, groupID } = toRefs(props)

const action = ref<string>('')
// 架构组表单
const formData = reactive<ArchGroupForm>({})
const dialog = reactive({
  title: '',
  visible: false,
})
const emit = defineEmits(['handle-query'])

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

function handleDelete(rowData: ArchGroupResult) {
  ElMessageBox.confirm(`确定要删除${rowData?.groupName}吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      try {
        const id: number = rowData.id ? rowData.id : 0
        const resp: AxiosResponse = await ArchGroupAPI.delete(id)
        const status = resp.data.status
        const msg = resp.data.msg
        if (status === 200) {
          ElMessage.success('删除成功')
          emit('handle-query')
        } else {
          console.log(msg)
          ElMessage.error(`删除失败:${msg}`)
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
function handleQuery() {
  emit('handle-query')
}
function handleToGraph(group: ArchGroupResult) {
  router.push({
    path: `/arch/graph/${group.id}`,
    query: {
      groupID: group.id,
      groupName: group.groupName,
    },
  })
}
</script>

<style lang="scss" scoped>
.group-list-wrapper {
  .name-wrapper {
    cursor: pointer;
    color: rgb(96, 98, 102);
    font-size: 15px;
  }
  // :deep(.el-card__body) {
  // }
  .icon-wrapper {
    text-align: right;
  }
  .plus-wrapper {
    margin-left: 10px;
    font-size: 35px;
    cursor: pointer;
  }
}
</style>
