<template>
  <div class="team-opt-wrapper">
    <el-dialog v-model="visible" draggable :title="title" width="500px" @close="handleClose">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="团队名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入团队名称" />
        </el-form-item>

        <el-form-item label="组长" prop="leader">
          <el-autocomplete
            v-model="formData.leader"
            :fetch-suggestions="querySearch"
            clearable
            placeholder="请输入组长"
          />
        </el-form-item>
        <el-form-item label="总监" prop="director">
          <el-autocomplete
            v-model="formData.director"
            :fetch-suggestions="querySearch"
            clearable
            placeholder="请输入总监"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="handleSubmit">确 定</el-button>
          <el-button @click="handleClose">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import TeamAPI, { type TeamForm } from '@/api/process/team'
import { computed, onMounted, reactive, ref, toRefs, type PropType } from 'vue'

import { ElMessage } from 'element-plus'
import { type AxiosResponse } from 'axios'
import { useUserStore } from '@/store'
import type { UserResult } from '@/api/userPerm/user'

defineOptions({
  name: 'TeamOpt',
})

const userStore = useUserStore()
const { getUserList } = userStore
const userList = computed(() => userStore.userList)

//props
const props = defineProps({
  formData: {
    type: Object as PropType<TeamForm>,
    required: true,
    default: () => ({}),
  },
  action: {
    type: String,
    required: true,
  },
})

const visible = defineModel('visible', {
  type: Boolean,
  required: true,
  default: true,
})

const formData = props.formData
// 小心「解构赋值」导致丢失响应式 用 toRefs 或 toRef 保持响应式。
const { action } = toRefs(props)
const emit = defineEmits(['submit'])
// ref
const formRef = ref()
const title = ref<string>(action.value === 'create' ? '创建团队' : '编辑团队')
// reactive
const rules = reactive({
  name: [{ required: true, message: '请输入团队名称', trigger: 'blur' }],
  leader: [{ required: true, message: '请输入组长', trigger: 'blur' }],
  director: [{ required: true, message: '请输入总监', trigger: 'blur' }],
})

onMounted(async () => {
  await getUserList({})
  // el-autocomplete 的 suggestions 需要 value:xxx的键值对
  userList.value.forEach((user) => {
    user.value = user.userName
  })
})

//method
function handleClose() {
  visible.value = false
  formRef.value.resetFields()
  formRef.value.clearValidate()
}

const querySearch = (queryString: string, cb: any) => {
  const results = queryString ? userList.value.filter(createFilter(queryString)) : userList.value
  // call callback function to return suggestions
  cb(results)
}
const createFilter = (queryString: string) => {
  return (user: UserResult) => {
    return user.userName.toLowerCase().indexOf(queryString.toLowerCase()) === 0
  }
}

// 提交表单
async function handleCreate() {
  try {
    const resp: AxiosResponse = await TeamAPI.create(formData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('创建成功')
      handleClose()

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
    const resp: AxiosResponse = await TeamAPI.update(formData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('编辑成功')
      handleClose()
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
    if (action.value === 'update') {
      handleUpdate()
    } else {
      handleCreate()
    }
  })
}
</script>
<style lang="scss" scoped></style>
