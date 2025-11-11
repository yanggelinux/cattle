import ImageNode from './ImageNode'
import underCloudImg from '@/assets/images/under_cloud.png'

// 云形状的图片节点
class UnderCloudNode extends ImageNode.view {
  getImageHref() {
    return underCloudImg
  }
}

export default {
  type: 'underCloud',
  view: UnderCloudNode,
  model: ImageNode.model,
}
