import ImageNode from './ImageNode'
import browserImg from '@/assets/images/browser.png'

class BrowserNode extends ImageNode.view {
  getImageHref() {
    return browserImg
  }
}

export default {
  type: 'browser',
  view: BrowserNode,
  model: ImageNode.model,
}
