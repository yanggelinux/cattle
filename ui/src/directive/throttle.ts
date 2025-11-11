import type { Directive } from 'vue'

export const throttleBtn: Directive = {
  mounted(el, binding) {
    let timer: NodeJS.Timeout | null = null
    const delay = binding.arg ? Number(binding.arg) : 1000 // 可通过 v-throttle:3000 设置时间
    const callback = binding.value

    el.addEventListener('click', () => {
      if (timer) return
      callback()
      timer = setTimeout(() => {
        timer = null
      }, delay)
    })
  },
}
