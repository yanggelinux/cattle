<template>
  <div class="diagram">
    <div class="diagram-main">
      <!-- 侧边栏拖拽节点 -->
      <ProcSidebar class="proc-sidebar" @dragInNode="dragInNode" v-if="isSilentMode === false" />
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
        <DataPanel
          :lf="lf"
          :nodeType="curNodeType"
          v-model:nodeID="curNodeID"
          v-if="showPanel === 'data'"
          @close="closePanel"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, nextTick, onBeforeUnmount, computed } from 'vue'
import LogicFlow from '@logicflow/core'
import { SelectionSelect, Menu, Control, Snapshot, DynamicGroup } from '@logicflow/extension'
import '@logicflow/core/es/index.css'
import '@logicflow/extension/lib/style/index.css'
import { GridBackgroundImage } from '@/views/Arch/GraphDraw/constant/index.ts'
import ProcSidebar from './components/ProcSidebar.vue'
import DataPanel from './components/DataPanel.vue'
import { type AxiosResponse } from 'axios'
import { ElMessage } from 'element-plus'
// import router from '@/router'
import ProcessAPI, { type ProcessResult, type ProcessForm } from '@/api/process/process'
import { useAuthStore } from '@/store/modules/auth.ts'
import { useRoute, useRouter } from 'vue-router'
import { registerCustomElement } from '@/views/Arch/GraphDraw/components/node'
const router = useRouter()

type GraphData = {
  nodes: any[]
  edges: any[]
}

const route = useRoute()
const processID = route.query.processID ? Number(route.query.processID) : 0
const processName = route.query.processName ? String(route.query.processName) : ''
const silentMode = route.query.silentMode ? Number(route.query.silentMode) : 0

const { isSuper } = useAuthStore()

const silent = computed(() => {
  // 审批中的不能编辑
  if (silentMode === 1) {
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
const showPanel = ref<string>('')
const curNodeID = ref<string>('')
const curNodeType = ref<string>('')
const loading = ref<boolean>(false)
const loadingText = ref<string>('loading...')
// 画布是否可编辑
const isSilentMode = ref<boolean>(silent.value)

// 审批
const messageInstance = ref<any>(null)

const properties = reactive<Record<string, any>>({})

//从路由中获取图ID，所有的图中点和边都必须有这个id,新创建的图id由保存后返回
// const graphKey = ref<string>(`graphData-${processID.value}`)

defineOptions({
  name: 'ProcessDraw',
})

// 执行自动存储操作，1000 * 60 * 10  10分组自动存储一次,状态是未审批，创建人必须是自己
// if (status.value === 0 && owner.value === userName) {
//   useAutoSave(doSave, () => lf?.value, 6000, false)
// }

onMounted(async () => {
  if (isSilentMode.value) {
    messageInstance.value = ElMessage.info({
      grouping: true,
      showClose: true,
      duration: 5000,
      message: '流程只读模式！',
    })
  }
  if (processID > 0) {
    await genGraphData()
  } else {
    const graphData: GraphData = {
      edges: [],
      nodes: [],
    }
    initLogicFlow(graphData)
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
  instance.setDefaultEdgeType('pro-polyline')

  instance.on('blank:click', async () => {
    if (showPanel.value === 'data') {
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

  instance.on('node:click', async ({ data }) => {
    curNodeID.value = data.id
    curNodeType.value = data.type
    instance.graphModel.textEditElement?.setElementState(1)
    showPanel.value = 'data'
    // eventToSave()
  })

  instance.on('node:dnd-add', async ({ data }) => {
    // 如果想拖拽自定义的图标组的情况下，根据type判断类型，
    // 如果图标组的情况下把图标组的数据 执行 addNode，addEdge 方法
    // 注意节点所在的坐标情况
    const nodeType = data.type
    let curDragNodeText = ''
    if (nodeType === 'procStart') {
      curDragNodeText = '开始节点'
    } else if (nodeType === 'procApproval') {
      curDragNodeText = '审批节点'
    } else if (nodeType === 'procEnd') {
      curDragNodeText = '结束节点'
    }
    const nodeID = data.id
    instance.updateText(nodeID, curDragNodeText)
    instance.setProperties(nodeID, { fontSize: 16 })
  })

  addControlItem(instance)
  addMenuItem(instance)

  instance.render(data)
  lf.value = instance
}

//生成图所需的数据
async function genGraphData() {
  await getGrephData()
}

async function getGrephData() {
  try {
    const resp: AxiosResponse = await ProcessAPI.getDetail(processID)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      const detailData: ProcessResult = resp.data.data
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

function addControlItem(instance: LogicFlow | any) {
  if (!isSilentMode.value) {
    instance.extension.control.addItem({
      key: 'save',
      iconClass: 'i-svg:save',
      title: '',
      text: '保存',
      onClick: async (instance: LogicFlow) => {
        handleSave(instance)
      },
    })
  }
  instance.extension.control.addItem({
    key: 'goback',
    iconClass: 'i-svg:back',
    title: '',
    text: '返回',
    onClick: async () => {
      handleGoBack()
    },
  })
}

function addMenuItem(instance: LogicFlow | any) {
  const nodeMenus: any[] = []
  const edgeMenus: any[] = []
  const graphMenus: any[] = []

  if (!isSilentMode.value) {
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

// 保存
async function handleSave(instance: LogicFlow) {
  await doSave(instance)
}

async function doSave(instance?: LogicFlow) {
  loading.value = true
  loadingText.value = '流程保存中...'
  if (!instance) {
    instance = lf?.value
  }
  const saveGraphData: any = instance?.getGraphData()
  const nodeData = saveGraphData.nodes
  const edgeData = saveGraphData.edges
  const saveForm: ProcessForm = {
    id: processID,
    name: processName,
    nodeData: nodeData,
    edgeData: edgeData,
  }
  try {
    const resp: AxiosResponse = await ProcessAPI.update(saveForm)
    const status = resp.data.status
    const msg = resp.data.msg
    if (status === 200) {
      ElMessage.success('存储成功')
      // 保存成功的时候清理一下缓存
      // sessionStorage.removeItem(graphKey.value)
    } else {
      console.log(msg)
      ElMessage.error(`存储失败,${msg}`)
    }
  } catch (err) {
    console.error(err)

    ElMessage.error('存储失败')
  } finally {
    loading.value = false
  }
}
function updateProperties() {
  Object.keys(properties).forEach((key) => delete properties[key])
  activeNodes.value.forEach((node) => Object.assign(properties, node.properties))
  activeEdges.value.forEach((edge) => Object.assign(properties, edge.properties))
}

function dragInNode(type: string) {
  lf.value?.dnd.startDrag({ type })
  //拖拽也保存数据
  // eventToSave()
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

$sidebar-width: 100px;
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

    .proc-sidebar {
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
