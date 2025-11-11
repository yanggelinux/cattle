import ImageNode from './ImageNode'
import integratedMachineImg from '@/assets/images/integrated_machine.png'

class IntegratedMachineNode extends ImageNode.view {
  getImageHref() {
    return integratedMachineImg
  }
}

export default {
  type: 'integratedMachine',
  view: IntegratedMachineNode,
  model: ImageNode.model,
}
