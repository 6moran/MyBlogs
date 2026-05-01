<template>
  <section class="category-manage-page">
    <div class="scroll-container">
    <header class="page-header">
      <div>
        <p class="page-kicker">分类管理</p>
        <h2>管理文章分类</h2>
        <p class="page-desc">管理博客文章的分类，方便读者浏览和检索。</p>
      </div>
      <button class="primary-btn" type="button">
        <span class="btn-icon">+</span>
        新建分类
      </button>
    </header>

    <section class="list-panel">
      <div class="toolbar">
        <label class="search-shell">
          <svg class="search-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="11" cy="11" r="8"></circle>
            <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
          </svg>
          <input type="text" class="search-input" placeholder="搜索分类名称" />
        </label>
      </div>

      <div class="category-list">
        <div class="empty-state">
          <h3>暂无分类</h3>
          <p>点击上方「新建分类」按钮创建第一个分类。</p>
          <button class="primary-btn" type="button">
            <span class="btn-icon">+</span>
            新建分类
          </button>
        </div>
      </div>
    </section>
    </div>
  </section>
</template>

<script setup>
</script>

<style scoped>
.category-manage-page {
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

.page-kicker {
  margin: 0 0 8px;
  color: var(--text-subtle);
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.page-header h2 {
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

.search-input {
  width: 100%;
  padding: 12px 14px 12px 42px;
  border: 1px solid var(--line);
  border-radius: 14px;
  background: var(--bg-soft);
  color: var(--text-main);
  outline: none;
  transition: 0.2s ease;
}

.search-input:focus {
  border-color: rgba(37, 99, 235, 0.4);
  box-shadow: 0 0 0 4px rgba(37, 99, 235, 0.1);
}

.category-list {
  display: grid;
  gap: 14px;
}

.primary-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 12px 18px;
  border: 0;
  border-radius: 14px;
  background: linear-gradient(135deg, #2563eb 0%, #1d4ed8 100%);
  color: #fff;
  font-weight: 700;
  cursor: pointer;
  transition: 0.2s ease;
}

.primary-btn:hover {
  transform: translateY(-1px);
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
</style>
