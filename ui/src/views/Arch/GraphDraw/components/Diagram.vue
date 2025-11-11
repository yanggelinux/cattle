<template>
  <div class="diagram">
    <div class="diagram-main">
      <!-- 侧边栏拖拽节点 -->
      <DiagramSidebar
        class="diagram-sidebar"
        @dragInNode="dragInNode"
        v-if="isSilentMode === false"
      />
      <!-- 画布容器 -->
      <div class="diagram-container">
        <div
          ref="diagramWrapperRef"
          class="diagram-wrapper"
          v-loading="loading"
          :element-loading-text="loadingText"
        >
          <div ref="diagramRef" class="lf-diagram"></div>
        </div>
      </div>
      <!-- 右侧属性面板 -->
      <div class="diagram-panel-wrapper" v-if="showPanel !== '' && isSilentMode === false">
        <PropertyPanel
          :lf="lf"
          :activeEdges="activeEdges"
          v-if="showPanel === 'property'"
          :onlyEdge="activeNodes.length === 0"
          :elementsStyle="properties"
          @setStyle="setStyle"
          @setZIndex="setZIndex"
          @close="closePanel"
        />
      </div>
      <div class="order-apply-panel" v-if="isSilentMode === true && curHasOrder === 1">
        <ApplyPanel
          v-if="applyVisible"
          :graphID="graphID"
          :graphName="graphName"
          v-model:curOrderID="curOrderID"
          v-model:curOrderName="curOrderName"
          v-model:graphNodeInfo="graphNodeInfo"
          v-model:orderVisible="orderVisible"
          @handleApprove="handleApprove"
        ></ApplyPanel>
      </div>
      <!-- 快照 -->
      <GraphRecord
        :graphID="graphID"
        v-model:visible="recordVisible"
        @select-record="handleSelectRecord"
      ></GraphRecord>
      <!-- 审批组件 -->
      <ApprovalPanel
        :graphID="graphID"
        :graphName="graphName"
        :imageData="imageData"
        :enabledImageData="enabledImageData"
        v-model:approvalVisible="approvalVisible"
        @handleApprove="handleApprove"
      ></ApprovalPanel>
      <ReviewPanel
        :status="status"
        :graphID="graphID"
        :graphKey="graphKey"
        :graphName="graphName"
        v-model:visable="reviewVisible"
      ></ReviewPanel>
      <LinkGraph
        v-if="linkVisible"
        v-model:visible="linkVisible"
        :graphID="linkGraphID"
        :groupID="linkGroupID"
        :nodeID="linkNodeID"
        @handleSubmit="handleLinkNode"
      ></LinkGraph>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, nextTick, toRefs, onBeforeUnmount, computed } from 'vue'
import LogicFlow from '@logicflow/core'
import {
  SelectionSelect,
  Menu,
  Control,
  Snapshot,
  DynamicGroup,
  type ToImageOptions,
} from '@logicflow/extension'
import '@logicflow/core/es/index.css'
import '@logicflow/extension/lib/style/index.css'
import { GridBackgroundImage } from '../constant/index.ts'
import DiagramSidebar from './DiagramSidebar.vue'
import PropertyPanel from './PropertyPanel.vue'
import GraphRecord from './GraphRecord.vue'
import ApprovalPanel from './ApprovalPanel.vue'
import ApplyPanel from './ApplyPanel.vue'
import ReviewPanel from './ReviewPanel.vue'
import LinkGraph from './LinkGraph.vue'
import { registerCustomElement } from './node/index.ts'
import { type AxiosResponse } from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'
// import router from '@/router'
import ArchGraphAPI, {
  type ArchGraphDetailResult,
  type ArchGraphRecordResult,
  type ArchGraphForm,
  type ArchGraphResult,
  type ArchGraphReviewQuery,
} from '@/api/arch/graph'
import { useAuthStore } from '@/store/modules/auth.ts'
import {
  type LabelValue,
  type StringLabelValuesMaping,
  type StringStringMaping,
} from '@/utils/constant.ts'
import { usedNodes, otherNodes, type Node } from '../constant/index'
import { useRouter } from 'vue-router'
import OrderAPI, { type OrderQuery } from '@/api/order/order.ts'
// import { useAutoSave } from './hooks/useAutoSave.ts'
const router = useRouter()

