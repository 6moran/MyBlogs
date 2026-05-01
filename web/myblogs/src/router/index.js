import { createRouter, createWebHistory } from 'vue-router'
import AdminArticleManageView from '../views/AdminArticleManageView.vue'
import AdminCategoryView from '../views/AdminCategoryView.vue'
import AdminTagView from '../views/AdminTagView.vue'
import AdminUserView from '../views/AdminUserView.vue'

const routes = [
  {
    path: '/',
    redirect: '/admin/articles'
  },
  {
    path: '/admin/articles',
    name: 'admin-article-manage',
    component: AdminArticleManageView
  },
  {
    path: '/admin/categories',
    name: 'admin-categories',
    component: AdminCategoryView
  },
  {
    path: '/admin/tags',
    name: 'admin-tags',
    component: AdminTagView
  },
  {
    path: '/admin/users',
    name: 'admin-users',
    component: AdminUserView
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
