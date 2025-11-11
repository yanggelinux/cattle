import { type RouteRecordRaw } from 'vue-router'

export const Layout = () => import('@/layout/index.vue')

export const subMenuRoutes: RouteRecordRaw[] = [
  {
    path: '/dashboard',
    component: () => import('@/views/Dashboard/index.vue'),
    name: 'Dashboard',
    meta: {
      title: '首页',
      icon: 'menu_dashboard',
      affix: true,
      keepAlive: false,
      hidden: false,
    },
  },

  {
    path: '/arch',
    name: 'Arch',
    redirect: '/arch/group',
    meta: {
      title: '架构图管理',
      icon: 'menu_process',
      alwaysShow: true,
      hidden: false,
    },
    children: [
      {
        path: 'group',
        name: 'ArchGroup',
        meta: {
          title: '架构图组',
          icon: 'menu_arch',
          keepAlive: false,
          alwaysShow: false,
          hidden: false,
        },
        component: () => import('@/views/Arch/Group/index.vue'),
      },

      {
        path: 'graph/:groupID',
        name: 'ArchGraph',
        meta: {
          title: '架构图',
          icon: 'menu_graph',
          keepAlive: false,
          alwaysShow: false,
          hidden: true,
        },
        component: () => import('@/views/Arch/Graph/index.vue'),
      },
      {
        path: 'graph-draw/:graphID',
        name: 'GraphDraw',
        meta: {
          title: '架构图绘制',
          icon: 'menu_graph',
          keepAlive: false,
          alwaysShow: false,
          hidden: true,
        },
        component: () => import('@/views/Arch/GraphDraw/index.vue'),
      },
    ],
  },
  {
    path: '/process',
    name: 'ProcessManage',
    redirect: '/process/list',
    meta: {
      title: '流程管理',
      icon: 'menu_process_manage',
      alwaysShow: true,
      hidden: false,
    },
    children: [
      {
        path: 'list',
        name: 'ProcessList',
        meta: {
          title: '流程列表',
          icon: 'menu_process_group',
          keepAlive: false,
          alwaysShow: false,
          hidden: false,
        },
        component: () => import('@/views/Process/List/index.vue'),
      },
      {
        path: 'team',
        name: 'ProcessTeam',
        meta: {
          title: '流程团队',
          icon: 'menu_process_group',
          keepAlive: false,
          alwaysShow: false,
          hidden: false,
        },
        component: () => import('@/views/Process/Team/index.vue'),
      },

      {
        path: 'draw/:processID',
        name: 'ProcessDraw',
        meta: {
          title: '绘制流程',
          icon: 'menu_graph',
          keepAlive: false,
          alwaysShow: false,
          hidden: true,
        },
        component: () => import('@/views/Process/Draw/index.vue'),
      },
    ],
  },
  {
    path: '/demand',
    name: 'DemandManage',
    redirect: '/demand',
    meta: {
      title: '请求管理',
      icon: 'menu_demand',
      keepAlive: false,
      alwaysShow: true,
      hidden: false,
    },
    children: [
      {
        path: 'list',
        name: 'Demand',
        meta: {
          title: '请求列表',
          icon: 'menu_demand',
          keepAlive: false,
          alwaysShow: false,
          hidden: false,
        },
        component: () => import('@/views/Demand/index.vue'),
      },
      {
        path: 'review',
        name: 'DemandReview',
        meta: {
          title: '请求评审',
          icon: 'menu_demand',
          keepAlive: false,
          alwaysShow: false,
          hidden: true,
        },
        component: () => import('@/views/Demand/Review.vue'),
      },
    ],
  },
  {
    path: '/process-order',
    name: 'Process',
    redirect: '/process-order/graph-list',
    meta: {
      title: '流程工单',
      icon: 'menu_process_order',
      alwaysShow: true,
      hidden: false,
    },
    children: [
      {
        path: 'graph-list',
        name: 'ArchGraphList',
        meta: {
          title: '资源工单申请',
          icon: 'menu_graph',
          keepAlive: false,
          alwaysShow: false,
          hidden: false,
        },
        component: () => import('@/views/Arch/Graph/List.vue'),
      },
      {
        path: 'apply-list',
        name: 'OrderApplyList',
        meta: {
          title: '工单申请',
          icon: 'menu_order_apply',
          keepAlive: false,
          alwaysShow: false,
          hidden: false,
        },
        component: () => import('@/views/ProcessOrder/Apply/List.vue'),
      },

      {
        path: 'my-todo',
        name: 'MyTodoOrder',
        meta: {
          title: '我的待办',
          icon: 'menu_order_todo',
          keepAlive: false,
          alwaysShow: false,
          hidden: false,
        },
        component: () => import('@/views/ProcessOrder/Order/MyTodo.vue'),
      },
      {
        path: 'my-create',
        name: 'MyCreateOrder',
        meta: {
          title: '我创建的',
          icon: 'menu_order_create',
          keepAlive: false,
          alwaysShow: false,
          hidden: false,
        },
        component: () => import('@/views/ProcessOrder/Order/MyCreate.vue'),
      },

      {
        path: 'apply/:orderID',
        name: 'OrderApply',
        meta: {
          title: '工单申请详情',
          icon: 'menu_order_apply',
          keepAlive: false,
          alwaysShow: false,
          hidden: true,
        },
        component: () => import('@/views/ProcessOrder/Apply/index.vue'),
      },

      {
        path: 'order',
        name: 'ProcessOrder',
        meta: {
          title: '工单列表',
          icon: 'menu_order_list',
          keepAlive: false,
          alwaysShow: false,
          hidden: false,
        },
        component: () => import('@/views/ProcessOrder/Order/index.vue'),
      },
      {
        path: 'approval',
        name: 'ProcessApproval',
        meta: {
          title: '工单审批',
          icon: 'menu_order_list',
          keepAlive: false,
          alwaysShow: false,
          hidden: true,
        },
        component: () => import('@/views/ProcessOrder/Approval/index.vue'),
      },
    ],
  },

  {
    path: '/order',
    name: 'OrderManage',
    redirect: '/order',
    meta: {
      title: '工单管理',
      icon: 'menu_order_manage',
      alwaysShow: true,
      hidden: false,
    },
    children: [
      {
        path: 'list',
        name: 'Order',
        meta: {
          title: '工单',
          icon: 'menu_order',
          keepAlive: false,
          alwaysShow: false,
          hidden: false,
        },
        component: () => import('@/views/Order/Order/index.vue'),
      },
      {
        path: 'group',
        name: 'OrderGroup',
        meta: {
          title: '工单组',
          icon: 'menu_order_group',
          keepAlive: false,
          alwaysShow: false,
          hidden: false,
        },
        component: () => import('@/views/Order/OrderGroup/index.vue'),
      },

      {
        path: 'field/:orderID',
        name: 'OrderField',
        meta: {
          title: '工单字段',
          icon: 'menu_graph',
          keepAlive: false,
          alwaysShow: false,
          hidden: true,
        },
        component: () => import('@/views/Order/OrderField/index.vue'),
      },
    ],
  },

  {
    path: '/user-permission',
    name: 'UserPermission',
    redirect: '/user-permission/user',
    meta: {
      title: '用户权限',
      icon: 'menu_user_manage',
      alwaysShow: true,
      hidden: false,
    },
    children: [
      {
        path: 'user',
        name: 'User',
        meta: {
          title: '用户管理',
          icon: 'menu_user',
          keepAlive: false,
          alwaysShow: false,
          hidden: false,
        },
        component: () => import('@/views/UserPermission/User/index.vue'),
      },
      {
        path: 'role',
        name: 'Role',
        meta: {
          title: '角色管理',
          icon: 'menu_role',
          keepAlive: false,
          alwaysShow: false,
          hidden: false,
        },
        component: () => import('@/views/UserPermission/Role/index.vue'),
      },
      {
        path: 'permission',
        name: 'Permission',
        meta: {
          title: '权限管理',
          icon: 'menu_perm',
          keepAlive: false,
          alwaysShow: false,
          hidden: false,
        },
        component: () => import('@/views/UserPermission/Permission/index.vue'),
      },
    ],
  },
]

export const menuRoutes: RouteRecordRaw[] = [
  {
    path: '/',
    name: '/',
    meta: { hidden: false },
    component: Layout,
    redirect: '/dashboard',
    children: [...subMenuRoutes],
  },
]