type GraphData = {
  nodes: any[]
  edges: any[]
}

const props = defineProps({
  graphID: {
    required: true,
    type: Number,
  },
  silentMode: {
    required: true,
    type: Number,
  },
  hasOrder: {
    required: true,
    type: Number,
  },
  status: {
    required: true,
    type: Number,
  },
  graphName: {
    required: true,
    type: String,
  },
  owner: {
    required: true,
    type: String,
  },
})
const { graphID, graphName, silentMode, status, hasOrder } = toRefs(props)
const { isSuper } = useAuthStore()

const silent = computed(() => {
  // 审批中的不能编辑
  if (silentMode.value === 1) {
    return true
  }
  if (status.value === 1) {
    return true
  }
  if (isSuper) {
    return false
  }
  return false
})
// refs
const diagramRef = ref<HTMLElement>()

const lf = ref<LogicFlow>()
const activeNodes = ref<any[]>([])
const activeEdges = ref<any[]>([])
const selectedNode = ref<any>({})
const showPanel = ref<string>('')
const imageData = ref<string>('')
const curImgHash = ref<string>('')
const curDragNodeText = ref<string>('')
const nodeType = ref<string>('')
const loading = ref<boolean>(false)
const loadingText = ref<string>('loading...')
const graphLabel = ref<string>('')
const graphKey = ref<string>('')
const approveStatus = ref<number[]>([0, 3])
// 画布是否可编辑
const isSilentMode = ref<boolean>(silent.value)
const isSelection = ref<boolean>(false)
const recordVisible = ref<boolean>(false)
const reviewVisible = ref<boolean>(false)
const hasReview = ref<boolean>(false)

// 生效图选择

const enabledGraphID = ref<number>(0)
const enabledImageData = ref<string>('')

// 审批
const approvalVisible = ref<boolean>(false)
const applyVisible = ref<boolean>(true)
const orderVisible = ref<boolean>(false)
const linkVisible = ref<boolean>(false)
const graphNodeInfo = ref<string>('')

const linkGroupID = ref<number>(0)
const linkGraphID = ref<number>(0)
const linkNodeID = ref<string>('')

const curOrderID = ref<number>(0)
const curOrderName = ref<string>('')
const messageInstance = ref<any>(null)

const curHasOrder = ref<number>(hasOrder.value)

const properties = reactive<Record<string, any>>({})
const nodeTypeNameMapping = reactive<StringStringMaping>({})
const orderTypeInfoMapping = ref<StringLabelValuesMaping>({})

//从路由中获取图ID，所有的图中点和边都必须有这个id,新创建的图id由保存后返回
// const graphKey = ref<string>(`graphData-${graphID.value}`)

defineOptions({
  name: 'Diagram',
})

// 执行自动存储操作，1000 * 60 * 10  10分组自动存储一次,状态是未审批，创建人必须是自己
// if (status.value === 0 && owner.value === userName) {
//   useAutoSave(doSave, () => lf?.value, 6000, false)
// }

onMounted(async () => {
  if (hasOrder.value == 1) {
    await getOrderNodes()
  }
  if (isSilentMode.value) {
    messageInstance.value = ElMessage.info({
      grouping: true,
      showClose: true,
      duration: 5000,
      message: '架构图只读模式！',
    })
  }
  genNodeTypeNameMapping()
  await genGraphData()
  if (!isSilentMode.value) {
    await getEnableGraphData()
  }
})
onBeforeUnmount(() => {
  // 在组件卸载前销毁实例，避免后续还去用已经移除的container
  lf.value?.destroy()
  if (isSilentMode.value && messageInstance) {
    messageInstance.value.close()
  }
})

