<template>
  <div class="opt-wrapper">
    <el-dialog draggable v-model="visible" :title="title" width="600px" @close="handleCloseDialog">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="工单组名" prop="name">
          <el-input v-model="formData.name" placeholder="请输入工单组名称" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number
            v-model="formData.sort"
            style="width: 100px"
            controls-position="right"
            :min="0"
          />
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
          <el-button @click="handleCloseDialog">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import OrderGroupAPI, { type OrderGroupForm } from '@/api/order/orderGroup'
import { reactive, ref, toRefs, type PropType } from 'vue'
import { ElMessage } from 'element-plus'
import { processStatusList } from '@/utils/constant'
import { type AxiosResponse } from 'axios'
defineOptions({
  name: 'OrderGroupOpt',
})
//props
const props = defineProps({
  formData: {
    type: Object as PropType<OrderGroupForm>,
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

const emit = defineEmits(['submit'])
const formData = props.formData
// 小心「解构赋值」导致丢失响应式 用 toRefs 或 toRef 保持响应式。
const { action } = toRefs(props)
// ref
const formRef = ref()
const title = ref<string>(action.value === 'create' ? '创建工单组' : '编辑工单组')
// reactive
const rules = reactive({
  name: [{ required: true, message: '请输入工单组名称', trigger: 'blur' }],
  sort: [{ required: true, message: '请输入排序', trigger: 'blur' }],
  status: [{ required: true, message: '请输入状态', trigger: 'blur' }],
})

//method
function handleCloseDialog() {
  visible.value = false
  formRef.value.resetFields()
  formRef.value.clearValidate()
  formData.id = undefined
}

// 提交表单
async function handleCreate() {
  try {
    const resp: AxiosResponse = await OrderGroupAPI.create(formData)
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
    const resp: AxiosResponse = await OrderGroupAPI.update(formData)
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
    if (action.value === 'update') {
      handleUpdate()
    } else {
      handleCreate()
    }
  })
}
</script>
<style lang="scss" scoped></style>
