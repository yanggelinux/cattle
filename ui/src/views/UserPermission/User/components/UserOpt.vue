<template>
  <div class="user-opt-wrapper">
    <el-dialog
      draggable
      v-model="dialog.visible"
      :title="dialog.title"
      width="500px"
      @close="handleCloseDialog"
    >
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="用户名称" prop="userName">
          <el-input v-model="formData.userName" placeholder="请输入用户名称" />
        </el-form-item>

        <el-form-item label="展示名称" prop="displayName">
          <el-input v-model="formData.displayName" placeholder="请输入展示名称" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input type="password" v-model="formData.password" placeholder="请输入密码" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="formData.email" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="部门" prop="deptName">
          <el-input v-model="formData.deptName" placeholder="请输入部门" />
        </el-form-item>
        <el-form-item label="角色" prop="roleIDs">
          <el-select
            clearable
            filterable
            multiple
            placement="bottom-end"
            v-model="formData.roleIDs"
            placeholder="请选择角色"
          >
            <el-option
              v-for="item in roleList"
              :value="item.id"
              :label="item.displayName"
              :key="item.id"
            />
          </el-select>
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
import UserAPI, { type UserForm } from '@/api/userPerm/user'
import { computed, onMounted, reactive, ref, toRefs, type PropType } from 'vue'
import { ElMessage } from 'element-plus'
import { useRoleStore } from '@/store/modules/role'
import { encryptPassword } from '@/utils/crypto.ts'

import { type AxiosResponse } from 'axios'
defineOptions({
  name: 'UserOpt',
})
//props
const props = defineProps({
  formData: {
    type: Object as PropType<UserForm>,
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
})
const emit = defineEmits(['submit'])

const roleStore = useRoleStore()
const { getRoleList } = roleStore
const roleList = computed(() => roleStore.roleList)
// ref
const formRef = ref()
// reactive
const rules = reactive({
  userName: [{ required: true, message: '请输入用户名称', trigger: 'blur' }],
  displayName: [{ required: true, message: '请输入展示名称', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  email: [{ required: true, message: '请输入邮箱', trigger: 'blur' }],
  deptName: [{ required: true, message: '请输入部门', trigger: 'blur' }],
  roleIDs: [{ required: true, message: '请选择角色', trigger: 'blur' }],
})

const formData = props.formData
const dialog = props.dialog
// 小心「解构赋值」导致丢失响应式 用 toRefs 或 toRef 保持响应式。
const { action } = toRefs(props)

onMounted(() => {
  getRoleList({})
})
//method
function handleCloseDialog() {
  dialog.visible = false
  formRef.value.resetFields()
  formRef.value.clearValidate()
  formData.id = undefined
}

// 提交用户表单
async function handleCreate() {
  try {
    const resp: AxiosResponse = await UserAPI.create(formData)
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
    const resp: AxiosResponse = await UserAPI.update(formData)
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

async function handleSubmit() {
  formRef.value?.validate(async (valid: boolean) => {
    if (!valid) return
    const password = formData.password
    // 进行加密操作
    if (password) {
      const { iv, cipher } = await encryptPassword(password)
      const ePassowrd = `${iv}@${cipher}`
      formData.password = ePassowrd
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
