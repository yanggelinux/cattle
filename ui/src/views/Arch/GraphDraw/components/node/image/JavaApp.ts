import ImageNode from './ImageNode'
import javayAppImg from '@/assets/images/java_app.png'
// import { h } from '@logicflow/core'

class JavaAppModel extends ImageNode.model {
  initNodeData(data: any) {
    super.initNodeData(data)
    this.width = 50
    this.height = 40
  }
}

// 交换机 的图片节点
class javaAppNode extends ImageNode.view {
  getImageHref() {
    return javayAppImg
  }
}

export default {
  type: 'javaApp',
  view: javaAppNode,
  model: JavaAppModel,
}
