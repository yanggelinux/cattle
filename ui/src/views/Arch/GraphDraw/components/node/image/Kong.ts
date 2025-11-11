import ImageNode from './ImageNode'
import kongImg from '@/assets/images/kong.png'

class KongModel extends ImageNode.model {
  initNodeData(data: any) {
    super.initNodeData(data)
    this.width = 50
    this.height = 40
  }
}
// KONG 的图片节点
class KongNode extends ImageNode.view {
  getImageHref() {
    return kongImg
  }
}

export default {
  type: 'kong',
  view: KongNode,
  model: KongModel,
}
