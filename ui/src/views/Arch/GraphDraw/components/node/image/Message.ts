import ImageNode from './ImageNode'
import messageIcon from '@/assets/images/message.png'

// 左上角ICON为消息的节点
class MessageNode extends ImageNode.view {
  getImageHref() {
    return messageIcon
  }
}

export default {
  type: 'message',
  view: MessageNode,
  model: ImageNode.model,
}
