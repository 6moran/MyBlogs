<template>
  <div :class="['markdown-workspace', { 'preview-hidden': !showPreview }]">
    <div class="markdown-panel">
      <div class="panel-head">
        <h3>Markdown 编辑</h3>
      </div>

      <div class="toolbar" @mousedown.prevent>
        <input
          ref="imageInput"
          type="file"
          accept="image/jpeg,image/png,image/gif,image/webp,image/svg+xml"
          style="display: none"
          @change="handleImageFileSelect"
        />

        <div class="heading-dropdown-wrap">
          <button
            class="toolbar-btn heading-trigger"
            type="button"
            title="段落与标题"
            @click="headingDropdownOpen = !headingDropdownOpen"
          >
            <span class="heading-label">H</span>
            <svg class="dropdown-arrow" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 12 15 18 9"/></svg>
          </button>
          <div v-if="headingDropdownOpen" class="heading-dropdown" @mousedown.prevent>
            <button
              v-for="opt in headingOptions"
              :key="opt.name"
              class="heading-option"
              type="button"
              @click="opt.action(); headingDropdownOpen = false"
            >
              <span :class="['option-label', opt.cls]">{{ opt.label }}</span>
            </button>
          </div>
        </div>

        <button
          v-for="item in toolbarButtons"
          :key="item.name"
          class="toolbar-btn"
          type="button"
          :title="item.tooltip"
          @click="item.action"
        >
          <svg class="btn-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" v-html="item.icon" />
        </button>

        <div class="toolbar-spacer"></div>

        <button
          class="toolbar-btn preview-toggle-icon"
          :class="{ 'active-preview': showPreview }"
          type="button"
          :title="showPreview ? '收起预览' : '展开预览'"
          @click="showPreview = !showPreview"
        >
          <svg class="btn-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z" />
            <circle v-if="showPreview" cx="12" cy="12" r="3" fill="currentColor" stroke="none" />
          </svg>
        </button>
      </div>

      <textarea
        ref="markdownInput"
        :value="modelValue"
        class="markdown-input"
        placeholder="在这里输入 Markdown 内容..."
        @input="handleInput"
        @keydown="handleKeydown"
        @paste="handlePaste"
        @drop="handleTextareaDrop"
        @dragover.prevent
      />
    </div>

    <div v-if="showPreview" class="preview-panel">
      <div class="panel-head">
        <h3>实时预览</h3>
      </div>
      <div class="preview-body" v-html="previewHtml"></div>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { marked } from 'marked'
import { uploadImage } from '../services/articleService'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue'])
const markdownInput = ref(null)
const imageInput = ref(null)
const showPreview = ref(true)
const headingDropdownOpen = ref(false)

const previewHtml = computed(() =>
  marked.parse(props.modelValue || '', {
    gfm: true,
    breaks: true
  })
)

const handleInput = (event) => {
  emit('update:modelValue', event.target.value)
}

const execReplace = (start, end, newText) => {
  const textarea = markdownInput.value
  if (!textarea) return false

  textarea.focus()
  textarea.setSelectionRange(start, end)

  if (document.execCommand('insertText', false, newText)) {
    return true
  }

  const value = textarea.value
  const result = value.substring(0, start) + newText + value.substring(end)
  textarea.value = result
  emit('update:modelValue', result)
  return false
}

const setCursorLater = (start, end) => {
  setTimeout(() => {
    const textarea = markdownInput.value
    if (!textarea) return
    textarea.focus()
    textarea.setSelectionRange(start, end ?? start)
  }, 0)
}

