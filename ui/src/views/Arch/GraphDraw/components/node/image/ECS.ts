import ImageNode from './ImageNode'
import ecsImg from '@/assets/images/ecs.png'

// 云形状的图片节点
class ECSNode extends ImageNode.view {
  getImageHref() {
    return ecsImg
  }
}

export default {
  type: 'ecs',
  view: ECSNode,
  model: ImageNode.model,
}
