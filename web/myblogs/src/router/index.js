import { createRouter, createWebHistory } from 'vue-router'
import AdminLayout from '../layouts/AdminLayout.vue'
import BlogLayout from '../layouts/BlogLayout.vue'
import AdminArticleManageView from '../views/AdminArticleManageView.vue'
import AdminTagView from '../views/AdminTagView.vue'
import AdminUserView from '../views/AdminUserView.vue'
import PostListView from '../views/blog/PostListView.vue'
import PostDetailView from '../views/blog/PostDetailView.vue'
import AboutView from '../views/blog/AboutView.vue'
import TagsView from '../views/blog/TagsView.vue'
import TagArticlesView from '../views/blog/TagArticlesView.vue'
import SearchView from '../views/blog/SearchView.vue'
import LoginView from '../views/LoginView.vue'

const routes = [
  {
    path: '/login',
    name: 'login',
    component: LoginView
  },
  {
    path: '/',
    component: BlogLayout,
    children: [
      {
        path: '',
        name: 'home',
        component: PostListView
      },
      {
        path: 'posts/:id',
        name: 'post-detail',
        component: PostDetailView,
        props: true
      },
      {
        path: 'about',
        name: 'about',
        component: AboutView
      },
      {
        path: 'tags',
        name: 'tags',
        component: TagsView
      },
      {
        path: 'tags/:id',
        name: 'tag-articles',
        component: TagArticlesView,
        props: true
      },
      {
        path: 'search',
        name: 'search',
        component: SearchView
      }
    ]
  },
  {
    path: '/admin',
    component: AdminLayout,
    redirect: '/admin/articles',
    children: [
      {
        path: 'articles',
        name: 'admin-article-manage',
        component: AdminArticleManageView
      },
      {
        path: 'tags',
        name: 'admin-tags',
        component: AdminTagView
      },
      {
        path: 'users',
        name: 'admin-users',
        component: AdminUserView
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')

  if (to.path.startsWith('/admin')) {
    if (!token) {
      next({ path: '/login', query: { redirect: to.fullPath } })
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router
