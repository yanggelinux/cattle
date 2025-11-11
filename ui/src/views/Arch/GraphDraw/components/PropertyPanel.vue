<template>
  <div class="diagram-property-panel">
    <div class="close-icon-wrapper">
      <el-icon class="close-icon" @click="handleClose"><Close /></el-icon>
    </div>
    <!-- 快捷样式（仅节点） -->
    <div v-if="!onlyEdge" class="setting-block">
      <div>快捷样式</div>
      <div class="short-styles">
        <div
          v-for="(item, idx) in shortStyles"
          :key="idx"
          class="style-swatch"
          :style="{
            backgroundColor: item.backgroundColor,
            borderColor: item.borderColor,
            borderWidth: item.borderWidth + 'px',
          }"
          @click="applyStyle(item)"
        />
      </div>
    </div>
    <!-- 样式设置 -->
    <div class="setting-block">
      <!-- 填充色 & 渐变色（仅节点） -->
      <div v-if="!onlyEdge" class="setting-item">
        <span>背景色</span>
        <el-color-picker
          show-alpha
          v-model="style.backgroundColor"
          @change="(c: any) => changeColorProperty(c, 'backgroundColor')"
        />
      </div>

      <!-- 线条链接样式 -->
      <div v-if="onlyEdge" class="setting-item">
        <span>连接样式</span>
        <el-select v-model="linetype">
          <el-option
            v-for="item in lineOptions"
            :key="item.value"
            :value="item.value"
            :label="item.label"
            @click="changeLineType(item.value)"
          />
        </el-select>
      </div>
      <!-- 线条样式 -->
      <div class="setting-item">
        <span>线条样式</span>
        <el-select
          v-model="style.borderStyle"
          @change="(val: any) => emitSetStyle({ borderStyle: val })"
        >
          <el-option value="hidden" label="不显示" />
          <el-option v-for="(b, i) in borderStyles" :key="i" :value="b.value" label="">
            <div class="border-style" :style="{ borderBottom: b.value, marginTop: '18px' }"></div>
          </el-option>
        </el-select>
      </div>

      <!-- 线条颜色 -->
      <div class="setting-item">
        <span>线条颜色</span>
        <el-color-picker
          v-model="style.borderColor"
          @change="(c: any) => changeColorProperty(c, 'borderColor')"
        />
      </div>

      <!-- 线条宽度 -->
      <div class="setting-item">
        <span>线条宽度</span>
        <el-select
          v-model="style.borderWidth"
          @change="(val: any) => emitSetStyle({ borderWidth: val })"
        >
          <el-option
            v-for="w in borderWidthOptions"
            :key="w"
            :label="`${w + 1}px`"
            :value="w + 1"
          />
        </el-select>
      </div>

      <!-- 文本颜色 -->
      <div class="setting-item">
        <span>文本颜色</span>
        <el-color-picker
          v-model="style.fontColor"
          @change="(c: any) => changeColorProperty(c, 'fontColor')"
        />
      </div>

      <!-- 文本大小 -->
      <div class="setting-item">
        <span>文本大小</span>
        <el-input-number
          v-model="style.fontSize"
          controls-position="right"
          :min="12"
          :max="30"
          @change="(val: any) => emitSetStyle({ fontSize: val })"
        />
      </div>

      <!-- 文本字体 -->
      <div class="setting-item">
        <span>字体</span>
        <el-select
          v-model="style.fontFamily"
          @change="(val: any) => emitSetStyle({ fontFamily: val })"
        >
          <el-option v-for="(f, i) in fontFamilies" :key="i" :label="f.label" :value="f.value" />
        </el-select>
      </div>

      <!-- 行高 -->
      <div class="setting-item">
        <span>行高</span>
        <el-select
          v-model="style.lineHeight"
          @change="(val: any) => emitSetStyle({ lineHeight: val })"
        >
          <el-option v-for="(lh, i) in lineHeightOptions" :key="i" :label="lh" :value="lh" />
        </el-select>
      </div>

      <!-- 文本对齐 -->
      <!-- <div class="setting-item">
        <span>对齐</span>
        <el-radio-group
          v-model="style.textAlign"
          size="small"
          @change="(val: any) => emitSetStyle({ textAlign: val })"
        >
          <el-radio-button value="left">左对齐</el-radio-button>
          <el-radio-button value="center">居中</el-radio-button>
          <el-radio-button value="right">右对齐</el-radio-button>
        </el-radio-group>
      </div> -->

      <!-- 文本样式 -->
      <div class="setting-item">
        <span>文本样式</span>
        <el-button size="small" @click="toggleFontWeight">B</el-button>
        <el-button size="small" @click="toggleTextDecoration">U</el-button>
        <el-button size="small" @click="toggleFontStyle">I</el-button>
      </div>

      <!-- Z-Index -->
      <div class="setting-item">
        <span>点边层级</span>
        <el-button size="small" @click="() => emitSetZIndex('top')">置顶</el-button>
        <el-button size="small" @click="() => emitSetZIndex('bottom')">置底</el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, watch, type PropType } from 'vue'
