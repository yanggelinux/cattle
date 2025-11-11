<template>
  <div class="permission-opt-wrapper">
    <el-drawer v-model="dialog.visible" :title="dialog.title" size="50%" @close="handleCloseDialog">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="父级菜单" prop="parentID">
          <el-tree-select
            v-model="formData.parentID"
            placeholder="选择上级菜单"
            :data="menuOptions"
            filterable
            check-strictly
            :render-after-expand="false"
          ></el-tree-select>
        </el-form-item>

        <el-form-item label="权限名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入权限名称" />
        </el-form-item>
        <el-form-item label="权限类型" prop="permType">
          <el-radio-group v-model="formData.permType" @change="handlePermTypeChange">
            <el-radio :value="1" size="small" border>菜单</el-radio>
            <el-radio :value="2" size="small" border>API</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="权限编码" prop="code">
          <el-input v-model="formData.code" placeholder="请输入权限编码" />
          <!-- <el-input v-if="curPermType === 2" v-model="formData.code" placeholder="请输入权限编码" /> -->
          <!-- <el-select v-else v-model="formData.code" clearable filterable>
            <el-option v-for="item in routeNames" :key="item" :value="item" :label="item" />
          </el-select> -->
        </el-form-item>
        <el-form-item
          v-if="curPermType === 2"
          label="权限URI"
          prop="uri"
          :required="curPermType == 2"
        >
          <el-input v-model="formData.uri" placeholder="请输入权限URI" />
        </el-form-item>
        <el-form-item
          v-if="curPermType === 2"
          label="权限方法"
          prop="method"
          :required="curPermType == 2"
        >
          <el-select v-model="formData.method" clearable filterable placeholder="请选择权限方式">
            <el-option v-for="item in metonds" :key="item" :value="item" :label="item" />
          </el-select>
        </el-form-item>

        <el-form-item label="是否生效" prop="isEnabled">
          <el-radio-group v-model="formData.isEnabled">
            <el-radio :value="1" size="small" border>生效</el-radio>
            <el-radio :value="0" size="small" border>失效</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number
            v-model="formData.sort"
            style="width: 100px"
            controls-position="right"
            :min="0"
          />
        </el-form-item>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="handleSubmit">确 定</el-button>
          <el-button @click="handleCloseDialog">取 消</el-button>
        </div>
      </template>
    </el-drawer>
  </div>
</template>
<script setup lang="ts">
import PermissionAPI, { type PermForm } from '@/api/userPerm/permission'
import { reactive, ref, toRefs, type PropType } from 'vue'
import { ElMessage } from 'element-plus'
import { type AxiosResponse } from 'axios'
import { type OptionType } from '@/types/global.d'

defineOptions({
  name: 'PermissionOpt',
})
//props
const props = defineProps({
  formData: {
    type: Object as PropType<PermForm>,
    required: true,
    default: () => ({}),
  },
  dialog: {
    type: Object,
    required: true,
    default: () => ({
      title: '',
      visible: false,
    }),
  },
  action: {
    type: String,
    required: true,
  },
  menuOptions: {
    type: Object as PropType<OptionType[]>,
    required: true,
  },
})

const curPermType = defineModel('curPermType', {
  type: Number as PropType<number | undefined>,
  required: true,
})

const emit = defineEmits(['submit'])
const metonds: string[] = ['GET', 'POST', 'PUT', 'DELETE']
// ref
const formRef = ref()
// reactive
const checkUri = (rule: any, value: any, callback: any) => {
  if (!value) {
    return callback(new Error('请输入权限URI'))
  }
  setTimeout(() => {
    callback()
  }, 10)
}
const checkMethod = (rule: any, value: any, callback: any) => {
  if (!value) {
    return callback(new Error('请输入请求方法'))
  }
  setTimeout(() => {
    callback()
  }, 10)
}

const rules = reactive({
  parentID: [{ required: true, message: '请选择父级菜单', trigger: 'blur' }],
  name: [{ required: true, message: '请输入权限名称', trigger: 'blur' }],
  code: [{ required: true, message: '请选择权限编码', trigger: 'blur' }],
  uri: [{ validator: checkUri, trigger: 'blur' }],
  method: [{ validator: checkMethod, trigger: 'blur' }],
  permType: [{ required: true, message: '请输选择权限类型', trigger: 'blur' }],
  isEnabled: [{ required: true, message: '请选择显示状态', trigger: 'change' }],
  sort: [{ required: true, message: '请输入排序序号', trigger: 'blur' }],
})

const formData = props.formData
const dialog = props.dialog
// 小心「解构赋值」导致丢失响应式 用 toRefs 或 toRef 保持响应式。
const { action } = toRefs(props)

//method
function handleCloseDialog() {
  dialog.visible = false
  formRef.value.resetFields()
  formRef.value.clearValidate()
  formData.id = undefined
}

// 菜单类型切换
function handlePermTypeChange(val: number) {
  curPermType.value = val
}

/**
 * 提交表单
 */
async function handleCreate() {
  try {
    const resp: AxiosResponse = await PermissionAPI.create(formData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('创建成功')
      handleCloseDialog()
      emit('submit')
    } else {
      console.log(msg)
      ElMessage.error(`创建失败${msg}`)
    }
  } catch (err) {
    console.error(err)
    ElMessage.error('创建失败')
  }
}

async function handleUpdate() {
  try {
    const resp: AxiosResponse = await PermissionAPI.update(formData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('编辑成功')
      handleCloseDialog()
      emit('submit')
    } else {
      console.log(msg)
      ElMessage.error('编辑失败')
    }
  } catch (err) {
    console.error(err)
    ElMessage.error('编辑失败')
  }
}

function handleSubmit() {
  formRef.value?.validate(async (valid: boolean) => {
    if (!valid) return
    if (curPermType.value === 1) {
      formData.method = 'GET'
      formData.uri = formData.code.toLowerCase()
    }
    if (action.value === 'update') {
      handleUpdate()
    } else {
      handleCreate()
    }
  })
}
</script>
<style lang="scss" scoped></style>
