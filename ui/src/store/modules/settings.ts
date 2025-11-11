import defaultSettings from '@/settings'
import { LayoutMode } from '@/enums/settings/layout.enum'
import { defineStore } from 'pinia'
import { ref, type Ref } from 'vue'
import { useStorage } from '@vueuse/core'

type SettingsValue = boolean | string

export const useSettingsStore = defineStore('setting', () => {
  // 基本设置
  const settingsVisible = ref(false)
  // 标签视图
  const tagsView = useStorage<boolean>('tagsView', defaultSettings.tagsView)
  // 侧边栏 Logo
  const sidebarLogo = useStorage<boolean>('sidebarLogo', defaultSettings.sidebarLogo)
  // 布局
  const layout = useStorage<LayoutMode>('layout', defaultSettings.layout as LayoutMode)

  // 设置更改函数
  const settingsMap: Record<string, Ref<SettingsValue>> = {
    tagsView,
    sidebarLogo,
    layout,
  }

  function changeSetting({ key, value }: { key: string; value: SettingsValue }) {
    const setting = settingsMap[key]
    if (setting) setting.value = value
  }

  function changeLayout(val: LayoutMode) {
    layout.value = val
  }

  return {
    settingsVisible,
    tagsView,
    sidebarLogo,
    layout,
    changeSetting,
    changeLayout,
  }
})