function initLogicFlow(data: GraphData) {
  let silentConfig = {}
  if (isSilentMode.value) {
    silentConfig = {
      stopZoomGraph: false,
      stopScrollGraph: false,
      stopMoveGraph: false,
      adjustEdge: false,
      adjustEdgeStartAndEnd: false,
      adjustNodePosition: false,
      hideAnchors: true,
      nodeSelectedOutline: true,
      nodeTextEdit: false,
      edgeTextEdit: false,
      nodeTextDraggable: false,
      edgeTextDraggable: false,
    }
  }
  const instance = new LogicFlow({
    ...silentConfig,
    container: diagramRef.value as HTMLElement,
    overlapMode: 1,
    // isSilentMode: isSilentMode.value,
    animation: true,
    keyboard: { enabled: true },
    plugins: [SelectionSelect, Menu, Control, Snapshot, DynamicGroup],
    grid: {
      visible: false,
      size: 5,
    },
    // 允许图标缩放
    allowResize: true,
    // 允许旋转图标
    allowRotate: true,
    // 不允许拖动画布
    stopMoveGraph: isSelection.value,
    background: {
      backgroundImage: GridBackgroundImage,
      backgroundRepeat: 'repeat',
    },
  })
  instance.setTheme({
    baseEdge: { strokeWidth: 1 },
    baseNode: { strokeWidth: 1 },
  })
  registerCustomElement(instance)
  instance.setDefaultEdgeType('pro-line')
  // 禁用全局CSS规则
  // instance.extension.snapshot.useGlobalRules = false

  instance.on('blank:click', async () => {
    if (showPanel.value === 'property') {
      showPanel.value = ''
    } else {
      // showPanel.value = 'property'
    }

    await nextTick()
    const { nodes, edges } = instance.getSelectElements()
    activeNodes.value = nodes
    activeEdges.value = edges
    updateProperties()
  })

  instance.on('selection:selected,node:click,edge:click', async () => {
    instance.graphModel.textEditElement?.setElementState(2)
    showPanel.value = 'property'
    await nextTick()
    const { nodes, edges } = instance.getSelectElements()
    activeNodes.value = nodes
    activeEdges.value = edges
    updateProperties()
    // eventToSave()
  })
  instance.on('node:dbclick', async ({ data }) => {
    selectedNode.value = data
    nodeType.value = data.type
    if (nodeType.value === 'archGraph') {
      instance.graphModel.textEditElement?.setElementState(1)
      if (data.properties) {
        const linkGraph = data.properties ? data.properties.linkGraph : null
        if (linkGraph) {
          handleToGraph(linkGraph)
        }
      }
      return
    }
    // showPanel.value = 'data'
    // eventToSave()
  })

  instance.on('node:mouseenter', async ({ data }) => {
    if (!(isSilentMode.value && hasOrder.value === 1 && status.value === 2)) {
      // 如果是架构图组件添加关联架构图功能
      return
    }
    const ntype = data.type
    const nodeMenus: any[] = []
    instance.setMenuConfig({
      nodeMenu: nodeMenus,
    })
    const orders: LabelValue[] = orderTypeInfoMapping.value[ntype]
    // orders 可能为空
    if (!orders) {
      return
    }

    orders.forEach((order) => {
      nodeMenus.push({
        text: order.label,
        callback(node: any) {
          curHasOrder.value = 1
          genNodeInfoByNode(node)
          curOrderID.value = order.value
          curOrderName.value = order.label
          orderVisible.value = true
        },
      })
    })
    instance.setMenuConfig({
      nodeMenu: nodeMenus,
    })
  })

  instance.on('edge:mouseenter', async () => {
    if (!(isSilentMode.value && hasOrder.value === 1 && status.value === 2)) {
      return
    }
    const edgeMenus: any[] = []
    instance.setMenuConfig({
      edgeMenu: edgeMenus,
    })
    const orders: LabelValue[] = orderTypeInfoMapping.value['line']
    // orders 可能为空
    if (!orders) {
      return
    }

    orders.forEach((order) => {
      edgeMenus.push({
        text: order.label,
        callback(edge: any) {
          curHasOrder.value = 1
          genNodeInfoByEdge(edge)
          curOrderID.value = order.value
          curOrderName.value = order.label
          orderVisible.value = true
        },
      })
    })
    instance.setMenuConfig({
      edgeMenu: edgeMenus,
    })
  })

  addControlItem(instance)
  addMenuItem(instance)

  instance.render(data)
  lf.value = instance
}

