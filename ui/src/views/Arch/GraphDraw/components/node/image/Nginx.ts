import ImageNode from './ImageNode'
import nginxImg from '@/assets/images/nginx.png'

// Nginx 的图片节点
class NginxNode extends ImageNode.view {
  getImageHref() {
    return nginxImg
  }
}

export default {
  type: 'nginx',
  view: NginxNode,
  model: ImageNode.model,
}
