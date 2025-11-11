<template>
  <el-config-provider :locale="locale" :size="size">
    <router-view />
  </el-config-provider>
</template>

<script setup lang="ts">
import { useAppStore } from '@/store'
import { ComponentSize } from '@/enums/settings/layout.enum'
import { computed } from 'vue'

const appStore = useAppStore()

const locale = computed(() => appStore.locale)
const size = computed(() => appStore.size as ComponentSize)
</script>
