<template>
  <div class="rich-editor">
    <div class="toolbar">
      <div class="heading-dropdown-wrap">
        <button
          :class="['toolbar-btn heading-trigger', { active: isHeadingActive }]"
          type="button"
          title="段落与标题"
          @click="headingDropdownOpen = !headingDropdownOpen"
        >
          <span class="heading-label">{{ currentHeadingLabel }}</span>
          <svg class="dropdown-arrow" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 12 15 18 9"/></svg>
        </button>
        <div v-if="headingDropdownOpen" class="heading-dropdown">
          <button
            v-for="opt in headingOptions"
            :key="opt.name"
            :class="['heading-option', { active: opt.active?.() }]"
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
        :class="['toolbar-btn', { active: item.active?.() }]"
        type="button"
        :title="item.tooltip"
        @click="item.action"
      >
        <svg v-if="item.icon" class="btn-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" v-html="item.icon" />
        <span v-if="item.label" class="btn-text">{{ item.label }}</span>
      </button>
    </div>

    <div class="editor-stage">
      <EditorContent v-if="editor" :editor="editor" class="editor-content" />
      <div v-else class="editor-placeholder">正在初始化富文本编辑器...</div>
    </div>

    <div v-if="linkDialogVisible" class="dialog-overlay" @click.self="linkDialogVisible = false">
      <div class="dialog-card">
        <h3 class="dialog-title">插入链接</h3>
        <div class="dialog-field">
          <label class="dialog-label">链接地址</label>
          <input
            v-model="linkUrl"
            class="dialog-input"
            type="url"
            placeholder="https://example.com"
            @keydown.enter="confirmLink"
          />
        </div>
        <div class="dialog-field">
          <label class="dialog-label">显示文字</label>
          <input
            v-model="linkText"
            class="dialog-input"
            type="text"
            placeholder="点击此处跳转"
            @keydown.enter="confirmLink"
          />
        </div>
        <div class="dialog-actions">
          <button class="dialog-btn cancel" type="button" @click="linkDialogVisible = false">取消</button>
          <button class="dialog-btn confirm" type="button" @click="confirmLink">确定</button>
        </div>
      </div>
    </div>

    <input
      ref="imageInput"
      type="file"
      accept="image/*"
      style="display: none"
      @change="handleImageUpload"
    />
  </div>
</template>

<script setup>
/* global defineProps, defineEmits */
import { computed, nextTick, onBeforeUnmount, ref, watch } from 'vue'
import { EditorContent, useEditor } from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'
import Underline from '@tiptap/extension-underline'
import Link from '@tiptap/extension-link'
import Image from '@tiptap/extension-image'
import { Markdown } from 'tiptap-markdown'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue'])
const applyingExternalContent = ref(false)

const linkDialogVisible = ref(false)
const linkUrl = ref('')
const linkText = ref('')
const imageInput = ref(null)
const headingDropdownOpen = ref(false)

const editor = useEditor({
  content: props.modelValue,
  extensions: [
    StarterKit,
    Underline,
    Link.configure({
      openOnClick: false,
      autolink: true,
      defaultProtocol: 'https'
    }),
    Image.configure({
      inline: false,
      allowBase64: true
    }),
    Markdown.configure({
      html: true,
      linkify: true,
      breaks: true,
      transformCopiedText: true
    })
  ],
  editorProps: {
    attributes: {
      class: 'tiptap-editor',
      spellcheck: 'false'
    }
  },
  onUpdate: ({ editor: instance }) => {
    if (applyingExternalContent.value) {
      return
    }

    emit('update:modelValue', instance.storage.markdown.getMarkdown())
  }
})

watch(
  () => props.modelValue,
  async (value) => {
    if (!editor.value) {
      return
    }

    const currentMarkdown = editor.value.storage.markdown.getMarkdown()

    if (value === currentMarkdown) {
      return
    }

    applyingExternalContent.value = true
    editor.value.commands.setContent(value || '')
    await nextTick()
    applyingExternalContent.value = false
  }
)

const openLinkDialog = () => {
  if (!editor.value) {
    return
  }

  const previousUrl = editor.value.getAttributes('link').href || ''
  const { from, to } = editor.value.state.selection
  const selectedText = editor.value.state.doc.textBetween(from, to, '')

  linkUrl.value = previousUrl
  linkText.value = selectedText
  linkDialogVisible.value = true
}

