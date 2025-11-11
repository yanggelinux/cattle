import { type RouteRecordRaw } from 'vue-router'

export const Layout = () => import('@/layout/index.vue')

export const publicRoutes: RouteRecordRaw[] = [
  {
    path: '/redirect',
    meta: { hidden: true },
    component: Layout,
    children: [
      {
        path: '/redirect/:path(.*)',
        component: () => import('@/views/Redirect/index.vue'),
      },
    ],
  },
  {
    path: '/login',
    name: 'Login',
    meta: { hidden: true },
    component: () => import('@/views/Login/index.vue'),
  },
  {
    path: '/401',
    name: '401',
    meta: { hidden: true },
    component: () => import('@/views/Error/401.vue'),
  },
  {
    path: '/404',
    name: '404',
    meta: { hidden: true },
    component: () => import('@/views/Error/404.vue'),
  },
]
