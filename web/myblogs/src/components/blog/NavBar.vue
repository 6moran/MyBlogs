<template>
  <header class="nav-bar">
    <div class="nav-inner">
      <RouterLink class="site-title" to="/">HwBlogs</RouterLink>

      <button class="hamburger" :class="{ open: menuOpen }" @click="menuOpen = !menuOpen" aria-label="菜单" :aria-expanded="menuOpen">
        <span></span>
        <span></span>
        <span></span>
      </button>

      <div class="nav-end">
        <nav class="nav-right" :class="{ open: menuOpen }">
          <RouterLink class="nav-link" to="/" @click="menuOpen = false">博客</RouterLink>
          <RouterLink class="nav-link" to="/tags" @click="menuOpen = false">标签</RouterLink>
          <RouterLink class="nav-link" to="/about" @click="menuOpen = false">关于</RouterLink>
          <ThemeToggle />
        </nav>
        <SearchBox />
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import ThemeToggle from './ThemeToggle.vue'
import SearchBox from './SearchBox.vue'

const menuOpen = ref(false)

function onKeyDown(e) {
  if (e.key === 'Escape') menuOpen.value = false
}

onMounted(() => document.addEventListener('keydown', onKeyDown))
onUnmounted(() => document.removeEventListener('keydown', onKeyDown))
</script>

<style scoped>
.nav-bar {
  position: sticky;
  top: 0;
  z-index: 100;
  background: var(--nav-bg);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  border-bottom: 1px solid var(--border);
  max-width: 900px;
  margin: 0 auto;
  width: 100%;
}

.nav-inner {
  display: flex;
  align-items: center;
  justify-content: space-between;
  max-width: 900px;
  margin: 0 auto;
  padding: 0 20px;
  height: 56px;
  position: relative;
}

.site-title {
  font-size: 18px;
  font-weight: 700;
  color: var(--text);
  text-decoration: none;
  transition: color 0.2s ease;
}

.site-title:hover {
  color: var(--link);
}

.nav-end {
  display: flex;
  align-items: center;
  gap: 16px;
}

.nav-right {
  display: flex;
  align-items: center;
  gap: 20px;
}

.nav-link {
  font-size: 15px;
  font-weight: 500;
  color: var(--text-secondary);
  text-decoration: none;
  transition: color 0.2s ease;
  padding: 4px 0;
}

.nav-link:hover {
  color: var(--link);
}

.nav-link.router-link-exact-active {
  color: var(--link);
}

/* Hamburger button - hidden on desktop */
.hamburger {
  display: none;
  background: none;
  border: none;
  cursor: pointer;
  padding: 4px;
  width: 34px;
  height: 34px;
  position: relative;
}

.hamburger span {
  display: block;
  width: 20px;
  height: 2px;
  background-color: var(--text);
  position: absolute;
  left: 7px;
  transition: transform 0.3s ease, opacity 0.3s ease;
}

.hamburger span:nth-child(1) {
  top: 9px;
}

.hamburger span:nth-child(2) {
  top: 16px;
}

.hamburger span:nth-child(3) {
  top: 23px;
}

/* Hamburger -> X animation */
.hamburger.open span:nth-child(1) {
  top: 16px;
  transform: rotate(45deg);
}

.hamburger.open span:nth-child(2) {
  opacity: 0;
}

.hamburger.open span:nth-child(3) {
  top: 16px;
  transform: rotate(-45deg);
}

/* Mobile layout */
@media (max-width: 767px) {
  .hamburger {
    display: block;
  }

  .nav-right {
    display: none;
    position: absolute;
    top: 56px;
    left: 0;
    right: 0;
    flex-direction: column;
    align-items: flex-start;
    gap: 0;
    padding: 12px 20px 16px;
    background: var(--nav-bg);
    backdrop-filter: blur(12px);
    -webkit-backdrop-filter: blur(12px);
    border-bottom: 1px solid var(--border);
    box-shadow: var(--shadow);
  }

  .nav-right.open {
    display: flex;
  }

  .nav-link {
    width: 100%;
    padding: 10px 0;
    font-size: 16px;
    border-bottom: 1px solid var(--border);
  }

  .nav-link:last-of-type {
    border-bottom: none;
  }

  .nav-right .theme-toggle {
    margin-top: 8px;
  }
}
</style>
