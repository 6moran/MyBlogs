<template>
  <section class="admin-page">
    <div class="admin-scroll">
      <header class="admin-header">
        <div>
          <span class="admin-label">标签管理</span>
          <h2>管理文章标签</h2>
          <p class="admin-desc">管理博客文章的标签，为内容添加关键词标记。</p>
        </div>
        <button class="admin-btn admin-btn-primary" type="button">
          <span class="admin-btn-icon">+</span>
          新建标签
        </button>
      </header>

      <section class="admin-panel">
        <div class="admin-toolbar">
          <label class="admin-search">
            <svg class="admin-search-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/></svg>
            <input type="text" class="admin-search-input" placeholder="搜索标签名称" />
          </label>
        </div>

        <div class="admin-empty">
          <h3>暂无标签</h3>
          <p>点击上方「新建标签」按钮创建第一个标签。</p>
          <button class="admin-btn admin-btn-primary" type="button">
            <span class="admin-btn-icon">+</span>
            新建标签
          </button>
        </div>
      </section>
    </div>
  </section>
</template>

<script setup>
</script>
