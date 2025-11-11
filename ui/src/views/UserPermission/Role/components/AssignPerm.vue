<template>
  <div class="assign-perm-container">
    <el-drawer
      v-model="permVisible"
      :title="'【' + checkedRole.displayName + '】权限分配'"
      size="30%"
    >
      <div class="flex-x-between">
        <el-input
          prefix-icon="Search"
          v-model="permKeywords"
          clearable
          placeholder="菜单权限名称"
        ></el-input>
        <div class="flex-center">
          <el-button class="expand-btn" type="primary" size="small" plain @click="togglePermTree">
            <template #icon>
              <Switch />
            </template>
            {{ isExpanded ? '收缩' : '展开' }}
          </el-button>
          <el-tooltip
            content="如果只需勾选菜单权限，不需要勾选子菜单或者按钮权限，请关闭父子联动"
            placement="bottom"
          >
            <el-checkbox
              v-model="parentChildLinked"
              class="ml-5"
              @change="handleparentChildLinkedChange"
            >
              父子联动
            </el-checkbox>
          </el-tooltip>
        </div>
      </div>

      <el-tree
        ref="permTreeRef"
        node-key="id"
        show-checkbox
        :data="permOptions"
        :filter-node-method="handlePermFilter"
        :default-expand-all="true"
        :check-strictly="!parentChildLinked"
      >
        <template #default="{ data }">
          <span class="icon-wrapper" :class="`i-svg:menu`" v-if="data.permType === 1"></span>
          <span class="icon-wrapper" :class="`i-svg:api`" v-else></span>
          {{ data.name }}
        </template>
      </el-tree>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="handleAssignPermSubmit">确 定</el-button>
          <el-button @click="permVisible = false">取 消</el-button>
        </div>
      </template>
    </el-drawer>
  </div>
</template>

<script setup lang="ts">
import PermissionAPI, {
  type RolePermData,
  type PermTreeResult,
  type RolePermQuery,
  type RolePermdata,
} from '@/api/userPerm/permission'
import { reactive, ref, watch, type PropType } from 'vue'
import { ElMessage } from 'element-plus'
import { type AxiosResponse } from 'axios'

defineOptions({
  name: 'AssignPerm',
})

// 选中的角色
export interface CheckedRole {
  id: number
  displayName: string
  isSuper: number
}
//props
const props = defineProps({
  checkedRole: {
    type: Object as PropType<CheckedRole>,
    required: true,
    default: () => ({}),
  },
})
const emit = defineEmits(['submit'])
//在 Vue 的响应式更新里，emit('update:xxx', newVal) 只是「发出修改意图」，
// 真正的 prop 值要等到父组件接收到事件、修改自己的状态、再重新渲染子组件后才会更新。
// 在同一个事件回调里马上读 props.xxx，还是旧值。
// nextTick 等下一次 DOM 更新周期 或者 watch 监听 prop 变化
const permVisible = defineModel('permVisible', {
  type: Boolean,
  required: true,
  default: true,
})

const permTreeRef = ref()
const permKeywords = ref<string>('')
const project = ref<string>('cattle')
const isExpanded = ref<boolean>(true)
const parentChildLinked = ref(false)

const permOptions = ref<PermTreeResult[]>([])

const rolePermParams = reactive<RolePermQuery>({
  roleID: 0,
  isSuper: 0,
  project: 'cattle',
})

const checkedRole = props.checkedRole
// 绝不 在子组件里直接修改 props.xxx ,想要修改v-model双向绑定
// 生命周期
watch(permKeywords, (val) => {
  permTreeRef.value!.filter(val)
})

watch(permVisible, (val) => {
  if (val) {
    handleGetPermInfo()
  }
})
// method
async function handleGetPermInfo() {
  const roleID = checkedRole.id
  if (roleID) {
    // 获取所有的菜单
    try {
      rolePermParams.roleID = roleID
      rolePermParams.isSuper = checkedRole?.isSuper
      rolePermParams.project = project.value
      const resp: AxiosResponse = await PermissionAPI.getRolePermList(rolePermParams)
      const status = resp.data.status
      const msg = resp.data.msg
      if (status === 200) {
        const rolePermData: RolePermdata = resp.data.data
        const rolePermIDList: number[] = rolePermData.rolePermIDList
        permOptions.value = rolePermData.allPermTreeData
        // 回显角色已拥有的菜单
        permTreeRef.value.setCheckedKeys(rolePermIDList)
      } else {
        console.log(msg)
      }
    } catch (err) {
      console.log(err)
    }
  }
}

// 分配菜单权限提交
async function handleAssignPermSubmit() {
  const roleID = checkedRole.id
  if (roleID) {
    const permIDList: number[] = permTreeRef
      .value!.getCheckedNodes(false, true)
      .map((node: any) => node.id)

    const rolePermData: RolePermData = {
      roleID: roleID,
      project: project.value,
      permIDList: permIDList,
    }
    try {
      const resp: AxiosResponse = await PermissionAPI.updateRolePerm(rolePermData)
      const status = resp.data.status
      const msg = resp.data.msg
      if (status === 200) {
        ElMessage.success('分配权限成功')
        permVisible.value = false
        emit('submit')
      } else {
        ElMessage.error('分配权限失败')
        console.log(msg)
      }
    } catch (err) {
      console.log(err)
    }
  }
}

// 展开/收缩 菜单权限树
function togglePermTree() {
  isExpanded.value = !isExpanded.value
  if (permTreeRef.value) {
    Object.values(permTreeRef.value.store.nodesMap).forEach((node: any) => {
      if (isExpanded.value) {
        node.expand()
      } else {
        node.collapse()
      }
    })
  }
}

// 父子菜单节点是否联动
function handleparentChildLinkedChange(val: any) {
  parentChildLinked.value = val
}

// 权限筛选

function handlePermFilter(
  value: string,
  data: {
    [key: string]: any
  }
) {
  if (!value) return true
  return data.name.includes(value)
}
</script>

<style lang="scss" scoped>
.assign-perm-container {
  .flex-x-between {
    margin-bottom: 10px;
    display: flex;
    justify-content: space-betwee;
    .flex-center {
      display: flex;
      margin-left: 10px;
      justify-content: space-around;
      .expand-btn {
        margin-top: 4px;
      }
    }
  }
  .icon-wrapper {
    margin-right: 5px;
  }
}
</style>
