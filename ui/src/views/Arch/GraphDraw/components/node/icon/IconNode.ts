import { h } from '@logicflow/core'
import RectNode from '../basic/RectNode'

class IconModel extends RectNode.model {
  initNodeData(data: any) {
    super.initNodeData(data)
    this.width = 70
    this.height = 70
  }
  getTextStyle() {
    const style = super.getTextStyle()
    // style.fontSize = 16
    style.x = this.x
    style.y = this.y + this.height / 2 + 20
    style.editable = true
    return style
  }
}

// 左上角带ICON的节点
class IconNode extends RectNode.view {
  getImageHref() {
    return
  }
  getResizeShape() {
    const { x, y, width, height } = this.props.model
    const href = this.getImageHref()
    const iconAttrs = {
      x: x - (1 / 2) * width,
      y: y - (1 / 2) * height, // icon在左上角
      width: width,
      height,
      href,
      // 根据宽高缩放
      preserveAspectRatio: 'none meet',
    }
    return h('g', {}, [h('image', { ...iconAttrs })])
  }
}

export default {
  type: 'icon-node',
  view: IconNode,
  model: IconModel,
}
