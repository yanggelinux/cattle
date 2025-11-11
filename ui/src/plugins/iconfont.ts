import type { App } from 'vue'
import IconSvg from '@/components/IconSvg/index.vue'

export function setupIconfont(app: App<Element>) {
  app.component('IconSvg', IconSvg)
}
