import ImageNode from './ImageNode'
import switchImg from '@/assets/images/switch.png'

// 交换机 的图片节点
class SwitchNode extends ImageNode.view {
  getImageHref() {
    return switchImg
  }
}

export default {
  type: 'switch',
  view: SwitchNode,
  model: ImageNode.model,
}
