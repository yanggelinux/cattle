<template>
  <div class="process-opt-wrapper">
    <el-dialog v-model="visible" draggable :title="title" width="500px" @close="handleClose">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="流程名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入流程名称" />
        </el-form-item>

        <el-form-item label="状态" prop="status">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.status"
            placeholder="请选择状态"
          >
            <el-option
              v-for="item in processStatusList"
              :value="item.value"
              :label="item.label"
              :key="item.value"
            />
          </el-select>
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
import ProcessAPI, { type ProcessForm, type ProcessOptResult } from '@/api/process/process'
import { reactive, ref, toRefs, type PropType } from 'vue'
import { processStatusList } from '@/utils/constant'

import { ElMessage } from 'element-plus'
import { type AxiosResponse } from 'axios'
import router from '@/router'

defineOptions({
  name: 'ProcessOpt',
})
//props
const props = defineProps({
  formData: {
    type: Object as PropType<ProcessForm>,
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
const title = ref<string>(action.value === 'create' ? '创建流程' : '编辑流程')
// reactive
const rules = reactive({
  name: [{ required: true, message: '请输入请求名称', trigger: 'blur' }],
})

//method
function handleClose() {
  visible.value = false
  formRef.value.resetFields()
  formRef.value.clearValidate()
}

// 提交表单
async function handleCreate() {
  try {
    formData.status = 1
    const processName = formData.name
    const resp: AxiosResponse = await ProcessAPI.create(formData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resultData: ProcessOptResult = resp.data.data
      const processID: number = resultData.id
      ElMessage.success('创建成功')
      handleClose()

      router.push({
        path: `/process/draw/${processID}`,
        query: {
          processID: processID,
          processName: processName,
        },
      })
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
    const resp: AxiosResponse = await ProcessAPI.update(formData)
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
