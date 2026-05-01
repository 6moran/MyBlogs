<template>
  <div id="app" class="admin-shell">
    <header class="topbar">
      <div class="brand-group">
        <div class="brand-mark">MB</div>
        <div>
          <p class="brand-subtitle">Personal Blog Admin</p>
          <h1 class="brand-title">MyBlogs 后台管理</h1>
        </div>
      </div>

      <div class="topbar-right">
        <div class="status-pill">系统运行正常</div>
        <div class="user-card">
          <span class="user-avatar">管</span>
          <div>
            <strong>管理员</strong>
            <small>个人博客后台</small>
          </div>
        </div>
      </div>
    </header>

    <div class="layout-grid">
      <aside class="sidebar">
        <div class="sidebar-panel">
          <nav class="nav-list">
            <RouterLink class="nav-item" to="/admin/articles">
              <strong>文章管理</strong>
              <small>增删、查看和编辑文章</small>
            </RouterLink>
            <RouterLink class="nav-item" to="/admin/categories">
              <strong>分类管理</strong>
              <small>管理文章分类</small>
            </RouterLink>
            <RouterLink class="nav-item" to="/admin/tags">
              <strong>标签管理</strong>
              <small>管理文章标签</small>
            </RouterLink>
            <RouterLink class="nav-item" to="/admin/users">
              <strong>用户管理</strong>
              <small>管理用户账号与权限</small>
            </RouterLink>
          </nav>
        </div>
      </aside>

      <main class="content-shell">
        <router-view />
      </main>
    </div>
  </div>
</template>

<style>
:root {
  color-scheme: light;
  --bg-page: #f5f7fb;
  --bg-panel: #ffffff;
  --bg-soft: #f3f6fb;
  --bg-sidebar: #1f2937;
  --accent: #2563eb;
  --accent-soft: rgba(37, 99, 235, 0.12);
  --line: rgba(15, 23, 42, 0.08);
  --line-strong: rgba(15, 23, 42, 0.16);
  --text-main: #111827;
  --text-subtle: #6b7280;
  --success: #15803d;
  --danger: #dc2626;
  --shadow: 0 18px 48px rgba(15, 23, 42, 0.08);
  --radius-xl: 24px;
  --radius-lg: 18px;
  --radius-md: 14px;
}

* {
  box-sizing: border-box;
}

html,
body,
#app {
  margin: 0;
  height: 100%;
  overflow: hidden;
}

body {
  font-family: "Segoe UI", "PingFang SC", "Hiragino Sans GB", sans-serif;
  color: var(--text-main);
  background: linear-gradient(180deg, #f8fafc 0%, #eef2f7 100%);
}

a {
  color: inherit;
  text-decoration: none;
}

button,
input,
select,
textarea {
  font: inherit;
}

.admin-shell {
  height: 100vh;
  padding: 20px 0 20px 20px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.topbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 18px 22px;
  margin-bottom: 18px;
  margin-right: 20px;
  border: 1px solid var(--line);
  border-radius: var(--radius-xl);
  background: rgba(255, 255, 255, 0.92);
  box-shadow: var(--shadow);
}

.brand-group,
.topbar-right,
.user-card {
  display: flex;
  align-items: center;
  gap: 14px;
}

.brand-mark {
  display: grid;
  place-items: center;
  width: 48px;
  height: 48px;
  border-radius: 14px;
  background: linear-gradient(135deg, #2563eb 0%, #1d4ed8 100%);
  color: #fff;
  font-weight: 800;
  letter-spacing: 0.08em;
}

.brand-subtitle,
.sidebar-label {
  margin: 0 0 4px;
  color: var(--text-subtle);
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.brand-title {
  margin: 0;
  font-size: 24px;
  font-weight: 700;
}

.status-pill {
  padding: 9px 14px;
  border-radius: 999px;
  background: rgba(21, 128, 61, 0.1);
  color: var(--success);
  font-size: 13px;
  font-weight: 700;
}

.user-card {
  padding: 8px 12px;
  border: 1px solid var(--line);
  border-radius: 16px;
  background: var(--bg-soft);
}

.user-avatar {
  display: grid;
  place-items: center;
  width: 36px;
  height: 36px;
  border-radius: 12px;
  background: var(--bg-sidebar);
  color: #fff;
  font-weight: 700;
}

.user-card strong,
.user-card small {
  display: block;
}

.user-card small {
  margin-top: 2px;
  color: var(--text-subtle);
  font-size: 12px;
}

.layout-grid {
  display: grid;
  grid-template-columns: 250px minmax(0, 1fr);
  gap: 18px;
  flex: 1;
  min-height: 0;
}

.sidebar-panel {
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 22px;
  border-radius: var(--radius-xl);
  background: linear-gradient(180deg, #1f2937 0%, #111827 100%);
  color: #f9fafb;
  box-shadow: var(--shadow);
}

.sidebar-header {
  margin-bottom: 18px;
}

.sidebar-text {
  margin: 0;
  color: rgba(249, 250, 251, 0.72);
  line-height: 1.7;
  font-size: 14px;
}

.nav-list {
  display: grid;
  gap: 10px;
}

.nav-item {
  display: grid;
  gap: 6px;
  padding: 16px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 18px;
  color: rgba(249, 250, 251, 0.82);
  background: rgba(255, 255, 255, 0.04);
  transition: 0.2s ease;
}

.nav-item strong {
  font-size: 16px;
}

.nav-item small {
  color: rgba(249, 250, 251, 0.66);
}

.nav-item:hover,
.nav-item.router-link-active {
  border-color: rgba(96, 165, 250, 0.5);
  background: rgba(37, 99, 235, 0.18);
  color: #fff;
}

.sidebar-foot {
  margin-top: auto;
  padding: 16px;
  border-radius: 18px;
  background: rgba(255, 255, 255, 0.05);
  color: rgba(249, 250, 251, 0.72);
  font-size: 13px;
  line-height: 1.7;
}

.sidebar-foot p {
  margin: 0;
}

.content-shell {
  min-width: 0;
  min-height: 0;
  overflow: hidden;
}

@media (max-width: 960px) {
  .layout-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 768px) {
  .admin-shell {
    padding: 14px 0 14px 14px;
  }

  .topbar {
    flex-direction: column;
    align-items: flex-start;
    padding: 16px;
    margin-right: 14px;
  }

  .topbar-right {
    width: 100%;
    flex-direction: column;
    align-items: stretch;
  }

  .status-pill {
    text-align: center;
  }

  .user-card {
    width: 100%;
    justify-content: flex-start;
  }
}
</style>
