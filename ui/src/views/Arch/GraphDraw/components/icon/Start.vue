<template>
  <!-- <svg class="svg-node">
    <g transform="translate(0.5,0.5)" style="visibility: visible">
      <ellipse
        cx="15.98"
        cy="14.96"
        rx="13.600000000000001"
        ry="13.600000000000001"
        fill="#ffffff"
        stroke="#409EFF"
        stroke-width="1.3"
        pointer-events="all"
      ></ellipse>
    </g>
  </svg> -->
  <svg class="svg-node" viewBox="0 0 32 32">
    <ellipse
      cx="16"
      cy="16"
      rx="13.6"
      ry="13.6"
      fill="#ffffff"
      stroke="#409EFF"
      stroke-width="1.3"
    />
  </svg>
</template>

<script setup lang="ts"></script>
<style scoped>
.svg-node {
  left: 1px;
  top: 1px;
  width: 60px;
  height: 60px;
  display: block;
  position: relative;
  overflow: hidden;
}
</style>