import { shortStyles, borderStyles, fontFamilies } from '../constant/index.ts'
import LogicFlow from '@logicflow/core'

defineOptions({
  name: 'PropertyPanel',
})
// Props & Emits
const props = defineProps({
  lf: {
    required: true,
    type: Object as PropType<LogicFlow | any>,
  },
  elementsStyle: {
    type: Object,
  },
  onlyEdge: {
    type: Boolean,
  },
  activeEdges: {
    type: Array<{ id: string }>,
  },
})

const lineOptions = [
  { value: 'pro-line', label: '直线' },
  { value: 'pro-polyline', label: '折线' },
  { value: 'pro-bezier', label: '曲线' },
]

const emit = defineEmits<{
  (e: 'close'): void
  (e: 'setStyle', style: Record<string, any>): void
  (e: 'setZIndex', z: 'top' | 'bottom'): void
  (e: 'changeLineType', value: string): void
}>()

// State
// const shortStyles:any = shortStyles
// const borderStyles:any = borderStyles
// const fontFamilies:any = fontFamilies
const borderWidthOptions = Array.from({ length: 11 }, (_, i) => i)
const lineHeightOptions = Array.from({ length: 5 }, (_, i) => 1 + i * 0.5)
const linetype = ref<string>('pro-line')

const { elementsStyle, lf } = props

const style = reactive({
  backgroundColor: '',
  gradientColor: '',
  borderColor: '',
  borderWidth: 1,
  borderStyle: '',
  fontSize: 12,
  fontColor: '',
  fontFamily: '',
  lineHeight: 1,
  textAlign: 'left',
  fontWeight: '',
  textDecoration: '',
  fontStyle: '',
})

// Sync props to local style
watch(
  () => elementsStyle,
  (val) => {
    Object.assign(style, val)
  },
  { immediate: true }
)

// Helpers
function handleClose() {
  emit('close')
}

function emitSetStyle(payload: Record<string, any>) {
  emit('setStyle', payload)
}

function emitSetZIndex(z: 'top' | 'bottom') {
  emit('setZIndex', z)
}

function applyStyle(item: any) {
  emitSetStyle(item)
}

function changeColorProperty(color: any, prop: string) {
  ;(style as any)[prop] = color
  emitSetStyle({ [prop]: color })
}

function toggleFontWeight() {
  const val = style.fontWeight === 'bold' ? 'normal' : 'bold'
  style.fontWeight = val
  emitSetStyle({ fontWeight: val })
}

function toggleTextDecoration() {
  const val = style.textDecoration === 'underline' ? '' : 'underline'
  style.textDecoration = val
  emitSetStyle({ textDecoration: val })
}

function toggleFontStyle() {
  const val = style.fontStyle === 'italic' ? 'normal' : 'italic'
  style.fontStyle = val
  emitSetStyle({ fontStyle: val })
}
function changeLineType(value: string) {
  lf.setDefaultEdgeType(value)
  if (props.activeEdges?.length) {
    props.activeEdges.forEach((edge) => {
      lf.graphModel.changeEdgeType(edge.id, value)
    })
  }
  emit('changeLineType', value)
}
</script>

<style lang="scss" scoped>
.diagram-property-panel {
  padding: 10px 20px 20px 20px;
  .close-icon-wrapper {
    text-align: right;
    .close-icon {
      cursor: pointer;
    }
  }

  .setting-block {
    overflow: hidden;
  }

  .short-styles {
    width: 240px;

    > div {
      width: 20px;
      height: 20px;
      margin: 0 10px 5px 0;
      box-sizing: border-box;
      float: left;
      border: 1px solid #fff;
      cursor: pointer;
    }
  }

  .border-style {
    width: 150px;
    height: 0;
    margin-top: 18px !important;
    border-bottom-width: 1px;
    border-bottom-color: black;
  }

  .setting-item {
    line-height: 12px;
    font-size: 12px;
    display: flex;
    vertical-align: middle;
    align-items: center;
    margin-top: 10px;

    > span {
      width: 50px;
      margin-right: 10px;
      text-align: right;
      flex-shrink: 0;
      flex-grow: 0;
    }
  }

  .border-color {
    width: 40px;
    height: 30px;
    display: inline-block;
    border: 1px solid #eaeaeb;
  }
}
</style>
