<template>
  <div class="post-list-view">
    <h1 class="page-heading">文章</h1>
    <div v-if="loading" class="loading">加载中...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <template v-else>
      <PostList :posts="posts" />
      <Pagination
        :current-page="currentPage"
        :total-pages="totalPages"
        :on-page-change="goToPage"
      />
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { getPublishedArticles } from '@/services/articleApiService'
import PostList from '@/components/blog/PostList.vue'
import Pagination from '@/components/blog/Pagination.vue'

const posts = ref([])
const currentPage = ref(1)
const totalPages = ref(1)
const loading = ref(true)
const error = ref('')
const PER_PAGE = 10

async function fetchArticles() {
  loading.value = true
  error.value = ''
  try {
    const res = await getPublishedArticles(currentPage.value, PER_PAGE)
    const data = res.data
    posts.value = (data.articles || []).map(a => ({
      id: a.id,
      title: a.title,
      excerpt: a.summary,
      date: a.created_at?.split('T')[0] || '',
      cover: a.cover_image,
      likes: a.like_count,
      comments: a.comment_count,
      tags: (a.tags || []).map(t => t.name)
    }))
    totalPages.value = Math.ceil((data.total || 0) / PER_PAGE) || 1
  } catch (e) {
    error.value = e.message || '加载失败'
  } finally {
    loading.value = false
  }
}

function goToPage(page) {
  if (page < 1 || page > totalPages.value) return
  currentPage.value = page
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

onMounted(fetchArticles)
watch(currentPage, fetchArticles)
</script>

<style scoped>
.page-heading {
  font-size: 24px;
  font-weight: 700;
  margin: 0 0 24px 0;
  color: var(--text);
}

.loading, .error {
  text-align: center;
  padding: 40px 0;
  color: var(--text-secondary);
}
</style>
