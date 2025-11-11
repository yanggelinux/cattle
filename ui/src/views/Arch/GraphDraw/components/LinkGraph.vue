<template>
  <div class="ArchGraph-opt-wrapper">
    <el-dialog
      draggable
      v-model="visible"
      title="关联架构图"
      width="500px"
      @close="handleCloseDialog"
    >
      <el-form ref="formRef" label-width="80px" :model="formData" :rules="rules">
        <el-form-item label-position="right" label="图组" prop="groupIDs">
          <el-cascader
            :props="cascaderProps"
            v-model="groupIDs"
            :options="archGroupTree"
            @change="handleQueryGraph"
            style="width: 400px"
          />
        </el-form-item>

        <el-form-item label-position="right" label="架构图" prop="graphID">
          <el-select v-model="formData.graphID" placeholder="请选择架构图">
            <el-option
              v-for="item in graphList"
              :key="item.id"
              :label="item.graphName"
              :value="item.id"
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
import ArchGraphAPI, {
  type ArchGraphQuery,
  type ArchGraphResult,
  type ArchGraphData,
} from '@/api/arch/graph'
import { computed, onMounted, reactive, ref, toRefs } from 'vue'
import { ElMessage } from 'element-plus'
import { type AxiosResponse } from 'axios'
import { useArchGroupStore } from '@/store/modules/archGroup'

defineOptions({
  name: 'GraphOpt',
})

//props
const props = defineProps({
  nodeID: {
    type: String,
    required: true,
  },
  graphID: {
    type: Number,
    required: true,
  },
  groupID: {
    type: Number,
    required: true,
  },
})

const visible = defineModel('visible', {
  type: Boolean,
  required: true,
  default: true,
})

const archGroupStore = useArchGroupStore()
const { getArchGroupTree, findPathByValue } = archGroupStore
const archGroupTree = computed(() => archGroupStore.archGroupTree)
const cascaderProps = {
  showPrefix: false,
  checkStrictly: true,
  checkOnClickNode: true,
}
// ref
const formRef = ref()
const groupIDs = ref<number[]>([])
const graphList = ref<ArchGraphResult[]>([])
const graphIDResult = ref<Map<number, ArchGraphResult>>(new Map<number, ArchGraphResult>())
const formData = reactive<any>({})
// reactive
const rules = reactive({
  // groupID: [{ required: true, message: '请选择图组', trigger: 'blur' }],
  graphID: [{ required: true, message: '请选择图', trigger: 'blur' }],
})

// 小心「解构赋值」导致丢失响应式 用 toRefs 或 toRef 保持响应式。
const { graphID, groupID, nodeID } = toRefs(props)

const emit = defineEmits<{
  (e: 'handleSubmit', nodeID: string, graph: ArchGraphResult): void
}>()

onMounted(async () => {
  if (graphID.value) {
    groupIDs.value = findPathByValue(archGroupStore.archGroupTree, groupID.value)
    await handleQueryGraph()
    formData.graphID = graphID.value
  }
  await getArchGroupTree({})
})
//method

async function handleQueryGraph() {
  try {
    const params: ArchGraphQuery = {}
    if (groupIDs.value.length > 0) {
      const groupID = groupIDs.value[groupIDs.value.length - 1]
      params.groupID = groupID
    }
    const resp: AxiosResponse = await ArchGraphAPI.getList(params)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      formData.graphID = null
      const resData: ArchGraphData = resp.data.data
      const retList = resData.retList
      graphList.value = resData.retList
      for (const ret of retList) {
        graphIDResult.value.set(ret.id, ret)
      }
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}
function handleCloseDialog() {
  formRef.value.resetFields()
  formRef.value.clearValidate()
  visible.value = false
}

function handleSubmit() {
  formRef.value?.validate(async (valid: boolean) => {
    if (!valid) return
    const id = formData.graphID
    if (id) {
      const graphInfo: ArchGraphResult | any = graphIDResult.value.get(id) || {}
      emit('handleSubmit', nodeID.value, graphInfo)
      ElMessage.success('关联架构图成功')
      handleCloseDialog()
    }
  })
}
</script>
<style lang="scss" scoped></style>