const wrapSelection = (before, after, defaultText) => {
  const textarea = markdownInput.value
  if (!textarea) return

  const start = textarea.selectionStart
  const end = textarea.selectionEnd
  const value = textarea.value
  const selectedText = value.substring(start, end)

  if (selectedText.length > 0 && selectedText.startsWith(before) && selectedText.endsWith(after) && selectedText.length >= before.length + after.length) {
    const unwrapped = selectedText.substring(before.length, selectedText.length - after.length)
    execReplace(start, end, unwrapped)
    setCursorLater(start, start + unwrapped.length)
    return
  }

  const beforeRegion = value.substring(Math.max(0, start - before.length), start)
  const afterRegion = value.substring(end, Math.min(value.length, end + after.length))

  if (beforeRegion === before && afterRegion === after) {
    const wrapStart = start - before.length
    const wrapEnd = end + after.length
    const inner = value.substring(start, end)
    execReplace(wrapStart, wrapEnd, inner)
    setCursorLater(wrapStart, wrapStart + inner.length)
    return
  }

  const beforePos = value.lastIndexOf(before, start - 1)
  const afterPos = value.indexOf(after, end)
  if (beforePos !== -1 && afterPos !== -1 && beforePos + before.length <= start && afterPos >= end) {
    const between = value.substring(beforePos + before.length, afterPos)
    const noNestedBefore = !between.includes(before)
    const noNestedAfter = !between.includes(after)
    if (noNestedBefore && noNestedAfter) {
      execReplace(beforePos, afterPos + after.length, between)
      setCursorLater(beforePos, beforePos + between.length)
      return
    }
  }

  const content = selectedText || defaultText
  const newText = before + content + after
  execReplace(start, end, newText)
  setCursorLater(start + before.length, start + before.length + content.length)
}

