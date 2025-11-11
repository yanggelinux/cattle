<template>
  <div class="demand-opt-wrapper">
    <el-drawer v-model="visible" :title="title" size="40%" @close="handleClose">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="请求名称" prop="name">
          <el-input
            :disabled="action === 'update'"
            v-model="formData.name"
            placeholder="请输入请求名称"
          />
          <span v-if="action === 'create'" style="font-size: 12px; color: red">
            *请求名称会自动添加时间戳和归属人信息*
          </span>
        </el-form-item>
        <el-form-item label="归属人" prop="owner">
          <el-input :disabled="true" v-model="formData.owner" placeholder="请输入归属人" />
        </el-form-item>
        <el-form-item label="请求类型" prop="demandType">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.demandType"
            placeholder="请选择请求类型"
          >
            <el-option
              v-for="item in demandTypeList"
              :value="item.value"
              :label="item.label"
              :key="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="业务组" prop="biz">
          <el-input v-model="formData.biz" placeholder="请输入业务组" />
        </el-form-item>
        <!-- <el-form-item label="OA单号" prop="orderNo">
          <el-input v-model="formData.orderNo" placeholder="请输入OA单号" />
          <span v-if="action === 'create'" style="font-size: 12px; color: red">
            *请求默认是标准生产变更，如果是非标准生产变更请填工单号*
          </span>
        </el-form-item> -->
        <el-form-item label="请求描述" prop="description">
          <el-input
            style="width: 1030px"
            :rows="15"
            type="textarea"
            v-model="formData.description"
            :placeholder="descTemplate"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button v-if="action === 'update'" type="success" @click="handleUpdate(1)">
            重新提交
          </el-button>
          <el-button type="primary" @click="handleSubmit">确 定</el-button>
          <el-button @click="handleClose">取 消</el-button>
        </div>
      </template>
    </el-drawer>
  </div>
</template>
<script setup lang="ts">
import DemandAPI, { type DemandForm } from '@/api/demand'
import { reactive, ref, toRefs, type PropType } from 'vue'
import { ElMessage } from 'element-plus'
import { type AxiosResponse } from 'axios'
import { demandTypeList } from '@/utils/constant'
import { getCurrentMinutesString } from '@/utils'
import { useAuthStore } from '@/store'

const descTemplate = `
请求描述
1、添加nginx转发配置示例：
upstream pensions_upstream {
    server pensions.bigdata.uat;
}
location ~ ^/test/ {
    proxy_pass http://pensions_upstream/;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header Host $http_host;
}
2、创建edas和ecs命名示例：
x-p-prod-业务线-应用名字-ecs
x-p-prod-业务线-应用名字-edas
`

defineOptions({
  name: 'DemandOpt',
})
//props
const props = defineProps({
  formData: {
    type: Object as PropType<DemandForm>,
    required: true,
    default: () => ({}),
  },
  action: {
    type: String,
    required: true,
  },
})

const visible = defineModel('visible', {
  type: Boolean,
  required: true,
  default: true,
})

const { userName } = useAuthStore()

const formData = props.formData
formData.owner = userName
// 小心「解构赋值」导致丢失响应式 用 toRefs 或 toRef 保持响应式。
const { action } = toRefs(props)
const emit = defineEmits(['submit'])
// ref
const formRef = ref()
const title = ref<string>(action.value === 'create' ? '创建请求' : '编辑请求')
// reactive
const rules = reactive({
  name: [{ required: true, message: '请输入请求名称', trigger: 'blur' }],
  // biz: [{ required: true, message: '请选择业务线', trigger: 'blur' }],
  owner: [{ required: true, message: '请输入归属人', trigger: 'blur' }],
  demandType: [{ required: true, message: '请选择请求类型', trigger: 'blur' }],
  description: [{ required: true, message: '请输入请求描述', trigger: 'blur' }],
})

//method
function handleClose() {
  visible.value = false
  formRef.value.resetFields()
  formRef.value.clearValidate()
}

// 提交表单
async function handleCreate() {
  try {
    const owner = formData.owner || ''
    const name = formData.name || ''
    if (!name.includes(owner)) {
      const timestr = getCurrentMinutesString()
      const newName = `${owner}-${name}-${timestr}`
      formData.name = newName
    }
    formData.status = 1
    const resp: AxiosResponse = await DemandAPI.create(formData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('创建成功')
      handleClose()
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

async function handleUpdate(reSubmit?: number) {
  try {
    const stat = formData.status
    // 审批失败的请求再次编辑可以改变状态重新评审
    if (stat === 3 && reSubmit === 1) {
      formData.status = 1
    }
    if (formData.orderNo && formData.orderNo.length < 10) {
      ElMessage.error('OA工单编号不能小于10')
      return
    }
    const resp: AxiosResponse = await DemandAPI.update(formData)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('编辑成功')
      handleClose()
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

async function handleSubmit() {
  formRef.value?.validate(async (valid: boolean) => {
    if (!valid) return
    if (formData.orderNo && formData.orderNo.length < 10) {
      ElMessage.error('OA工单编号不能小于10')
      return
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
