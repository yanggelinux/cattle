<!-- <template>
  <svg>
    <g>
      <g transform="translate(0.5,0.5)" style="visibility: visible">
        <rect
          x="0.98"
          y="7.84"
          width="29.4"
          height="12.74"
          fill="none"
          stroke="white"
          pointer-events="stroke"
          visibility="hidden"
          stroke-width="9"
        ></rect>
        <rect
          x="0.98"
          y="7.84"
          width="29.4"
          height="12.74"
          fill="none"
          stroke="none"
          pointer-events="all"
        ></rect>
      </g>
      <g style="">
        <clipPath id="mx-clip-2-9-28-15-0">
          <rect x="2" y="9" width="28" height="15"></rect>
        </clipPath>
        <g
          fill="#000000"
          font-family="Helvetica"
          clip-path="url(https://app.diagrams.net/#mx-clip-2-9-28-15-0)"
          font-size="5.88px"
        >
          <text x="3.92" y="16.66">Text Node</text>
        </g>
      </g>
    </g>
  </svg>
</template>

<script setup lang="ts"></script> -->

<template>
  <svg width="30" height="30" xmlns="http://www.w3.org/2000/svg">
    <!-- 虚线边框矩形 -->
    <rect
      x="1"
      y="1"
      width="30"
      height="28"
      fill="none"
      stroke="#000000"
      stroke-width="0.3"
      stroke-dasharray="2 1"
    />

    <!-- 居中文本 -->
    <text
      x="15"
      y="15"
      fill="#000000"
      font-family="Helvetica"
      font-size="10px"
      text-anchor="middle"
      dominant-baseline="middle"
    >
      Text
    </text>
  </svg>
</template>

<script setup lang="ts"></script>
