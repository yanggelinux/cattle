import { h } from '@logicflow/core'
import GroupNode from '../basic/GroupNode'

class TriangleGroupModel extends GroupNode.model {
  initNodeData(data: any) {
    super.initNodeData(data)
    this.width = 500
    this.height = 200
  }
  getTextStyle() {
    const style = super.getTextStyle()
    style.x = this.x
    style.y = this.y + this.height / 2 + 20
    style.editable = false
    return style
  }
  getNodeStyle() {
    const style = super.getNodeStyle()
    return style
  }
}

class TriangleGroupNode extends GroupNode.view {
  getImageHref() {
    return
  }
  getShape() {
    const { x, y, width, height } = this.props.model
    const style = this.props.model.getNodeStyle()
    const attrs = {
      ...style,
      x,
      y,
      width,
      height,
      points: [
        [x - width / 2, y + height / 2],
        [x - width / 2, y - height / 2],
        [x + width / 2, y],
      ],
    }
    return h('g', {}, [h('polygon', { ...attrs })])
  }
}

export default {
  type: 'triangleGroup',
  view: TriangleGroupNode,
  model: TriangleGroupModel,
}