function dealNodeText(text: string): string {
  const texts = text.split('\n')
  return texts.join(',')
}

function genNodeTypeNameMapping() {
  const nodes: Node[] = [...usedNodes, ...otherNodes]
  for (const node of nodes) {
    nodeTypeNameMapping[node.type] = node.text
  }
}

function genNodeInfoByNode(node: any) {
  const text = node.text
  if (text) {
    const name = dealNodeText(text.value)
    const info = `节点类型:${nodeTypeNameMapping[node.type]},名称:${name}`
    graphNodeInfo.value = info
    return
  }
  graphNodeInfo.value = ''
}

function genNodeInfoByEdge(edge: any) {
  let sourceInfo = ''
  let targetInfo = ''
  const sourceID = edge.sourceNodeId
  const targetID = edge.targetNodeId
  const sourceNode = lf.value?.getNodeDataById(sourceID)
  const targetNode = lf.value?.getNodeDataById(targetID)
  const sourceText = sourceNode?.text
  const targetText = targetNode?.text
  if (sourceText) {
    const sourceName = dealNodeText(sourceText.value)
    sourceInfo = `节点类型:${nodeTypeNameMapping[sourceNode.type]},名称:${sourceName}`
  }
  if (targetText) {
    const targetName = dealNodeText(targetText.value)
    targetInfo = `节点类型:${nodeTypeNameMapping[targetNode.type]},名称:${targetName}`
  }
  const info = `源节点:${sourceInfo} -----> 目标节点:${targetInfo}`
  graphNodeInfo.value = info
}

//生成图所需的数据
async function genGraphData() {
  //首先从 sessionStorage 获取数据

  // const localGraphData = sessionStorage.getItem(graphKey.value)
  // if (localGraphData) {
  //   const initialData = JSON.parse(localGraphData)
  //   initLogicFlow(initialData)
  //   return
  // }
  //从后端获取数据并init
  await getGrephData()
}

