import ImageNode from './ImageNode'
import img from '@/assets/images/tencent_cloud.png'

// 左上角ICON为消息的节点
class TencentCloudNode extends ImageNode.view {
  getImageHref() {
    return img
  }
}

export default {
  type: 'tencentCloud',
  view: TencentCloudNode,
  model: ImageNode.model,
}
