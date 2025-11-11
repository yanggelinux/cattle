import ImageNode from './ImageNode'
import oceanbaseImg from '@/assets/images/oceanbase.png'

// OceanBase 的图片节点
class OceanBaseNode extends ImageNode.view {
  getImageHref() {
    return oceanbaseImg
  }
}

export default {
  type: 'oceanbase',
  view: OceanBaseNode,
  model: ImageNode.model,
}
