import { getShapeStyleFuction, getTextStyleFunction } from '../getShapeStyleUtil'
import { RectNodeModel, RectNode } from '@logicflow/core'

// 矩形
class RectNewModel extends RectNodeModel {
  initNodeData(data: any) {
    super.initNodeData(data)
    this.width = 180
    this.height = 80
  }

  setToBottom() {
    this.zIndex = 0
  }

  getNodeStyle() {
    const style = super.getNodeStyle()
    const properties = this.getProperties()
    // style.stroke = '#67C23A'
    return getShapeStyleFuction(style, properties)
  }

  getTextStyle() {
    const style = super.getTextStyle()
    const properties = this.getProperties()
    return getTextStyleFunction(style, properties)
  }
}

export default {
  type: 'procApproval',
  view: RectNode,
  model: RectNewModel,
}
