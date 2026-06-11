<template>
  <section class="admin-page">
    <div class="admin-scroll">
      <header class="admin-header">
        <div>
          <span class="admin-label">文章管理</span>
          <h2>管理博客文章</h2>
          <p class="admin-desc">查看文章列表，处理草稿、编辑内容，或发布新的博客文章。</p>
        </div>
        <button class="admin-btn admin-btn-primary" type="button" @click="openCreateModal">
          <span class="admin-btn-icon">+</span>
          新建文章
        </button>
      </header>

      <section class="admin-stats">
        <article v-for="card in overviewCards" :key="card.label" class="admin-stat">
          <span class="admin-stat-label">{{ card.label }}</span>
          <strong class="admin-stat-value">{{ card.value }}</strong>
        </article>
      </section>

      <section class="admin-panel">
        <div class="admin-toolbar">
          <label class="admin-search">
            <svg class="admin-search-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/></svg>
            <input
              v-model="searchKeyword"
              type="text"
              class="admin-search-input"
              placeholder="搜索标题、标签或内容"
            />
          </label>
          <div class="admin-filters">
            <button
              v-for="filter in statusFilters"
              :key="filter.value"
              type="button"
              :class="['admin-filter-btn', { active: statusFilter === filter.value }]"
              @click="statusFilter = filter.value"
            >
              {{ filter.label }} ({{ filter.count }})
            </button>
          </div>
        </div>

        <div class="admin-list">
          <article
            v-for="article in filteredArticles"
            :key="article.id"
            class="admin-list-item"
          >
            <div class="admin-list-body">
              <div style="display: flex; align-items: center; gap: 10px; margin-bottom: 4px;">
                <h3 class="admin-list-title">{{ article.title }}</h3>
                <span :class="['admin-status', statusTone(article.status)]">{{ article.status }}</span>
              </div>
              <p class="admin-list-excerpt">{{ article.excerpt }}</p>
              <div class="admin-list-meta">
                <span class="admin-meta-tag">{{ formatDate(article.createdAt) }}</span>
                <span v-for="tag in article.tags" :key="`${article.id}-${tag}`" class="admin-meta-tag">{{ tag }}</span>
              </div>
            </div>
            <div class="admin-list-actions">
              <button class="admin-btn admin-btn-ghost" type="button" @click="editArticle(article)">编辑</button>
              <button class="admin-btn admin-btn-ghost" type="button" style="color: var(--danger);" @click="deleteArticle(article.id)">删除</button>
            </div>
          </article>

          <div v-if="filteredArticles.length === 0" class="admin-empty">
            <h3>暂无匹配文章</h3>
            <p>当前筛选条件下没有找到文章，可以调整关键词或直接新建一篇文章。</p>
            <button class="admin-btn admin-btn-primary" type="button" @click="openCreateModal">新建文章</button>
          </div>
        </div>
      </section>
    </div>

    <!-- Modal -->
    <div v-if="showCreateModal" class="admin-modal-overlay" @wheel.prevent>
      <div :class="['admin-modal', { 'admin-modal-visible': modalVisible }]" @wheel.stop>
        <div class="admin-modal-header">
          <h2>{{ isEditing ? '编辑文章' : '创建文章' }}</h2>
          <button class="admin-modal-close" type="button" aria-label="关闭" @click="closeCreateModal">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>
          </button>
        </div>

        <div class="admin-modal-body">
          <div class="admin-field">
            <label class="admin-field-label required" for="modal-title">文章标题</label>
            <input id="modal-title" v-model.trim="modalForm.title" class="admin-input admin-input-title" type="text" placeholder="请输入文章标题" />
          </div>

          <div class="admin-form-row">
            <div class="admin-field">
              <label class="admin-field-label" for="modal-tags">标签</label>
              <input id="modal-tags" v-model.trim="modalForm.tags" class="admin-input" type="text" placeholder="等待后端数据" disabled />
            </div>
          </div>

          <div class="admin-field">
            <label class="admin-field-label">封面图</label>
            <div
              :class="['cover-upload', { 'cover-dragover': coverDragover, 'cover-loading': coverUploading }]"
              @click="triggerCoverUpload"
              @dragover.prevent="coverDragover = true"
              @dragleave.prevent="coverDragover = false"
              @drop.prevent="handleCoverDrop"
            >
              <input ref="coverInput" type="file" accept="image/*" style="display: none" @change="handleCoverChange" />
              <template v-if="coverUploading">
                <div class="cover-spinner"></div>
                <span class="cover-text">正在上传...</span>
              </template>
              <template v-else-if="modalForm.coverImage">
                <img :src="modalForm.coverImage" class="cover-preview" alt="封面预览" />
                <button class="cover-remove" type="button" @click.stop="removeCover">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>
                </button>
              </template>
              <template v-else>
                <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" style="color: var(--text-muted);"><rect x="3" y="3" width="18" height="18" rx="2"/><circle cx="8.5" cy="8.5" r="1.5"/><polyline points="21 15 16 10 5 21"/></svg>
                <span class="cover-text">点击或拖拽上传封面图</span>
              </template>
            </div>
          </div>

          <div class="admin-field">
            <label class="admin-field-label" for="modal-excerpt">摘要</label>
            <textarea id="modal-excerpt" v-model.trim="modalForm.excerpt" class="admin-textarea" rows="3" placeholder="留空将自动从正文截取第一句话作为摘要" />
          </div>

          <div class="admin-field">
            <label class="admin-field-label">正文内容</label>
            <ArticleEditor v-model="modalForm.content" />
          </div>
        </div>

        <div class="admin-modal-footer">
          <button class="admin-btn admin-btn-secondary" type="button" @click="closeCreateModal">取消</button>
          <button class="admin-btn admin-btn-primary" type="button" :disabled="isSubmitting" @click="saveArticle">
            <span v-if="isSubmitting" class="admin-spinner"></span>
            {{ isSubmitting ? '保存中...' : '保存' }}
          </button>
          <button class="admin-btn admin-btn-success" type="button" :disabled="isSubmitting" @click="publishArticle">发布</button>
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
  status: '草稿',
  tags: '',
  coverImage: '',
  excerpt: ''
})