const insertBlock = (prefix, defaultText) => {
  const textarea = markdownInput.value
  if (!textarea) return

  const value = textarea.value
  const start = textarea.selectionStart
  const end = textarea.selectionEnd

  let lineStart = value.lastIndexOf('\n', start - 1) + 1
  let lineEnd = value.indexOf('\n', end)
  if (lineEnd === -1) lineEnd = value.length

  const currentLine = value.substring(lineStart, lineEnd)
  const existingPrefix = currentLine.match(/^(#{1,6}\s|>\s|- |\d+\.\s)/)

  let content
  if (start !== end) {
    content = value.substring(start, end)
  } else if (existingPrefix) {
    content = currentLine.substring(existingPrefix[1].length) || defaultText
  } else {
    content = currentLine.trim() || defaultText
  }

  const newPrefix = (existingPrefix && existingPrefix[1] === prefix) ? '' : prefix
  const newLine = newPrefix + content

  execReplace(lineStart, lineEnd, newLine)
  setCursorLater(lineStart + newPrefix.length, lineStart + newPrefix.length + content.length)
}

const headingOptions = [
  { name: 'paragraph', label: '正文', cls: 'cls-p', action: () => insertBlock('', '') },
  { name: 'h1', label: '标题 1', cls: 'cls-h1', action: () => insertBlock('# ', '一级标题') },
  { name: 'h2', label: '标题 2', cls: 'cls-h2', action: () => insertBlock('## ', '二级标题') },
  { name: 'h3', label: '标题 3', cls: 'cls-h3', action: () => insertBlock('### ', '三级标题') },
  { name: 'h4', label: '标题 4', cls: 'cls-h4', action: () => insertBlock('#### ', '四级标题') },
  { name: 'h5', label: '标题 5', cls: 'cls-h5', action: () => insertBlock('##### ', '五级标题') }
]

const toolbarButtons = [
  {
    name: 'bold',
    tooltip: '粗体 (Ctrl+B)',
    icon: '<path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/><path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/>',
    action: () => wrapSelection('**', '**', '粗体文字')
  },
  {
    name: 'italic',
    tooltip: '斜体 (Ctrl+I)',
    icon: '<line x1="19" y1="4" x2="10" y2="4"/><line x1="14" y1="20" x2="5" y2="20"/><line x1="15" y1="4" x2="9" y2="20"/>',
    action: () => wrapSelection('*', '*', '斜体文字')
  },
  {
    name: 'underline',
    tooltip: '下划线 (Ctrl+U)',
    icon: '<path d="M6 3v7a6 6 0 0 0 6 6 6 6 0 0 0 6-6V3"/><line x1="4" y1="21" x2="20" y2="21"/>',
    action: () => wrapSelection('<u>', '</u>', '下划线文字')
  },
  {
    name: 'strike',
    tooltip: '删除线 (Ctrl+Shift+S)',
    icon: '<path d="M16 4H9a3 3 0 0 0-3 3v0a3 3 0 0 0 3 3h6"/><line x1="4" y1="12" x2="20" y2="12"/><path d="M15 12a3 3 0 1 1 0 6H8"/>',
    action: () => wrapSelection('~~', '~~', '删除线文字')
  },
  {
    name: 'bulletList',
    tooltip: '无序列表 (Ctrl+Shift+8)',
    icon: '<line x1="8" y1="6" x2="21" y2="6"/><line x1="8" y1="12" x2="21" y2="12"/><line x1="8" y1="18" x2="21" y2="18"/><circle cx="4" cy="6" r="1" fill="currentColor"/><circle cx="4" cy="12" r="1" fill="currentColor"/><circle cx="4" cy="18" r="1" fill="currentColor"/>',
    action: () => insertBlock('- ', '列表项')
  },
  {
    name: 'orderedList',
    tooltip: '有序列表 (Ctrl+Shift+7)',
    icon: '<line x1="10" y1="6" x2="21" y2="6"/><line x1="10" y1="12" x2="21" y2="12"/><line x1="10" y1="18" x2="21" y2="18"/><text x="3" y="8" font-size="8" fill="currentColor" stroke="none" font-weight="bold">1</text><text x="3" y="14" font-size="8" fill="currentColor" stroke="none" font-weight="bold">2</text><text x="3" y="20" font-size="8" fill="currentColor" stroke="none" font-weight="bold">3</text>',
    action: () => insertBlock('1. ', '列表项')
  },
  {
    name: 'blockquote',
    tooltip: '引用 (Ctrl+Shift+B)',
    icon: '<path d="M3 21c3 0 7-1 7-5V5c0-1.25-.756-2.017-2-2-4 0-5 3-5 5s1 3 2 3c1 0 1-1 1-1"/><path d="M15 21c3 0 7-1 7-5V5c0-1.25-.757-2.017-2-2-4 0-5 3-5 5s1 3 2 3c1 0 1-1 1-1"/>',
    action: () => insertBlock('> ', '引用内容')
  },
  {
    name: 'codeBlock',
    tooltip: '代码块 (Ctrl+Alt+C)',
    icon: '<polyline points="16 18 22 12 16 6"/><polyline points="8 6 2 12 8 18"/>',
    action: () => wrapSelection('\n```\n', '\n```\n', '代码')
  },
  {
    name: 'link',
    tooltip: '链接 (Ctrl+K)',
    icon: '<path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/>',
    action: () => wrapSelection('[', '](https://example.com)', '链接文字')
  },
  {
    name: 'image',
    tooltip: '图片',
    icon: '<rect x="3" y="3" width="18" height="18" rx="2" ry="2"/><circle cx="8.5" cy="8.5" r="1.5"/><polyline points="21 15 16 10 5 21"/>',
    action: () => imageInput.value?.click()
  },
  {
    name: 'horizontalRule',
    tooltip: '分割线',
    icon: '<line x1="3" y1="12" x2="21" y2="12"/>',
    action: () => {
      const textarea = markdownInput.value
      if (!textarea) return
      const start = textarea.selectionStart
      const beforeText = textarea.value.substring(0, start)
      const needNewline = beforeText.length > 0 && !beforeText.endsWith('\n')
      const insert = (needNewline ? '\n' : '') + '\n---\n'
      execReplace(start, textarea.selectionEnd, insert)
      setCursorLater(start + insert.length, start + insert.length)
    }
  }
]

const insertUploadingPlaceholder = () => {
  const textarea = markdownInput.value
  if (!textarea) return

  const start = textarea.selectionStart
  const end = textarea.selectionEnd
  const placeholder = '![上传中...]()'

  execReplace(start, end, placeholder)
  setCursorLater(start + placeholder.length, start + placeholder.length)

  return { start, placeholderLen: placeholder.length }
}

const replaceUploadingPlaceholder = (start, placeholderLen, imageUrl, alt) => {
  const textarea = markdownInput.value
  if (!textarea) return

  const value = textarea.value
  const placeholder = value.substring(start, start + placeholderLen)

  if (placeholder !== '![上传中...]()') {
    const idx = value.lastIndexOf('![上传中...]()')
    if (idx !== -1) {
      const replacement = `![${alt}](${imageUrl})`
      execReplace(idx, idx + placeholder.length, replacement)
      setCursorLater(idx + replacement.length, idx + replacement.length)
    }
    return
  }

  const replacement = `![${alt}](${imageUrl})`
  execReplace(start, start + placeholderLen, replacement)
  setCursorLater(start + replacement.length, start + replacement.length)
}

const uploadAndInsertImage = async (file) => {
  const { start, placeholderLen } = insertUploadingPlaceholder()

  try {
    const url = await uploadImage(file)
    const alt = file.name.replace(/\.[^/.]+$/, '')
    replaceUploadingPlaceholder(start, placeholderLen, url, alt)
  } catch (error) {
    const textarea = markdownInput.value
    if (textarea) {
      const value = textarea.value
      const idx = value.indexOf('![上传中...]()')
      if (idx !== -1) {
        const failedText = `![上传失败]()`
        execReplace(idx, idx + '![上传中...]()'.length, failedText)
        setCursorLater(idx + failedText.length, idx + failedText.length)
      }
    }
    window.alert(`图片上传失败：${error.message}`)
  }
}

const handleImageFileSelect = (e) => {
  const file = e.target.files?.[0]
  if (!file) return
  uploadAndInsertImage(file)
  e.target.value = ''
}

const handlePaste = (e) => {
  const items = e.clipboardData?.items
  if (!items) return

  for (const item of items) {
    if (item.type.startsWith('image/')) {
      e.preventDefault()
      const file = item.getAsFile()
      if (file) {
        uploadAndInsertImage(file)
      }
      return
    }
  }
}

const handleTextareaDrop = (e) => {
  const file = e.dataTransfer?.files?.[0]
  if (file && file.type.startsWith('image/')) {
    e.preventDefault()
    uploadAndInsertImage(file)
  }
}

const handleKeydown = (e) => {
  const ctrl = e.ctrlKey || e.metaKey
  if (!ctrl) return

  const shift = e.shiftKey
  const alt = e.altKey
  const key = e.key.toLowerCase()
  const code = e.code

  if (key === 'b' && !shift && !alt) {
    e.preventDefault()
    wrapSelection('**', '**', '粗体文字')
  } else if (key === 'i' && !shift && !alt) {
    e.preventDefault()
    wrapSelection('*', '*', '斜体文字')
  } else if (key === 'u' && !shift && !alt) {
    e.preventDefault()
    wrapSelection('<u>', '</u>', '下划线文字')
  } else if (key === 's' && shift && !alt) {
    e.preventDefault()
    wrapSelection('~~', '~~', '删除线文字')
  } else if (code === 'Digit8' && shift && !alt) {
    e.preventDefault()
    insertBlock('- ', '列表项')
  } else if (code === 'Digit7' && shift && !alt) {
    e.preventDefault()
    insertBlock('1. ', '列表项')
  } else if (key === 'b' && shift && !alt) {
    e.preventDefault()
    insertBlock('> ', '引用内容')
  } else if (code === 'KeyC' && alt && !shift) {
    e.preventDefault()
    wrapSelection('\n```\n', '\n```\n', '代码')
  } else if (key === 'k' && !shift && !alt) {
    e.preventDefault()
    wrapSelection('[', '](https://example.com)', '链接文字')
  }
}

const closeHeadingDropdown = (e) => {
  if (!e.target.closest('.heading-dropdown-wrap')) {
    headingDropdownOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', closeHeadingDropdown)
})

onUnmounted(() => {
  document.removeEventListener('click', closeHeadingDropdown)
})
</script>

<style scoped>
.markdown-workspace {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0;
  transition: 0.3s ease;
}

.markdown-workspace.preview-hidden {
  grid-template-columns: 1fr;
}

.markdown-panel,
.preview-panel {
  display: grid;
  grid-template-rows: auto auto minmax(420px, 1fr);
  border: 1px solid var(--line);
  border-radius: 22px;
  background: rgba(255, 255, 255, 0.82);
  overflow: hidden;
}

.preview-panel {
  border-left: 0;
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
}

.markdown-panel {
  border-right: 0;
  border-top-right-radius: 0;
  border-bottom-right-radius: 0;
}

.preview-hidden .markdown-panel {
  border-right: 1px solid var(--line);
  border-top-right-radius: 22px;
  border-bottom-right-radius: 22px;
}

.panel-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 14px 18px;
  border-bottom: 1px solid var(--line);
  background: #fcf7ef;
}

.panel-head h3 {
  margin: 0;
  font-size: 15px;
}

.toolbar {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 2px;
  padding: 6px 8px;
  border-bottom: 1px solid var(--line);
  background: rgba(255, 255, 255, 0.9);
}

.toolbar-spacer {
  flex: 1;
}

.preview-toggle-icon {
  margin-left: 4px;
}

.preview-toggle-icon.active-preview {
  color: var(--accent);
}

.heading-dropdown-wrap {
  position: relative;
}

.heading-trigger {
  width: auto;
  padding: 0 8px;
  gap: 2px;
}

.heading-label {
  font-size: 13px;
  font-weight: 700;
}

.dropdown-arrow {
  width: 10px;
  height: 10px;
  flex-shrink: 0;
}

.heading-dropdown {
  position: absolute;
  top: calc(100% + 4px);
  left: 0;
  z-index: 100;
  min-width: 130px;
  padding: 4px;
  border: 1px solid var(--line);
  border-radius: 12px;
  background: #fff;
  box-shadow: 0 8px 28px rgba(15, 23, 42, 0.12);
}

.heading-option {
  display: block;
  width: 100%;
  padding: 6px 12px;
  border: 0;
  border-radius: 8px;
  background: transparent;
  color: var(--text-main);
  text-align: left;
  cursor: pointer;
  transition: 0.15s ease;
}

.heading-option:hover {
  background: var(--bg-soft);
}

.option-label {
  display: block;
}

.cls-p {
  font-size: 13px;
  font-weight: 400;
}

.cls-h1 {
  font-size: 20px;
  font-weight: 700;
  line-height: 1.3;
}

.cls-h2 {
  font-size: 17px;
  font-weight: 700;
  line-height: 1.3;
}

.cls-h3 {
  font-size: 15px;
  font-weight: 700;
  line-height: 1.3;
}

.cls-h4 {
  font-size: 13px;
  font-weight: 700;
  line-height: 1.3;
}

.cls-h5 {
  font-size: 12px;
  font-weight: 700;
  line-height: 1.3;
}

.toolbar-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border: 1px solid transparent;
  border-radius: 8px;
  background: transparent;
  color: var(--text-subtle);
  cursor: pointer;
  transition: 0.15s ease;
}

