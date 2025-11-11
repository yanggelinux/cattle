import CircleGroup from './CircleGroup'

class EllipseGroupModel extends CircleGroup.model {
  initNodeData(data: any) {
    super.initNodeData(data)
    this.width = 300
    this.height = 200
  }
}

export default {
  type: 'ellipseGroup',
  view: CircleGroup.view,
  model: EllipseGroupModel,
}