const statusFilters = computed(() => [
  { label: '全部', value: 'all', count: articles.value.length },
  { label: '草稿', value: '草稿', count: articles.value.filter(a => a.status === '草稿').length },
  { label: '已发布', value: '已发布', count: articles.value.filter(a => a.status === '已发布').length }
])

const overviewCards = computed(() => [
  { label: '文章总数', value: `${articles.value.length}` },
  { label: '草稿数', value: `${articles.value.filter(a => a.status === '草稿').length}` },
  { label: '已发布', value: `${articles.value.filter(a => a.status === '已发布').length}` }
])

const filteredArticles = computed(() => {
  const keyword = searchKeyword.value.trim().toLowerCase()
  return [...articles.value]
    .filter((article) => {
      if (statusFilter.value !== 'all' && article.status !== statusFilter.value) return false
      if (!keyword) return true
      const text = [article.title, article.content, article.status, ...article.tags].join(' ').toLowerCase()
      return text.includes(keyword)
    })
    .sort((a, b) => new Date(b.updatedAt || b.createdAt) - new Date(a.updatedAt || a.createdAt))
})

const formatDate = (dateString) =>
  new Date(dateString).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' })

const triggerCoverUpload = () => { if (!coverUploading.value) coverInput.value?.click() }
const handleCoverChange = (e) => { const file = e.target.files?.[0]; if (file) uploadCoverFile(file) }
const handleCoverDrop = (e) => { coverDragover.value = false; const file = e.dataTransfer?.files?.[0]; if (file?.type.startsWith('image/')) uploadCoverFile(file) }

const uploadCoverFile = async (file) => {
  coverUploading.value = true
  try {
    modalForm.coverImage = await uploadImage(file)
  } catch (error) {
    window.alert(`封面图上传失败：${error.message}`)
  } finally {
    coverUploading.value = false
    if (coverInput.value) coverInput.value.value = ''
  }
}

const removeCover = () => { modalForm.coverImage = ''; if (coverInput.value) coverInput.value.value = '' }

const resetModalForm = () => {
  modalForm.title = ''
  modalForm.content = defaultContent
  modalForm.status = '草稿'
  modalForm.tags = ''
  modalForm.coverImage = ''
  modalForm.excerpt = ''
  if (coverInput.value) coverInput.value.value = ''
}

const openCreateModal = () => {
  isEditing.value = false
  currentArticleId.value = null
  resetModalForm()
  showCreateModal.value = true
  requestAnimationFrame(() => { modalVisible.value = true })
}

const closeCreateModal = () => {
  modalVisible.value = false
  setTimeout(() => { showCreateModal.value = false }, 250)
}

