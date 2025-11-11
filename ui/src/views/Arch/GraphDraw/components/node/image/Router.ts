import ImageNode from './ImageNode'
import routerImg from '@/assets/images/router.png'

// 路由器 的图片节点
class RouterNode extends ImageNode.view {
  getImageHref() {
    return routerImg
  }
}

export default {
  type: 'router',
  view: RouterNode,
  model: ImageNode.model,
}
