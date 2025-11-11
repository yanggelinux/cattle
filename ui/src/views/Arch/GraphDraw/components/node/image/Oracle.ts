import ImageNode from './ImageNode'
import oracleImg from '@/assets/images/oracle.png'

// Oracle 的图片节点
class OracleNode extends ImageNode.view {
  getImageHref() {
    return oracleImg
  }
}

export default {
  type: 'oracle',
  view: OracleNode,
  model: ImageNode.model,
}
