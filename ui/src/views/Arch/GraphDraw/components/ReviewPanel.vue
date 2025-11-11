<!--  -->
<template>
  <div class="graph-review-wrapper">
    <el-drawer v-model="visable" title="架构图评审" width="500px" @close="handleClose">
      <div class="review-content">
        <div class="content-wrapper" v-if="reviewStatus.includes(status)">
          <el-form ref="formRef" :model="formData" :rules="rules" style="width: 100%">
            <el-form-item label-position="top" label="评审意见:" prop="content">
              <el-input
                :rows="4"
                type="textarea"
                v-model="formData.content"
                placeholder="评审意见"
              />
            </el-form-item>
          </el-form>
          <div class="btn-wrapper">
            <el-button icon="plus" @click="handleSubmit" type="primary">提交评审</el-button>
          </div>
        </div>

        <el-divider content-position="left">评审内容</el-divider>
        <div class="refresh-btn-wrapper">
          <el-button @click="handleQuery" type="primary" size="small" icon="Refresh" circle />
        </div>
        <el-table
          row-key="id"
          v-loading="loading"
          size="small"
          :data="reviewList"
          style="width: 100%"
        >
          <el-table-column prop="createdTime" label="时间" width="130">
            <template #default="scope">
              <span>
                {{ scope.row.createdTime }}
              </span>
            </template>
          </el-table-column>
          <el-table-column prop="reviewer" label="评审人" width="100">
            <template #default="scope">
              <span>
                {{ scope.row.reviewer }}
              </span>
            </template>
          </el-table-column>
          <el-table-column prop="notify_party" label="被通知人" width="100">
            <template #default="scope">
              <span>
                {{ scope.row.notifyParty.join(',') }}
              </span>
            </template>
          </el-table-column>
          <el-table-column prop="content" label="意见">
            <template #default="scope">
              <div v-for="(txt, idx) in scope.row.content.split('\n')" :key="idx">{{ txt }}</div>
            </template>
          </el-table-column>

          <el-table-column fixed="right" label="操作" width="60">
            <template #default="scope">
              <el-button
                type="danger"
                size="small"
                icon="delete"
                @click="handleDelete(scope.row)"
                circle
              ></el-button>
            </template>
          </el-table-column>
        </el-table>
        <Pagination
          v-if="total > 0"
          v-model:total="total"
          v-model:page="queryParams.page"
          v-model:limit="queryParams.pageSize"
          size="small"
          @pagination="handleQuery"
        />
      </div>
    </el-drawer>
  </div>
</template>

<script setup lang="ts">
import ArchGraphAPI, {
  type ArchGraphReviewForm,
  type ArchGraphReviewQuery,
  type ArchGraphReviewResult,
} from '@/api/arch/graph'
import { onMounted, reactive, ref, toRefs } from 'vue'
import type { AxiosResponse } from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({
  name: 'ReviewPanel',
})

const props = defineProps({
  graphID: {
    required: true,
    type: Number,
  },
  graphKey: {
    required: true,
    type: String,
  },
  graphName: {
    required: true,
    type: String,
  },
  status: {
    required: true,
    type: Number,
  },
})
const visable = defineModel('visable', {
  type: Boolean,
  required: true,
  default: true,
})

const { graphID, graphKey, graphName, status } = toRefs(props)
const formRef = ref()
const reviewList = ref<ArchGraphReviewResult[]>([])
const loading = ref<boolean>(false)
const reviewStatus = ref<number[]>([0, 3])
const total = ref<number>(0)
const queryParams = reactive<ArchGraphReviewQuery>({
  graphID: graphID.value,
  page: 1,
  pageSize: 10,
})

const formData = reactive<ArchGraphReviewForm>({})
const rules = reactive({
  content: [{ required: true, message: '评审意见不能为空', trigger: 'blur' }],
})

onMounted(() => {
  handleQuery()
})

async function handleQuery() {
  try {
    const resp: AxiosResponse = await ArchGraphAPI.getReviewList(queryParams)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData = resp.data.data
      reviewList.value = resData.retList
      total.value = resData.total
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

function handleDelete(rowData: ArchGraphReviewResult) {
  ElMessageBox.confirm(`确定要删除吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      loading.value = true
      try {
        const id: number = rowData.id ? rowData.id : 0
        const resp: AxiosResponse = await ArchGraphAPI.deleteReview(id)
        const status = resp.data.status
        const msg = resp.data.msg
        if (status === 200) {
          ElMessage.success('删除成功')
          handleQuery()
        } else {
          console.log(msg)
          ElMessage.error('删除失败')
        }
      } catch (err) {
        console.error(err)
        ElMessage.error('删除失败')
      } finally {
        loading.value = false
      }
    })
    .catch(() => {
      // 用户点击“取消” 或 关闭弹窗，什么都不做即可避免报错
    })
}

function handleClose() {
  visable.value = false
}

// 提交申请
async function handleSubmit() {
  formRef.value?.validate(async (valid: boolean) => {
    if (!valid) return
    try {
      const orderData: ArchGraphReviewForm = {
        graphID: graphID.value,
        graphKey: graphKey.value,
        graphName: graphName.value,
        notifyParty: formData.notifyParty || [],
        content: formData.content,
      }
      const resp: AxiosResponse = await ArchGraphAPI.createReview(orderData)
      const status = resp.data.status
      const msg = resp.data.msg
      if (status === 200) {
        ElMessage.success('提交成功')
        formRef.value.resetFields()
        handleQuery()
      } else {
        console.log(msg)
        ElMessage.error('提交失败')
      }
    } catch (err) {
      ElMessage.error('提交失败')
      console.log(err)
    }
  })
}
</script>

<style lang="scss" scoped>
.graph-review-wrapper {
  height: 100%;
  display: flex;
  flex-direction: column;
  .review-content {
    :deep(.el-divider) {
      margin-bottom: 10px;
    }
    .refresh-btn-wrapper {
      text-align: right;
      margin-right: 30px;
    }
    .btn-wrapper {
      margin-top: 10px;
      text-align: right;
      // margin-right: 20px;
    }
  }
}
</style>
