import ImageNode from './ImageNode'
import encryptMachineImg from '@/assets/images/encrypt_machine.png'
// import { h } from '@logicflow/core'

// 加密机 的图片节点
class EncryptMachineNode extends ImageNode.view {
  getImageHref() {
    return encryptMachineImg
  }
}

export default {
  type: 'encryptMachine',
  view: EncryptMachineNode,
  model: ImageNode.model,
}
