<template>
  <div class="search-wrapper" ref="wrapperRef">
    <div class="search-bar" :class="{ open: visible }">
      <div class="input-slide">
        <input
          ref="inputRef"
          v-model="query"
          class="search-input"
          type="text"
          placeholder="搜索文章..."
          @keydown.escape="close"
          @keydown.enter="goSearch"
        />
      </div>
      <button class="search-btn" @click="toggle" aria-label="搜索">
        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="11" cy="11" r="8"/>
          <line x1="21" y1="21" x2="16.65" y2="16.65"/>
        </svg>
      </button>
    </div>
    <div v-if="visible && query.trim()" class="search-results">
      <div v-if="results.length === 0" class="no-results">没有找到相关文章</div>
      <template v-else>
        <RouterLink
          v-for="post in previewResults"
          :key="post.id"
          :to="`/posts/${post.id}`"
          class="result-item"
          @click="close"
        >
          <span class="result-title">{{ post.title }}</span>
          <span class="result-date">{{ post.date }}</span>
        </RouterLink>
        <RouterLink
          v-if="results.length > 5"
          :to="{ name: 'search', query: { q: query } }"
          class="view-all"
          @click="close"
        >
          查看全部 {{ results.length }} 条结果
        </RouterLink>
      </template>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, nextTick, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { getPublishedArticles } from '@/services/articleApiService'

const router = useRouter()

const visible = ref(false)
const query = ref('')
const inputRef = ref(null)
const wrapperRef = ref(null)
const allPosts = ref([])

async function fetchAllPosts() {
  try {
    const res = await getPublishedArticles(1, 100)
    const data = res.data
    allPosts.value = (data.articles || []).map(a => ({
      id: a.id,
      title: a.title,
      excerpt: a.summary,
      date: a.created_at?.split('T')[0] || '',
    }))
  } catch {
    allPosts.value = []
  }
}

const results = computed(() => {
  const q = query.value.trim().toLowerCase()
  if (!q) return []
  return allPosts.value.filter(
    p =>
      p.title.toLowerCase().includes(q) ||
      p.excerpt.toLowerCase().includes(q)
  )
})

const previewResults = computed(() => results.value.slice(0, 5))

function toggle() {
  if (visible.value) {
    close()
  } else {
    visible.value = true
    if (allPosts.value.length === 0) fetchAllPosts()
    nextTick(() => inputRef.value?.focus())
  }
}

function close() {
  visible.value = false
  setTimeout(() => { query.value = '' }, 300)
}

function goSearch() {
  const q = query.value.trim()
  if (q) {
    router.push({ name: 'search', query: { q } })
    close()
  }
}

function onClickOutside(e) {
  if (wrapperRef.value && !wrapperRef.value.contains(e.target)) {
    close()
  }
}

onMounted(() => document.addEventListener('mousedown', onClickOutside))
onUnmounted(() => document.removeEventListener('mousedown', onClickOutside))
</script>

<style scoped>
.search-wrapper {
  position: relative;
}

.search-bar {
  display: flex;
  align-items: center;
  width: 34px;
  height: 34px;
  transition: width 0.3s ease;
  flex-shrink: 0;
  border: 1px solid transparent;
  border-radius: 8px;
}

.search-bar.open {
  width: 194px;
  border-color: var(--border);
  background: var(--bg);
}

.search-bar.open:focus-within {
  border-color: var(--link);
}

.input-slide {
  flex: 1;
  min-width: 0;
  overflow: hidden;
}

.search-input {
  width: 100%;
  height: 100%;
  padding: 0 10px;
  border: none;
  outline: none;
  font-size: 14px;
  background: transparent;
  color: var(--text);
}

.search-input::placeholder {
  color: var(--text-muted);
}

.search-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 34px;
  height: 34px;
  background: none;
  border: 1px solid var(--border);
  border-radius: 8px;
  cursor: pointer;
  color: var(--text-secondary);
  transition: border-color 0.2s ease, color 0.2s ease;
  padding: 0;
  flex-shrink: 0;
}

.search-btn:hover {
  border-color: var(--link);
  color: var(--link);
}

.search-bar.open .search-btn {
  border: none;
  border-radius: 0 8px 8px 0;
}

.search-results {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  margin-top: 4px;
  background: var(--bg-card);
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-md);
  z-index: 199;
  max-height: 320px;
  overflow-y: auto;
}

.no-results {
  padding: 20px;
  text-align: center;
  color: var(--text-muted);
  font-size: 14px;
}

.result-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  text-decoration: none;
  color: var(--text);
  transition: background 0.15s;
}

.result-item:hover {
  background: var(--bg-secondary);
}

.result-title {
  font-size: 14px;
  font-weight: 500;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.result-date {
  flex-shrink: 0;
  font-size: 13px;
  color: var(--text-muted);
}

.view-all {
  display: block;
  padding: 10px 16px;
  text-align: center;
  font-size: 13px;
  color: var(--link);
  text-decoration: none;
  border-top: 1px solid var(--border-light);
  transition: background 0.15s;
}

.view-all:hover {
  background: var(--bg-secondary);
}
</style>
