import ImageNode from './ImageNode'
import doubaoImg from '@/assets/images/doubao.png'
// import { h } from '@logicflow/core'

// 左上角ICON为消息的节点
class DoubaoNode extends ImageNode.view {
  getImageHref() {
    return doubaoImg
  }
}

export default {
  type: 'doubao',
  view: DoubaoNode,
  model: ImageNode.model,
}
