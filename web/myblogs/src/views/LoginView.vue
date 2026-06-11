<template>
  <div class="login-container">
    <div class="login-card">
      <div class="login-header">
        <h1>MyBlogs</h1>
        <p>登录到您的账户</p>
      </div>

      <div class="login-tabs">
        <button
          :class="['tab-btn', { active: activeTab === 'admin' }]"
          @click="activeTab = 'admin'"
        >
          管理员登录
        </button>
        <button
          :class="['tab-btn', { active: activeTab === 'github' }]"
          @click="activeTab = 'github'"
        >
          GitHub 登录
        </button>
      </div>

      <div v-if="activeTab === 'admin'" class="login-form">
        <form @submit.prevent="handleAdminLogin">
          <div class="form-group">
            <label for="username">用户名</label>
            <input
              id="username"
              v-model="username"
              type="text"
              placeholder="请输入用户名"
              required
            />
          </div>

          <div class="form-group">
            <label for="password">密码</label>
            <input
              id="password"
              v-model="password"
              type="password"
              placeholder="请输入密码"
              required
            />
          </div>

          <div class="form-group captcha-group">
            <label for="captcha">验证码</label>
            <div class="captcha-input-wrapper">
              <input
                id="captcha"
                v-model="captchaCode"
                type="text"
                placeholder="请输入验证码"
                required
                maxlength="4"
              />
              <img
                v-if="captchaImage"
                :src="captchaImage"
                class="captcha-image"
                @click="refreshCaptcha"
                title="点击刷新验证码"
              />
            </div>
          </div>

          <button type="submit" class="login-btn" :disabled="loading">
            {{ loading ? '登录中...' : '登录' }}
          </button>
        </form>
      </div>

      <div v-else class="login-form github-login">
        <p class="github-desc">使用 GitHub 账户快速登录</p>
        <button class="github-btn" @click="handleGitHubLogin">
          <svg class="github-icon" viewBox="0 0 24 24" width="20" height="20">
            <path fill="currentColor" d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
          </svg>
          使用 GitHub 登录
        </button>
      </div>

      <div v-if="error" class="error-message">
        {{ error }}
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { login, getCaptcha, refreshCaptcha } from '@/services/authService'
import { useUserStore } from '@/stores/user'

export default {
  name: 'LoginView',
  setup() {
    const router = useRouter()
    const route = useRoute()
    const userStore = useUserStore()

    const activeTab = ref('admin')
    const username = ref('')
    const password = ref('')
    const captchaCode = ref('')
    const captchaID = ref('')
    const captchaImage = ref('')
    const loading = ref(false)
    const error = ref('')

    const loadCaptcha = async () => {
      try {
        const res = await getCaptcha()
        if (res.code === 200) {
          captchaID.value = res.data.captcha_id
          captchaImage.value = res.data.captcha_image
        }
      } catch (e) {
        console.error('获取验证码失败:', e)
      }
    }

    const refreshCaptchaImage = async () => {
      try {
        const res = await refreshCaptcha(captchaID.value)
        if (res.code === 200) {
          captchaID.value = res.data.captcha_id
          captchaImage.value = res.data.captcha_image
          captchaCode.value = ''
        }
      } catch (e) {
        console.error('刷新验证码失败:', e)
        await loadCaptcha()
      }
    }

    onMounted(() => {
      const token = route.query.token
      const errorParam = route.query.error

      if (token) {
        userStore.setToken(token)
        userStore.fetchUserInfo()
        router.push('/admin')
      } else if (errorParam) {
        error.value = decodeURIComponent(errorParam)
      } else {
        loadCaptcha()
      }
    })

    const handleAdminLogin = async () => {
      if (!username.value || !password.value || !captchaCode.value) {
        error.value = '请输入用户名、密码和验证码'
        return
      }

      loading.value = true
      error.value = ''

      try {
        const res = await login(username.value, password.value, captchaID.value, captchaCode.value)
        if (res.code === 200) {
          userStore.setToken(res.data.token)
          userStore.setUserInfo(res.data.user)
          router.push('/admin')
        } else {
          error.value = res.message || '登录失败'
          await refreshCaptchaImage()
        }
      } catch (e) {
        error.value = e.message || '登录失败，请稍后重试'
        await refreshCaptchaImage()
      } finally {
        loading.value = false
      }
    }

    const handleGitHubLogin = () => {
      window.location.href = '/api/front/auth/github'
    }

    return {
      activeTab,
      username,
      password,
      captchaCode,
      captchaImage,
      loading,
      error,
      handleAdminLogin,
      handleGitHubLogin,
      refreshCaptcha: refreshCaptchaImage
    }
  }
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--bg);
  padding: 40px 20px;
}

