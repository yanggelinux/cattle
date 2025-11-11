<template>
  <div>
    <el-steps :active="activeIndex" align-center finish-status="success">
      <el-step v-for="(step, index) in stepList" :key="index" :icon="stepIconMap[step.type][index]">
        <template #title>
          <div style="display: block">
            <span v-for="title in step.title" :key="title">
              {{ title }}
            </span>
          </div>
        </template>
        <template #description>
          <div v-for="desc in step.description" :key="desc">
            {{ desc }}
          </div>
        </template>
      </el-step>
    </el-steps>
  </div>
</template>
<script setup lang="ts">
import { ref, toRefs, onMounted, type PropType } from 'vue'
import { type ProcessNode } from '@/api/process/process'

interface StepNode {
  title: string[]
  description: string[]
  type: string
}

defineOptions({
  name: 'ProcessSteps',
})
//props
const props = defineProps({
  process: {
    type: Object as PropType<ProcessNode[]>,
    required: true,
  },
  activeIndex: {
    type: Number,
    required: true,
    default: 0,
  },
})
const stepList = ref<StepNode[]>([])
const { process, activeIndex } = toRefs(props)
const stepIconMap: any = {
  procStart: { 0: 'BellFilled' },
  procApproval: {
    0: 'Stamp',
    1: 'Stamp',
    2: 'Stamp',
    3: 'Stamp',
    4: 'Stamp',
    5: 'Stamp',
    6: 'Stamp',
    7: 'Stamp',
    8: 'Stamp',
    9: 'Stamp',
    10: 'Stamp',
  },
  procEnd: {
    1: 'SuccessFilled',
    2: 'SuccessFilled',
    3: 'SuccessFilled',
    4: 'SuccessFilled',
    5: 'SuccessFilled',
    6: 'SuccessFilled',
    7: 'SuccessFilled',
    8: 'SuccessFilled',
    9: 'SuccessFilled',
    10: 'SuccessFilled',
  },
}

onMounted(() => {
  genStepList()
})

function genStepList() {
  process.value.forEach((item) => {
    const descriptions: string[] = []
    if (item.type === 'procStart' && item.approvalInfo.length > 0) {
      for (const info of item.approvalInfo) {
        if (info.approverName) {
          descriptions.push(`操作人:${info.approverName}`)
        }
      }
    }
    if (item.type === 'procApproval' && item.approvalInfo.length > 0) {
      if (item.approver) {
        descriptions.push(`分配审批人:${item.approver}`)
      }
      for (const info of item.approvalInfo) {
        if (info.roleName) {
          descriptions.push(`审批角色:${info.roleName}`)
        }

        if (info.approverName) {
          descriptions.push(`审批人:${info.approverName}`)
        }
      }
    }
    const step: StepNode = {
      title: item.name.split(',') || [item.name],
      description: descriptions,
      type: item.type,
    }
    stepList.value.push(step)
  })
}
</script>
