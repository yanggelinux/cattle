<template>
  <template v-if="icon">
    <el-icon v-if="isElIcon" class="el-icon">
      <component :is="iconComponent" />
    </el-icon>
    <!-- <IconSvg v-else-if="isIconFont" :icon-name="icon" /> -->
    <div v-else :class="`i-svg:${icon}`" />
  </template>
  <template v-else>
    <div :class="`i-svg:user`" />
  </template>
  <!-- 菜单标题 -->
  <span v-if="title" class="ml-1">{{ title }}</span>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  icon?: string
  title?: string
}>()

const isElIcon = computed(() => props.icon?.startsWith('el-icon'))
// const isIconFont = computed(() => props.icon?.startsWith('icon'))
const iconComponent = computed(() => props.icon?.replace('el-icon-', ''))
</script>

<style lang="scss" scoped>
.el-icon {
  margin-right: 5px;
  width: 14px;
  height: 14px;
  color: currentcolor !important;
}

// .icon-svg {
//   margin-right: 5px;
//   width: 14px;
//   height: 14px;
//   color: currentcolor !important;
// }

[class^='i-svg:'] {
  width: 14px;
  height: 14px;
  color: currentcolor !important;
}

.hideSidebar {
  .el-sub-menu,
  .el-menu-item {
    .el-icon {
      margin-left: 20px;
    }
  }
  .el-icon {
    margin-left: 15px;
    width: 24px !important;
    height: 24px !important;
    color: currentcolor;
  }
  [class^='i-svg:'] {
    margin-left: 15px;
    width: 24px !important;
    height: 24px !important;
    color: currentcolor;
  }
}
</style>