.login-card {
  background: var(--bg-card);
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  padding: 40px;
  width: 400px;
  box-shadow: var(--shadow-md);
  overflow: hidden;
}

.login-header {
  text-align: center;
  margin-bottom: 32px;
}

.login-header h1 {
  font-size: 28px;
  font-weight: 700;
  color: var(--text);
  margin: 0 0 8px 0;
}

.login-header p {
  color: var(--text-secondary);
  font-size: 14px;
  margin: 0;
}

.login-tabs {
  display: flex;
  gap: 4px;
  margin-bottom: 24px;
  background: var(--bg-secondary);
  border-radius: var(--radius);
  padding: 4px;
}

.tab-btn {
  flex: 1;
  padding: 10px 16px;
  border: none;
  background: transparent;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  color: var(--text-secondary);
  border-radius: var(--radius);
  transition: all 0.2s;
}

.tab-btn:hover {
  color: var(--text);
}

.tab-btn.active {
  background: var(--bg-card);
  color: var(--text);
  box-shadow: var(--shadow);
}

.login-form {
  margin-top: 24px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  font-size: 14px;
  font-weight: 500;
  color: var(--text);
}

.form-group input {
  width: 100%;
  padding: 12px 16px;
  border: 1px solid var(--border);
  border-radius: var(--radius);
  font-size: 14px;
  background: var(--bg);
  color: var(--text);
  transition: border-color 0.2s;
  box-sizing: border-box;
}

.form-group input:focus {
  outline: none;
  border-color: var(--link);
}

.form-group input::placeholder {
  color: var(--text-muted);
}

.captcha-group {
  margin-bottom: 20px;
}

.captcha-input-wrapper {
  display: flex;
  gap: 10px;
  align-items: center;
}

.captcha-input-wrapper input {
  flex: 1;
}

.captcha-image {
  width: 120px;
  height: 40px;
  border: 1px solid var(--border);
  border-radius: var(--radius);
  cursor: pointer;
  transition: opacity 0.2s;
}

.captcha-image:hover {
  opacity: 0.8;
}

.login-btn {
  width: 100%;
  padding: 12px;
  background: var(--link);
  color: white;
  border: none;
  border-radius: var(--radius);
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s;
}

.login-btn:hover:not(:disabled) {
  background: var(--link-hover);
}

.login-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.github-login {
  text-align: center;
  padding: 20px 0;
}

.github-desc {
  color: var(--text-secondary);
  margin-bottom: 20px;
}

.github-btn {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  padding: 12px 24px;
  background: var(--text);
  color: var(--bg);
  border: none;
  border-radius: var(--radius);
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
  transition: opacity 0.2s;
}

.github-btn:hover {
  opacity: 0.9;
}

.github-icon {
  vertical-align: middle;
}

.error-message {
  margin-top: 16px;
  padding: 12px;
  background: var(--bg-secondary);
  color: var(--link);
  border: 1px solid var(--border);
  border-radius: var(--radius);
  font-size: 14px;
  text-align: center;
  flex-shrink: 0;
}
</style>
