<template>
  <div class="admin-shell">
    <aside class="sidebar">
      <div class="sidebar-top">
        <RouterLink class="brand" to="/">
          <span class="brand-mark">MB</span>
          <span class="brand-text">MyBlogs</span>
        </RouterLink>
      </div>

      <nav class="sidebar-nav">
        <RouterLink
          v-for="item in navItems"
          :key="item.to"
          class="nav-item"
          :to="item.to"
        >
          <span class="nav-icon" v-html="item.icon"></span>
          <span class="nav-label">{{ item.label }}</span>
        </RouterLink>
      </nav>

      <div class="sidebar-bottom">
        <RouterLink class="nav-item back-link" to="/">
          <span class="nav-icon">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M15 18l-6-6 6-6"/></svg>
          </span>
          <span class="nav-label">返回前台</span>
        </RouterLink>
      </div>
    </aside>

    <div class="main-area">
      <header class="topbar">
        <div class="topbar-left">
          <h1 class="page-title">{{ currentPageTitle }}</h1>
        </div>
        <div class="topbar-right">
          <span class="status-dot"></span>
          <span class="status-text">运行正常</span>
          <div class="user-badge">
            <span class="user-avatar">管</span>
            <span class="user-name">管理员</span>
          </div>
        </div>
      </header>

      <main class="content-area">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const navItems = [
  {
    to: '/admin/articles',
    label: '文章',
    icon: '<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/><polyline points="10 9 9 9 8 9"/></svg>'
  },
  {
    to: '/admin/tags',
    label: '标签',
    icon: '<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.59 13.41l-7.17 7.17a2 2 0 0 1-2.83 0L2 12V2h10l8.59 8.59a2 2 0 0 1 0 2.82z"/><line x1="7" y1="7" x2="7.01" y2="7"/></svg>'
  },
  {
    to: '/admin/users',
    label: '用户',
    icon: '<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M23 21v-2a4 4 0 0 0-3-3.87"/><path d="M16 3.13a4 4 0 0 1 0 7.75"/></svg>'
  }
]

const pageTitles = {
  'admin-article-manage': '文章管理',
  'admin-tags': '标签管理',
  'admin-users': '用户管理'
}

const currentPageTitle = computed(() => pageTitles[route.name] || '后台管理')
</script>

<style scoped>
.admin-shell {
  display: flex;
  height: 100vh;
  background: var(--bg-page);
  overflow: hidden;
}

/* ===== Sidebar ===== */
.sidebar {
  display: flex;
  flex-direction: column;
  width: 220px;
  flex-shrink: 0;
  background: var(--bg-sidebar);
  border-right: 1px solid var(--line);
  overflow-y: auto;
}

.sidebar-top {
  padding: 20px 16px;
  border-bottom: 1px solid var(--line);
}

.brand {
  display: flex;
  align-items: center;
  gap: 10px;
  text-decoration: none;
}

.brand-mark {
  display: grid;
  place-items: center;
  width: 34px;
  height: 34px;
  border-radius: var(--radius-lg);
  background: linear-gradient(135deg, #c0392b 0%, #a93226 100%);
  color: #fff;
  font-size: 13px;
  font-weight: 800;
  letter-spacing: 0.04em;
}

.brand-text {
  font-size: 16px;
  font-weight: 700;
  color: var(--text-main);
}

/* ===== Nav ===== */
.sidebar-nav {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 2px;
  padding: 12px 8px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 12px;
  border-radius: var(--radius-lg);
  color: var(--text-subtle);
  text-decoration: none;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.15s;
}

.nav-item:hover {
  background: var(--bg-soft);
  color: var(--text-main);
}

.nav-item.router-link-active {
  background: var(--accent-soft);
  color: var(--accent);
  font-weight: 600;
}

.nav-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 20px;
  height: 20px;
  flex-shrink: 0;
}

.nav-label {
  white-space: nowrap;
}

/* ===== Sidebar Bottom ===== */
.sidebar-bottom {
  padding: 8px;
  border-top: 1px solid var(--line);
}

.back-link {
  color: var(--text-muted);
}

.back-link:hover {
  color: var(--text-main);
}

/* ===== Main Area ===== */
.main-area {
  flex: 1;
  display: flex;
  flex-direction: column;
  min-width: 0;
  overflow: hidden;
}

/* ===== Topbar ===== */
.topbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 14px 28px;
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(12px);
  border-bottom: 1px solid var(--line);
  flex-shrink: 0;
}

.page-title {
  margin: 0;
  font-family: var(--font-serif);
  font-size: 20px;
  font-weight: 700;
  color: var(--text-main);
}

.topbar-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.status-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: var(--success);
}

.status-text {
  font-size: 13px;
  color: var(--text-subtle);
  font-weight: 500;
}

.user-badge {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 5px 12px 5px 5px;
  border: 1px solid var(--line);
  border-radius: 999px;
  margin-left: 4px;
}

.user-avatar {
  display: grid;
  place-items: center;
  width: 28px;
  height: 28px;
  border-radius: 50%;
  background: linear-gradient(135deg, #c0392b 0%, #a93226 100%);
  color: #fff;
  font-size: 12px;
  font-weight: 700;
}

.user-name {
  font-size: 13px;
  font-weight: 600;
  color: var(--text-main);
}

/* ===== Content ===== */
.content-area {
  flex: 1;
  min-height: 0;
  overflow: hidden;
  padding: 24px 28px;
}

/* ===== Responsive ===== */
@media (max-width: 960px) {
  .sidebar {
    width: 60px;
  }

  .brand-text,
  .nav-label,
  .status-text,
  .user-name {
    display: none;
  }

  .sidebar-top {
    padding: 16px 13px;
    display: flex;
    justify-content: center;
  }

  .brand {
    justify-content: center;
  }

  .nav-item {
    justify-content: center;
    padding: 10px;
  }

  .sidebar-bottom .nav-item {
    padding: 10px;
  }

  .user-badge {
    padding: 5px;
  }

  .content-area {
    padding: 20px;
  }
}

@media (max-width: 640px) {
  .sidebar {
    display: none;
  }

  .topbar {
    padding: 12px 16px;
  }

  .content-area {
    padding: 16px;
  }
}
</style>
