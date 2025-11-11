import ImageNode from './ImageNode'
import tianyanchaImg from '@/assets/images/tianyancha.png'

// 左上角ICON为消息的节点
class TianyanchaNode extends ImageNode.view {
  getImageHref() {
    return tianyanchaImg
  }
}

export default {
  type: 'tianyancha',
  view: TianyanchaNode,
  model: ImageNode.model,
}
