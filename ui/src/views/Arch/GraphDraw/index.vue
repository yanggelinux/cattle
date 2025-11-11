<template>
  <div class="graph-draw-container">
    <div class="graph-draw-main">
      <Diagram
        v-if="graphID !== 0"
        :graphID="graphID"
        :graphName="graphName"
        :silentMode="silentMode"
        :hasOrder="hasOrder"
        :owner="owner"
        :status="status"
      ></Diagram>
    </div>
  </div>
</template>

<script setup lang="ts">
import Diagram from './components/Diagram.vue'
import { useRoute } from 'vue-router'

defineOptions({
  name: 'GraphDraw',
})
const route = useRoute()
const graphID = route.query.graphID ? Number(route.query.graphID) : 0
const graphName = route.query.graphName ? String(route.query.graphName) : ''
const silentMode = route.query.silentMode ? Number(route.query.silentMode) : 0
const hasOrder = route.query.hasOrder ? Number(route.query.hasOrder) : 0
const status = route.query.status ? Number(route.query.status) : 0
const owner = route.query.owner ? String(route.query.owner) : ''
</script>

<style lang="scss" scoped>
.graph-draw-container {
  height: 100%;
  width: 100%;
  touch-action: none;
  .graph-draw-main {
    height: 100%;
    width: 100%;
  }
}
</style>
