<template>
  <section class="article-manage-page">
    <div class="scroll-container">
    <header class="page-header">
      <div>
        <p class="page-kicker">文章管理</p>
        <h2>管理博客文章</h2>
        <p class="page-desc">在这里查看文章列表，处理草稿、编辑内容，或发布新的博客文章。</p>
      </div>

      <button class="primary-btn" type="button" @click="openCreateModal">
        <span class="btn-icon">+</span>
        新建文章
      </button>
    </header>

    <section class="stats-row">
      <article
        v-for="card in overviewCards"
        :key="card.label"
        class="stat-card"
      >
        <span class="stat-label">{{ card.label }}</span>
        <strong>{{ card.value }}</strong>
      </article>
    </section>

    <section class="list-panel">
      <div class="toolbar">
        <label class="search-shell">
          <svg class="search-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="11" cy="11" r="8"></circle>
            <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
          </svg>
          <input
            v-model="searchKeyword"
            type="text"
            class="search-input"
            placeholder="搜索标题、分类、标签或内容"
          />
        </label>

        <div class="filter-row">
          <button
            v-for="filter in statusFilters"
            :key="filter.value"
            type="button"
            :class="['filter-chip', { active: statusFilter === filter.value }]"
            @click="statusFilter = filter.value"
          >
            {{ filter.label }} ({{ filter.count }})
          </button>
        </div>
      </div>

      <div class="article-list">
        <article
          v-for="article in filteredArticles"
          :key="article.id"
          class="article-card"
        >
          <div class="article-main">
            <div class="article-top">
              <div>
                <h3 class="article-title">{{ article.title }}</h3>
                <p class="article-excerpt">{{ article.excerpt }}</p>
              </div>

              <span :class="['status-pill', statusTone(article.status)]">
                {{ article.status }}
              </span>
            </div>

            <div class="meta-row">
              <span class="meta-pill category-pill">{{ article.category }}</span>
              <span class="meta-pill">创建于 {{ formatDate(article.createdAt) }}</span>
              <span class="meta-pill">更新于 {{ formatDate(article.updatedAt || article.createdAt) }}</span>
            </div>

            <div class="tag-row">
              <span
                v-for="tag in article.tags"
                :key="`${article.id}-${tag}`"
                class="tag-pill"
              >
                {{ tag }}
              </span>
            </div>
          </div>

          <div class="article-actions">
            <button class="action-btn edit-btn" type="button" @click="editArticle(article)">编辑</button>
            <button class="action-btn delete-btn" type="button" @click="deleteArticle(article.id)">删除</button>
          </div>
        </article>

        <div v-if="filteredArticles.length === 0" class="empty-state">
          <h3>暂无匹配文章</h3>
          <p>当前筛选条件下没有找到文章，可以调整关键词或直接新建一篇文章。</p>
          <button class="primary-btn" type="button" @click="openCreateModal">新建文章</button>
        </div>
      </div>
    </section>

    </div>

    <div v-if="showCreateModal" class="modal-overlay" @wheel.prevent>
      <div :class="['modal-content', { 'modal-content-visible': modalVisible }]" @wheel.stop>
        <div class="modal-header">
          <h2>{{ isEditing ? '编辑文章' : '创建文章' }}</h2>

          <button
            class="close-btn"
            type="button"
            aria-label="关闭弹窗"
            @click="closeCreateModal"
          >
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <line x1="18" y1="6" x2="6" y2="18"></line>
              <line x1="6" y1="6" x2="18" y2="18"></line>
            </svg>
          </button>
        </div>

        <div class="modal-body">
          <div class="form-group">
            <label class="field-label required" for="modal-title">文章标题</label>
            <input
              id="modal-title"
              v-model.trim="modalForm.title"
              class="title-input"
              type="text"
              placeholder="请输入文章标题"
            />
          </div>

          <div class="form-row">
            <div class="form-group">
              <label class="field-label" for="modal-category">分类</label>
              <select
                id="modal-category"
                v-model="modalForm.category"
                class="select-input"
                disabled
              >
                <option value="">等待后端数据</option>
              </select>
            </div>

            <div class="form-group">
              <label class="field-label" for="modal-tags">标签</label>
              <input
                id="modal-tags"
                v-model.trim="modalForm.tags"
                class="tags-input"
                type="text"
                placeholder="等待后端数据"
                disabled
              />
            </div>
          </div>

          <div class="form-group">
            <label class="field-label">封面图</label>
            <div
              :class="['cover-upload', { 'cover-upload-dragover': coverDragover, 'cover-upload-loading': coverUploading }]"
              @click="triggerCoverUpload"
              @dragover.prevent="coverDragover = true"
              @dragleave.prevent="coverDragover = false"
              @drop.prevent="handleCoverDrop"
            >
              <input
                ref="coverInput"
                type="file"
                accept="image/jpeg,image/png,image/gif,image/webp,image/svg+xml"
                style="display: none"
                @change="handleCoverChange"
              />
              <template v-if="coverUploading">
                <div class="cover-loading-spinner"></div>
                <span class="cover-upload-text">正在上传...</span>
              </template>
              <template v-else-if="modalForm.coverImage">
                <img :src="modalForm.coverImage" class="cover-preview" alt="封面预览" />
                <button
                  class="cover-remove"
                  type="button"
                  @click.stop="removeCover"
                >
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>
                </button>
              </template>
              <template v-else>
                <svg class="cover-upload-icon" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"/><circle cx="8.5" cy="8.5" r="1.5"/><polyline points="21 15 16 10 5 21"/></svg>
                <span class="cover-upload-text">点击或拖拽上传封面图</span>
              </template>
            </div>
          </div>

          <div class="form-group">
            <label class="field-label" for="modal-excerpt">摘要</label>
            <textarea
              id="modal-excerpt"
              v-model.trim="modalForm.excerpt"
              class="excerpt-input"
              rows="3"
              placeholder="留空将自动从正文截取第一句话作为摘要"
            />
          </div>

          <div class="form-group">
            <label class="field-label">正文内容</label>
            <ArticleEditor v-model="modalForm.content" />
          </div>
        </div>

        <div class="modal-footer">
          <button class="secondary-btn" type="button" @click="closeCreateModal">
            取消
          </button>
          <button
            class="primary-btn"
            type="button"
            :disabled="isSubmitting"
            @click="saveArticle"
          >
            <span v-if="isSubmitting" class="loading-spinner"></span>
            {{ isSubmitting ? '保存中...' : '保存' }}
          </button>
          <button
            class="primary-btn publish-btn"
            type="button"
            :disabled="isSubmitting"
            @click="publishArticle"
          >
            发布
          </button>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { computed, reactive, ref, watch } from 'vue'
