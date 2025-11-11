import { h } from '@logicflow/core'
import RectNode from '../basic/RectNode'

// 图片-基础节点
class ImageModel extends RectNode.model {
  initNodeData(data: any) {
    super.initNodeData(data)
    this.width = 40
    this.height = 40
  }
  getTextStyle() {
    const style = super.getTextStyle()
    // style.fontSize = 20
    style.x = this.x
    style.y = this.y + this.height / 2 + 10
    style.editable = false
    return style
  }
}

class ImageNode extends RectNode.view {
  getImageHref() {
    return
  }
  getShape() {
    const { x, y, width, height } = this.props.model
    const href = this.getImageHref()
    const attrs = {
      x: x - (1 / 2) * width,
      y: y - (1 / 2) * height,
      width: width,
      height,
      href,
      // 根据宽高缩放
      preserveAspectRatio: 'none meet',
    }
    return h('g', {}, [h('image', { ...attrs })])
  }
}

export default {
  type: 'image-node',
  view: ImageNode,
  model: ImageModel,
}
