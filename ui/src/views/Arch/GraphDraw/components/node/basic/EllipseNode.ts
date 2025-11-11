import { EllipseNodeModel, EllipseNode } from '@logicflow/core'

// 椭圆
class EllipseNewModel extends EllipseNodeModel {
  initNodeData(data: any) {
    super.initNodeData(data)
    this.rx = 60
    this.ry = 30
  }
  getNodeStyle() {
    const style = super.getNodeStyle()
    return { ...style }
  }
}
export default {
  type: 'pro-ellipse',
  view: EllipseNode,
  model: EllipseNewModel,
}
