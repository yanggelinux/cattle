import type { App } from 'vue'
import {
  createRouter,
  createWebHashHistory,
  createWebHistory,
  type RouteRecordRaw,
} from 'vue-router'
import { menuRoutes } from './menuRoute'
import { publicRoutes } from './publicRoute'

// 静态路由
export const constantRoutes: RouteRecordRaw[] = [...publicRoutes, ...menuRoutes]

/**
 * 创建路由
 */
// const router = createRouter({
//   history: createWebHistory(),
//   routes: constantRoutes,
//   // 刷新时，滚动条位置还原
//   scrollBehavior: () => ({ left: 0, top: 0 }),
// })

function genRouter() {
  if (import.meta.env.MODE === 'development') {
    //开发环境
    //import.meta.env.VITE_APP_BASE_API
    return createRouter({
      history: createWebHashHistory('/'),
      routes: constantRoutes,
      // 刷新时，滚动条位置还原
      scrollBehavior: () => ({ left: 0, top: 0 }),
    })
  } else {
    return createRouter({
      history: createWebHistory('/'),
      routes: constantRoutes,
      // 刷新时，滚动条位置还原
      scrollBehavior: () => ({ left: 0, top: 0 }),
    })
  }
}

const router = genRouter()

// 全局注册 router
export function setupRouter(app: App<Element>) {
  app.use(router)
}

export default router
