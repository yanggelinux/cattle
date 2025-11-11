<template>
  <div class="evaluation-wrapper">
    <el-drawer v-model="visible" :title="title" size="40%" @close="handleClose">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="应用运维组" prop="opsEvaluation">
          <el-select
            clearable
            placement="bottom-end"
            v-model="formData.opsEvaluation"
            placeholder="应用运维组"
          >
            <el-option v-for="item in evaluationList" :value="item" :label="item" :key="item" />
          </el-select>
        </el-form-item>
        <el-form-item
          v-if="formData.opsEvaluation === '不满意'"
          :required="formData.opsEvaluation === '不满意'"
          label-position="right"
          label="不满意原因:"
          prop="opsReason"
        >
          <el-input
            style="width: 1030px"
            :rows="3"
            type="textarea"
            v-model="formData.opsReason"
            placeholder="不满意原因"
          />
        </el-form-item>
        <el-form-item label="资源管理组" prop="resEvaluation">
          <el-select
            clearable
            placement="bottom-end"
            v-model="formData.resEvaluation"
            placeholder="资源管理组"
          >
            <el-option v-for="item in evaluationList" :value="item" :label="item" :key="item" />
          </el-select>
        </el-form-item>
        <el-form-item
          v-if="formData.resEvaluation === '不满意'"
          :required="formData.resEvaluation === '不满意'"
          label-position="right"
          label="不满意原因:"
          prop="resReason"
        >
          <el-input
            style="width: 1030px"
            :rows="3"
            type="textarea"
            v-model="formData.resReason"
            placeholder="不满意原因"
          />
        </el-form-item>
        <el-form-item label="安全支撑组" prop="netEvaluation">
          <el-select
            clearable
            placement="bottom-end"
            v-model="formData.netEvaluation"
            placeholder="安全支撑组"
          >
            <el-option v-for="item in evaluationList" :value="item" :label="item" :key="item" />
          </el-select>
        </el-form-item>
        <el-form-item
          v-if="formData.netEvaluation === '不满意'"
          :required="formData.netEvaluation === '不满意'"
          label-position="right"
          label="不满意原因:"
          prop="netReason"
        >
          <el-input
            style="width: 1030px"
            :rows="3"
            type="textarea"
            v-model="formData.netReason"
            placeholder="不满意原因"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="handleSubmit">确 定</el-button>
          <el-button @click="handleClose">取 消</el-button>
        </div>
      </template>
    </el-drawer>
  </div>
</template>
<script setup lang="ts">
import DemandAPI, { type EvaluateDemandForm } from '@/api/demand'
import { reactive, ref, type PropType } from 'vue'
import { ElMessage } from 'element-plus'
import { type AxiosResponse } from 'axios'
import { evaluationList } from '@/utils/constant'

defineOptions({
  name: 'Evaluation',
})
//props
const props = defineProps({
  formData: {
    type: Object as PropType<EvaluateDemandForm>,
    required: true,
    default: () => ({}),
  },
})

const visible = defineModel('visible', {
  type: Boolean,
  required: true,
  default: true,
})

const formData = props.formData
// 小心「解构赋值」导致丢失响应式 用 toRefs 或 toRef 保持响应式。
const emit = defineEmits(['submit'])
// ref
const formRef = ref()
const title = ref<string>('请求评价')
// reactive
const rules = reactive({
  // opsEvaluation: [{ required: true, message: '请选择请是否满意', trigger: 'blur' }],
  // resEvaluation: [{ required: true, message: '请选择请是否满意', trigger: 'blur' }],
  // netEvaluation: [{ required: true, message: '请选择请是否满意', trigger: 'blur' }],
})

//method
function handleClose() {
  visible.value = false
  formRef.value.resetFields()
  formRef.value.clearValidate()
}

// 提交表单
async function handleEvaluate() {
  try {
    if (
      (formData.opsEvaluation === '不满意' &&
        formData.opsReason &&
        formData.opsReason?.length < 20) ||
      (formData.opsEvaluation === '不满意' && !formData.opsReason)
    ) {
      ElMessage.warning('不满意原因不能少于20字')
      return
    }
    if (
      (formData.resEvaluation === '不满意' &&
        formData.resReason &&
        formData.resReason?.length < 20) ||
      (formData.resEvaluation === '不满意' && !formData.resReason)
    ) {
      ElMessage.warning('不满意原因不能少于20字')
      return
    }
    if (
      (formData.netEvaluation === '不满意' &&
        formData.netReason &&
        formData.netReason?.length < 20) ||
      (formData.netEvaluation === '不满意' && !formData.netReason)
    ) {
      ElMessage.warning('不满意原因不能少于20字')
      return
    }
    if (!formData.opsEvaluation && !formData.resEvaluation && !formData.netEvaluation) {
      formData.opsEvaluation = '满意'
      formData.resEvaluation = '满意'
      formData.netEvaluation = '满意'
    }

    const resp: AxiosResponse = await DemandAPI.evaluate(formData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('评价成功')
      handleClose()
      emit('submit')
    } else {
      console.log(msg)
      ElMessage.error(`评价失败${msg}`)
    }
  } catch (err) {
    console.error(err)
    ElMessage.error('评价失败')
  }
}

async function handleSubmit() {
  formRef.value?.validate(async (valid: boolean) => {
    if (!valid) return
    handleEvaluate()
  })
}
</script>
<style lang="scss" scoped></style>
