import _ from 'lodash'

/**
 * useThrottleAsync
 * 支持异步函数和参数传递
 *
 * @param fn 需要节流的异步函数
 * @param wait 节流时间（毫秒）
 * @param options lodash throttle options
 * @returns 返回节流后的异步函数
 */
export function useThrottleAsync<T extends (...args: any[]) => Promise<any>>(
  fn: T,
  wait: number = 2000,
  options: { leading?: boolean; trailing?: boolean } = { leading: true, trailing: false }
) {
  // 使用 lodash throttle 创建节流函数
  const throttledFn = _.throttle(
    async (...args: Parameters<T>): Promise<ReturnType<T> | void> => {
      try {
        return await fn(...args)
      } catch (err) {
        console.error('节流异步函数执行失败:', err)
      }
    },
    wait,
    options
  )

  return throttledFn
}