const confirmLink = () => {
  if (!editor.value) {
    return
  }

  let url = linkUrl.value.trim()
  if (!url) {
    return
  }

  if (!/^https?:\/\//i.test(url)) {
    url = 'https://' + url
  }

  if (linkText.value.trim()) {
    editor.value.chain().focus().insertContent({
      type: 'text',
      text: linkText.value.trim(),
      marks: [{ type: 'link', attrs: { href: url } }]
    }).run()
  } else {
    editor.value.chain().focus().extendMarkRange('link').setLink({ href: url }).run()
  }

  linkDialogVisible.value = false
  linkUrl.value = ''
  linkText.value = ''
}

const triggerImageUpload = () => {
  imageInput.value?.click()
}

const handleImageUpload = (event) => {
  const file = event.target.files?.[0]
  if (!file) {
    return
  }

  if (!file.type.startsWith('image/')) {
    return
  }

  const reader = new FileReader()
  reader.onload = (e) => {
    const src = e.target?.result
    if (src && editor.value) {
      editor.value.chain().focus().setImage({ src }).run()
    }
  }
  reader.readAsDataURL(file)

  event.target.value = ''
}

const headingOptions = computed(() => {
  if (!editor.value) return []

  return [
    {
      name: 'paragraph',
      label: '正文',
      cls: 'cls-p',
      action: () => editor.value.chain().focus().setParagraph().run(),
      active: () => editor.value.isActive('paragraph')
    },
    ...[1, 2, 3, 4, 5].map(level => ({
      name: `h${level}`,
      label: `标题 ${level}`,
      cls: `cls-h${level}`,
      action: () => editor.value.chain().focus().toggleHeading({ level }).run(),
      active: () => editor.value.isActive('heading', { level })
    }))
  ]
})

const isHeadingActive = computed(() => {
  if (!editor.value) return false
  return [1, 2, 3, 4, 5].some(l => editor.value.isActive('heading', { level: l }))
})

const currentHeadingLabel = computed(() => {
  if (!editor.value) return 'H'
  for (let l = 1; l <= 5; l++) {
    if (editor.value.isActive('heading', { level: l })) return `H${l}`
  }
  return 'H'
})

const closeHeadingDropdown = (e) => {
  if (!e.target.closest('.heading-dropdown-wrap')) {
    headingDropdownOpen.value = false
  }
}

if (typeof document !== 'undefined') {
  document.addEventListener('click', closeHeadingDropdown)
}

onBeforeUnmount(() => {
  if (editor.value) {
    editor.value.destroy()
  }
  if (typeof document !== 'undefined') {
    document.removeEventListener('click', closeHeadingDropdown)
  }
})

