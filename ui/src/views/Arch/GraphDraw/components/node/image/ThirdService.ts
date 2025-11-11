import ImageNode from './ImageNode'
import thirdServiceImg from '@/assets/images/third_service.png'

// 左上角ICON为消息的节点
class ThirdServiceNode extends ImageNode.view {
  getImageHref() {
    return thirdServiceImg
  }
}

export default {
  type: 'thirdService',
  view: ThirdServiceNode,
  model: ImageNode.model,
}
