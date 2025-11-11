<template>
  <div class="order-field-container app-container">
    <div class="order-field-header">
      <el-row>
        <el-col :span="12">
          <div class="info-wrapper">
            <span class="icon-wrapper" :class="`i-svg:field`"></span>
            <el-text class="info-item">工单名称：</el-text>
            <el-tag>{{ orderName }}</el-tag>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="btn-wrapper">
            <el-button
              v-hasPerm="['/order-field:post']"
              type="success"
              icon="plus"
              @click="handleOpenDialog('create')"
            >
              创建字段
            </el-button>
            <el-button icon="Back" @click="handleGoBack">返回</el-button>
          </div>
        </el-col>
      </el-row>
    </div>
    <div class="order-field-main">
      <div class="filter-wrapper">
        <el-collapse accordion>
          <el-collapse-item>
            <template #title>
              <span class="filter-text">
                表格显示字段选择
                <el-icon><CaretBottom /></el-icon>
              </span>
            </template>
            <el-checkbox-group class="filter-group" v-model="checkboxVal" :min="1">
              <el-checkbox v-for="item in tableHeadOptions" :key="item" :value="item">
                <span>{{ tableItemMapping.get(item) }}</span>
              </el-checkbox>
            </el-checkbox-group>
          </el-collapse-item>
        </el-collapse>
      </div>
      <div class="table-wrapper">
        <el-table
          ref="dataTableRef"
          v-loading="loading"
          :data="orderFieldList"
          highlight-current-row
          borderField:true
        >
          <el-table-column label="序号" type="index" :index="indexMethod" width="60" />
          <el-table-column
            v-for="item in tableHead"
            :key="item"
            :label="tableItemMapping.get(item)"
            :prop="item"
          >
            <template #default="scope">
              <span v-if="item === 'status'">{{ processStatusMapping[scope.row[item]] }}</span>
              <span v-else-if="item === 'component'">
                {{ componentMapping[scope.row[item]] }}
              </span>
              <span v-else-if="item === 'verRule'">
                {{ verRuleMapping[scope.row[item]] }}
              </span>
              <span v-else-if="item === 'isRequired'">
                {{ yesOrNoMapping[scope.row[item]] }}
              </span>
              <span v-else-if="item === 'isTitle'">
                {{ yesOrNoMapping[scope.row[item]] }}
              </span>
              <span v-else>{{ scope.row[item] }}</span>
            </template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" width="220">
            <template #default="scope">
              <el-button
                v-hasPerm="['/order-field:put']"
                type="primary"
                size="small"
                link
                icon="edit"
                @click="handleOpenDialog('update', scope.row)"
              >
                编辑
              </el-button>
              <el-button
                v-hasPerm="['/order-field:delete']"
                type="danger"
                size="small"
                link
                icon="delete"
                @click="handleDelete(scope.row)"
              >
                删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
    <!-- 表单弹窗 -->
    <OrderFieldOpt
      :formData="formData"
      :action="action"
      :orderID="orderID"
      v-model:visible="visible"
      @submit="handleQuery"
    ></OrderFieldOpt>
  </div>
</template>

<script setup lang="ts">
import OrderFieldAPI, {
  type OrderFieldQuery,
  type OrderFieldResult,
  type OrderFieldData,
  type OrderFieldForm,
} from '@/api/order/orderField'
import { onMounted, reactive, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useRoute } from 'vue-router'
import { type AxiosResponse } from 'axios'
import OrderFieldOpt from './components/OrderFieldOpt.vue'
import {
  processStatusMapping,
  componentMapping,
  yesOrNoMapping,
  verRuleMapping,
} from '@/utils/constant'
import { indexMethod } from '@/utils/view'
import router from '@/router'

defineOptions({
  name: 'OrderField',
})
const route = useRoute()
const orderID = route.query.orderID ? Number(route.query.orderID) : 0
const orderName = route.query.orderName ? String(route.query.orderName) : ''

