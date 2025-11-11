import GroupNode from '../basic/GroupNode'

class NetworkGroupModel extends GroupNode.model {
  initNodeData(data: any) {
    super.initNodeData(data)
    this.width = 500
    this.height = 200
    this.isShowAnchor = true
    this.collapsible = false
    // this.radius = 20
  }

  setAttributes() {
    super.setAttributes()
    // this.radius = 10
  }

  getTextStyle() {
    const style = super.getTextStyle()
    style.x = this.x - this.width / 2 + 60
    style.y = this.y - this.height / 2 + 20
    // style.value = '互联网'
    // style.draggable = true
    return style
  }
  getNodeStyle() {
    const style = super.getNodeStyle()
    return style
  }
}

class NetworkGroupNode extends GroupNode.view {
  getImageHref() {
    return
  }
}

export default {
  type: 'networkGroup',
  view: NetworkGroupNode,
  model: NetworkGroupModel,
}
