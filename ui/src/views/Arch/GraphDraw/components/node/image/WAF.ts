import ImageNode from './ImageNode'
import wafImg from '@/assets/images/waf.png'
// import { h } from '@logicflow/core'

// RDS 的图片节点
class WAFNode extends ImageNode.view {
  getImageHref() {
    return wafImg
  }
}

export default {
  type: 'waf',
  view: WAFNode,
  model: ImageNode.model,
}
