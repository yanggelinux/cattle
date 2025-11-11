import ImageNode from './ImageNode'
import csbImg from '@/assets/images/csb.png'

// CSB 的图片节点
class CSBNode extends ImageNode.view {
  getImageHref() {
    return csbImg
  }
}

export default {
  type: 'csb',
  view: CSBNode,
  model: ImageNode.model,
}
