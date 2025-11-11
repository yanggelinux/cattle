<!-- components/IconSvg.vue -->
<template>
  <svg aria-hidden="true" :class="className" :style="iconStyle">
    <use :xlink:href="`#${iconName}`" />
  </svg>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

defineOptions({
  name: 'IconSvg',
})
const props = defineProps({
  iconName: { type: String, required: true }, // 图标名称（如 icon-xxx）
  className: { type: String, default: 'icon-svg' }, // 自定义类名
  color: { type: String, default: '' },
  width: { type: String, default: '1em' },
  height: { type: String, default: '1em' },
})
interface iconStyle {
  width: string
  height: string
  color?: string
}
const iconStyle = ref<iconStyle>({ width: props.width, height: props.height })
onMounted(() => {
  if (props.color !== '') {
    iconStyle.value.color = props.color
  }
})
</script>
<style lang="scss" scoped>
.icon-svg {
  fill: currentcolor;
  overflow: hidden;
}
</style>
