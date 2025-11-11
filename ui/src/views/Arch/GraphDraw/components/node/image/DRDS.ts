import ImageNode from './ImageNode'
import drdsImg from '@/assets/images/drds.png'

// DRDS 的图片节点
class DRDSNode extends ImageNode.view {
  getImageHref() {
    return drdsImg
  }
}

export default {
  type: 'drds',
  view: DRDSNode,
  model: ImageNode.model,
}
