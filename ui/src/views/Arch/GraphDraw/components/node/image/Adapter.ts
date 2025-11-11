import ImageNode from './ImageNode'
import adapterImg from '@/assets/images/adapter.png'

// 左上角ICON为消息的节点
class AdapterNode extends ImageNode.view {
  getImageHref() {
    return adapterImg
  }
}

export default {
  type: 'adapter',
  view: AdapterNode,
  model: ImageNode.model,
}
