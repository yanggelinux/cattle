import ImageNode from './ImageNode'
import cdnImg from '@/assets/images/cdn.png'

// 左上角ICON为消息的节点
class CDNNode extends ImageNode.view {
  getImageHref() {
    return cdnImg
  }
}

export default {
  type: 'cdn',
  view: CDNNode,
  model: ImageNode.model,
}
