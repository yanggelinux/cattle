import ImageNode from './ImageNode'
import mpassImg from '@/assets/images/mpass.png'

// 左上角ICON为消息的节点
class MpassNode extends ImageNode.view {
  getImageHref() {
    return mpassImg
  }
}

export default {
  type: 'mpass',
  view: MpassNode,
  model: ImageNode.model,
}
