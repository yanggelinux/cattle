import ImageNode from './ImageNode'

// 云形状的图片节点
class CloudNode extends ImageNode.view {
  getImageHref() {
    return '@/assets/images/cloud.png'
  }
}

export default {
  type: 'cloud',
  view: CloudNode,
  model: ImageNode.model,
}
