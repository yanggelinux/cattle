import ImageNode from './ImageNode'
import firewallImg from '@/assets/images/firewall.png'
// import { h } from '@logicflow/core'

// 交换机 的图片节点
class FirewallNode extends ImageNode.view {
  getImageHref() {
    return firewallImg
  }
}

export default {
  type: 'firewall',
  view: FirewallNode,
  model: ImageNode.model,
}
