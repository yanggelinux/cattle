import { onMounted, onUnmounted, ref } from 'vue'

export function useAutoSave<T>(
  saveFn: (params: T) => Promise<void>, // 保存函数
  getParams: () => T, // 每次执行时获取参数
  interval: number = 10000, // 默认 10 秒
  runImmediately: boolean = true // 是否立即执行一次
) {
  const isActive = ref(false)
  const isRunning = ref(false) // 是否正在执行中
  let timer: number | undefined
  let currentInterval = interval

  // 内部执行函数，带防重入
  const execute = async () => {
    if (isRunning.value) {
      console.warn('上一次保存还未完成，跳过本次执行')
      return
    }
    isRunning.value = true
    try {
      await saveFn(getParams())
    } catch (err) {
      console.error('保存失败:', err)
    } finally {
      isRunning.value = false
    }
  }

  // 启动
  const start = () => {
    if (isActive.value) return
    isActive.value = true
    if (runImmediately) {
      execute()
    }
    timer = window.setInterval(() => {
      execute()
    }, currentInterval)
  }

  // 停止
  const stop = () => {
    if (timer) {
      clearInterval(timer)
      timer = undefined
    }
    isActive.value = false
  }

  // 修改间隔
  const setIntervalTime = (newInterval: number) => {
    currentInterval = newInterval
    if (isActive.value) {
      stop()
      start()
    }
  }

  onMounted(() => start())
  onUnmounted(() => stop())

  return {
    isActive,
    isRunning,
    start,
    stop,
    setIntervalTime,
  }
}
