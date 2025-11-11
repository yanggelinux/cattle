<template>
  <div class="diagram-sidebar">
    <!-- 基本节点 -->
    <section>
      <h1 class="node-category-title">基础组件</h1>
      <div
        v-for="item in generalNodes"
        :key="item.type"
        class="default-node-wrapper"
        @mousedown.prevent="dragInNode(item.type, '')"
      >
        <component :is="item.component" class="svg-node" />
      </div>
    </section>
    <!-- 分组节点 -->
    <section>
      <h1 class="node-category-title">分组组件</h1>

      <div
        v-for="item in groupNodes"
        :key="item.type"
        class="default-node-wrapper"
        @mousedown.prevent="dragInNode(item.type, '')"
      >
        <component :is="item.component" class="svg-node" />
      </div>
    </section>
    <!--架构图-->
    <section>
      <h1 class="node-category-title">架构图</h1>
      <div
        v-for="item in graphNodes"
        :key="item.type"
        class="image-node-wrapper"
        @mousedown.prevent="dragInNode(item.type, item.text)"
      >
        <div class="image-node" :class="`i-svg:${item.class}`"></div>
        <div class="image-label">{{ item.text }}</div>
      </div>
    </section>
    <section>
      <h1 class="node-category-title">常用组件</h1>
      <div
        v-for="item in usedNodes"
        :key="item.type"
        class="image-node-wrapper"
        @mousedown.prevent="dragInNode(item.type, item.text)"
      >
        <div class="image-node" :class="`i-svg:${item.class}`"></div>
        <div class="image-label">{{ item.text }}</div>
      </div>
    </section>

    <!-- 图标节点 -->
    <section>
      <h1 class="node-category-title">其它组件</h1>
      <div
        v-for="item in otherNodes"
        :key="item.type"
        class="icon-node-wrapper"
        @mousedown.prevent="dragInNode(item.type, item.text)"
      >
        <div class="icon-node" :class="`i-svg:${item.class}`"></div>
        <div class="icon-label">{{ item.text }}</div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// ICON 组件按需引入
import IconCircle from './icon/Circle.vue'
import IconRect from './icon/Rect.vue'
import IconRectRadius from './icon/RectRadius.vue'
import IconTriangle from './icon/Triangle.vue'
import IconEllipse from './icon/Ellipse.vue'
import IconDiamond from './icon/Diamond.vue'
import IconCylinde from './icon/Cylinde.vue'
import IconActor from './icon/Actor.vue'
import IconParallelogram from './icon/Parallelogram.vue'
import IconText from './icon/Text.vue'

import IconPentagon from './icon/Pentagon.vue'
import IconHexagon from './icon/Hexagon.vue'
import IconSeptagon from './icon/Septagon.vue'
import IconHeptagon from './icon/Heptagon.vue'
import IconTrapezoid from './icon/Trapezoid.vue'
import { usedNodes, otherNodes, graphNodes } from '../constant/index'

defineOptions({
  name: 'DiagramSidebar',
})
// 定义事件
const emit = defineEmits<{
  (e: 'dragInNode', type: string, text: string): void
}>()

// 通用节点列表
const generalNodes = [
  { type: 'pro-rect', component: IconRect },
  { type: 'rect-radius', component: IconRectRadius },
  { type: 'pro-circle', component: IconCircle },
  { type: 'pro-ellipse', component: IconEllipse },
  { type: 'pro-diamond', component: IconDiamond },
  { type: 'triangle', component: IconTriangle },
  { type: 'cylinde', component: IconCylinde },
  { type: 'parallelogram', component: IconParallelogram },

  { type: 'pentagon', component: IconPentagon },
  { type: 'hexagon', component: IconHexagon },
  { type: 'septagon', component: IconSeptagon },
  { type: 'heptagon', component: IconHeptagon },
  { type: 'trapezoid', component: IconTrapezoid },
  { type: 'actor', component: IconActor },
  { type: 'pro-text', component: IconText },
]

const groupNodes = [
  { type: 'rectGroup', component: IconRect, text: '矩形分组' },
  { type: 'circleGroup', component: IconCircle, text: '圆性分组' },
  { type: 'ellipseGroup', component: IconEllipse, text: '椭圆分组' },
  { type: 'triangleGroup', component: IconTriangle, text: '三角形分组' },
]

// 拖拽入画布
function dragInNode(type: string, text: string) {
  emit('dragInNode', type, text)
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

    .default-node-wrapper {
      width: 30px;
      height: 30px;
      margin: 3px;
      display: inline-block;

      .svg-node {
        left: 1px;
        top: 1px;
        width: 32px;
        height: 30px;
        display: block;
        position: relative;
        overflow: hidden;
      }
      &:hover {
        background: $icon-background;
        border-radius: 3px;
      }
    }

    .network-node-wrapper {
      display: inline-flex;
      flex-direction: column;
      align-items: center;
      width: 50px;
      cursor: pointer;
      font-size: 12px;
      text-align: center;
      .svg-node {
        left: 1px;
        top: 1px;
        width: 32px;
        height: 30px;
        display: block;
        position: relative;
        overflow: hidden;
      }
      .network-label {
        margin-top: 4px;
        line-height: 1.2;
        color: #555;
        word-break: break-word;
      }
      &:hover {
        background: $icon-background;
        border-radius: 3px;
      }
    }

    .image-node-wrapper {
      display: inline-flex;
      flex-direction: column;
      align-items: center;
      width: 50px;
      cursor: pointer;
      font-size: 12px;
      text-align: center;
      .image-node {
        display: inline-block;
        height: 20px;
        width: 50px;
        margin: 5px 10px 2px 10px;
        cursor: pointer;
      }
      .image-label {
        margin-top: 4px;
        line-height: 1.2;
        color: #555;
        word-break: break-word;
      }
      &:hover {
        background: $icon-background;
        border-radius: 3px;
      }
    }

    .image-node-wrapper-hover {
      display: inline-flex;
      justify-content: center;
      .image-node-hover {
        display: inline-block;
      }
    }

    .icon-node-wrapper {
      display: inline-flex;
      flex-direction: column;
      align-items: center;
      width: 50px;
      // margin: 10px;
      cursor: pointer;
      font-size: 12px;
      text-align: center;

      .icon-node {
        display: inline-block;
        width: 50px;
        margin: 5px 10px 2px 10px;
        cursor: pointer;
        height: 20px;
      }

      .icon-label {
        margin-top: 4px;
        color: #555;
        line-height: 1.2;
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
