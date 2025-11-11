import ImageNode from './ImageNode'
import internalFaceImg from '@/assets/images/internal_face.png'

// 内部人脸服务 的图片节点
class InternalFaceNode extends ImageNode.view {
  getImageHref() {
    return internalFaceImg
  }
}

export default {
  type: 'internalFace',
  view: InternalFaceNode,
  model: ImageNode.model,
}
