import ImageNode from './ImageNode'
import emailImg from '@/assets/images/email.png'

// 左上角ICON为消息的节点
class EmailNode extends ImageNode.view {
  getImageHref() {
    return emailImg
  }
}

export default {
  type: 'email',
  view: EmailNode,
  model: ImageNode.model,
}
