import GroupNode from '../basic/GroupNode'

class RectGroupModel extends GroupNode.model {
  initNodeData(data: any) {
    super.initNodeData(data)
    this.width = 500
    this.height = 200
    this.isShowAnchor = true
    this.collapsible = false
  }
  getTextStyle() {
    const style = super.getTextStyle()
    // style.x = this.x
    // style.y = this.y + this.height / 2 + 20
    style.x = this.x - this.width / 2 + 60
    style.y = this.y - this.height / 2 + 20
    // style.editable = true
    // style.draggable = true
    return style
  }
  getNodeStyle() {
    const style = super.getNodeStyle()
    return style
  }
}

class RectGroupNode extends GroupNode.view {
  getImageHref() {
    return
  }
}

export default {
  type: 'rectGroup',
  view: RectGroupNode,
  model: RectGroupModel,
}
