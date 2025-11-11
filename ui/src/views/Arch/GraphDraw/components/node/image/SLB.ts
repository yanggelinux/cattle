import ImageNode from './ImageNode'
import slbImg from '@/assets/images/slb.png'

// SLB 的图片节点
class SLBNode extends ImageNode.view {
  getImageHref() {
    return slbImg
  }
}

export default {
  type: 'slb',
  view: SLBNode,
  model: ImageNode.model,
}
