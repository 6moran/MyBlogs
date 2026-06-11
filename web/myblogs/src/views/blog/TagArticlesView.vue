<template>
  <div class="tag-articles-page">
    <h1 class="page-title">{{ tagName }}</h1>
    <div v-if="loading" class="loading">加载中...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <template v-else>
      <PostList :posts="posts" />
      <div v-if="posts.length === 0" class="empty">
        <p>暂无相关文章</p>
        <RouterLink to="/tags">返回标签</RouterLink>
      </div>
      <Pagination
        v-if="totalPages > 1"
        :current-page="currentPage"
        :total-pages="totalPages"
        :on-page-change="goToPage"
      />
    </template>
    <div class="back-link">
      <RouterLink to="/tags">← 返回标签</RouterLink>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { getPublishedArticles } from '@/services/articleApiService'
import { getTags } from '@/services/tagService'
import PostList from '@/components/blog/PostList.vue'
import Pagination from '@/components/blog/Pagination.vue'

const props = defineProps({
  id: { type: String, required: true }
})

const posts = ref([])
const tagName = ref('')
const currentPage = ref(1)
const totalPages = ref(1)
const loading = ref(true)
const error = ref('')
const PER_PAGE = 10

async function fetchData() {
  loading.value = true
  error.value = ''
  try {
    const [tagsRes, articlesRes] = await Promise.all([
      getTags(),
      getPublishedArticles(currentPage.value, PER_PAGE, Number(props.id))
    ])
    const tag = (tagsRes.data || []).find(t => t.id === Number(props.id))
    tagName.value = tag ? tag.name : '标签'

    const data = articlesRes.data
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

onMounted(fetchData)
watch(() => props.id, () => {
  currentPage.value = 1
  fetchData()
})
watch(currentPage, fetchData)
</script>

<style scoped>
.page-title {
  font-size: 24px;
  font-weight: 700;
  margin: 0 0 24px;
  color: var(--text);
}

.empty {
  text-align: center;
  padding: 40px 0;
  color: var(--text-secondary);
}

.empty a {
  color: var(--link);
  margin-top: 8px;
  display: inline-block;
}

.back-link {
  margin-top: 32px;
  padding-top: 20px;
  border-top: 1px solid var(--border);
}

.back-link a {
  color: var(--link);
  font-size: 14px;
}

.loading, .error {
  text-align: center;
  padding: 40px 0;
  color: var(--text-secondary);
}
</style>
