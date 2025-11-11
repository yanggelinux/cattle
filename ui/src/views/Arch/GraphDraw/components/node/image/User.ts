import ImageNode from './ImageNode'

// 图片-用户节点
class UserNode extends ImageNode.view {
  getImageHref() {
    return '@/assets/images/user.png'
  }
}

export default {
  type: 'user',
  view: UserNode,
  model: ImageNode.model,
}
