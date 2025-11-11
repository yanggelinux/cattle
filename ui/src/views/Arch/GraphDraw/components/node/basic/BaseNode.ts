import { BaseNode, BaseNodeModel } from '@logicflow/core'
import type { JSX } from 'preact'

class BaseNewNode extends BaseNode {
  getShape(): JSX.Element | null {
    throw new Error('Method not implemented.')
  }
}

class BaseNewModel extends BaseNodeModel {
  setAttributes() {
    this.fill = 'red'
  }
}

export default {
  type: 'BaseNode',
  view: BaseNewNode,
  model: BaseNewModel,
}
