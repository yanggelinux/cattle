<template>
  <div class="role-opt-wrapper">
    <el-dialog
      draggable
      v-model="dialog.visible"
      :title="dialog.title"
      width="500px"
      @close="handleCloseDialog"
    >
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="角色名称" prop="roleName">
          <el-input v-model="formData.roleName" placeholder="请输入角色名称" />
        </el-form-item>

        <el-form-item label="展示名称" prop="displayName">
          <el-input v-model="formData.displayName" placeholder="请输入展示名称" />
        </el-form-item>

        <el-form-item label="超级管理员" prop="isSuper">
          <el-radio-group v-model="formData.isSuper">
            <el-radio :value="1">是</el-radio>
            <el-radio :value="0">否</el-radio>
          </el-radio-group>
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
import RoleAPI, { type RoleForm } from '@/api/userPerm/role'
import { reactive, ref, toRefs, type PropType } from 'vue'
import { ElMessage } from 'element-plus'
import { type AxiosResponse } from 'axios'
defineOptions({
  name: 'RoleOpt',
})
//props
const props = defineProps({
  formData: {
    type: Object as PropType<RoleForm>,
    required: true,
    default: () => ({
      isSper: 1,
    }),
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
})
const emit = defineEmits(['submit'])
// ref
const formRef = ref()
// reactive
const rules = reactive({
  roleName: [{ required: true, message: '请输入角色名称', trigger: 'blur' }],
  displayName: [{ required: true, message: '请输入展示名称', trigger: 'blur' }],
  isSuper: [{ required: true, message: '请选择是否超管', trigger: 'blur' }],
})

const formData = props.formData
const dialog = props.dialog
// 小心「解构赋值」导致丢失响应式 用 toRefs 或 toRef 保持响应式。
const { action } = toRefs(props)

//method
function handleCloseDialog() {
  dialog.visible = false
  formRef.value.resetFields()
  formRef.value.clearValidate()
  formData.id = undefined
  formData.isSuper = 0
}

// 提交角色表单
async function handleCreate() {
  try {
    const resp: AxiosResponse = await RoleAPI.create(formData)
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
    const resp: AxiosResponse = await RoleAPI.update(formData)
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
