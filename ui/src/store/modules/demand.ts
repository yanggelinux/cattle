import { ref } from 'vue'
import { store } from '@/store'
import DemandAPI, { type DemandQuery, type DemandData, type DemandResult } from '@/api/demand'

import { defineStore } from 'pinia'
import { type AxiosResponse } from 'axios'

export const useDemandStore = defineStore('demand', () => {
  const demandList = ref<DemandResult[]>([])
  const total = ref<number>(0)

  async function getDemandList(queryParams: DemandQuery) {
    try {
      const resp: AxiosResponse = await DemandAPI.getList(queryParams)
      const status = resp.data.status
      const msg = resp.data.msg
      if (status === 200) {
        const demandData: DemandData = resp.data.data
        demandList.value = demandData.retList
        total.value = demandData.total
      } else {
        console.log(msg)
      }
    } catch (err) {
      console.log(err)
    }
  }
  return {
    getDemandList,
    demandList,
    total,
  }
})
export function useDemandStoreHook() {
  return useDemandStore(store)
}
