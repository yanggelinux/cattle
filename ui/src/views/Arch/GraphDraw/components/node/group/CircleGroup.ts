import { h } from '@logicflow/core'
import GroupNode from '../basic/GroupNode'

class CircleGroupModel extends GroupNode.model {
  initNodeData(data: any) {
    super.initNodeData(data)
    this.width = 300
    this.height = 300
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

class CircleGroupNode extends GroupNode.view {
  getImageHref() {
    return
  }
  getShape() {
    const { x, y, width, height } = this.props.model
    const style = this.props.model.getNodeStyle()

    const ellipseBAttrs = {
      ...style,
      cx: x,
      cy: y,
      rx: (1 / 2) * width,
      ry: (1 / 2) * height,
      width,
      height,
    }

    return h('g', {}, [
      h('ellipse', {
        ...ellipseBAttrs,
      }),
    ])
  }
}

export default {
  type: 'circleGroup',
  view: CircleGroupNode,
  model: CircleGroupModel,
}
