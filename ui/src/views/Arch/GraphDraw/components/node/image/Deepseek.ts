import ImageNode from './ImageNode'
import deepseekImg from '@/assets/images/deepseek.png'
// import { h } from '@logicflow/core'

// 左上角ICON为消息的节点
class DeepseekNode extends ImageNode.view {
  getImageHref() {
    return deepseekImg
  }
}

export default {
  type: 'deepseek',
  view: DeepseekNode,
  model: ImageNode.model,
}
