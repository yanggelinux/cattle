<template>
  <div>
    <el-steps :active="activeIndex" align-center finish-status="success">
      <el-step
        v-for="(step, index) in stepList"
        :key="index"
        :title="step.title"
        :description="step.description"
        :icon="stepIconMap[step.opt][index]"
      ></el-step>
    </el-steps>
  </div>
</template>
<script setup lang="ts">
import { ref, toRefs, onMounted, type PropType } from 'vue'
import { type ProcessNode } from '@/api/demand'

interface StepNode {
  title: string
  description: string
  opt: string
}

defineOptions({
  name: 'ProcessSteps',
})
//props
const props = defineProps({
  orderProcess: {
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
const { orderProcess, activeIndex } = toRefs(props)
const stepIconMap: any = {
  apply: { 0: 'BellFilled' },
  approve: {
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
  complete: {
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
  orderProcess.value.forEach((item) => {
    let description = ''
    if (item.opt === 'apply' && item.approverName.length > 0) {
      description += `申请人:${item.approverName}\n`
    }
    if (item.opt === 'approve' && item.approverGroup.length > 0) {
      description += `审批团队:${item.approverGroup}\n`
    }
    if (item.opt === 'approve' && item.approverName.length > 0) {
      description += `实际审批人:${item.approverName}\n`
    }
    const step: StepNode = {
      title: item.name,
      description: description,
      opt: item.opt,
    }
    stepList.value.push(step)
  })
}
</script>
