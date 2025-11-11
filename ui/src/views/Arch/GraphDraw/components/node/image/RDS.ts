import ImageNode from './ImageNode'
import rdsImg from '@/assets/images/rds.png'

// RDS 的图片节点
class RDSNode extends ImageNode.view {
  getImageHref() {
    return rdsImg
  }
}

export default {
  type: 'rds',
  view: RDSNode,
  model: ImageNode.model,
}
