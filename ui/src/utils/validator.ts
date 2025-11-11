import type { FormInstance } from 'element-plus'

export function calculateDisplayLength(str: string): number {
  return str.trim().length
}

export function validateIP(ipString: string, label: string): [boolean, string] {
  const ipRegex =
    /^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/

  if (ipString.includes('/') && ipString.includes('^')) {
    return [
      ipString.split('^').every((ip) => {
        const [addr, mask] = ip.split('/')
        return ipRegex.test(addr) && mask && +mask >= 0 && +mask <= 32
      }),
      `${label}格式不正确！${ipString}`,
    ]
  }

  if (ipString.includes('-') && ipString.includes('^')) {
    return [
      ipString.split('^').every((segment) => {
        if (segment.includes('-')) {
          const [start, end] = segment.split('-')
          return ipRegex.test(start) && ipRegex.test(end) && start !== end
        } else {
          return ipRegex.test(segment)
        }
      }),
      `${label}格式不正确！${ipString}`,
    ]
  }

  if (ipString.includes('/')) {
    const [ip, mask] = ipString.split('/')
    return [ipRegex.test(ip) && +mask >= 0 && +mask <= 32, `${label}格式不正确！${ipString}`]
  }

  if (ipString.includes('-')) {
    const [start, end] = ipString.split('-')
    return [
      ipRegex.test(start) && ipRegex.test(end) && start !== end,
      `${label}格式不正确！${ipString}`,
    ]
  }

  if (ipString.includes('^')) {
    return [ipString.split('^').every((ip) => ipRegex.test(ip)), `${label}格式不正确！${ipString}`]
  }

  return [ipRegex.test(ipString), `${label}格式不正确！${ipString}`]
}

// 将 validate 包装成 Promise
function validateForm(formRef: FormInstance | undefined): Promise<boolean> {
  return new Promise((resolve) => {
    formRef?.validate((valid: boolean) => {
      resolve(valid)
    })
  })
}

/**
 * 顺序校验多个表单
 * @param forms 表单 ref 数组
 * @returns 最终是否全部校验通过
 */
export async function validateFormsSequentially(
  forms: (FormInstance | undefined)[]
): Promise<boolean> {
  let allValid = true

  for (let i = 0; i < forms.length; i++) {
    const valid = await validateForm(forms[i])
    if (!valid) allValid = false
  }

  return allValid
}
