import ImageNode from './ImageNode'
import img from '@/assets/images/arch_graph.png'

class ArchGroupModel extends ImageNode.model {
  initNodeData(data: any) {
    super.initNodeData(data)
    this.width = 40
    this.height = 40
  }
  getTextStyle() {
    const style = super.getTextStyle()
    style.x = this.x
    style.y = this.y + this.height / 2 + 10
    style.editable = false
    return style
  }
}

// 云形状的图片节点
class ArchGraphNode extends ImageNode.view {
  getImageHref() {
    return img
  }
}

export default {
  type: 'archGraph',
  view: ArchGraphNode,
  model: ArchGroupModel,
}
