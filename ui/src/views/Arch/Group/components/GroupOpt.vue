<template>
  <div class="ArchGroup-opt-wrapper">
    <el-dialog
      draggable
      v-model="dialog.visible"
      :title="dialog.title"
      width="500px"
      @close="handleCloseDialog"
    >
      <el-form ref="formRef" :model="formData" :rules="rules">
        <el-form-item label-position="left" label-width="65px" label="组名称" prop="groupName">
          <el-input v-model="formData.groupName" placeholder="请输入组名称" />
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
import ArchGroupAPI, { type ArchGroupForm } from '@/api/arch/group'
import { reactive, ref, toRefs, type PropType } from 'vue'
import { ElMessage } from 'element-plus'
import { type AxiosResponse } from 'axios'
defineOptions({
  name: 'GroupOpt',
})
//props
const props = defineProps({
  formData: {
    type: Object as PropType<ArchGroupForm>,
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
  parentID: {
    type: Number,
    required: true,
  },
})
const emit = defineEmits(['submit'])
// ref
const formRef = ref()
// reactive
const rules = reactive({
  groupName: [{ required: true, message: '请输入组名称', trigger: 'blur' }],
})

const formData = props.formData
const dialog = props.dialog
// 小心「解构赋值」导致丢失响应式 用 toRefs 或 toRef 保持响应式。
const { action, parentID } = toRefs(props)

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
    formData.parentID = parentID.value
    const resp: AxiosResponse = await ArchGroupAPI.create(formData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('创建成功')
      handleCloseDialog()
      emit('submit')
    } else {
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
    formData.parentID = parentID.value
    const resp: AxiosResponse = await ArchGroupAPI.update(formData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('编辑成功')
      handleCloseDialog()
      emit('submit')
    } else {
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
    if (action.value === 'update') {
      handleUpdate()
    } else {
      handleCreate()
    }
  })
}
</script>
<style lang="scss" scoped></style>
