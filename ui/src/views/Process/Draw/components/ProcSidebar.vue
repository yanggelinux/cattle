<template>
  <div class="diagram-sidebar">
    <section>
      <h1 class="node-category-title">流程节点</h1>
      <div
        v-for="item in generalNodes"
        :key="item.type"
        class="image-node-wrapper"
        @mousedown.prevent="dragInNode(item.type)"
      >
        <component :is="item.component" class="svg-node" />
        <div class="image-label">{{ item.text }}</div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// ICON 组件按需引入
import IconStart from '@/views/Arch/GraphDraw/components/icon/Start.vue'
import IconEnd from '@/views/Arch/GraphDraw/components/icon/End.vue'
import IconApproval from '@/views/Arch/GraphDraw/components/icon/Approval.vue'

defineOptions({
  name: 'ProcSidebar',
})
// 定义事件
const emit = defineEmits<{
  (e: 'dragInNode', type: string): void
}>()

// 通用节点列表
const generalNodes = [
  { type: 'procStart', component: IconStart, text: '开始节点' },
  { type: 'procApproval', component: IconApproval, text: '审批节点' },
  { type: 'procEnd', component: IconEnd, text: '结束节点' },
]

// 拖拽入画布
function dragInNode(type: string) {
  emit('dragInNode', type)
}
</script>

<style lang="scss" scoped>
$icon-background: rgb(216.8, 235.6, 255);
.diagram-sidebar {
  user-select: none;
  height: 100vh; /* 或指定固定高度，比如 600px */
  overflow-y: auto;
  padding: 10px;
  box-sizing: border-box;

  section {
    .node-category-title {
      margin: 10px;
      font-size: 14px;
      display: block;
      border-bottom: 1px solid #e5e5e5;
      line-height: 30px;
      // margin-bottom: 10px;
    }

    & .node-category {
      border-bottom: 1px solid #e5e5e5;
    }
    .image-node-wrapper {
      cursor: pointer;
      font-size: 12px;
      // text-align: center;
      .svg-node {
        margin: 0px 0px 0px 0px;
        cursor: pointer;
      }
      .image-label {
        margin-top: 4px;
        line-height: 1.2;
        color: #555;
        margin-left: 10px;
        word-break: break-word;
      }
      &:hover {
        background: $icon-background;
        border-radius: 3px;
      }
    }
  }
}
.diagram-sidebar::-webkit-scrollbar {
  width: 6px;
}
.diagram-sidebar::-webkit-scrollbar-thumb {
  background-color: rgba(0, 0, 0, 0.15);
  border-radius: 3px;
}
</style>
