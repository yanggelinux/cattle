<template>
  <div class="diagram-data-panel">
    <div class="close-icon-wrapper">
      <el-icon class="close-icon" @click="handleClose"><Close /></el-icon>
    </div>
    <div class="setting-block">
      <div class="settting-title">
        流程节点：
        <el-tag>{{ nodeName }}</el-tag>
      </div>
      <div>
        <el-form ref="formRef" label-position="top" :model="formData" :rules="rules">
          <el-form-item label-width="85px" label="节点名称" prop="name">
            <el-input v-model="formData.name" placeholder="请输入节点名称" />
          </el-form-item>
          <el-form-item
            label-width="85px"
            label="角色"
            prop="role"
            v-if="nodeType === 'procApproval'"
          >
            <el-select
              clearable
              filterable
              placement="bottom-end"
              v-model="formData.role"
              placeholder="请选择角色"
            >
              <el-option
                v-for="item in roleList"
                :value="item.roleName"
                :label="item.displayName"
                :key="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item
            label-width="85px"
            label="审批通过类型"
            prop="approvalType"
            v-if="nodeType === 'procApproval'"
          >
            <el-select
              clearable
              filterable
              placement="bottom-end"
              v-model="formData.approvalType"
              placeholder="请选择审批通过类型"
            >
              <el-option
                v-for="item in approvalInfoList"
                :value="item.value"
                :label="item.label"
                :key="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item
            v-if="nodeType === 'procApproval'"
            label-width="85px"
            label="审批时可编辑"
            prop="approvalEdit"
          >
            <el-select
              clearable
              filterable
              placement="bottom-end"
              v-model="formData.approvalEdit"
              placeholder="请选择是否审批时可编辑"
            >
              <el-option
                v-for="item in yesOrNoList"
                :value="item.value"
                :label="item.label"
                :key="item.value"
              />
            </el-select>
          </el-form-item>
        </el-form>
      </div>
      <div class="setting-btn">
        <el-button type="primary" @click="handleSubmit">确 定</el-button>
        <el-button @click="handleClose">取 消</el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type LogicFlow from '@logicflow/core'
import { computed, onMounted, reactive, ref, watch, type PropType } from 'vue'
import type { ProcessNodeForm } from '@/api/process/process'
import { useRoleStore } from '@/store'
import { approvalInfoList, yesOrNoList } from '@/utils/constant'
import { ElMessage } from 'element-plus'

const roleStore = useRoleStore()
const { getRoleList, roleNameMapping } = roleStore
const roleList = computed(() => roleStore.roleList)

const props = defineProps({
  lf: {
    required: true,
    type: Object as PropType<LogicFlow | any>,
  },
  nodeType: {
    type: String,
    required: true,
    default: '',
  },
})

const { lf } = props
// 选择框已选择

// 已经选择table数据

defineOptions({
  name: 'DataPanel',
})

const formRef = ref()
const nodeName = ref<string>('')

const rules = reactive({
  name: [{ required: true, message: '请输入请求名称', trigger: 'blur' }],
  role: [{ required: true, message: '请输选择审批角色', trigger: 'blur' }],
})

const formData = reactive<ProcessNodeForm>({
  name: '',
  role: '',
  roleName: '',
  approvalType: 0,
  approvalEdit: 0,
})

const emit = defineEmits<{
  (e: 'close'): void
}>()

const nodeID = defineModel('nodeID', {
  type: String,
  required: true,
  default: '',
})

onMounted(async () => {
  await getRoleList({})
})

watch(
  () => nodeID.value,
  () => {
    const selectedNode = lf.getNodeModelById(nodeID.value)
    nodeName.value = selectedNode.text.value
    const nodeInfo = selectedNode.properties.nodeInfo
    if (nodeInfo) {
      Object.assign(formData, nodeInfo)
    } else {
      formData.name = ''
      formData.role = ''
    }

    // 查询数据
  },
  { immediate: true }
)

function setNodeInfo() {
  const text = formData.name
  if (formData.role) {
    formData.roleName = roleNameMapping.get(formData.role) || ''
  }

  lf.updateText(nodeID.value, text)
  lf.setProperties(nodeID.value, { nodeInfo: formData })
}

async function handleSubmit() {
  formRef.value?.validate(async (valid: boolean) => {
    if (!valid) return
    // 在设计设置node 属性信息
    setNodeInfo()
    handleClose()
    ElMessage.success('设置成功')
  })
}

function handleClose() {
  emit('close')
}
</script>

<style lang="scss" scoped>
.diagram-data-panel {
  padding: 10px 20px 20px 20px;
  .close-icon-wrapper {
    text-align: right;
    .close-icon {
      cursor: pointer;
    }
  }

  .setting-block {
    overflow: hidden;
    .settting-title {
      margin-bottom: 10px;
    }
    .setting-btn {
      text-align: right;
    }
  }
}
</style>
