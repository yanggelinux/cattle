import ImageNode from './ImageNode'
import img from '@/assets/images/sls.png'

// 左上角ICON为消息的节点
class SlsNode extends ImageNode.view {
  getImageHref() {
    return img
  }
}

export default {
  type: 'sls',
  view: SlsNode,
  model: ImageNode.model,
}
