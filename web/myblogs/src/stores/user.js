import { defineStore } from 'pinia'
import { getGitHubUserInfo } from '@/services/authService'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('token') || null,
    userInfo: null
  }),

  getters: {
    isLoggedIn: state => !!state.token,
    isAdmin: state => state.userInfo?.role === 1
  },

  actions: {
    setToken(token) {
      this.token = token
      localStorage.setItem('token', token)
    },

    setUserInfo(info) {
      this.userInfo = info
    },

    async fetchUserInfo() {
      if (!this.token) return
      try {
        const res = await getGitHubUserInfo()
        if (res.code === 200) {
          this.userInfo = res.data
        }
      } catch {
        this.logout()
      }
    },

    logout() {
      this.token = null
      this.userInfo = null
      localStorage.removeItem('token')
      localStorage.removeItem('github_code')
    },

    handleGitHubCallback() {
      const urlParams = new URLSearchParams(window.location.search)
      const code = urlParams.get('code')
      if (code) {
        localStorage.setItem('github_code', code)
        window.history.replaceState({}, document.title, window.location.pathname)
      }
    }
  }
})
