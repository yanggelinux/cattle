import ImageNode from './ImageNode'
import redisImg from '@/assets/images/redis.png'

// Redis 的图片节点
class RedisNode extends ImageNode.view {
  getImageHref() {
    return redisImg
  }
}

export default {
  type: 'redis',
  view: RedisNode,
  model: ImageNode.model,
}
