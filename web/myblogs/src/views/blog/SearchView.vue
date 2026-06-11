<template>
  <div class="search-view">
    <h1 class="page-heading">搜索：{{ query }}</h1>
    <div v-if="loading" class="loading">搜索中...</div>
    <template v-else>
      <p class="result-count" v-if="posts.length">找到 {{ total }} 篇文章</p>
      <PostList v-if="posts.length" :posts="posts" />
      <div v-else class="no-results">
        <p>没有找到相关文章</p>
        <RouterLink to="/" class="back-link">返回首页</RouterLink>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { getPublishedArticles } from '@/services/articleApiService'
import PostList from '@/components/blog/PostList.vue'

const route = useRoute()
const query = ref('')
const posts = ref([])
const total = ref(0)
const loading = ref(false)

async function search() {
  const q = query.value.trim()
  if (!q) {
    posts.value = []
    return
  }
  loading.value = true
  try {
    const res = await getPublishedArticles(1, 50, 0, q)
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
    total.value = data.total || 0
  } catch {
    posts.value = []
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  query.value = route.query.q || ''
  search()
})

watch(() => route.query.q, (val) => {
  query.value = val || ''
  search()
})
</script>

<style scoped>
.page-heading {
  font-size: 24px;
  font-weight: 700;
  margin: 0 0 8px 0;
  color: var(--text);
}

.result-count {
  font-size: 14px;
  color: var(--text-muted);
  margin: 0 0 24px 0;
}

.no-results {
  padding: 60px 0;
  text-align: center;
}

.no-results p {
  color: var(--text-muted);
  font-size: 16px;
  margin-bottom: 16px;
}

.back-link {
  color: var(--link);
  text-decoration: none;
  font-size: 14px;
}

.back-link:hover {
  color: var(--link-hover);
}

.loading {
  text-align: center;
  padding: 40px 0;
  color: var(--text-secondary);
}
</style>
