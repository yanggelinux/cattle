import ImageNode from './ImageNode'
import img from '@/assets/images/ministry.png'

// 左上角ICON为消息的节点
class MinistryNode extends ImageNode.view {
  getImageHref() {
    return img
  }
}

export default {
  type: 'ministry',
  view: MinistryNode,
  model: ImageNode.model,
}
