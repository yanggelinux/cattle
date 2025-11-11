<!--  -->
<template>
  <div class="graph-enable-wrapper">
    <el-dialog
      draggable
      v-model="visible"
      title="存在已经审批通过生效的架构图，是否选择对其编辑？"
      width="500px"
      @close="handleClose"
    >
      <div class="image-wrapper">
        <el-image
          style="width: 300px; height: 300px"
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
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="visible = false">取消</el-button>
          <el-button type="primary" @click="handleSelectRecord">确认</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import ArchGraphAPI, { type ArchGraphSelectForm } from '@/api/arch/graph'
import type { AxiosResponse } from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, toRefs } from 'vue'

defineOptions({
  name: 'EnabledRecord',
})

const props = defineProps({
  graphID: {
    required: true,
    type: Number,
  },
  recordID: {
    required: true,
    type: Number,
  },
  imageData: {
    required: true,
    type: String,
  },
})

const loading = ref<boolean>(false)
const { graphID, recordID, imageData } = toRefs(props)
const emit = defineEmits(['select-record'])
const visible = defineModel('visible', {
  type: Boolean,
  required: true,
  default: true,
})

async function handleSelectRecord() {
  ElMessageBox.confirm(`确定要选择审批生效的图吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      try {
        loading.value = true
        const formData: ArchGraphSelectForm = {
          id: recordID.value,
          graphID: graphID.value,
        }
        const resp: AxiosResponse = await ArchGraphAPI.selectRecord(formData)
        const status = resp.data.status
        const msg = resp.data.msg
        if (status === 200) {
          ElMessage.success('选择成功')
          // 选择成功后重新获取后端数据，render
          emit('select-record')
          visible.value = false
        } else {
          console.log(msg)
          ElMessage.error('选择失败')
        }
      } catch (err) {
        ElMessage.error('选择失败')
        console.log(err)
      }
    })
    .catch(() => {
      // 用户点击“取消” 或 关闭弹窗，什么都不做即可避免报错
    })
}
function handleClose() {
  visible.value = false
}
</script>

<style lang="scss" scoped>
.graph-enable-wrapper {
  .image-wrapper {
    display: flex;
    justify-content: center;
  }
}
</style>
