import { ref } from 'vue'
import { store } from '@/store'
import RoleAPI, { type RoleQuery, type RoleData, type RoleResult } from '@/api/userPerm/role'

import { defineStore } from 'pinia'
import { type AxiosResponse } from 'axios'

export const useRoleStore = defineStore('role', () => {
  const roleList = ref<RoleResult[]>([])
  const total = ref<number>(0)
  const roleNameMapping = ref<Map<string, string>>(new Map<string, string>())

  // 全局获取角色信息
  async function getRoleList(queryParams: RoleQuery) {
    try {
      const resp: AxiosResponse = await RoleAPI.getList(queryParams)
      const status = resp.data.status
      const msg = resp.data.msg
      if (status === 200) {
        const roleData: RoleData = resp.data.data
        const retList = roleData.retList
        roleList.value = retList
        total.value = roleData.total
        for (const ret of retList) {
          roleNameMapping.value.set(ret.roleName, ret.displayName)
        }
      } else {
        console.log(msg)
      }
    } catch (err) {
      console.log(err)
    }
  }

  return {
    getRoleList,
    roleList,
    total,
    roleNameMapping,
  }
})
export function useRoleStoreHook() {
  return useRoleStore(store)
}
