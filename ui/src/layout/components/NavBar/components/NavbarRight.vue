<template>
  <div class="navbar-right">
    <!-- 用户头像（个人中心、注销登录等） -->
    <el-dropdown trigger="click">
      <div class="user-profile">
        <img class="user-profile__avatar" src="@/assets/images/user-info.png" />
        <span class="user-profile__name">{{ displayName }}</span>
      </div>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item icon="Message">邮箱:{{ email }}</el-dropdown-item>
          <el-dropdown-item icon="House">部门:{{ deptName }}</el-dropdown-item>
          <el-dropdown-item icon="User">角色:{{ roleDisplayNames }}</el-dropdown-item>
          <el-dropdown-item divided @click="logout">注销登录</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </div>
</template>
<script setup lang="ts">
import { useAuthStore, useTagsViewStore } from '@/store'
import { ElMessageBox } from 'element-plus'
import { useRouter } from 'vue-router'

const authStore = useAuthStore()
const tagsViewStore = useTagsViewStore()

const router = useRouter()

const { displayName, email, deptName, roleDisplayNames } = authStore

/**
 * 注销登录
 */
function logout() {
  ElMessageBox.confirm('确定注销并退出系统吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    lockScroll: false,
  })
    .then(() => {
      authStore.logout()
      router.push(`/login`)
      tagsViewStore.delAllViews()
    })
    .catch(() => {
      // 用户点击“取消” 或 关闭弹窗，什么都不做即可避免报错
    })
}
</script>

<style lang="scss" scoped>
$border-color: #e8eaee;
.navbar-right {
  display: flex;
  align-items: center;
  justify-content: center;

  & > * {
    display: inline-block;
    min-width: 40px;
    height: $navbar-height;
    line-height: $navbar-height;
    color: var(--el-text-color);
    text-align: center;
    cursor: pointer;
  }
  .user-profile {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
    padding: 0 13px;

    &__avatar {
      width: 32px;
      height: 32px;
      border-radius: 50%;
    }

    &__name {
      margin-left: 10px;
    }
    &:hover {
      background: rgb(0 0 0 / 10%);
    }
  }
}
</style>