.toolbar-btn:hover {
  background: var(--bg-soft);
  color: var(--text-main);
  border-color: var(--line);
}

.toolbar-btn:active {
  transform: scale(0.92);
  background: rgba(0, 0, 0, 0.06);
}

.btn-icon {
  width: 16px;
  height: 16px;
  flex-shrink: 0;
}

.markdown-input {
  width: 100%;
  min-height: 420px;
  padding: 18px 20px;
  border: 0;
  resize: vertical;
  outline: none;
  background: transparent;
  color: var(--text-main);
  line-height: 1.75;
  font-family: 'Cascadia Code', 'Fira Code', 'JetBrains Mono', Consolas, monospace;
  font-size: 14px;
}

.preview-body {
  min-height: 420px;
  padding: 18px 20px;
  overflow: auto;
  line-height: 1.8;
}

.preview-body :deep(h1),
.preview-body :deep(h2),
.preview-body :deep(h3) {
  margin: 1.2em 0 0.6em;
  line-height: 1.3;
}

.preview-body :deep(p) {
  margin: 0.75em 0;
}

.preview-body :deep(ul),
.preview-body :deep(ol) {
  padding-left: 1.5em;
}

.preview-body :deep(pre) {
  padding: 14px 16px;
  border-radius: 16px;
  background: #18211f;
  color: #edf5f1;
  overflow-x: auto;
}

.preview-body :deep(blockquote) {
  margin: 1em 0;
  padding-left: 1em;
  border-left: 4px solid #d6dfda;
  color: #52615d;
}

.preview-body :deep(img) {
  max-width: 100%;
  border-radius: 14px;
}

@media (max-width: 1024px) {
  .markdown-workspace {
    grid-template-columns: 1fr;
  }

  .markdown-panel {
    border-right: 1px solid var(--line);
    border-radius: 22px;
  }

  .preview-panel {
    border-left: 1px solid var(--line);
    border-radius: 22px;
  }
}
</style>
