<template>
  <div class="opt-wrapper">
    <el-drawer v-model="visible" :title="title" size="40%" @close="handleCloseDialog">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="140px">
        <el-form-item label="字段名" prop="name">
          <!-- <el-input v-model="formData.name" placeholder="请输入字段名称" /> -->
          <el-autocomplete
            v-model="formData.name"
            :fetch-suggestions="querySearch"
            @select="handleNameSelect"
            clearable
            placeholder="请输入字段名称"
          />
        </el-form-item>

        <el-form-item label="字段标识" prop="key">
          <template #label>
            <el-tooltip
              content="字段标识是工单内容唯一的，使用英文加下划线，例如用户名的唯一标识是user_name,不要写成userName"
              placement="top"
              effect="dark"
            >
              <el-text>
                字段标识
                <el-icon><InfoFilled /></el-icon>
                :
              </el-text>
            </el-tooltip>
          </template>
          <el-input v-model="formData.key" placeholder="请输入字段标识" />
        </el-form-item>

        <el-form-item label="渲染组件" prop="component">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.component"
            placeholder="请选择渲染组件"
          >
            <el-option
              v-for="item in componentList"
              :value="item.value"
              :label="item.label"
              :key="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="枚举值"
          prop="enum"
          v-if="formData.component === 'select' || formData.component === 'multipleSelect'"
        >
          <template #label>
            <el-tooltip content="多个枚举值用换行分隔" placement="top" effect="dark">
              <el-text>
                枚举值
                <el-icon><InfoFilled /></el-icon>
                :
              </el-text>
            </el-tooltip>
          </template>
          <el-input
            type="textarea"
            :autosize="{ minRows: 1, maxRows: 10 }"
            v-model="formData.enum"
            placeholder="请输入枚举值,多个枚举值用换行分隔"
          />
        </el-form-item>
        <el-form-item label="校验规则" prop="verRule">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.verRule"
            placeholder="请选择校验规则"
          >
            <el-option
              v-for="item in verRuleList"
              :value="item.value"
              :label="item.label"
              :key="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="是否必填" prop="isRequired">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.isRequired"
            placeholder="请选择是否必填"
          >
            <el-option
              v-for="item in yesOrNoList"
              :value="item.value"
              :label="item.label"
              :key="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="审批时可编辑" prop="isEdit">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.isEdit"
            placeholder="请选择是否审批时可编辑"
          >
            <el-option
              v-for="item in yesOrNoList"
              :value="item.value"
              :label="item.label"
              :key="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="审批不通过清空" prop="isClear">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.isClear"
            placeholder="请选择是否审批不通过清空"
          >
            <el-option
              v-for="item in yesOrNoList"
              :value="item.value"
              :label="item.label"
              :key="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="占位文本" prop="placeholder">
          <el-input v-model="formData.placeholder" placeholder="请输入占位文本" />
        </el-form-item>
        <el-form-item label="默认值" prop="defaultVal">
          <el-input v-model="formData.defaultVal" placeholder="请输入默认值" />
        </el-form-item>
        <el-form-item label="依赖展示字段" prop="displayField">
          <template #label>
            <el-tooltip
              content="依赖展示字段表示该字段依赖哪个字段展示，不填默认展示"
              placement="top"
              effect="dark"
            >
              <el-text>
                依赖展示字段
                <el-icon><InfoFilled /></el-icon>
                :
              </el-text>
            </el-tooltip>
          </template>
          <el-input v-model="formData.displayField" placeholder="请输入依赖展示字段" />
        </el-form-item>
        <el-form-item label="依赖展示字段值" prop="displayVal" v-if="formData.displayField">
          <template #label>
            <el-tooltip
              content="依赖展示字段只表示该字段依赖哪个字段的值展示，选择‘依赖展示字段’后次字段才生效"
              placement="top"
              effect="dark"
            >
              <el-text>
                依赖展示字段值
                <el-icon><InfoFilled /></el-icon>
                :
              </el-text>
            </el-tooltip>
          </template>
          <el-input v-model="formData.displayVal" placeholder="请输入依赖展示字段值" />
        </el-form-item>

        <el-form-item label="字段分组" prop="groupName">
          <template #label>
            <el-tooltip
              content="默认字段分组为''，如果要对字段分组请输入分组名称"
              placement="top"
              effect="dark"
            >
              <el-text>
                字段分组
                <el-icon><InfoFilled /></el-icon>
                :
              </el-text>
            </el-tooltip>
          </template>
          <el-input v-model="formData.groupName" placeholder="字段分组" />
        </el-form-item>

        <el-form-item label="排序" prop="sort">
          <el-input-number
            v-model="formData.sort"
            style="width: 100px"
            controls-position="right"
            :min="0"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select
            clearable
            filterable
            placement="bottom-end"
            v-model="formData.status"
            placeholder="请选择状态"
          >
            <el-option
              v-for="item in processStatusList"
              :value="item.value"
              :label="item.label"
              :key="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="字段描述" prop="description">
          <el-input
            :autosize="{ minRows: 2, maxRows: 4 }"
            type="textarea"
            v-model="formData.description"
            placeholder="请输入字段描述"
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
import OrderFieldAPI, {
  type OrderFieldData,
  type OrderFieldForm,
  type OrderFieldQuery,
  type OrderFieldResult,
} from '@/api/order/orderField'
import { onMounted, reactive, ref, toRefs, type PropType } from 'vue'
import { ElMessage } from 'element-plus'
import { processStatusList, componentList, verRuleList, yesOrNoList } from '@/utils/constant'
import { type AxiosResponse } from 'axios'
defineOptions({
  name: 'OrderFieldOpt',
})
//props
const props = defineProps({
  formData: {
    type: Object as PropType<OrderFieldForm>,
    required: true,
    default: () => ({}),
  },
  orderID: {
    type: Number,
    required: true,
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

const emit = defineEmits(['submit'])
const formData = props.formData
const orderID = props.orderID
// 小心「解构赋值」导致丢失响应式 用 toRefs 或 toRef 保持响应式。
const { action } = toRefs(props)
// ref
const formRef = ref()
const orderFiels = ref<OrderFieldResult[]>([])

const title = ref<string>(action.value === 'create' ? '创建字段' : '编辑字段')

// reactive
const rules = reactive({
  name: [{ required: true, message: '请输入字段名称', trigger: 'blur' }],
  key: [{ required: true, message: '请输入字段Key', trigger: 'blur' }],
  component: [{ required: true, message: '请选择渲染组件', trigger: 'blur' }],
  verRule: [{ required: true, message: '请选择校验规则', trigger: 'blur' }],
  isRequired: [{ required: true, message: '请选择是否必填', trigger: 'blur' }],
  isTitle: [{ required: true, message: '请选择校是否标题自动', trigger: 'blur' }],
  sort: [{ required: true, message: '请输入排序', trigger: 'blur' }],
  status: [{ required: true, message: '请选择状态', trigger: 'blur' }],
})

onMounted(async () => {
  await getOrderFields()
})

//method

async function getOrderFields() {
  try {
    const params: OrderFieldQuery = {}
    const resp: AxiosResponse = await OrderFieldAPI.getList(params)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: OrderFieldData = resp.data.data
      // const retList = resData.retList
      // for (const ret of retList) {
      //   ret.value = `${ret.name}(${ret.orderName})`
      // }
      orderFiels.value = resData.retList
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

const querySearch = (queryString: string, cb: any) => {
  // 立即执行过滤,不要有异步操作
  const query = queryString.trim().toLowerCase()

  const results = query
    ? orderFiels.value.filter((field) => field.name.toLowerCase().includes(query))
    : orderFiels.value

  const suggestions = results.map((field) => ({
    ...field,
    value: `${field.name}(${field.orderName})`, // 覆盖原有的 value
  }))

  // 确保同步调用 cb
  cb(suggestions)
}

function handleNameSelect(val: any) {
  // 注意这里一定要修改orderID
  val.orderID = orderID
  const id = formData.id
  Object.assign(formData, val)
  // 注意更新的情况，id还是使用原来的id要不然会更新错误
  if (action.value === 'update') {
    formData.id = id
  }
}
function handleCloseDialog() {
  visible.value = false
  formRef.value.resetFields()
  formRef.value.clearValidate()
  formData.id = undefined
}

// 提交表单
async function handleCreate() {
  try {
    const resp: AxiosResponse = await OrderFieldAPI.create(formData)
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
    const resp: AxiosResponse = await OrderFieldAPI.update(formData)
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

async function handleSubmit() {
  formRef.value?.validate(async (valid: boolean) => {
    if (!valid) return
    if (action.value === 'update') {
      handleUpdate()
    } else {
      handleCreate()
    }
    // 重新获取一遍
    await getOrderFields()
  })
}
</script>
<style lang="scss" scoped></style>
