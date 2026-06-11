import { ref } from 'vue'

const STORAGE_KEY = 'blog-theme'

function getSystemTheme() {
  if (typeof window !== 'undefined' && window.matchMedia) {
    return window.matchMedia('(prefers-color-scheme: dark)').matches
      ? 'dark'
      : 'light'
  }
  return 'light'
}

function getInitialTheme() {
  if (typeof window !== 'undefined') {
    const stored = localStorage.getItem(STORAGE_KEY)
    if (stored === 'dark' || stored === 'light') {
      return stored
    }
  }
  return getSystemTheme()
}

const theme = ref(getInitialTheme())

function applyTheme(value) {
  if (typeof window !== 'undefined') {
    document.documentElement.setAttribute('data-theme', value)
  }
}

function toggle() {
  theme.value = theme.value === 'dark' ? 'light' : 'dark'
  localStorage.setItem(STORAGE_KEY, theme.value)
  applyTheme(theme.value)
}

let initialized = false

export function useTheme() {
  if (!initialized) {
    applyTheme(theme.value)

    if (typeof window !== 'undefined' && window.matchMedia) {
      const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
      mediaQuery.addEventListener('change', (e) => {
        if (!localStorage.getItem(STORAGE_KEY)) {
          theme.value = e.matches ? 'dark' : 'light'
          applyTheme(theme.value)
        }
      })
    }

    initialized = true
  }

  return { theme, toggle }
}
