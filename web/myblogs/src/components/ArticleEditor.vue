<template>
  <section class="article-editor-card">
    <div class="editor-top">
      <div>
        <p class="eyebrow">Editor</p>
        <h2>内容编辑器</h2>
      </div>

      <div class="mode-switch">
        <button
          :class="['mode-btn', { active: mode === 'rich' }]"
          type="button"
          @click="mode = 'rich'"
        >
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 20h9"/><path d="M16.5 3.5a2.121 2.121 0 1 1 3 3L7 19l-4 1 1-4L16.5 3.5z"/></svg>
          富文本
        </button>
        <button
          :class="['mode-btn', { active: mode === 'markdown' }]"
          type="button"
          @click="mode = 'markdown'"
        >
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="16 18 22 12 16 6"/><polyline points="8 6 2 12 8 18"/></svg>
          Markdown
        </button>
      </div>
    </div>

    <div class="editor-meta">
      <span>当前模式：{{ modeLabel }}</span>
      <span>字符数：{{ characterCount }}</span>
      <span>预计阅读：{{ readingMinutes }} 分钟</span>
    </div>

    <RichTextEditor
      v-if="mode === 'rich'"
      :model-value="modelValue"
      @update:model-value="emit('update:modelValue', $event)"
    />
    <MarkdownWorkspace
      v-else
      :model-value="modelValue"
      @update:model-value="emit('update:modelValue', $event)"
    />
  </section>
</template>

<script setup>
/* global defineProps, defineEmits */
import { computed, ref } from 'vue'
import MarkdownWorkspace from './MarkdownWorkspace.vue'
import RichTextEditor from './RichTextEditor.vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue'])

const mode = ref('rich')

const plainTextContent = computed(() =>
  (props.modelValue || '')
    .replace(/```[\s\S]*?```/g, ' ')
    .replace(/`[^`]*`/g, ' ')
    .replace(/!\[[^\]]*]\([^)]*\)/g, ' ')
    .replace(/\[[^\]]*]\([^)]*\)/g, ' ')
    .replace(/[#>*_\-\n\r]/g, ' ')
    .replace(/\s+/g, ' ')
    .trim()
)

const characterCount = computed(() => plainTextContent.value.length)

const readingMinutes = computed(() => {
  const count = Math.max(plainTextContent.value.length, 1)
  return Math.max(1, Math.ceil(count / 350))
})

const modeLabel = computed(() => (mode.value === 'rich' ? '富文本编辑' : 'Markdown 编辑'))
</script>

<style scoped>
.article-editor-card {
  display: grid;
  gap: 18px;
  padding: 24px;
  border: 1px solid var(--line);
  border-radius: 24px;
  background: rgba(255, 255, 255, 0.54);
}

.editor-top {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 18px;
}

.eyebrow {
  margin: 0 0 8px;
  color: #9e5b3d;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.14em;
  text-transform: uppercase;
}

.editor-top h2 {
  margin: 0;
  font-family: Georgia, "Source Han Serif SC", "Songti SC", serif;
  font-size: clamp(24px, 2vw, 32px);
  line-height: 1.16;
}

.mode-switch {
  display: flex;
  gap: 0;
  border: 1px solid var(--line);
  border-radius: 14px;
  background: rgba(255, 255, 255, 0.7);
  overflow: hidden;
}

.mode-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 10px 18px;
  border: 0;
  border-radius: 0;
  background: transparent;
  color: var(--text-subtle);
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  transition: 0.2s ease;
  position: relative;
}

.mode-btn:first-child {
  border-right: 1px solid var(--line);
}

.mode-btn:hover:not(.active) {
  background: rgba(37, 99, 235, 0.04);
  color: var(--text-main);
}

.mode-btn.active {
  background: var(--accent);
  color: #fff;
}

.mode-btn.active svg {
  stroke: #fff;
}

.editor-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  color: var(--text-subtle);
  font-size: 14px;
}

.editor-meta span {
  padding: 8px 12px;
  border-radius: 999px;
  background: #f5ede1;
}

@media (max-width: 1024px) {
  .editor-top {
    flex-direction: column;
  }

  .mode-switch {
    width: 100%;
  }

  .mode-btn {
    flex: 1;
  }
}
</style>