import ArticleEditor from '../components/ArticleEditor.vue'
import { uploadImage } from '../services/articleService'

const defaultContent = `# 新文章标题

在这里开始编写你的博客内容。
`

const toExcerpt = (content) => {
  const plainText = (content || '')
    .replace(/```[\s\S]*?```/g, ' ')
    .replace(/`[^`]*`/g, ' ')
    .replace(/!\[[^\]]*]\([^)]*\)/g, ' ')
    .replace(/\[[^\]]*]\([^)]*\)/g, ' ')
    .replace(/[#>*_\-\n\r]/g, ' ')
    .replace(/\s+/g, ' ')
    .trim()

  if (!plainText) return '暂无摘要'

  const sentenceEnd = plainText.match(/[。！？.!?\n]/)
  if (sentenceEnd) {
    return plainText.slice(0, sentenceEnd.index + 1).trim()
  }

  return plainText.length > 100 ? `${plainText.slice(0, 100)}...` : plainText
}

const articles = ref([
  {
    id: 1,
    title: 'Vue 3 项目结构整理',
    content: `# Vue 3 项目结构整理

记录一下最近在整理 Vue 3 项目结构时的一些做法和想法。
`,
    category: '技术',
    status: '已发布',
    tags: ['Vue', '前端'],
    createdAt: '2026-04-20T10:00:00Z',
    updatedAt: '2026-04-21T09:30:00Z',
    excerpt: '记录一下最近在整理 Vue 3 项目结构时的一些做法和想法。'
  },
  {
    id: 2,
    title: '四月阅读记录',
    content: `# 四月阅读记录

整理这个月读过的几本书和一些简单的感受。
`,
    category: '生活',
    status: '草稿',
    tags: ['阅读', '记录'],
    createdAt: '2026-04-19T15:30:00Z',
    updatedAt: '2026-04-22T06:15:00Z',
    excerpt: '整理这个月读过的几本书和一些简单的感受。'
  },
  {
    id: 3,
    title: 'Go 学习笔记',
    content: `# Go 学习笔记

把最近练习 Go 时遇到的几个知识点简单记下来。
`,
    category: '学习记录',
    status: '草稿',
    tags: ['Go', '后端'],
    createdAt: '2026-04-18T09:15:00Z',
    updatedAt: '2026-04-18T09:15:00Z',
    excerpt: '把最近练习 Go 时遇到的几个知识点简单记下来。'
  }
])

const searchKeyword = ref('')
const statusFilter = ref('all')
const showCreateModal = ref(false)
const modalVisible = ref(false)
const isEditing = ref(false)
const currentArticleId = ref(null)
const isSubmitting = ref(false)
const coverInput = ref(null)
const coverDragover = ref(false)
const coverUploading = ref(false)

watch(showCreateModal, (visible) => {
  if (visible) {
    const bodyScrollbarWidth = window.innerWidth - document.documentElement.clientWidth
    document.body.style.overflow = 'hidden'
    if (bodyScrollbarWidth > 0) {
      document.body.style.paddingRight = `${bodyScrollbarWidth}px`
    }
  } else {
    document.body.style.overflow = ''
    document.body.style.paddingRight = ''
  }
})

const modalForm = reactive({
  title: '',
  content: defaultContent,
  category: '',
  status: '草稿',
  tags: '',
  coverImage: '',
  excerpt: ''
})

const statusFilters = computed(() => [
  {
    label: '全部',
    value: 'all',
    count: articles.value.length
  },
  {
    label: '草稿',
    value: '草稿',
    count: articles.value.filter(article => article.status === '草稿').length
  },
  {
    label: '已发布',
    value: '已发布',
    count: articles.value.filter(article => article.status === '已发布').length
  }
])

const overviewCards = computed(() => [
  {
    label: '文章总数',
    value: `${articles.value.length} 篇`
  },
  {
    label: '草稿数',
    value: `${articles.value.filter(article => article.status === '草稿').length} 篇`
  },
  {
    label: '已发布',
    value: `${articles.value.filter(article => article.status === '已发布').length} 篇`
  }
])

const filteredArticles = computed(() => {
  const keyword = searchKeyword.value.trim().toLowerCase()

  return [...articles.value]
    .filter((article) => {
      if (statusFilter.value !== 'all' && article.status !== statusFilter.value) {
        return false
      }

      if (!keyword) {
        return true
      }

      const searchableText = [
        article.title,
        article.content,
        article.category,
        article.status,
        ...article.tags
      ]
        .join(' ')
        .toLowerCase()

      return searchableText.includes(keyword)
    })
    .sort((a, b) => new Date(b.updatedAt || b.createdAt) - new Date(a.updatedAt || a.createdAt))
})

const formatDate = (dateString) =>
  new Date(dateString).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })

const triggerCoverUpload = () => {
  if (coverUploading.value) return
  coverInput.value?.click()
}

const handleCoverChange = (e) => {
  const file = e.target.files?.[0]
  if (!file) return
  uploadCoverFile(file)
}

const handleCoverDrop = (e) => {
  coverDragover.value = false
  const file = e.dataTransfer?.files?.[0]
  if (!file || !file.type.startsWith('image/')) return
  uploadCoverFile(file)
}

const uploadCoverFile = async (file) => {
  coverUploading.value = true
  try {
    const url = await uploadImage(file)
    modalForm.coverImage = url
  } catch (error) {
    window.alert(`封面图上传失败：${error.message}`)
  } finally {
    coverUploading.value = false
    if (coverInput.value) {
      coverInput.value.value = ''
    }
  }
}

const removeCover = () => {
  modalForm.coverImage = ''
  if (coverInput.value) {
    coverInput.value.value = ''
  }
}

const resetModalForm = () => {
  modalForm.title = ''
  modalForm.content = defaultContent
  modalForm.category = ''
  modalForm.status = '草稿'
  modalForm.tags = ''
  modalForm.coverImage = ''
  modalForm.excerpt = ''
  if (coverInput.value) {
    coverInput.value.value = ''
  }
}

const openCreateModal = () => {
  isEditing.value = false
  currentArticleId.value = null
  resetModalForm()
  showCreateModal.value = true

  requestAnimationFrame(() => {
    modalVisible.value = true
  })
}

const closeCreateModal = () => {
  modalVisible.value = false

  setTimeout(() => {
    showCreateModal.value = false
  }, 220)
}

const editArticle = (article) => {
  isEditing.value = true
  currentArticleId.value = article.id
  modalForm.title = article.title
  modalForm.content = article.content
  modalForm.category = article.category
  modalForm.status = article.status
  modalForm.tags = article.tags.join(', ')
  modalForm.coverImage = article.coverImage || ''
  modalForm.excerpt = article.excerpt || ''
  showCreateModal.value = true

  requestAnimationFrame(() => {
    modalVisible.value = true
  })
}

const deleteArticle = (id) => {
  if (!window.confirm('确定要删除这篇文章吗？删除后将无法恢复。')) {
    return
  }

  articles.value = articles.value.filter(article => article.id !== id)
}

const saveArticle = async () => {
  if (!modalForm.title) {
    window.alert('请输入文章标题。')
    return
  }

  if (!modalForm.content.trim()) {
    window.alert('正文内容不能为空。')
    return
  }

  isSubmitting.value = true

  try {
    await new Promise(resolve => setTimeout(resolve, 600))

    const tagsArray = modalForm.tags
      .split(',')
      .map(tag => tag.trim())
      .filter(Boolean)

    const payload = {
      title: modalForm.title,
      content: modalForm.content,
      category: modalForm.category || '未分类',
      status: modalForm.status,
      tags: tagsArray,
      coverImage: modalForm.coverImage,
      excerpt: modalForm.excerpt || toExcerpt(modalForm.content),
      updatedAt: new Date().toISOString()
    }

    if (isEditing.value) {
      const index = articles.value.findIndex(article => article.id === currentArticleId.value)

      if (index !== -1) {
        articles.value[index] = {
          ...articles.value[index],
          ...payload
        }
      }
    } else {
      const nextId = articles.value.length
        ? Math.max(...articles.value.map(article => article.id)) + 1
        : 1

      articles.value.unshift({
        id: nextId,
        createdAt: new Date().toISOString(),
        ...payload
      })
    }

    closeCreateModal()
  } catch (error) {
    window.alert('保存失败，请稍后再试。')
  } finally {
    isSubmitting.value = false
  }
}

const publishArticle = async () => {
  if (!modalForm.title) {
    window.alert('请输入文章标题。')
    return
  }

  if (!modalForm.content.trim()) {
    window.alert('正文内容不能为空。')
    return
  }

  isSubmitting.value = true

  try {
    await new Promise(resolve => setTimeout(resolve, 600))

    const tagsArray = modalForm.tags
      .split(',')
      .map(tag => tag.trim())
      .filter(Boolean)

    const payload = {
      title: modalForm.title,
      content: modalForm.content,
      category: modalForm.category || '未分类',
      status: '已发布',
      tags: tagsArray,
      coverImage: modalForm.coverImage,
      excerpt: modalForm.excerpt || toExcerpt(modalForm.content),
      updatedAt: new Date().toISOString()
    }

    if (isEditing.value) {
      const index = articles.value.findIndex(article => article.id === currentArticleId.value)

      if (index !== -1) {
        articles.value[index] = {
          ...articles.value[index],
          ...payload
        }
      }
    } else {
      const nextId = articles.value.length
        ? Math.max(...articles.value.map(article => article.id)) + 1
        : 1

      articles.value.unshift({
        id: nextId,
        createdAt: new Date().toISOString(),
        ...payload
      })
    }

    closeCreateModal()
  } catch (error) {
    window.alert('发布失败，请稍后再试。')
  } finally {
    isSubmitting.value = false
  }
}

const statusTone = (status) => (status === '已发布' ? 'is-published' : 'is-draft')
</script>

<style scoped>
.article-manage-page {
  height: 100%;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.scroll-container {
  flex: 1;
  min-height: 0;
  overflow-y: auto;
  scrollbar-gutter: stable;
  -webkit-overflow-scrolling: touch;
  overscroll-behavior-y: contain;
  display: grid;
  gap: 18px;
  align-content: start;
  padding: 0 20px 30px 0;
}

.page-header,
.list-panel {
  padding: 24px;
  border: 1px solid var(--line);
  border-radius: var(--radius-xl);
  background: rgba(255, 255, 255, 0.94);
  box-shadow: var(--shadow);
}

.page-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 18px;
}

.page-kicker,
.modal-kicker,
.stat-label {
  margin: 0 0 8px;
  color: var(--text-subtle);
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.field-label {
  display: block;
  margin: 0 0 8px;
  color: var(--text-main);
  font-size: 14px;
  font-weight: 600;
  letter-spacing: 0.01em;
}

.field-label.required::before {
  content: '* ';
  color: var(--danger);
}

.publish-btn {
  background: linear-gradient(135deg, #15803d 0%, #166534 100%);
}

.page-header h2,
.modal-header h2 {
  margin: 0;
  font-size: 30px;
  line-height: 1.2;
}

.page-desc {
  max-width: 620px;
  margin: 12px 0 0;
  color: var(--text-subtle);
  line-height: 1.7;
}

.stats-row {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 16px;
}

.stat-card {
  padding: 20px 22px;
  border: 1px solid var(--line);
  border-radius: var(--radius-lg);
  background: rgba(255, 255, 255, 0.9);
  box-shadow: 0 10px 30px rgba(15, 23, 42, 0.05);
}

.stat-card strong {
  display: block;
  font-size: 28px;
  line-height: 1.1;
}

.toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  margin-bottom: 20px;
}

.search-shell {
  position: relative;
  display: flex;
  align-items: center;
  width: min(360px, 100%);
}

.search-icon {
  position: absolute;
  left: 14px;
  color: var(--text-subtle);
}

.search-input,
.title-input,
.select-input,
.tags-input {
  width: 100%;
  border: 1px solid var(--line);
  color: var(--text-main);
  outline: none;
  transition: 0.2s ease;
}

.search-input {
  padding: 12px 14px 12px 42px;
  border-radius: 14px;
  background: var(--bg-soft);
}

.search-input:focus,
.title-input:focus,
.select-input:focus,
.tags-input:focus {
  border-color: rgba(37, 99, 235, 0.4);
  box-shadow: 0 0 0 4px rgba(37, 99, 235, 0.1);
}

.filter-row {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.filter-chip {
  padding: 10px 14px;
  min-width: 96px;
  text-align: center;
  border: 1px solid var(--line);
  border-radius: 999px;
  background: #fff;
  color: var(--text-subtle);
  cursor: pointer;
  transition: 0.2s ease;
}

.filter-chip:hover,
.filter-chip.active {
  border-color: rgba(37, 99, 235, 0.28);
  background: var(--accent-soft);
  color: var(--accent);
}

.article-list {
  display: grid;
  gap: 14px;
}

.article-card {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 18px;
  padding: 20px;
  border: 1px solid var(--line);
  border-radius: var(--radius-lg);
  background: #fff;
}

.article-main {
  min-width: 0;
  flex: 1;
}

.article-top {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 14px;
  margin-bottom: 14px;
}

.article-title {
  margin: 0 0 10px;
  font-size: 22px;
  line-height: 1.3;
}

.article-excerpt {
  margin: 0;
  color: var(--text-subtle);
  line-height: 1.7;
}

.status-pill,
.meta-pill,
.tag-pill {
  display: inline-flex;
  align-items: center;
  padding: 7px 12px;
  border-radius: 999px;
  font-size: 13px;
}

.status-pill {
  font-weight: 700;
  white-space: nowrap;
}

.status-pill.is-published {
  background: rgba(21, 128, 61, 0.1);
  color: var(--success);
}

.status-pill.is-draft {
  background: rgba(245, 158, 11, 0.12);
  color: #b45309;
}

.meta-row,
.tag-row {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.meta-row {
  margin-bottom: 12px;
}

.meta-pill {
  background: var(--bg-soft);
  color: var(--text-subtle);
}

.category-pill {
  color: var(--accent);
  background: var(--accent-soft);
}

.tag-pill {
  background: #f8fafc;
  color: #4b5563;
}

.article-actions {
  display: flex;
  gap: 10px;
}

.primary-btn,
.secondary-btn,
.action-btn,
.close-btn {
  border: 0;
  cursor: pointer;
  transition: 0.2s ease;
}

.primary-btn,
.secondary-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 12px 18px;
  border-radius: 14px;
  font-weight: 700;
}

.primary-btn {
  background: linear-gradient(135deg, #2563eb 0%, #1d4ed8 100%);
  color: #fff;
}

.primary-btn:hover,
.secondary-btn:hover,
.action-btn:hover,
.close-btn:hover {
  transform: translateY(-1px);
}

.secondary-btn {
  border: 1px solid var(--line);
  background: #fff;
  color: var(--text-main);
}

.action-btn {
  padding: 10px 14px;
  border-radius: 12px;
  background: var(--bg-soft);
  color: var(--text-main);
  font-weight: 600;
}

.edit-btn:hover {
  background: rgba(37, 99, 235, 0.12);
  color: var(--accent);
}

.delete-btn:hover {
  background: rgba(220, 38, 38, 0.1);
  color: var(--danger);
}

.btn-icon {
  font-size: 20px;
  line-height: 1;
}

.empty-state {
  padding: 42px 20px;
  border: 1px dashed var(--line-strong);
  border-radius: var(--radius-lg);
  text-align: center;
  background: #fff;
}

.empty-state h3 {
  margin: 0 0 10px;
  font-size: 24px;
}

.empty-state p {
  max-width: 420px;
  margin: 0 auto 20px;
  color: var(--text-subtle);
  line-height: 1.7;
}

.modal-overlay {
  position: fixed;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  background: rgba(15, 23, 42, 0.36);
  z-index: 1000;
}

.modal-content {
  width: min(1440px, 96vw);
  min-width: 860px;
  min-height: 560px;
  max-height: 94vh;
  border-radius: var(--radius-xl);
  background: #fff;
  box-shadow: 0 28px 80px rgba(15, 23, 42, 0.2);
  overflow: hidden;
  transform: translateY(12px);
  opacity: 0;
  transition: 0.22s ease;
  display: flex;
  flex-direction: column;
}

.modal-content-visible {
  transform: translateY(0);
  opacity: 1;
}

.modal-header,
.modal-body,
.modal-footer {
  padding: 24px 28px;
}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid var(--line);
}

.close-btn {
  display: grid;
  place-items: center;
  width: 40px;
  height: 40px;
  border-radius: 12px;
  background: var(--bg-soft);
  color: var(--text-subtle);
}

.modal-body {
  flex: 1;
  min-height: 0;
  overflow: auto;
}

.form-group {
  margin-bottom: 22px;
}

.form-row {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 16px;
  margin-bottom: 22px;
}

.form-row .form-group {
  margin-bottom: 0;
}

.title-input {
  padding: 16px 18px;
  border-radius: 14px;
  background: #fff;
  font-size: 18px;
}

.select-input,
.tags-input {
  padding: 13px 15px;
  border-radius: 14px;
  background: #fff;
}

.cover-upload {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 10px;
  width: 320px;
  min-height: 140px;
  border: 2px dashed var(--line);
  border-radius: 16px;
  background: var(--bg-soft);
  cursor: pointer;
  transition: 0.2s ease;
  overflow: hidden;
}

.cover-upload:hover {
  border-color: rgba(37, 99, 235, 0.35);
  background: rgba(37, 99, 235, 0.03);
}

.cover-upload-dragover {
  border-color: var(--accent);
  background: rgba(37, 99, 235, 0.06);
}

.cover-upload-loading {
  pointer-events: none;
  opacity: 0.7;
}

.cover-loading-spinner {
  width: 28px;
  height: 28px;
  border: 3px solid var(--line);
  border-top-color: var(--accent);
  border-radius: 50%;
  animation: spin 0.9s linear infinite;
}

.cover-upload-icon {
  color: var(--text-subtle);
}

.cover-upload-text {
  color: var(--text-subtle);
  font-size: 14px;
}

.cover-preview {
  width: 100%;
  height: 160px;
  object-fit: cover;
  border-radius: 14px;
}

.cover-remove {
  position: absolute;
  top: 10px;
  right: 10px;
  display: grid;
  place-items: center;
  width: 30px;
  height: 30px;
  border: 0;
  border-radius: 10px;
  background: rgba(0, 0, 0, 0.5);
  color: #fff;
  cursor: pointer;
  transition: 0.15s ease;
}

.cover-remove:hover {
  background: rgba(220, 38, 38, 0.8);
}

.excerpt-input {
  width: 100%;
  padding: 13px 15px;
  border: 1px solid var(--line);
  border-radius: 14px;
  background: #fff;
  color: var(--text-main);
  font-size: 14px;
  line-height: 1.7;
  resize: vertical;
  outline: none;
  transition: 0.2s ease;
  font-family: inherit;
}

.excerpt-input:focus {
  border-color: rgba(37, 99, 235, 0.4);
  box-shadow: 0 0 0 4px rgba(37, 99, 235, 0.1);
}

.modal-footer {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 12px;
  border-top: 1px solid var(--line);
}

.loading-spinner {
  width: 16px;
  height: 16px;
  border: 2px solid rgba(255, 255, 255, 0.4);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin 0.9s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

@media (max-width: 1024px) {
  .toolbar,
  .page-header,
  .article-card,
  .article-top {
    flex-direction: column;
    align-items: flex-start;
  }

  .stats-row,
  .form-row {
    grid-template-columns: 1fr;
  }

  .search-shell {
    width: 100%;
  }
}

@media (max-width: 768px) {
  .page-header,
  .list-panel,
  .stat-card {
    padding: 18px;
  }

  .page-header h2,
  .modal-header h2 {
    font-size: 24px;
  }

  .modal-overlay {
    padding: 8px;
  }

  .modal-content {
    min-width: unset;
    min-height: unset;
    width: 100%;
    max-height: 100vh;
    border-radius: var(--radius-lg);
  }

  .modal-header,
  .modal-body,
  .modal-footer {
    padding: 18px;
  }

  .modal-footer {
    flex-direction: column-reverse;
    align-items: stretch;
  }
}
</style>
