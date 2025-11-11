<template>
  <div class="hamburger-wrapper" @click="toggleClick">
    <el-icon size="30" v-if="isActive"><Fold /></el-icon>
    <el-icon size="30" v-else><Expand /></el-icon>
  </div>
</template>

<script setup lang="ts">
defineProps({
  isActive: { type: Boolean, required: true },
})

const emit = defineEmits(['toggleClick'])

function toggleClick() {
  emit('toggleClick')
}
</script>

<style scoped lang="scss">
.hamburger-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 15px;
  cursor: pointer;

  .hamburger {
    vertical-align: middle;
    transform: scaleX(-1);
    transition: transform 0.3s ease;

    &--white {
      color: #fff;
    }

    &.is-active {
      transform: scaleX(1);
    }
  }
}
</style>
