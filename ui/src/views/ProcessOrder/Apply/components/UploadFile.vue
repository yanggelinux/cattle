<template>
  <div class="upload-file-wrapper">
    <el-upload
      v-loading="loading"
      element-loading-text="上传中..."
      v-model:file-list="fileList"
      ref="uploadRef"
      class="upload-demo"
      drag
      action=""
      :http-request="handleUpload"
      :on-exceed="handleExceed"
      :show-file-list="true"
      :limit="1"
      :auto-upload="true"
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">
        Drop file here or
        <em>click to upload</em>
      </div>
    </el-upload>
  </div>
</template>
<script setup lang="ts">
import UploadAPI, { type UploadFileInfo, type UploadFileResult } from '@/api/upload'
import { ElMessage, type UploadInstance, type UploadUserFile } from 'element-plus'
import { type AxiosResponse } from 'axios'
import { onMounted, ref, toRefs } from 'vue'

defineOptions({
  name: 'UploadFile',
})

const props = defineProps({
  fieldKey: {
    type: String,
    required: true,
    default: () => '',
  },
  fieldVal: {
    type: String,
    required: true,
    default: () => '',
  },
  groupName: {
    type: String,
    required: false,
    default: () => '',
  },

  idx: {
    type: Number,
    required: false,
    default: () => 0,
  },
})

const { fieldKey, fieldVal, groupName, idx } = toRefs(props)

// 小心「解构赋值」导致丢失响应式 用 toRefs 或 toRef 保持响应式。
const uploadRef = ref<UploadInstance>()
const loading = ref<boolean>(false)
const fileList = ref<UploadUserFile[]>([])

const emit = defineEmits<{
  (e: 'handleUpload', fileInfo: UploadFileInfo): void
}>()

function handleExceed() {
  ElMessage.warning('只支持上传一个文件')
}

onMounted(() => {
  if (fieldKey.value.length > 0) {
    const fileUrl = fieldVal.value
    const fileUrls = fieldVal.value.split('/')
    const fileName = fileUrls[fileUrls.length - 1]
    if (fileName.trim().length > 0) {
      const file: UploadUserFile = {
        name: fileName,
        url: fileUrl,
      }
      fileList.value = [file]
    }
  }
})

async function handleUpload(param: any) {
  const formData = new FormData()
  formData.append('file', param.file)
  loading.value = true
  try {
    const resp: AxiosResponse = await UploadAPI.uploadFile(formData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const retData: UploadFileResult = resp.data.data
      const fileName = retData.fileName
      const fileUrl = retData.fileUrl
      const file: UploadUserFile = {
        name: fileName,
        url: fileUrl,
      }
      fileList.value = [file]
      ElMessage.success('上传成功')
      const fileInfo: UploadFileInfo = {
        fieldKey: fieldKey.value,
        fieldVal: fileUrl,
        groupName: groupName.value,
        idx: idx.value,
      }
      emit('handleUpload', fileInfo)
      loading.value = false
    } else {
      console.log(msg)
      ElMessage.error(`上传失败:${msg}`)
    }
  } catch (err) {
    console.log(err)
    ElMessage.error('上传失败')
  } finally {
    loading.value = false
  }
}
</script>
<style lang="scss" scoped>
.upload-file-wrapper {
  width: 100%;
}
</style>