const editArticle = (article) => {
  isEditing.value = true
  currentArticleId.value = article.id
  modalForm.title = article.title
  modalForm.content = article.content
  modalForm.status = article.status
  modalForm.tags = article.tags.join(', ')
  modalForm.coverImage = article.coverImage || ''
  modalForm.excerpt = article.excerpt || ''
  showCreateModal.value = true
  requestAnimationFrame(() => { modalVisible.value = true })
}

const deleteArticle = (id) => {
  if (!window.confirm('确定要删除这篇文章吗？删除后将无法恢复。')) return
  articles.value = articles.value.filter(a => a.id !== id)
}

const buildPayload = () => ({
  title: modalForm.title,
  content: modalForm.content,
  tags: modalForm.tags.split(',').map(t => t.trim()).filter(Boolean),
  coverImage: modalForm.coverImage,
  excerpt: modalForm.excerpt || toExcerpt(modalForm.content),
  updatedAt: new Date().toISOString()
})

const saveArticle = async () => {
  if (!modalForm.title) { window.alert('请输入文章标题。'); return }
  if (!modalForm.content.trim()) { window.alert('正文内容不能为空。'); return }
  isSubmitting.value = true
  try {
    await new Promise(r => setTimeout(r, 600))
    const payload = buildPayload()
    if (isEditing.value) {
      const idx = articles.value.findIndex(a => a.id === currentArticleId.value)
      if (idx !== -1) articles.value[idx] = { ...articles.value[idx], ...payload }
    } else {
      const nextId = articles.value.length ? Math.max(...articles.value.map(a => a.id)) + 1 : 1
      articles.value.unshift({ id: nextId, createdAt: new Date().toISOString(), status: '草稿', ...payload })
    }
    closeCreateModal()
  } catch { window.alert('保存失败，请稍后再试。') }
  finally { isSubmitting.value = false }
}

const publishArticle = async () => {
  if (!modalForm.title) { window.alert('请输入文章标题。'); return }
  if (!modalForm.content.trim()) { window.alert('正文内容不能为空。'); return }
  isSubmitting.value = true
  try {
    await new Promise(r => setTimeout(r, 600))
    const payload = { ...buildPayload(), status: '已发布' }
    if (isEditing.value) {
      const idx = articles.value.findIndex(a => a.id === currentArticleId.value)
      if (idx !== -1) articles.value[idx] = { ...articles.value[idx], ...payload }
    } else {
      const nextId = articles.value.length ? Math.max(...articles.value.map(a => a.id)) + 1 : 1
      articles.value.unshift({ id: nextId, createdAt: new Date().toISOString(), ...payload })
    }
    closeCreateModal()
  } catch { window.alert('发布失败，请稍后再试。') }
  finally { isSubmitting.value = false }
}

const statusTone = (status) => status === '已发布' ? 'admin-status-published' : 'admin-status-draft'
</script>

<style scoped>
/* ===== Cover Upload ===== */
.cover-upload {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 280px;
  min-height: 130px;
  border: 1.5px dashed var(--line-strong);
  border-radius: var(--radius-xl);
  background: var(--bg-soft);
  cursor: pointer;
  transition: all 0.2s;
  overflow: hidden;
}

.cover-upload:hover {
  border-color: rgba(192, 57, 43, 0.3);
  background: rgba(192, 57, 43, 0.02);
}

.cover-dragover {
  border-color: var(--accent);
  background: rgba(192, 57, 43, 0.04);
}

.cover-loading {
  pointer-events: none;
  opacity: 0.6;
}

.cover-spinner {
  width: 24px;
  height: 24px;
  border: 2.5px solid var(--line);
  border-top-color: var(--accent);
  border-radius: 50%;
  animation: admin-spin 0.8s linear infinite;
}

.cover-text {
  font-size: 13px;
  color: var(--text-subtle);
}

.cover-preview {
  width: 100%;
  height: 140px;
  object-fit: cover;
  border-radius: calc(var(--radius-xl) - 2px);
}

.cover-remove {
  position: absolute;
  top: 8px;
  right: 8px;
  display: grid;
  place-items: center;
  width: 28px;
  height: 28px;
  border: 0;
  border-radius: var(--radius);
  background: rgba(0, 0, 0, 0.45);
  color: #fff;
  cursor: pointer;
  transition: background 0.15s;
}

.cover-remove:hover {
  background: rgba(192, 57, 43, 0.85);
}
</style>
