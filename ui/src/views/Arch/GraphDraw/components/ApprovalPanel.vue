<!--  -->
<template>
  <div class="graph-approval-wrapper">
    <el-drawer v-model="approvalVisible" title="架构图提交审批" width="500px" @close="handleClose">
      <el-form ref="formRef" :model="formData" :rules="rules" style="max-width: 500px">
        <el-form-item label="架构图名:" label-position="left">
          <el-tag type="info">{{ graphName }}</el-tag>
        </el-form-item>
        <el-form-item label="审批类型:" label-position="left">
          <el-tag type="primary">{{ processOrderTypeMapping[orderType] }}</el-tag>
        </el-form-item>

        <el-form-item v-if="enabledImageData.length > 0" label="生效架构图:" label-position="top">
          <div class="image-wrapper">
            <el-image
              style="width: 450px; height: 300px"
              :src="enabledImageData"
              :zoom-rate="1.2"
              :max-scale="7"
              :min-scale="0.2"
              :preview-src-list="[enabledImageData, imageData]"
              show-progress
              :initial-index="4"
              fit="cover"
            />
          </div>
        </el-form-item>
        <el-form-item label="审批架构图:" label-position="top">
          <div class="image-wrapper">
            <el-image
              style="width: 450px; height: 300px"
              :src="imageData"
              :zoom-rate="1.2"
              :max-scale="7"
              :min-scale="0.2"
              :preview-src-list="[imageData]"
              show-progress
              :initial-index="4"
              fit="cover"
            />
          </div>
        </el-form-item>
        <el-form-item label-position="top" label="请求名称" prop="demandName">
          <el-select
            v-model="formData.demandName"
            placeholder="请选择请求名称"
            clearable
            filterable
            style="width: 100%"
          >
            <el-option
              v-for="item in demandList"
              :key="item.id"
              :label="item.name"
              :value="item.name"
            />
          </el-select>
        </el-form-item>
        <el-form-item label-position="top" label="架构图描述:" prop="description">
          <el-input
            style="width: 1030px"
            :rows="5"
            type="textarea"
            v-model="formData.description"
            placeholder="架构图描述"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="approvalVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确认</el-button>
        </div>
      </template>
    </el-drawer>
  </div>
</template>

<script setup lang="ts">
import ProcessOrderAPI, { type ProcessApplyForm, type ProcessOrderForm } from '@/api/process/order'
import type { AxiosResponse } from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'
import { computed, onMounted, reactive, ref, toRefs } from 'vue'
import { processOrderTypeMapping } from '@/utils/constant'
import type { DemandQuery } from '@/api/demand'
import { useAuthStore } from '@/store/modules/auth'
import { useDemandStore } from '@/store'
import { getCurrentMinutesString } from '@/utils'

defineOptions({
  name: 'ApprovalPanel',
})

const props = defineProps({
  graphID: {
    required: true,
    type: Number,
  },
  graphName: {
    required: true,
    type: String,
  },
  imageData: {
    required: true,
    type: String,
  },
  enabledImageData: {
    required: true,
    type: String,
  },
})

const approvalVisible = defineModel('approvalVisible', {
  type: Boolean,
  required: true,
  default: false,
})

const orderType = computed(() => {
  if (enabledImageData.value.length > 0) {
    return 2
  }
  return 1
})

const authStore = useAuthStore()
const { userName } = authStore

const demandStore = useDemandStore()
const { getDemandList } = demandStore
const demandList = computed(() => demandStore.demandList)

const initDesc: string = `1、请求方:
2、开发请求:
3、评审内容:`
const { graphID, graphName, imageData, enabledImageData } = toRefs(props)
const formRef = ref()

const formData = reactive<ProcessOrderForm>({
  description: initDesc,
})
const rules = reactive({
  demandName: [{ required: true, message: '请求名称不能为空', trigger: 'change' }],
  description: [{ required: true, message: '请输入组架构图审批描述', trigger: 'blur' }],
})

const emit = defineEmits<{
  (e: 'handleApprove', orderType: number): void
}>()

onMounted(async () => {
  const params: DemandQuery = {
    status: 2,
  }
  await getDemandList(params)
})

function handleSubmit() {
  formRef.value?.validate(async (valid: boolean) => {
    if (!valid) return
    handleApply()
  })
}

async function handleApply() {
  ElMessageBox.confirm(`确定要提交审批架构图吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      try {
        const timestr = getCurrentMinutesString()
        const title: string = `${userName}-${graphName.value}-${timestr}`
        const orderData: ProcessApplyForm = {
          title: title,
          graphName: graphName.value,
          graphID: graphID.value,
          imageData: imageData.value,
          enabledImageData: enabledImageData.value,
          orderType: orderType.value,
          description: formData.description,
          demandName: formData.demandName,
          owner: userName,
        }
        const resp: AxiosResponse = await ProcessOrderAPI.create(orderData)
        const status = resp.data.status
        const msg = resp.data.msg
        if (status === 200) {
          ElMessage.success('提交成功')
          // 选择成功后重新获取后端数据，render
          emit('handleApprove', orderType.value)
          approvalVisible.value = false
        } else {
          console.log(msg)
          ElMessage.error('提交失败')
        }
      } catch (err) {
        ElMessage.error('提交失败')
        console.log(err)
      }
    })
    .catch(() => {
      // 用户点击“取消” 或 关闭弹窗，什么都不做即可避免报错
    })
}
function handleClose() {
  approvalVisible.value = false
}
</script>

<style lang="scss" scoped>
.graph-approval-wrapper {
  :deep(.el-divider__text) {
    font-size: 16px;
  }
  .image-wrapper {
    display: flex;
    justify-content: center;
  }
}
</style>
