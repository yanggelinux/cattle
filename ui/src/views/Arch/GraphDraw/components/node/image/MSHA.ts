import ImageNode from './ImageNode'
import mshaImg from '@/assets/images/msha.png'

// 云形状的图片节点
class MSHANode extends ImageNode.view {
  getImageHref() {
    return mshaImg
  }
}

export default {
  type: 'msha',
  view: MSHANode,
  model: ImageNode.model,
}
