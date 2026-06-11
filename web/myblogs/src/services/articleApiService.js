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

export function getPublishedArticles(page = 1, size = 10, tagID = 0, keyword = '') {
  return api.get('/front/articles', {
    params: { page, size, tag_id: tagID, keyword }
  })
}

export function getArticleDetail(id) {
  return api.get(`/front/articles/${id}`)
}

export function likeArticle(id) {
  return api.post(`/front/articles/${id}/like`)
}
