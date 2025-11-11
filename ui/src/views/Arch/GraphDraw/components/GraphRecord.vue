<!--  -->
<template>
  <div class="graph-record-wrapper">
    <el-drawer v-model="visible" title="快照选择" width="500px" @close="handleClose">
      <el-table
        v-loading="loading"
        :element-loading-text="loadingText"
        ref="dataTableRef"
        :data="recordList"
        highlight-current-row
        border:true
      >
        <el-table-column label="快照时间" prop="createdTime" />
        <el-table-column label="快照" prop="imageData">
          <template #default="scope">
            <el-image
              style="width: 100px; height: 100px"
              :src="scope.row.imageData"
              :zoom-rate="1.2"
              :max-scale="7"
              :min-scale="0.2"
              :preview-src-list="[scope.row.imageData]"
              show-progress
              :initial-index="4"
              fit="cover"
            />
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="120">
          <template #default="scope">
            <el-button
              type="primary"
              size="small"
              link
              icon="position"
              @click="handleSelectRecord(scope.row)"
            >
              选择快照
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-drawer>
  </div>
</template>

<script setup lang="ts">
import ArchGraphAPI, {
  type ArchGraphRecordQuery,
  type ArchGraphRecordResult,
  type ArchGraphRecordData,
  type ArchGraphSelectForm,
} from '@/api/arch/graph'
import type { AxiosResponse } from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, toRefs, watch } from 'vue'

defineOptions({
  name: 'GraphRecord',
})

const props = defineProps({
  graphID: {
    required: true,
    type: Number,
  },
})

const recordList = ref<ArchGraphRecordResult[]>([])
const loading = ref<boolean>(false)
const loadingText = ref<string>('loading...')
const { graphID } = toRefs(props)
const emit = defineEmits(['select-record'])
const visible = defineModel('visible', {
  type: Boolean,
  required: true,
  default: true,
})

// onMounted(() => {
//   // 查询数据
//   handleQuery()
// })

// 监控 visible 变化
watch(
  () => visible.value,
  (newVal) => {
    if (newVal) {
      // 弹框打开时执行
      handleQuery()
    }
  }
)

async function handleQuery() {
  try {
    const queryParams: ArchGraphRecordQuery = {
      graphID: graphID.value,
    }
    const resp: AxiosResponse = await ArchGraphAPI.getRecordList(queryParams)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const respData: ArchGraphRecordData = resp.data.data
      recordList.value = respData.retList
      // 生成表格数据
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}
async function handleSelectRecord(row: ArchGraphRecordResult) {
  ElMessageBox.confirm(`确定要选择${row?.createdTime}快照吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      try {
        loading.value = true
        const formData: ArchGraphSelectForm = {
          id: row.id,
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
      } finally {
        loading.value = false
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

<style lang="scss" scoped></style>
