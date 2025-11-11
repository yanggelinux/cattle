import axios, { type InternalAxiosRequestConfig, type AxiosResponse } from 'axios'
// import router from '@/router'
import {
  getAccessToken,
  getUserID,
  getUserName,
  getAuthorization,
  getIsSuper,
  getDeptName,
} from '@/utils/auth'
import { ElMessage } from 'element-plus'
import { useAuthStore } from '@/store'
import router from '@/router'

function genBaseUrl(): string {
  let baseURL: string = ''
  if (import.meta.env.MODE === 'development') {
    //开发环境
    baseURL = import.meta.env.VITE_APP_BASE_API
  } else {
    const { protocol, host } = document.location
    const url: string = `${protocol}//${host}/cattle`
    baseURL = url
  }
  return baseURL
}

// 创建 axios 实例
const service = axios.create({
  baseURL: genBaseUrl(),
  timeout: 50000,
  headers: { 'Content-Type': 'application/json;charset=utf-8' },
})

// 请求拦截器
service.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const token = getAccessToken()
    const userID = getUserID()
    const userName = getUserName()
    const authorization = getAuthorization()
    const deptName = getDeptName()
    const isSuper = getIsSuper()
    if (token !== null) {
      // 让每个请求携带token-- ['X-Token']为自定义key 请根据实际情况自行修改
      config.headers['X-Token'] = token
    }
    if (userID) {
      config.headers['X-UserID'] = userID
    }
    if (userName) {
      config.headers['X-Username'] = encodeURI(userName)
    }
    if (deptName) {
      config.headers['X-DeptName'] = encodeURI(deptName)
    }
    if (authorization) {
      config.headers['X-Authorization'] = encodeURI(authorization)
    }
    if (isSuper) {
      config.headers['X-Super'] = isSuper
    }
    return config
  },
  (error) => {
    console.log(error)
    Promise.reject(error)
  }
)

// 响应拦截器
service.interceptors.response.use(
  (response: AxiosResponse) => {
    // 如果响应是二进制流，则直接返回，用于下载文件、Excel 导出等
    const resp: AxiosResponse = response
    const status: number = resp.data.status
    const msg: string = resp.data.msg
    if ([10001, 10002, 10003, 10004, 10004, 401].includes(status)) {
      // 如果token 过期将登陆状态 设置为false,并跳转到登陆页面
      if (!['/login', '/403', '/404'].includes(location.pathname)) {
        const authStore = useAuthStore()
        authStore.clearSessionAndCache().then(() => {
          router.push('/login')
        })
      }
      ElMessage.error(msg || '系统出错')
      return Promise.reject(new Error(msg || 'Error'))
    }
    return response
  },
  async (error: any) => {
    console.log('err ' + error)
    return Promise.reject(error.message)
  }
)

export default service
