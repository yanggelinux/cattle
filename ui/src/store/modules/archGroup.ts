import { ref } from 'vue'
import { store } from '@/store'
import ArchGroupAPI, {
  type ArchGroupQuery,
  type ArchGroupData,
  type ArchGroupResult,
  type GroupTreeNode,
} from '@/api/arch/group'

import { defineStore } from 'pinia'
import { type AxiosResponse } from 'axios'

export const useArchGroupStore = defineStore('archGroup', () => {
  const archGroupList = ref<ArchGroupResult[]>([])
  const archGroupTree = ref<GroupTreeNode[]>([])

  // 全局获取角色信息
  async function getArchGroupList(queryParams: ArchGroupQuery) {
    try {
      const resp: AxiosResponse = await ArchGroupAPI.getList(queryParams)
      const status = resp.data.status
      const msg = resp.data.msg
      if (status === 200) {
        const data: ArchGroupData = resp.data.data
        archGroupList.value = data.retList
      } else {
        console.log(msg)
      }
    } catch (err) {
      console.log(err)
    }
  }
  async function getArchGroupTree(queryParams: ArchGroupQuery) {
    try {
      const resp: AxiosResponse = await ArchGroupAPI.getList(queryParams)
      const status = resp.data.status
      const msg = resp.data.msg
      if (status === 200) {
        const data: ArchGroupData = resp.data.data
        const retList = data.retList
        archGroupTree.value = buildGroupTree(retList)
      } else {
        console.log(msg)
      }
    } catch (err) {
      console.log(err)
    }
  }

  function buildGroupTree(list: ArchGroupResult[]): GroupTreeNode[] {
    const map = new Map<number, GroupTreeNode>()

    // 先把所有节点放进 map
    list.forEach((item) => {
      map.set(item.id, {
        label: item.groupName,
        value: item.id,
        children: [],
      })
    })

    const roots: GroupTreeNode[] = []

    list.forEach((item) => {
      const node = map.get(item.id)!
      if (item.parentID === 0) {
        // 根节点
        roots.push(node)
      } else {
        // 非根节点，挂到父节点下面
        const parent = map.get(item.parentID)
        if (parent) {
          parent.children = parent.children || []
          parent.children.push(node)
        }
      }
    })

    return roots
  }

  function findPathByValue(tree: GroupTreeNode[], targetValue: number): number[] {
    const path: number[] = []

    function dfs(nodes: GroupTreeNode[], currentPath: number[]): boolean {
      for (const node of nodes) {
        const newPath = [...currentPath, node.value]

        if (node.value === targetValue) {
          path.push(...newPath)
          return true
        }

        if (node.children && node.children.length > 0) {
          if (dfs(node.children, newPath)) {
            return true
          }
        }
      }
      return false
    }

    dfs(tree, [])
    return path
  }

  function getNodePathMap(tree: GroupTreeNode[]): Map<number, string[]> {
    const result = new Map<number, string[]>()

    const dfs = (node: GroupTreeNode, path: string[]) => {
      const currentPath = [...path, node.label]
      // 不管是不是叶子节点，直接记录
      result.set(node.value, currentPath)

      if (node.children && node.children.length > 0) {
        node.children.forEach((child) => dfs(child, currentPath))
      }
    }

    tree.forEach((node) => dfs(node, []))
    return result
  }

  return {
    getArchGroupList,
    getArchGroupTree,
    findPathByValue,
    getNodePathMap,
    archGroupList,
    archGroupTree,
  }
})
export function useArchGroupStoreHook() {
  return useArchGroupStore(store)
}