const toolbarButtons = computed(() => {
  if (!editor.value) {
    return []
  }

  return [
    {
      name: 'bold',
      tooltip: '粗体 (Ctrl+B)',
      icon: '<path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/><path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/>',
      action: () => editor.value.chain().focus().toggleBold().run(),
      active: () => editor.value.isActive('bold')
    },
    {
      name: 'italic',
      tooltip: '斜体 (Ctrl+I)',
      icon: '<line x1="19" y1="4" x2="10" y2="4"/><line x1="14" y1="20" x2="5" y2="20"/><line x1="15" y1="4" x2="9" y2="20"/>',
      action: () => editor.value.chain().focus().toggleItalic().run(),
      active: () => editor.value.isActive('italic')
    },
    {
      name: 'underline',
      tooltip: '下划线 (Ctrl+U)',
      icon: '<path d="M6 3v7a6 6 0 0 0 6 6 6 6 0 0 0 6-6V3"/><line x1="4" y1="21" x2="20" y2="21"/>',
      action: () => editor.value.chain().focus().toggleUnderline().run(),
      active: () => editor.value.isActive('underline')
    },
    {
      name: 'strike',
      tooltip: '删除线 (Ctrl+Shift+S)',
      icon: '<path d="M16 4H9a3 3 0 0 0-3 3v0a3 3 0 0 0 3 3h6"/><line x1="4" y1="12" x2="20" y2="12"/><path d="M15 12a3 3 0 1 1 0 6H8"/>',
      action: () => editor.value.chain().focus().toggleStrike().run(),
      active: () => editor.value.isActive('strike')
    },
    {
      name: 'bulletList',
      tooltip: '无序列表 (Ctrl+Shift+8)',
      icon: '<line x1="8" y1="6" x2="21" y2="6"/><line x1="8" y1="12" x2="21" y2="12"/><line x1="8" y1="18" x2="21" y2="18"/><circle cx="4" cy="6" r="1" fill="currentColor"/><circle cx="4" cy="12" r="1" fill="currentColor"/><circle cx="4" cy="18" r="1" fill="currentColor"/>',
      action: () => editor.value.chain().focus().toggleBulletList().run(),
      active: () => editor.value.isActive('bulletList')
    },
    {
      name: 'orderedList',
      tooltip: '有序列表 (Ctrl+Shift+7)',
      icon: '<line x1="10" y1="6" x2="21" y2="6"/><line x1="10" y1="12" x2="21" y2="12"/><line x1="10" y1="18" x2="21" y2="18"/><text x="3" y="8" font-size="8" fill="currentColor" stroke="none" font-weight="bold">1</text><text x="3" y="14" font-size="8" fill="currentColor" stroke="none" font-weight="bold">2</text><text x="3" y="20" font-size="8" fill="currentColor" stroke="none" font-weight="bold">3</text>',
      action: () => editor.value.chain().focus().toggleOrderedList().run(),
      active: () => editor.value.isActive('orderedList')
    },
    {
      name: 'blockquote',
      tooltip: '引用 (Ctrl+Shift+B)',
      icon: '<path d="M3 21c3 0 7-1 7-5V5c0-1.25-.756-2.017-2-2-4 0-5 3-5 5s1 3 2 3c1 0 1-1 1-1"/><path d="M15 21c3 0 7-1 7-5V5c0-1.25-.757-2.017-2-2-4 0-5 3-5 5s1 3 2 3c1 0 1-1 1-1"/>',
      action: () => editor.value.chain().focus().toggleBlockquote().run(),
      active: () => editor.value.isActive('blockquote')
    },
    {
      name: 'codeBlock',
      tooltip: '代码块 (Ctrl+Alt+C)',
      icon: '<polyline points="16 18 22 12 16 6"/><polyline points="8 6 2 12 8 18"/>',
      action: () => editor.value.chain().focus().toggleCodeBlock().run(),
      active: () => editor.value.isActive('codeBlock')
    },
    {
      name: 'link',
      tooltip: '链接 (Ctrl+K)',
      icon: '<path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/>',
      action: openLinkDialog,
      active: () => editor.value.isActive('link')
    },
    {
      name: 'image',
      tooltip: '上传图片',
      icon: '<rect x="3" y="3" width="18" height="18" rx="2" ry="2"/><circle cx="8.5" cy="8.5" r="1.5"/><polyline points="21 15 16 10 5 21"/>',
      action: triggerImageUpload
    },
    {
      name: 'horizontalRule',
      tooltip: '分割线',
      icon: '<line x1="3" y1="12" x2="21" y2="12"/>',
      action: () => editor.value.chain().focus().setHorizontalRule().run()
    }
  ]
})
</script>

<style scoped>
.rich-editor {
  display: grid;
  gap: 14px;
  position: relative;
}

.toolbar {
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
  padding: 8px 10px;
  border: 1px solid var(--line);
  border-radius: 16px;
  background: rgba(255, 255, 255, 0.9);
}

.heading-dropdown-wrap {
  position: relative;
}

.heading-trigger {
  width: auto;
  padding: 0 10px;
  gap: 3px;
}

.heading-label {
  font-size: 14px;
  font-weight: 700;
}

.dropdown-arrow {
  width: 12px;
  height: 12px;
  flex-shrink: 0;
}

.heading-dropdown {
  position: absolute;
  top: calc(100% + 6px);
  left: 0;
  z-index: 100;
  min-width: 140px;
  padding: 6px;
  border: 1px solid var(--line);
  border-radius: 14px;
  background: #fff;
  box-shadow: 0 12px 36px rgba(15, 23, 42, 0.14);
}

.heading-option {
  display: block;
  width: 100%;
  padding: 8px 14px;
  border: 0;
  border-radius: 10px;
  background: transparent;
  color: var(--text-main);
  text-align: left;
  cursor: pointer;
  transition: 0.15s ease;
}

.heading-option:hover {
  background: var(--bg-soft);
}

.heading-option.active {
  background: var(--accent-soft);
  color: var(--accent);
}

.option-label {
  display: block;
}

.cls-p {
  font-size: 14px;
  font-weight: 400;
}

.cls-h1 {
  font-size: 22px;
  font-weight: 700;
  line-height: 1.3;
}

.cls-h2 {
  font-size: 19px;
  font-weight: 700;
  line-height: 1.3;
}

