import ImageNode from './ImageNode'
import sliderImg from '@/assets/images/slider.png'
// 左上角ICON为消息的节点
class SliderNode extends ImageNode.view {
  getImageHref() {
    return sliderImg
  }
}

export default {
  type: 'slider',
  view: SliderNode,
  model: ImageNode.model,
}
