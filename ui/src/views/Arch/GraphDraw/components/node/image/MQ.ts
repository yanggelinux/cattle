import ImageNode from './ImageNode'
import mqImg from '@/assets/images/mq.png'

// MQ 的图片节点
class MQNode extends ImageNode.view {
  getImageHref() {
    return mqImg
  }
}

export default {
  type: 'mq',
  view: MQNode,
  model: ImageNode.model,
}
