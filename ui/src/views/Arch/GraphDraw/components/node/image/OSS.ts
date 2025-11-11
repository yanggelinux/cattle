import ImageNode from './ImageNode'
import ossImg from '@/assets/images/oss.png'

// OSS 的图片节点
class OSSNode extends ImageNode.view {
  getImageHref() {
    return ossImg
  }
}

export default {
  type: 'oss',
  view: OSSNode,
  model: ImageNode.model,
}
