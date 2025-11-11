<template>
  <div class="ArchGraph-opt-wrapper">
    <el-dialog
      draggable
      v-model="dialog.visible"
      :title="dialog.title"
      width="500px"
      @close="handleCloseDialog"
    >
      <el-form ref="formRef" label-width="80px" :model="formData" :rules="rules">
        <el-form-item label-position="right" label="图名称" prop="graphName">
          <el-input v-model="formData.graphName" placeholder="请输入图名称" />
        </el-form-item>

        <el-form-item label-position="right" label="归属人" prop="owner" v-if="action === 'update'">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.owner"
            placeholder="请选择归属人"
          >
            <el-option
              v-for="item in userList"
              :value="item.userName"
              :label="item.userName"
              :key="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label-position="right"
          label="图组"
          prop="groupID"
          :required="true"
          v-if="action === 'update'"
        >
          <el-cascader
            :props="cascaderProps"
            v-model="groupIDs"
            :options="archGroupTree"
            style="width: 400px"
          />
        </el-form-item>
        <el-form-item label="是否共享" prop="isShared">
          <el-radio-group v-model="formData.isShared">
            <el-radio :value="1" size="small" border>共享</el-radio>
            <el-radio :value="0" size="small" border>不共享</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="isSuper === 1" label-position="right" label="状态" prop="status">
          <el-select v-model="formData.status" placeholder="请选择状态">
            <el-option
              v-for="item in graphStatusList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label-position="right" label="图标签" prop="graphLabel">
          <el-input
            v-model="formData.graphLabel"
            placeholder="图标签可为空，输入标签确保全局唯一"
          />
        </el-form-item>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="handleSubmit">确 定</el-button>
          <el-button @click="handleCloseDialog">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import ArchGraphAPI, { type ArchGraphForm, type ArchGraphOptResult } from '@/api/arch/graph'
import { computed, onMounted, reactive, ref, toRefs, type PropType } from 'vue'
import { ElMessage } from 'element-plus'
import { type AxiosResponse } from 'axios'
import router from '@/router'
import { useUserStore } from '@/store/modules/user'
import { useAuthStore } from '@/store/modules/auth'
import { useArchGroupStore } from '@/store/modules/archGroup'
import { graphStatusList } from '@/utils/constant'

defineOptions({
  name: 'GraphOpt',
})

//props
const props = defineProps({
  formData: {
    type: Object as PropType<ArchGraphForm>,
    required: true,
    default: () => ({}),
  },
  dialog: {
    type: Object,
    required: true,
    default: () => ({
      title: '',
      visible: false,
    }),
  },
  action: {
    type: String,
    required: true,
  },
  groupID: {
    type: Number,
    required: true,
  },
})
const emit = defineEmits(['submit'])
const userStore = useUserStore()
const { getUserList } = userStore
const { userName, isSuper } = useAuthStore()
const userList = computed(() => userStore.userList)

const archGroupStore = useArchGroupStore()
const { getArchGroupTree, findPathByValue } = archGroupStore
const archGroupTree = computed(() => archGroupStore.archGroupTree)
console.log(archGroupStore.archGroupTree)
const cascaderProps = {
  showPrefix: false,
  checkStrictly: true,
  checkOnClickNode: true,
}
// ref
const formRef = ref()
const groupIDs = ref<number[]>([])
// reactive
const rules = reactive({
  graphName: [{ required: true, message: '请输入图名称', trigger: 'blur' }],
  owner: [{ required: true, message: '请选择归属人', trigger: 'blur' }],
  // groupID: [{ required: true, message: '请选择图组', trigger: 'blur' }],
})

const formData = props.formData
const dialog = props.dialog
// 小心「解构赋值」导致丢失响应式 用 toRefs 或 toRef 保持响应式。
const { action } = toRefs(props)
const groupID = props.groupID

onMounted(async () => {
  if (action.value === 'update' && formData.groupID) {
    groupIDs.value = findPathByValue(archGroupStore.archGroupTree, formData.groupID)
  }
  await getUserList({})
  await getArchGroupTree({})
})
//method
function handleCloseDialog() {
  dialog.visible = false
  formRef.value.resetFields()
  formRef.value.clearValidate()
  formData.id = undefined
}

// 提交角色表单
async function handleCreate() {
  try {
    formData.groupID = Number(groupID)
    const { graphName, status } = formData
    formData.owner = userName
    const resp: AxiosResponse = await ArchGraphAPI.create(formData)
    const statusCode = resp.data.status
    const msg = resp.data.msg
    if (statusCode === 200) {
      const resultData: ArchGraphOptResult = resp.data.data
      const graphID: number = resultData.id
      ElMessage.success('创建成功')
      handleCloseDialog()
      // 创建成功后跳转到图绘制页面
      router.push({
        path: `/arch/graph-draw/${graphID}`,
        query: {
          graphID: graphID,
          graphName: graphName,
          owner: userName,
          status: status,
        },
      })
      emit('submit')
    } else {
      if (status === 40005) {
        ElMessage.error('创建失败,图标签重复')
        return
      }
      console.log(msg)
      ElMessage.error(`创建失败${msg}`)
    }
  } catch (err) {
    console.error(err)
    ElMessage.error('创建失败')
  }
}

async function handleUpdate() {
  try {
    // formData.groupID = Number(groupID)
    const resp: AxiosResponse = await ArchGraphAPI.update(formData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('编辑成功')
      handleCloseDialog()
      emit('submit')
    } else {
      if (status === 40005) {
        ElMessage.error('编辑失败,图标签重复')
        return
      }
      console.log(msg)
      ElMessage.error('编辑失败')
    }
  } catch (err) {
    console.error(err)
    ElMessage.error('编辑失败')
  }
}

function handleSubmit() {
  formRef.value?.validate(async (valid: boolean) => {
    if (!valid) return
    if (groupIDs.value.length > 0) {
      const groupID = groupIDs.value[groupIDs.value.length - 1]
      formData.groupID = groupID
    }
    if (action.value === 'update') {
      handleUpdate()
    } else {
      handleCreate()
    }
  })
}
</script>
<style lang="scss" scoped></style>
