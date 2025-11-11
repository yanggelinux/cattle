import ImageNode from './ImageNode'
import domainImg from '@/assets/images/domain.png'
// import { h } from '@logicflow/core'

// Oracle 的图片节点
class DomainNode extends ImageNode.view {
  getImageHref() {
    return domainImg
  }
}

export default {
  type: 'domain',
  view: DomainNode,
  model: ImageNode.model,
}
