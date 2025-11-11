/**
 * Check if an element has a class
 * @param {HTMLElement} ele
 * @param {string} cls
 * @returns {boolean}
 */
export function hasClass(ele: HTMLElement, cls: string) {
  return !!ele.className.match(new RegExp('(\\s|^)' + cls + '(\\s|$)'))
}

/**
 * Add class to element
 * @param {HTMLElement} ele
 * @param {string} cls
 */
export function addClass(ele: HTMLElement, cls: string) {
  if (!hasClass(ele, cls)) ele.className += ' ' + cls
}

/**
 * Remove class from element
 * @param {HTMLElement} ele
 * @param {string} cls
 */
export function removeClass(ele: HTMLElement, cls: string) {
  if (hasClass(ele, cls)) {
    const reg = new RegExp('(\\s|^)' + cls + '(\\s|$)')
    ele.className = ele.className.replace(reg, ' ')
  }
}

/**
 * 判断是否是外部链接
 *
 * @param {string} path
 * @returns {Boolean}
 */
export function isExternal(path: string) {
  const isExternal = /^(https?:|http?:|mailto:|tel:)/.test(path)
  return isExternal
}

export const setTitle = (title: any) => {
  title = title ? '-' + title : ''
  window.document.title = '架构流程平台' + title || '架构流程平台'
}

export function removePrefix(str: string, prefix: string): string {
  return str.startsWith(prefix) ? str.slice(prefix.length) : str
}

export function removeSuffix(str: string, suffix: string): string {
  return str.endsWith(suffix) ? str.slice(0, -suffix.length) : str
}

export function removeAt<T>(arr: T[], index: number): void {
  if (index >= 0 && index < arr.length) {
    arr.splice(index, 1)
  }
}
export function removeAtImmutable<T>(arr: T[], index: number): T[] {
  if (index < 0 || index >= arr.length) {
    return [...arr] // 返回原数组的浅拷贝
  }

  return [...arr.slice(0, index), ...arr.slice(index + 1)]
}

/**
 * 获取当前时间的字符串，格式为 yyyyMMddHH，精确到小时
 * @param useUTC 是否使用 UTC 时间（默认 false = 本地时间）
 * @returns 格式化后的时间字符串
 */
export function getCurrentHourString(useUTC: boolean = false): string {
  const now = new Date()

  const year = useUTC ? now.getUTCFullYear() : now.getFullYear()
  const month = (useUTC ? now.getUTCMonth() : now.getMonth()) + 1
  const day = useUTC ? now.getUTCDate() : now.getDate()
  const hour = useUTC ? now.getUTCHours() : now.getHours()

  // 补零格式化
  const pad = (n: number) => n.toString().padStart(2, '0')

  return `${year}${pad(month)}${pad(day)}${pad(hour)}`
}

export function getCurrentMinutesString(useUTC: boolean = false): string {
  const now = new Date()

  const year = useUTC ? now.getUTCFullYear() : now.getFullYear()
  const month = (useUTC ? now.getUTCMonth() : now.getMonth()) + 1
  const day = useUTC ? now.getUTCDate() : now.getDate()
  const hour = useUTC ? now.getUTCHours() : now.getHours()
  const minutes = useUTC ? now.getUTCHours() : now.getMinutes()

  // 补零格式化
  const pad = (n: number) => n.toString().padStart(2, '0')

  return `${year}${pad(month)}${pad(day)}${pad(hour)}${pad(minutes)}`
}

export function toDateTimeString(date: Date): string {
  const pad = (n: number) => n.toString().padStart(2, '0')

  const year = date.getFullYear()
  const month = pad(date.getMonth() + 1)
  const day = pad(date.getDate())
  const hours = pad(date.getHours())
  const minutes = pad(date.getMinutes())
  const seconds = pad(date.getSeconds())

  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
}

/**
 * 将IP字符串转为数字
 * @param ip IPv4地址字符串
 */
function ipToNumber(ip: string): number {
  return ip
    .split('.')
    .map(Number)
    .reduce((acc, val) => (acc << 8) + val)
}

/**
 * IP地址排序函数
 * @param ips IP地址数组
 * @returns 排序后的IP地址数组
 */
export function sortIPs(ips: string[]): string[] {
  return ips.sort((a, b) => ipToNumber(a) - ipToNumber(b))
}

// 去掉机器后面的编号
export function stripNumbers(input: string): string {
  const parts = input.split('-')
  const last = parts[parts.length - 1]
  if (/^\d+$/.test(last)) {
    parts.pop()
  }
  return parts.join('-')
}
