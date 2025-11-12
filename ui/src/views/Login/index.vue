<template>
  <div class="login-container">
    <div class="login-header"></div>
    <div class="login-main">
      <div class="login-content">
        <div class="login-content-top">
          <img class="image-logo" src="../../assets/logo.png" height="50px" width="50px" />
          <div class="top-title">
            <span>流程架构平台</span>
          </div>
        </div>
        <div class="login-content-main">
          <div class="desc-wrapper">
            <span>流程架构平台&nbsp;最权威的流程架构</span>
          </div>
          <el-form
            :model="loginForm"
            :rules="rules"
            ref="loginFormRef"
            label-width="30px"
            class="loginForm"
          >
            <el-form-item label prop="userName" class="username">
              <el-tooltip class="item" effect="dark" content="请输入用户名" placement="top-start">
                <el-input
                  prefix-icon="User"
                  v-model="loginForm.userName"
                  name="userName"
                  size="large"
                  placeholder="用户名"
                ></el-input>
              </el-tooltip>
            </el-form-item>
            <el-form-item label prop="password" class="password">
              <el-tooltip class="item" effect="dark" content="请输入密码" placement="top-start">
                <el-input
                  prefix-icon="Lock"
                  type="password"
                  name="password"
                  size="large"
                  placeholder="密码"
                  show-password
                  v-model="loginForm.password"
                  @keyup="checkCapslock"
                  @keyup.enter="handleLoginSubmit"
                ></el-input>
              </el-tooltip>
            </el-form-item>
            <el-button
              :loading="loading"
              type="primary"
              size="large"
              class="btn-login"
              @click.prevent="handleLoginSubmit"
            >
              登录
            </el-button>
          </el-form>
        </div>
      </div>
    </div>
    <div class="login-footer">
      <span>Copyright © 2025 技术支持</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { type LoginFormData } from '@/api/auth'
import router from '@/router'
import type { FormInstance } from 'element-plus'
import { ElMessage } from 'element-plus'
import { useAuthStore } from '@/store'
import { computed, nextTick, ref } from 'vue'

const authStore = useAuthStore()

const loginFormRef = ref<FormInstance>()

const loading = ref(false) // 按钮 loading 状态
const isCapslock = ref(false) // 是否大写锁定

const loginForm = ref<LoginFormData>({
  userName: '',
  password: '',
})

const rules = computed(() => {
  return {
    userName: [
      {
        required: true,
        trigger: 'blur',
        message: '请输入用户名',
      },
    ],
    password: [
      {
        required: true,
        trigger: 'blur',
        message: '请输入密码',
      },
      {
        min: 6,
        message: '密码长度不能小于6位',
        trigger: 'blur',
      },
    ],
  }
})

// 登录
function handleLoginSubmit() {
  loginFormRef.value?.validate(async (valid: boolean) => {
    if (!valid) return
    loading.value = true
    try {
      const { password, userName } = loginForm.value
      if (!password) {
        return
      }
      // const { iv, cipher } = await encryptPassword(password)
      // const ePassowrd = `${iv}@${cipher}`
      const loginData: LoginFormData = {
        userName: userName,
        password: password,
      }
      await authStore.login(loginData)
      await nextTick()
      // 跳转到登录前的页面
      await router.push({ path: '/dashboard', replace: true })
    } catch (err: any) {
      console.error(err.message)
      ElMessage.error(err.message)
    } finally {
      loading.value = false
    }
  })
}

/**
 * 解析 redirect 字符串 为 path 和  queryParams
 *
 * @returns { path: string, queryParams: Record<string, string> } 解析后的 path 和 queryParams
 */
// function parseRedirect(): {
//   path: string
//   queryParams: Record<string, string>
// } {
//   const query: LocationQuery = route.query
//   const redirect = (query.redirect as string) ?? '/'

//   const url = new URL(redirect, window.location.origin)
//   const path = url.pathname
//   const queryParams: Record<string, string> = {}

//   url.searchParams.forEach((value, key) => {
//     queryParams[key] = value
//   })

//   return { path, queryParams }
// }

// 检查输入大小写
function checkCapslock(event: KeyboardEvent) {
  // 防止浏览器密码自动填充时报错
  if (event instanceof KeyboardEvent) {
    isCapslock.value = event.getModifierState('CapsLock')
  }
}
</script>

<style lang="scss" scoped>
$imgPath: '../../assets/images';
$title-btn-color: #1b9dfc;
$text-color: rgba(0, 0, 0, 0.45);
.login-container {
  height: 100%;
  background: url('@/assets/images/login-bg.jpg') no-repeat;
  background-size: 100% 100%;
  display: flex;
  flex-direction: column;
  min-height: 686px;
  .login-header {
    flex: 0 0 50px;
  }
  .login-main {
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
    .login-content {
      width: 396px;
      height: 396px;
      text-align: center;
      overflow: hidden;
      .login-content-top {
        display: flex;
        height: 50px;
        margin-left: 20px;
        .top-title {
          font-size: 38px;
          line-height: 50px;
          color: rgba(0, 0, 0, 0.85);
          font-family:
            Avenir,
            Helvetica Neue,
            Arial,
            Helvetica,
            sans-serif;
          text-align: center;
        }
        .image-logo {
          margin-left: 45px;
          margin-right: 10px;
          text-align: center;
        }
      }

      .login-content-main {
        .desc-wrapper {
          font-size: 14px;
          color: $text-color;
          margin-top: 12px;
          margin-bottom: 40px;
          text-align: center;
        }
        .username {
          margin-top: 20px;
        }
        .password {
          margin-top: 60px;
        }
        .btn-login {
          cursor: pointer;
          margin-top: 35px;
          margin-left: 30px;
          font-size: 16px;
          color: white;
          background-color: $title-btn-color;
          height: 38px;
          width: 366px;
          line-height: 38px;
          text-align: center;
          border: 0;
          border-radius: 10px;
          outline: none;
        }
      }
    }
  }
  .login-footer {
    flex: 0 0 50px;
    line-height: 50px;
    text-align: center;
    color: $text-color;
    font-size: 14px;
  }
}
</style>
