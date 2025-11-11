import ImageNode from './ImageNode'
import esImg from '@/assets/images/es.png'

// 左上角ICON为消息的节点
class EsNode extends ImageNode.view {
  getImageHref() {
    return esImg
  }
}

export default {
  type: 'es',
  view: EsNode,
  model: ImageNode.model,
}
