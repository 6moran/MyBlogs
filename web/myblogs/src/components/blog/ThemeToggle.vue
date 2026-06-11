<template>
  <button class="theme-toggle" @click="toggle" :title="theme === 'dark' ? '切换到浅色模式' : '切换到深色模式'" :aria-label="theme === 'dark' ? '切换到浅色模式' : '切换到深色模式'">
    <span v-if="theme === 'dark'">&#9728;</span>
    <span v-else>&#9790;</span>
  </button>
</template>

<script setup>
import { useTheme } from '@/composables/useTheme'

const { theme, toggle } = useTheme()
</script>

<style scoped>
.theme-toggle {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 34px;
  height: 34px;
  background: none;
  border: 1px solid var(--border);
  border-radius: 8px;
  cursor: pointer;
  font-size: 18px;
  color: var(--text);
  transition: border-color 0.2s ease, color 0.2s ease;
  padding: 0;
}

.theme-toggle:hover {
  border-color: var(--link);
  color: var(--link);
}
</style>
