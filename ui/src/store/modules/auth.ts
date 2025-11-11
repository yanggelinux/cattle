import { ref } from 'vue'
import type { RouteRecordRaw } from 'vue-router'
import { store } from '@/store'
import router from '@/router'
import { subMenuRoutes } from '@/router/menuRoute'
import { publicRoutes } from '@/router/publicRoute'
import { useStorage } from '@vueuse/core'
import AuthAPI, {
  type LoginFormData,
  type LoginResult,
  type GetRolePermParams,
  type RolePermResult,
} from '@/api/auth'
import { defineStore } from 'pinia'
import { type AxiosResponse } from 'axios'

// useStorage 是基于 ref 和 Vue 的响应式机制构建的，在使用 setUserInfo() 方法中你连续对多个 useStorage 变量赋值，这在 Vue 中是响应式的，
// 但不是立即同步写入到 localStorage 的。如果你之后立刻跳转或刷新页面，有可能来不及存入 localStorage。
export const useAuthStore = defineStore('auth', () => {
  const userID = useStorage<number>('userID', 0)
  const userName = useStorage<string>('userName', '')
  const displayName = useStorage<string>('displayName', '')
  const email = useStorage<string>('email', '')
  const deptName = useStorage<string>('deptName', '')
  const project = useStorage<string>('project', '')
  const token = useStorage<string>('token', '')
  const authorization = useStorage<string>('authorization', '')
  const isSuper = useStorage<number>('isSuper', 0)
  const roleNames = useStorage<string>('roleNames', '')
  const roleDisplayNames = useStorage<string>('roleDisplayNames', '')
  const menus = useStorage<string[]>('menus', [])
  const uris = useStorage<string[]>('uris', [])

  const isPermsLoaded = ref(false)
  const routes = ref<RouteRecordRaw[]>([])
  // 路由是否加载完成
  const isRoutesLoaded = ref(false)

  async function login(loginData: LoginFormData): Promise<void> {
    try {
      const resp: AxiosResponse = await AuthAPI.login(loginData)
      const status: number = resp.data.status
      const msg: string = resp.data.msg
      if (status === 200) {
        const userInfo: LoginResult = resp.data.data
        setUserInfo(userInfo)
        // 登录后重新生成路由数据
        genRoutes(subMenuRoutes)
      } else {
        // 也可以抛出错误，以便调用方处理
        throw new Error(msg)
      }
    } catch (error) {
      // 抛出异常由调用者处理
      throw error
    }
  }

  /**
   * 登出
   */
  function logout() {
    clearSessionAndCache()
    router.push('/login')
  }

  // 获取用户的角色权限
  async function getRolePermissions() {
    // 获取用户 ID
    const getRolePermParams: GetRolePermParams = { userID: userID.value }
    try {
      const resp: AxiosResponse = await AuthAPI.getRolePerm(getRolePermParams)
      const status: number = resp.data.status
      const msg: string = resp.data.msg
      if (status === 200) {
        const rolePerm: RolePermResult = resp.data.data
        setRoleRerm(rolePerm)
        isPermsLoaded.value = true
      } else {
        console.log(msg)
      }
    } catch (error: any) {
      console.log('获取权限信息失败', error)
    }
  }

  async function generateRoutes() {
    try {
      await getRolePermissions()
      // 生成路由数据
      genRoutes(subMenuRoutes)
      return routes.value
    } catch (error: any) {
      console.log('生成路由信息失败', error)
    }
  }

  function genRoutes(menuRoutes: RouteRecordRaw[]) {
    let newMenuRoutes: RouteRecordRaw[] = []
    // 超级管理员拥有所有权限
    if (isSuper.value === 1) {
      newMenuRoutes = menuRoutes
    } else {
      newMenuRoutes = getMenuList(menuRoutes)
    }
    // 更新路由状态
    routes.value = [...newMenuRoutes, ...publicRoutes]
    isRoutesLoaded.value = true
  }

  function getMenuList(routerMenus: RouteRecordRaw[] = []): RouteRecordRaw[] {
    //处理路由信息生成菜单列表
    const menuList: RouteRecordRaw[] = []
    routerMenus.forEach((item: any) => {
      if (!item.meta.hidden && menus.value.includes(item.name)) {
        const newItem = { ...item }
        if (item.children) {
          delete newItem.children
          newItem.children = getMenuList(item.children)
        }
        menuList.push(newItem)
      }
    })
    return menuList
  }
  /**
   * 重置路由
   */
  const resetRouter = () => {
    // 清空本地存储的路由和菜单数据
    routes.value = []
    isRoutesLoaded.value = false
  }

  function setRoleRerm(rolePerm: RolePermResult) {
    roleNames.value = rolePerm.roleNames
    roleDisplayNames.value = rolePerm.roleDisplayNames
    isSuper.value = rolePerm.isSuper
    menus.value = rolePerm.menus
    uris.value = rolePerm.uris
    //
    localStorage.setItem('isSuper', String(rolePerm.isSuper))
    localStorage.setItem('roleNames', rolePerm.roleNames)
    localStorage.setItem('roleDisplayNames', rolePerm.roleDisplayNames)
    localStorage.setItem('menus', JSON.stringify(rolePerm.menus))
    localStorage.setItem('uris', JSON.stringify(rolePerm.uris))
  }

  function setUserInfo(userInfo: LoginResult) {
    userID.value = userInfo.userID
    userName.value = userInfo.userName
    displayName.value = userInfo.displayName
    email.value = userInfo.email
    deptName.value = userInfo.deptName
    project.value = userInfo.project
    token.value = userInfo.token
    authorization.value = userInfo.authorization
    roleNames.value = userInfo.roleNames
    roleDisplayNames.value = userInfo.roleDisplayNames
    isSuper.value = userInfo.isSuper
    menus.value = userInfo.menus
    uris.value = userInfo.uris
    // 再次填入
    localStorage.setItem('userID', String(userInfo.userID))
    localStorage.setItem('userName', userInfo.userName)
    localStorage.setItem('displayName', userInfo.displayName)
    localStorage.setItem('email', userInfo.email)
    localStorage.setItem('deptName', userInfo.deptName)
    localStorage.setItem('project', userInfo.project)
    localStorage.setItem('token', userInfo.token)
    localStorage.setItem('authorization', userInfo.authorization)
    localStorage.setItem('isSuper', String(userInfo.isSuper))
    localStorage.setItem('roleNames', userInfo.roleNames)
    localStorage.setItem('roleDisplayNames', userInfo.roleDisplayNames)
    localStorage.setItem('menus', JSON.stringify(userInfo.menus))
    localStorage.setItem('uris', JSON.stringify(userInfo.uris))
  }

  function clearUserInfo() {
    localStorage.removeItem('userID')
    localStorage.removeItem('userName')
    localStorage.removeItem('displayName')
    localStorage.removeItem('email')
    localStorage.removeItem('deptName')
    localStorage.removeItem('project')
    localStorage.removeItem('token')
    localStorage.removeItem('authorization')
    localStorage.removeItem('isSuper')
    localStorage.removeItem('roleNames')
    localStorage.removeItem('roleDisplayNames')
    localStorage.removeItem('menus')
    localStorage.removeItem('uris')
  }

  function setProject(projectName: string) {
    project.value = projectName
    localStorage.setItem('project', projectName)
  }
  /**
   * 清除用户会话和缓存
   */
  function clearSessionAndCache() {
    return new Promise<void>((resolve) => {
      clearUserInfo()
      resetRouter()
      resolve()
    })
  }

  return {
    userID,
    userName,
    displayName,
    email,
    deptName,
    project,
    token,
    authorization,
    isSuper,
    roleNames,
    roleDisplayNames,
    menus,
    uris,
    isPermsLoaded,
    isRoutesLoaded,
    routes,
    login,
    logout,
    generateRoutes,
    resetRouter,
    getRolePermissions,
    clearSessionAndCache,
    setProject,
  }
})
export function useAuthStoreHook() {
  return useAuthStore(store)
}