async function getGrephData() {
  try {
    const resp: AxiosResponse = await ArchGraphAPI.getDetail(graphID.value)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const detailData: ArchGraphDetailResult = resp.data.data
      curImgHash.value = detailData.imageHash
      graphLabel.value = detailData.graphLabel
      graphKey.value = detailData.graphKey
      const newGraphData: GraphData = {
        edges: detailData.edgeData,
        nodes: detailData.nodeData,
      }
      initLogicFlow(newGraphData)
    } else {
      console.log(msg)
      ElMessage.error(`获取图数据失败:${msg}`)
    }
  } catch (err) {
    console.log(err)
    ElMessage.error(`获取图数据失败:${err}`)
  }
}
// 获取审批生效的图信息
async function getEnableGraphData() {
  try {
    const resp: AxiosResponse = await ArchGraphAPI.getEnabled(graphID.value)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const enabledData: ArchGraphRecordResult = resp.data.data
      enabledGraphID.value = enabledData.graphID
      enabledImageData.value = enabledData.imageData
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

async function getOrderNodes() {
  try {
    const params: OrderQuery = {
      orderType: 3,
    }
    const resp: AxiosResponse = await OrderAPI.getOrderNodeList(params)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData = resp.data.data
      orderTypeInfoMapping.value = resData
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

function addControlItem(instance: LogicFlow | any) {
  if (!isSilentMode.value) {
    instance.extension.control.addItem({
      key: 'selection',
      iconClass: 'i-svg:selection',

      title: '',
      text: '框选',
      onClick: () => {
        handleSelection(true)
      },
    })
    instance.extension.control.addItem({
      key: 'not-selection',
      iconClass: 'i-svg:not_selection',

      title: '',
      text: '不框选',
      onClick: () => {
        handleSelection(false)
      },
    })
    instance.extension.control.addItem({
      key: 'save',
      iconClass: 'i-svg:save',
      title: '',
      text: '保存',
      onClick: async (instance: LogicFlow) => {
        handleSave(instance)
      },
    })

    instance.extension.control.addItem({
      key: 'snapshot',
      iconClass: 'i-svg:snapshot',

      title: '',
      text: '快照',
      onClick: () => {
        recordVisible.value = true
      },
    })
    if (approveStatus.value.includes(status.value)) {
      instance.extension.control.addItem({
        key: 'approve',
        iconClass: 'i-svg:menu_approve',
        title: '',
        text: '审批',
        onClick: async (instance: LogicFlow) => {
          // 在这里判断一下有没有评审记录
          if (!hasReview.value) {
            await checkHasReview()
            if (hasReview.value) {
              await getImgBase64(instance)
              approvalVisible.value = true
            } else {
              ElMessage.warning('架构图还没有评审请先评审，添加评审记录！')
            }
          } else {
            await getImgBase64(instance)
            approvalVisible.value = true
          }
        },
      })
    }

    instance.extension.control.addItem({
      key: 'review',
      iconClass: 'i-svg:review',
      title: '',
      text: '评审',
      onClick: async () => {
        reviewVisible.value = true
      },
    })
  }
  // 导出画布为图片
  instance.extension.control.addItem({
    key: 'output-image',
    iconClass: 'i-svg:export',
    title: '',
    text: '导出',
    onClick: async (instance: LogicFlow) => {
      await downLoadImage(instance)
    },
  })

  instance.extension.control.addItem({
    key: 'goback',
    iconClass: 'i-svg:back',
    title: '',
    text: '返回',
    onClick: async () => {
      handleGoBack()
    },
  })
  if (isSilentMode.value && hasOrder.value === 1) {
    instance.extension.control.addItem({
      key: 'goback',
      iconClass: 'i-svg:toggle',
      title: '',
      text: '工单切换',
      onClick: async () => {
        toggleOrderPanel()
      },
    })
  }
}

function toggleOrderPanel() {
  if (curHasOrder.value === 1) {
    curHasOrder.value = 0
  } else {
    curHasOrder.value = 1
  }
}

function addMenuItem(instance: LogicFlow | any) {
  const nodeMenus: any[] = []
  const edgeMenus: any[] = []
  const graphMenus: any[] = []

  if (!isSilentMode.value) {
    // graphMenus.push({
    //   text: '清空画布',
    //   callback() {
    //     instance.clearData()
    //     // sessionStorage.removeItem(graphKey.value)
    //   },
    // })
    instance.extension.menu.addMenuConfig({
      nodeMenu: nodeMenus,
      edgeMenu: edgeMenus,
      graphMenu: graphMenus,
    })
    // 对架构图组件添加架构图关联
    instance.setMenuByType({
      type: 'archGraph',
      menu: [
        {
          text: '关联架构图',
          className: 'lf-menu-item',
          callback(node: any) {
            linkNodeID.value = node.id
            const linkGraph = node.properties.linkGraph
            if (linkGraph) {
              linkGraphID.value = linkGraph.id
              linkGroupID.value = linkGraph.groupID
            }
            linkVisible.value = true
          },
        },
        {
          text: '跳转架构图',
          className: 'lf-menu-item',
          callback(node: any) {
            const linkGraph = node.properties.linkGraph
            if (linkGraph) {
              handleToGraph(linkGraph)
            } else {
              ElMessage.info('没有关联架构图')
            }
          },
        },
        {
          text: '编辑',
          className: 'lf-menu-item',
          callback: (node: any) => {
            instance.graphModel.setElementStateById(node.id, 2)
          },
        },
        {
          text: '删除',
          className: 'lf-menu-item',
          callback: (node: any) => {
            instance.graphModel.deleteNode(node.id)
          },
        },
      ],
    })
  }
}

function handleLinkNode(nodeID: string, graph: ArchGraphResult) {
  const linkNodeModel = lf.value?.getNodeModelById(nodeID)
  if (linkNodeModel) {
    linkNodeModel.setProperty('linkGraph', graph)
    lf.value?.graphModel.updateText(nodeID, graph.graphName)
  } else {
    ElMessage.error('关联架构图信息失败')
  }
}

function handleToGraph(linkGraph: ArchGraphResult | any) {
  router.push({
    path: `/arch/graph-draw/${linkGraph.id}`,
    query: {
      graphID: linkGraph.id,
      graphName: linkGraph.graphName,
      owner: linkGraph.owner,
      silentMode: silentMode.value,
      status: status.value,
      hasOrder: hasOrder.value,
    },
  })
}

async function checkHasReview() {
  try {
    const params: ArchGraphReviewQuery = {
      graphID: graphID.value,
    }
    const resp: AxiosResponse = await ArchGraphAPI.getReviewList(params)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const resData = resp.data.data
      const total = resData.total
      if (total > 0) {
        hasReview.value = true
      }
    } else {
      console.log(msg)
    }
  } catch (err) {
    console.log(err)
  }
}

async function downLoadImage(instance: LogicFlow) {
  ElMessageBox.confirm(`确定要导出架构图吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(async () => {
      try {
        loading.value = true
        loadingText.value = '图片导出中...'
        const params: ToImageOptions = {
          fileType: 'png', // 可选：'png'、'webp'、'jpeg'、'svg'
          backgroundColor: 'white',
          partial: false, // false: 导出所有元素，true: 只导出可见区域
        }
        await instance.getSnapshot(graphName.value, params)
        return
      } finally {
        loading.value = false
      }
    })
    .catch(() => {
      // 用户点击“取消” 或 关闭弹窗，什么都不做即可避免报错
    })
}

async function getImgBase64(instance?: LogicFlow) {
  if (!instance) {
    instance = lf?.value
  }
  const { data: base64 } = await instance?.getSnapshotBase64('#ffffff', 'png', {
    partial: false,
  })
  imageData.value = base64
}
// 保存
async function handleSave(instance: LogicFlow) {
  if (status.value === 2) {
    ElMessageBox.confirm(`对审批通过的架构图保存会生成一个状态是未审批的新图`, '提示', {
      confirmButtonText: '保存',
      cancelButtonText: `取消`,
      distinguishCancelAndClose: true,
      type: 'success',
      lockScroll: false,
    })
      .then(async () => {
        await doSave(instance, 'handle')
        // 保存后回退到前一页面
        handleGoBack()
      })
      .catch(() => {
        // 用户点击“取消” 或 关闭弹窗，什么都不做即可避免报错
      })
  } else {
    // 其它情况
    await doSave(instance, 'handle')
  }
}

async function doSave(instance?: LogicFlow, action: string = 'auto') {
  if (action === 'handle') {
    loading.value = true
    loadingText.value = '架构图保存中...'
  }

  if (!instance) {
    instance = lf?.value
  }
  await getImgBase64(instance)
  const saveGraphData: any = instance?.getGraphData()
  const nodeData = saveGraphData.nodes
  const edgeData = saveGraphData.edges
  const saveForm: ArchGraphForm = {
    id: graphID.value,
    nodeData: nodeData,
    edgeData: edgeData,
    imageData: imageData.value,
    action: 'save',
    status: status.value,
  }
  try {
    const resp: AxiosResponse = await ArchGraphAPI.save(saveForm)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      if (action === 'handle') {
        ElMessage.success('存储成功')
      }
      // 保存成功的时候清理一下缓存
      // sessionStorage.removeItem(graphKey.value)
    } else {
      console.log(msg)
      if (action === 'handle') {
        ElMessage.error('存储失败')
      }
    }
  } catch (err) {
    console.error(err)
    if (action === 'handle') {
      ElMessage.error('存储失败')
    }
  } finally {
    if (action === 'handle') {
      loading.value = false
    }
  }
}

// 监听到实践变化数据暂时存储在前端
// function eventToSave() {
//   const curGraphData = lf.value?.getGraphData()
//   const curJsonGraphData = JSON.stringify(curGraphData)
//   sessionStorage.setItem(graphKey.value, curJsonGraphData)
// }

function updateProperties() {
  Object.keys(properties).forEach((key) => delete properties[key])
  activeNodes.value.forEach((node) => Object.assign(properties, node.properties))
  activeEdges.value.forEach((edge) => Object.assign(properties, edge.properties))
}

function dragInNode(type: string, text: string) {
  curDragNodeText.value = text
  lf.value?.dnd.startDrag({ type })
  //拖拽也保存数据
  // eventToSave()
}

function setStyle(style: Record<string, any>) {
  activeNodes.value.forEach(({ id }) => lf.value?.setProperties(id, style))
  activeEdges.value.forEach(({ id }) => lf.value?.setProperties(id, style))
  updateProperties()
}
// 选择快照后重新render data
function handleSelectRecord() {
  // 先删除画布
  lf.value?.destroy()
  // 再从后端回去render
  getGrephData()
}

function handleSelection(selection: boolean) {
  if (selection) {
    isSelection.value = true
    lf.value?.openSelectionSelect()
    ElMessage.info('进入框选模式')
  } else {
    isSelection.value = false
    lf.value?.closeSelectionSelect()
    lf.value?.graphModel.editConfigModel.updateEditConfig({
      stopMoveGraph: false,
    })
    ElMessage.info('退出框选模式')
  }
}
// 提交成功后跳转到工单列表
async function handleApprove(orderType: number) {
  // 再执行一次保存操作
  if (orderType <= 2) {
    if (lf.value) {
      await doSave(lf.value)
    }
  }
  router.push({
    path: `/process-order/order`,
    query: {
      // graphName: graphName.value,
    },
  })
}

function setZIndex(type: 'top' | 'bottom') {
  activeNodes.value.forEach(({ id }) => lf.value?.setElementZIndex(id, type))
  activeEdges.value.forEach(({ id }) => lf.value?.setElementZIndex(id, type))
}
function closePanel() {
  showPanel.value = ''
}
function handleGoBack() {
  router.go(-1)
}
</script>

<style lang="scss" scoped>
@use 'sass:color';

$sidebar-width: 185px;
$panel-width: 300px;
$scrollbar-bg: #fff;
$scrollbar-thumb: #c9c9c9;
$scrollbar-thumb-hover: #b5b5b5;
$border-color: #dadce0;
$shadow-color: #838284;

.diagram {
  // overflow-y: scroll;
  overflow-y: hidden;
  width: 100%;
  height: 100%;

  *,
  *::before,
  *::after {
    box-sizing: border-box;
  }

  &-main {
    display: flex;
    width: 100%;
    height: 100%;

    .diagram-sidebar {
      width: $sidebar-width;
      height: 100%;
      border-right: 1px solid $border-color;
      padding: 10px;
    }
    .order-apply-panel {
      background-color: #fff;
      height: 100%;
      width: 600px;
      border-left: 1px solid $border-color;
      padding: 20px;
    }

    .diagram-container {
      flex: 1;
      .diagram-wrapper {
        width: 100%;
        height: 100%;
        .lf-diagram {
          box-shadow: 0 0 4px $shadow-color;
          width: 100%;
          height: 100%;
        }
      }
    }
  }
  &-panel-wrapper {
    width: $panel-width;
    background: #fff;
    height: calc(100% - 70px);
    position: absolute;
    right: 0;
    top: 70px;
    border-left: 1px solid $border-color;
    border-top: 1px solid $border-color;
    border-radius: 5px;
    overflow-y: auto;
  }

  // 修复背景与网格对齐
  :deep(.diagram .lf-background) {
    left: -9px;
  }
}
</style>
