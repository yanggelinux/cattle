import { globalIgnores } from 'eslint/config'
import { defineConfigWithVueTs, vueTsConfigs } from '@vue/eslint-config-typescript'
import pluginVue from 'eslint-plugin-vue'
import skipFormatting from '@vue/eslint-config-prettier/skip-formatting'

import pluginPrettier from 'eslint-plugin-prettier' // ✅ 引入 prettier 插件
import prettierConfig from 'eslint-config-prettier' // ✅ 用于关闭 ESLint 与 Prettier 冲突的格式化规则

// To allow more languages other than `ts` in `.vue` files, uncomment the following lines:
// import { configureVueProject } from '@vue/eslint-config-typescript'
// configureVueProject({ scriptLangs: ['ts', 'tsx'] })
// More info at https://github.com/vuejs/eslint-config-typescript/#advanced-setup

export default defineConfigWithVueTs(
  {
    name: 'app/files-to-lint',
    files: ['**/*.{ts,mts,tsx,vue}'],
  },

  globalIgnores(['**/dist/**', '**/dist-ssr/**', '**/coverage/**']),

  pluginVue.configs['flat/essential'],
  vueTsConfigs.recommended,
  skipFormatting,

  prettierConfig, // ✅ 禁用 ESLint 中的格式化规则（防止与 Prettier 冲突）

  {
    // 在这里添加规则
    rules: {
      // ✅ 使用 prettier 插件并将其问题作为错误处理
      'prettier/prettier': ['error'],
      '@typescript-eslint/no-explicit-any': 'off', // 允许使用 any
      '@typescript-eslint/no-empty-function': 'off', // 允许空函数
      '@typescript-eslint/no-empty-object-type': 'off', // 允许空对象类型
      'vue/multi-word-component-names': 'off', //关闭要求组件名必须由多个单词组成（如 UserList 而非 User）
    },
    plugins: {
      prettier: pluginPrettier,
    },
  }
)
