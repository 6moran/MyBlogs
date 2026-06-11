<template>
  <nav v-if="totalPages > 1" class="pagination" aria-label="分页导航">
    <button
      class="page-btn"
      :disabled="currentPage <= 1"
      @click="onPageChange(currentPage - 1)"
      aria-label="上一页"
    >
      &lsaquo;
    </button>

    <template v-for="page in displayPages" :key="page">
      <span v-if="page === '...'" class="page-ellipsis">...</span>
      <button
        v-else
        class="page-btn"
        :class="{ active: page === currentPage }"
        @click="onPageChange(page)"
        :aria-label="`第 ${page} 页`"
        :aria-current="page === currentPage ? 'page' : undefined"
      >
        {{ page }}
      </button>
    </template>

    <button
      class="page-btn"
      :disabled="currentPage >= totalPages"
      @click="onPageChange(currentPage + 1)"
      aria-label="下一页"
    >
      &rsaquo;
    </button>
  </nav>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  currentPage: {
    type: Number,
    required: true,
  },
  totalPages: {
    type: Number,
    required: true,
  },
  onPageChange: {
    type: Function,
    required: true,
  },
})

const displayPages = computed(() => {
  const total = props.totalPages
  const current = props.currentPage
  const pages = []

  if (total <= 7) {
    for (let i = 1; i <= total; i++) pages.push(i)
    return pages
  }

  // Always show first page
  pages.push(1)

  if (current > 3) {
    pages.push('...')
  }

  // Pages around current
  const start = Math.max(2, current - 1)
  const end = Math.min(total - 1, current + 1)
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }

  if (current < total - 2) {
    pages.push('...')
  }

  // Always show last page
  pages.push(total)

  return pages
})
</script>

<style scoped>
.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 4px;
  margin-top: 32px;
}

.page-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-width: 36px;
  height: 36px;
  padding: 0 8px;
  border: 1px solid var(--border);
  border-radius: 6px;
  background: transparent;
  color: var(--text);
  font-size: 14px;
  cursor: pointer;
  transition: background-color 0.2s ease, color 0.2s ease, border-color 0.2s ease;
}

.page-btn:hover:not(:disabled):not(.active) {
  background-color: var(--bg-secondary);
}

.page-btn.active {
  background-color: var(--link);
  color: #ffffff;
  border-color: var(--link);
}

.page-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.page-ellipsis {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-width: 36px;
  height: 36px;
  color: var(--text-secondary);
  font-size: 14px;
  user-select: none;
}
</style>