.cls-h3 {
  font-size: 16px;
  font-weight: 700;
  line-height: 1.3;
}

.cls-h4 {
  font-size: 14px;
  font-weight: 700;
  line-height: 1.3;
}

.cls-h5 {
  font-size: 13px;
  font-weight: 700;
  line-height: 1.3;
}

.toolbar-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  border: 1px solid transparent;
  border-radius: 10px;
  background: transparent;
  color: var(--text-subtle);
  cursor: pointer;
  transition: 0.18s ease;
}

.toolbar-btn:hover {
  background: var(--bg-soft);
  color: var(--text-main);
  border-color: var(--line);
}

.toolbar-btn.active {
  background: var(--accent-soft);
  color: var(--accent);
  border-color: rgba(37, 99, 235, 0.2);
}

.btn-icon {
  width: 18px;
  height: 18px;
  flex-shrink: 0;
}

.editor-stage {
  min-height: 420px;
  border: 1px solid var(--line);
  border-radius: 22px;
  background: rgba(255, 255, 255, 0.86);
  overflow: hidden;
}

.editor-placeholder {
  display: grid;
  place-items: center;
  min-height: 420px;
  color: var(--text-subtle);
}

:deep(.editor-content) {
  min-height: 420px;
}

:deep(.tiptap-editor) {
  min-height: 420px;
  padding: 22px 24px;
  outline: none;
  color: var(--text-main);
  line-height: 1.8;
}

:deep(.tiptap-editor p.is-editor-empty:first-child::before) {
  content: "开始撰写正文，或者把现有 Markdown 粘贴进来继续打磨。";
  color: #9aa5a1;
  float: left;
  height: 0;
  pointer-events: none;
}

:deep(.tiptap-editor h1),
:deep(.tiptap-editor h2),
:deep(.tiptap-editor h3) {
  margin: 1.2em 0 0.6em;
  line-height: 1.3;
}

:deep(.tiptap-editor p) {
  margin: 0.75em 0;
}

:deep(.tiptap-editor ul),
:deep(.tiptap-editor ol) {
  padding-left: 1.4em;
}

:deep(.tiptap-editor blockquote) {
  margin: 1em 0;
  padding-left: 1em;
  border-left: 4px solid #d6dfda;
  color: #52615d;
}

:deep(.tiptap-editor pre) {
  padding: 14px 16px;
  border-radius: 14px;
  background: #18211f;
  color: #edf5f1;
  overflow-x: auto;
}

:deep(.tiptap-editor a) {
  color: #1b6d5a;
  text-decoration: underline;
}

:deep(.tiptap-editor img) {
  display: block;
  max-width: 100%;
  margin: 1em 0;
  border-radius: 16px;
}

.dialog-overlay {
  position: fixed;
  inset: 0;
  z-index: 9999;
  display: grid;
  place-items: center;
  background: rgba(15, 23, 42, 0.4);
  backdrop-filter: blur(4px);
}

.dialog-card {
  width: min(440px, 92vw);
  padding: 28px;
  border-radius: 20px;
  background: #fff;
  box-shadow: 0 24px 64px rgba(15, 23, 42, 0.18);
}

.dialog-title {
  margin: 0 0 22px;
  font-size: 18px;
  font-weight: 700;
  color: var(--text-main);
}

.dialog-field {
  margin-bottom: 16px;
}

.dialog-label {
  display: block;
  margin-bottom: 6px;
  font-size: 13px;
  font-weight: 600;
  color: var(--text-subtle);
}

.dialog-input {
  width: 100%;
  padding: 10px 14px;
  border: 1px solid var(--line);
  border-radius: 12px;
  background: #fafafa;
  color: var(--text-main);
  font-size: 14px;
  outline: none;
  transition: 0.18s ease;
  box-sizing: border-box;
}

.dialog-input:focus {
  border-color: var(--accent);
  box-shadow: 0 0 0 3px rgba(37, 99, 235, 0.1);
}

.dialog-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 22px;
}

.dialog-btn {
  padding: 9px 22px;
  border: 1px solid var(--line);
  border-radius: 12px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: 0.18s ease;
}

.dialog-btn.cancel {
  background: #fff;
  color: var(--text-subtle);
}

.dialog-btn.cancel:hover {
  background: #f5f5f5;
}

.dialog-btn.confirm {
  background: var(--accent);
  color: #fff;
  border-color: transparent;
}

.dialog-btn.confirm:hover {
  opacity: 0.9;
}
</style>
