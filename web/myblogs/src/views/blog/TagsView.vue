<template>
  <div class="tags-page">
    <h1 class="page-title">标签</h1>
    <div v-if="loading" class="loading">加载中...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <div v-else class="tag-cloud">
      <RouterLink
        v-for="tag in tagsWithSize"
        :key="tag.id"
        :to="`/tags/${tag.id}`"
        class="tag-item"
        :style="{ fontSize: tag.fontSize + 'px', opacity: tag.opacity }"
      >
        {{ tag.name }}
      </RouterLink>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { getTags } from '@/services/tagService'

const tags = ref([])
const loading = ref(true)
const error = ref('')

const tagsWithSize = computed(() => {
  if (tags.value.length === 0) return []
  const maxCount = Math.max(...tags.value.map(t => t.count))
  const minCount = Math.min(...tags.value.map(t => t.count))
  return tags.value
    .map(tag => {
      const ratio = maxCount === minCount ? 0.5 : (tag.count - minCount) / (maxCount - minCount)
      return {
        ...tag,
        fontSize: 14 + Math.round(ratio * 14),
        opacity: 0.55 + ratio * 0.45,
      }
    })
    .sort((a, b) => b.count - a.count)
})

onMounted(async () => {
  try {
    const res = await getTags()
    tags.value = res.data || []
  } catch (e) {
    error.value = e.message || '加载失败'
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.tags-page {
  padding: 0;
}

.page-title {
  font-size: 24px;
  font-weight: 700;
  margin: 0 0 24px;
  color: var(--text);
}

.tag-cloud {
  display: flex;
  flex-wrap: wrap;
  gap: 12px 14px;
}

.tag-item {
  display: inline-block;
  padding: 4px 14px;
  color: var(--text-secondary);
  text-decoration: none;
  font-weight: 600;
  white-space: nowrap;
  line-height: 1.4;
  border-radius: var(--radius);
  border: 1px solid var(--border-light);
  transition: color 0.2s ease, border-color 0.2s ease;
}

.tag-item:visited {
  color: var(--text-secondary);
}

.tag-item:hover {
  color: var(--link);
  border-color: var(--link);
}

.loading, .error {
  text-align: center;
  padding: 40px 0;
  color: var(--text-secondary);
}
</style>
