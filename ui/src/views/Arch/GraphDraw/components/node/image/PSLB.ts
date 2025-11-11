import ImageNode from './ImageNode'
import img from '@/assets/images/p_slb.png'

class PSLBNode extends ImageNode.view {
  getImageHref() {
    return img
  }
}

export default {
  type: 'pslb',
  view: PSLBNode,
  model: ImageNode.model,
}
