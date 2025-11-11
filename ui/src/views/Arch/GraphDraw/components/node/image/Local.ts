import ImageNode from './ImageNode'
import drdsImg from '@/assets/images/local_gov.png'

// DRDS 的图片节点
class LocalNode extends ImageNode.view {
  getImageHref() {
    return drdsImg
  }
}

export default {
  type: 'local',
  view: LocalNode,
  model: ImageNode.model,
}
