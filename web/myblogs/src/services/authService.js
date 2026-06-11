import axios from 'axios'

const api = axios.create({
  baseURL: '/api',
  timeout: 10000
})

api.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

api.interceptors.response.use(
  response => response.data,
  error => {
    const message = error.response?.data?.message || '请求失败'
    return Promise.reject(new Error(message))
  }
)

export function login(username, password, captchaId, captchaCode) {
  return api.post('/front/auth/login', { username, password, captcha_id: captchaId, captcha_code: captchaCode })
}

export function generateCaptcha() {
  return api.get('/front/auth/captcha')
}

export function getCaptcha() {
  return api.get('/front/auth/captcha')
}

export function refreshCaptcha(captchaId) {
  return api.post('/front/auth/captcha/refresh', { captcha_id: captchaId })
}

export function getGitHubLoginURL() {
  return `${api.defaults.baseURL}/front/auth/github`
}

export function getGitHubUserInfo() {
  return api.get('/front/user/me')
}

export { api }
