import ImageNode from './ImageNode'
import gbaseImg from '@/assets/images/hbase.png'

// GBase 的图片节点
class HBaseNode extends ImageNode.view {
  getImageHref() {
    return gbaseImg
  }
}

export default {
  type: 'hbase',
  view: HBaseNode,
  model: ImageNode.model,
}
