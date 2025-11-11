import ImageNode from './ImageNode'
import sslvpnImg from '@/assets/images/sslvpn.png'

// SSLVPN 的图片节点
class SSLVPNNode extends ImageNode.view {
  getImageHref() {
    return sslvpnImg
  }
}

export default {
  type: 'sslvpn',
  view: SSLVPNNode,
  model: ImageNode.model,
}
