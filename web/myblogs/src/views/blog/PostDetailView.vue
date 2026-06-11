<template>
  <div class="post-detail-view" v-if="post">
    <h1 class="post-title">{{ post.title }}</h1>
    <div class="post-meta">
      <time class="post-date" :datetime="post.date">{{ post.date }}</time>
      <span class="meta-item">{{ post.viewCount }} 阅读</span>
      <span class="meta-item">{{ post.likeCount }} 点赞</span>
    </div>
    <div class="post-tags" v-if="post.tags.length">
      <span v-for="tag in post.tags" :key="tag.id" class="tag-badge">{{ tag.name }}</span>
    </div>
    <div class="markdown-body" v-html="renderedContent"></div>
    <CommentSection :articleID="post.id" />
    <nav class="post-nav">
      <router-link to="/" class="back-link">&larr; 返回列表</router-link>
    </nav>
  </div>
  <div v-else-if="loading" class="loading">加载中...</div>
  <div v-else class="not-found">
    <h1 class="not-found-title">404</h1>
    <p class="not-found-message">文章不存在</p>
    <router-link to="/" class="back-link">返回首页</router-link>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { marked } from 'marked'
import { getArticleDetail } from '@/services/articleApiService'
import CommentSection from '@/components/blog/CommentSection.vue'

const props = defineProps({
  id: {
    type: String,
    required: true,
  },
})

const post = ref(null)
const loading = ref(true)

const renderedContent = computed(() => {
  if (!post.value) return ''
  return marked(post.value.content)
})

async function fetchArticle() {
  loading.value = true
  try {
    const res = await getArticleDetail(Number(props.id))
    const d = res.data
    post.value = {
      id: d.id,
      title: d.title,
      summary: d.summary,
      content: d.content,
      cover: d.cover_image,
      date: d.created_at?.split('T')[0] || '',
      viewCount: d.view_count,
      likeCount: d.like_count,
      tags: d.tags || []
    }
  } catch {
    post.value = null
  } finally {
    loading.value = false
  }
}

onMounted(fetchArticle)
watch(() => props.id, fetchArticle)
</script>

<style scoped>
.post-detail-view {
  max-width: 780px;
  margin: 0 auto;
  padding: 24px 0 48px;
}

.post-title {
  font-size: 36px;
  font-weight: 700;
  margin: 0 0 8px 0;
  color: var(--text);
  line-height: 1.3;
}

.post-meta {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 14px;
  color: var(--text-secondary);
  margin-bottom: 16px;
}

.meta-item {
  color: var(--text-muted);
}

.post-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 24px;
}

.tag-badge {
  font-size: 12px;
  padding: 2px 10px;
  border-radius: 4px;
  background: var(--bg-secondary);
  color: var(--text-secondary);
  border: 1px solid var(--border-light);
}

.post-nav {
  margin-top: 48px;
  padding-top: 16px;
  border-top: 1px solid var(--border);
}

.back-link {
  color: var(--link);
  text-decoration: none;
  font-size: 15px;
  transition: color 0.2s ease;
}

.back-link:hover {
  color: var(--link-hover);
}

.loading {
  text-align: center;
  padding: 80px 0;
  color: var(--text-secondary);
}

/* 404 page */
.not-found {
  max-width: 780px;
  margin: 0 auto;
  padding: 120px 0 48px;
  text-align: center;
}

.not-found-title {
  font-size: 64px;
  font-weight: 700;
  color: var(--text);
  margin: 0 0 12px 0;
}

.not-found-message {
  font-size: 18px;
  color: var(--text-secondary);
  margin: 0 0 32px 0;
}
</style>
