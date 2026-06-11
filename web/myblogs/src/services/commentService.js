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

export function getComments(articleID, page = 1, size = 10) {
  return api.get(`/front/articles/${articleID}/comments`, {
    params: { page, size }
  })
}

export function createComment(articleID, content, parentID = 0) {
  return api.post(`/front/articles/${articleID}/comments`, {
    content,
    parent_id: parentID
  })
}
