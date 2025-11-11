import type { App } from 'vue'

import { hasPerm } from './permission'
import { throttleBtn } from './throttle'

// 全局注册 directive
export function setupDirective(app: App<Element>) {
  // 使 v-hasPerm 在所有组件中都可用
  app.directive('hasPerm', hasPerm)
  app.directive('throttleBtn', throttleBtn)
}