const tableItemMapping: Map<string, string> = new Map<string, string>([
  ['name', '字段名'],
  ['key', '字段Key'],
  ['component', '渲染组件'],
  ['placeholder', '占位文本'],
  ['verRule', '校验规则'],
  ['defaultVal', '默认值'],
  ['isRequired', '是否必填'],
  ['isTitle', '是否标题字段'],
  ['displayField', '依赖展示字段'],
  ['displayVal', '依赖展示字段值'],
  ['description', '描述'],
  ['enum', '枚举值'],
  ['groupName', '字段分组'],
  ['sort', '排序'],
  ['status', '状态'],
  ['updatedTime', '更新时间'],
  ['createdTime', '创建时间'],
])
const tableHead = ref<string[]>([
  'name',
  'key',
  'component',
  'isRequired',
  'groupName',
  'sort',
  'status',
  'createdTime',
])
const tableHeadOptions = ref<string[]>([
  'name',
  'key',
  'component',
  'placeholder',
  'verRule',
  'defaultVal',
  'isRequired',
  'isTitle',
  'displayField',
  'displayVal',
  'enum',
  'description',
  'groupName',
  'sort',
  'status',
  'updatedTime',
  'createdTime',
])
const checkboxVal = ref<string[]>(tableHead.value)

//ref
const loading = ref<boolean>(false)
const visible = ref<boolean>(false)
// 编辑还是新增
const action = ref<string>('')
const orderFieldList = ref<OrderFieldResult[]>([])

// reactive
const queryParams = reactive<OrderFieldQuery>({
  orderID: orderID,
})

// 表单
const formData = reactive<OrderFieldForm>({
  orderID: orderID,
})
// 生命周期
onMounted(() => {
  handleQuery()
})
watch(
  () => checkboxVal.value,
  (valArr) => {
    tableHead.value = tableHeadOptions.value.filter((i) => valArr.indexOf(i) >= 0)
  },
  { immediate: true }
)
// method
// 查询
async function handleQuery() {
  try {
    const resp: AxiosResponse = await OrderFieldAPI.getList(queryParams)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData: OrderFieldData = resp.data.data
      orderFieldList.value = resData.retList
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}
// 打开弹窗
async function handleOpenDialog(opt: string, rowData?: OrderFieldResult) {
  visible.value = true
  action.value = opt
  if (opt === 'update') {
    Object.assign(formData, rowData)
  }
}

function handleDelete(rowData: OrderFieldResult) {
  ElMessageBox.confirm(`确定要删除${rowData?.name}吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      loading.value = true
      try {
        const id: number = rowData.id ? rowData.id : 0
        const resp: AxiosResponse = await OrderFieldAPI.delete(id)
        const status = resp.data.status
        const msg = resp.data.msg
        if (status === 200) {
          ElMessage.success('删除成功')
          handleQuery()
        } else {
          console.log(msg)
          ElMessage.error('删除失败')
        }
      } catch (err) {
        console.error(err)
        ElMessage.error('删除失败')
      } finally {
        loading.value = false
      }
    })
    .catch(() => {
      // 点击“取消” 或 关闭弹窗，什么都不做即可避免报错
    })
}
function handleGoBack() {
  router.go(-1)
}
</script>

<style lang="scss" scoped>
.order-field-container {
  .order-field-header {
    background-color: #fff;
    height: 60px;
    line-height: 60px;
    padding: 0px 20px 0px 20px;
    margin-bottom: 10px;
    border: 1px solid var(--el-border-color-light);
    border-radius: 4px;
    box-shadow: var(--el-box-shadow-light);
    .info-wrapper {
      display: flex;
      align-items: center;
      .icon-wrapper {
        font-size: 40px;
      }
      .info-item {
        // margin-right: 10px;
        margin-left: 10px;
      }
    }

    .btn-wrapper {
      text-align: right;
      .btn-space-wrapper {
        margin-top: 10px !important;
      }
    }
  }
  .order-field-main {
    .filter-wrapper {
      text-align: left;
      .filter-text {
        margin-left: 10px;
        font-size: 14px;
        color: #409eff;
      }
      .filter-group {
        margin-left: 10px;
      }
    }
    .table-wrapper {
      margin-top: 3px;
    }
  }
}
</style>
