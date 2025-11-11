import { ref } from 'vue'
import { store } from '@/store'
import UserAPI, { type UserQuery, type UserData, type UserResult } from '@/api/userPerm/user'

import { defineStore } from 'pinia'
import { type AxiosResponse } from 'axios'

export const useUserStore = defineStore('user', () => {
  const userList = ref<UserResult[]>([])
  const total = ref<number>(0)

  async function getUserList(queryParams: UserQuery) {
    try {
      const resp: AxiosResponse = await UserAPI.getList(queryParams)
      const status = resp.data.status
      const msg = resp.data.msg
      if (status === 200) {
        const userData: UserData = resp.data.data
        userList.value = userData.retList
        total.value = userData.total
      } else {
        console.log(msg)
      }
    } catch (err) {
      console.log(err)
    }
  }
  return {
    getUserList,
    userList,
    total,
  }
})
export function useUserStoreHook() {
  return useUserStore(store)
}
