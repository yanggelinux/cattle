import ImageNode from './ImageNode'
import arbitrationImg from '@/assets/images/arbitration.png'

// 左上角ICON为消息的节点
class ArbitrationNode extends ImageNode.view {
  getImageHref() {
    return arbitrationImg
  }
}

export default {
  type: 'arbitration',
  view: ArbitrationNode,
  model: ImageNode.model,
}
