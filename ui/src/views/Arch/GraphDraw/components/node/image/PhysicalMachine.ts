import ImageNode from './ImageNode'
import physicalMachineImg from '@/assets/images/physical_machine.png'

class PhysicalMachineNode extends ImageNode.view {
  getImageHref() {
    return physicalMachineImg
  }
}

export default {
  type: 'physicalMachine',
  view: PhysicalMachineNode,
  model: ImageNode.model,
}
