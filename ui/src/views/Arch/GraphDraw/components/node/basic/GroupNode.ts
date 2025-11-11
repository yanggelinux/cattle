import { dynamicGroup } from '@logicflow/extension'
import { getShapeStyleFuction, getTextStyleFunction } from '../getShapeStyleUtil'

class GroupNodeModel extends dynamicGroup.model {
  setToBottom() {
    this.zIndex = 0
  }

  getNodeStyle() {
    const style = super.getNodeStyle()
    const properties = this.getProperties()
    return getShapeStyleFuction(style, properties)
  }

  getTextStyle() {
    const style = super.getTextStyle()
    const properties = this.getProperties()
    return getTextStyleFunction(style, properties)
  }
  getAnchorStyle() {
    const style: any = super.getAnchorStyle()
    style.stroke = '#000000'
    style.r = 4
    style.fill = '#FFFFFF'
    style.hover.r = 10
    style.hover.fill = '#949494'
    style.hover.stroke = '#949494'
    style.hover.fillOpacity = 0.5
    return style
  }
}

export default {
  type: 'baseGroup',
  view: dynamicGroup.view,
  model: GroupNodeModel,
}
