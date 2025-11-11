import ImageNode from './ImageNode'
import gbaseImg from '@/assets/images/gbase.png'

// GBase 的图片节点
class GBaseNode extends ImageNode.view {
  getImageHref() {
    return gbaseImg
  }
}

export default {
  type: 'gbase',
  view: GBaseNode,
  model: ImageNode.model,
}
