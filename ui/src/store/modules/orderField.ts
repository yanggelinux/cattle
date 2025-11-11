import { ref } from 'vue'
import { store } from '@/store'
import OrderFieldAPI, {
  type OrderFieldQuery,
  type OrderFieldData,
  type OrderFieldResult,
} from '@/api/order/orderField'

import { defineStore } from 'pinia'
import { type AxiosResponse } from 'axios'

export const useOrderFieldStore = defineStore('orderField', () => {
  const orderFieldList = ref<OrderFieldResult[]>([])
  const total = ref<number>(0)

  async function getOrderFieldList(queryParams: OrderFieldQuery) {
    try {
      const resp: AxiosResponse = await OrderFieldAPI.getList(queryParams)
      const status = resp.data.status
      const msg = resp.data.msg
      if (status === 200) {
        const orderFieldData: OrderFieldData = resp.data.data
        orderFieldList.value = orderFieldData.retList
        total.value = orderFieldData.total
      } else {
        console.log(msg)
      }
    } catch (err) {
      console.log(err)
    }
  }
  return {
    getOrderFieldList,
    orderFieldList,
    total,
  }
})
export function useOrderFieldStoreHook() {
  return useOrderFieldStore(store)
}
