import ImageNode from './ImageNode'
import telephoneImg from '@/assets/images/telephone.png'

// 左上角ICON为消息的节点
class TelephoneNode extends ImageNode.view {
  getImageHref() {
    return telephoneImg
  }
}

export default {
  type: 'telephone',
  view: TelephoneNode,
  model: ImageNode.model,
}
