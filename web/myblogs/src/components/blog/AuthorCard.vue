<template>
  <div class="author-card">
    <div class="author-avatar">
      <span class="avatar-text">Hw</span>
    </div>
    <div class="author-info">
      <h2 class="author-name">HwBlogs</h2>
      <p class="author-bio">记录编程学习笔记和技术思考，用文字沉淀知识。</p>
      <div class="author-links">
        <a href="https://github.com" target="_blank" rel="noopener" class="social-link">GitHub</a>
      </div>
    </div>
  </div>
</template>

<style scoped>
.author-card {
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 28px 0;
  border-bottom: 1px solid var(--border);
  margin-bottom: 32px;
}

.author-avatar {
  flex-shrink: 0;
  width: 56px;
  height: 56px;
  border-radius: 50%;
  background: var(--bg-secondary);
  border: 1px solid var(--border);
  display: grid;
  place-items: center;
}

.avatar-text {

  font-size: 18px;
  font-weight: 700;
  color: var(--text);
}

.author-info {
  min-width: 0;
}

.author-name {
  margin: 0 0 4px;
  font-size: 18px;
  font-weight: 700;

  color: var(--text);
}

.author-bio {
  margin: 0 0 8px;
  font-size: 14px;
  color: var(--text-secondary);
  line-height: 1.6;
}

.author-links {
  display: flex;
  gap: 12px;
}

.social-link {
  font-size: 13px;
  color: var(--text-muted);
  text-decoration: none;
  transition: color 0.2s ease;
}

.social-link:hover {
  color: var(--link);
}

@media (max-width: 600px) {
  .author-card {
    gap: 16px;
    padding: 20px 0;
  }

  .author-avatar {
    width: 44px;
    height: 44px;
  }

  .avatar-text {
    font-size: 15px;
  }
}
</style>
