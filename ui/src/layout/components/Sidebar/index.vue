<template>
  <div :class="{ 'has-logo': sidebarLogo }">
    <!-- 顶部布局顶部 || 左侧布局左侧 -->
    <SidebarLogo
      v-if="sidebarLogo"
      :collapse="isSidebarCollapsed"
      v-model:layoutModeValue="settingsStore.layout"
      @update:model-value="handleLayoutChange"
    />
    <el-scrollbar>
      <SidebarMenu :data="routes" base-path="" />
    </el-scrollbar>

    <!-- 顶部布局导航 -->
    <NavbarRight v-if="isTopLayout" />
  </div>
</template>

<script setup lang="ts">
import { LayoutMode } from '@/enums/settings/layout.enum'
import { useSettingsStore, useAppStore, useAuthStore } from '@/store'
import NavbarRight from '@/layout/components/NavBar/components/NavbarRight.vue'
import SidebarLogo from '@/layout/components/Sidebar/components/SidebarLogo.vue'
import SidebarMenu from '@/layout/components/Sidebar/components/SidebarMenu.vue'
import { computed } from 'vue'

defineOptions({
  name: 'SideBar',
})

const appStore = useAppStore()
const settingsStore = useSettingsStore()
const authStore = useAuthStore()

const sidebarLogo = computed(() => settingsStore.sidebarLogo)
const layout = computed(() => settingsStore.layout)
const isTopLayout = computed(() => layout.value === LayoutMode.TOP)
const isSidebarCollapsed = computed(() => !appStore.sidebar.opened)

const { routes } = authStore

const handleLayoutChange = (layout: LayoutMode) => {
  settingsStore.changeLayout(layout)
}
</script>

<style lang="scss" scoped>
.has-logo {
  .el-scrollbar {
    height: calc(100vh - $navbar-height);
  }
}
</style>
