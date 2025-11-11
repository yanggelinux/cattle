<template>
  <div class="logo">
    <transition enter-active-class="animate__animated animate__fadeInLeft">
      <div :key="+collapse" class="logo-wrapper" @click="handleLayoutChange(layoutMode)">
        <img :src="logo" class="logo-img" />
        <span v-if="!collapse" class="title">架构流程平台</span>
      </div>
    </transition>
  </div>
</template>

<script lang="ts" setup>
import logo from '@/assets/logo.png'
import { LayoutMode } from '@/enums/settings/layout.enum'
import { useSettingsStore } from '@/store'
import { computed } from 'vue'

defineProps({
  collapse: {
    type: Boolean,
    required: true,
  },
})

const layoutModeValue = defineModel<LayoutMode>('layoutModeValue', {
  required: true,
  default: () => LayoutMode.LEFT,
})

const settingsStore = useSettingsStore()

const layoutMode = computed(() => settingsStore.layout)

function handleLayoutChange(layout: LayoutMode) {
  if (layout === LayoutMode.LEFT) {
    layoutModeValue.value = LayoutMode.TOP
  } else {
    layoutModeValue.value = LayoutMode.LEFT
  }
}
</script>

<style lang="scss" scoped>
.logo {
  width: 100%;
  height: $navbar-height;
  background-color: $sidebar-logo-background;
  .logo-wrapper {
    cursor: pointer;
    height: 40px;
    line-height: 40px;
    display: flex;
    .logo-img {
      height: 25px;
      width: 25px;
      margin: 13px 0px 0px 15px;
    }
    .title {
      flex-shrink: 0; /* 防止容器在空间不足时缩小 */
      margin: 5px 0px 50px 10px;
      font-size: 23px;
      font-weight: bold;
      color: #409eff;
      display: inline-block;
    }
  }
}

.layout-top,
.layout-mix {
  .logo {
    width: $sidebar-width;
  }

  &.hideSidebar {
    .logo {
      width: $sidebar-width-collapsed;
    }
  }
}
</style>
