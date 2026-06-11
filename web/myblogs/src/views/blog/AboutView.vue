<template>
  <div class="about-view">
    <div class="about-header">
      <h1 class="page-title">关于</h1>
      <div class="author-card">
        <div class="author-avatar">
          <span class="avatar-text">Hw</span>
        </div>
        <div class="author-info">
          <h3 class="author-name">HwBlogs</h3>
          <p class="author-bio">记录编程学习笔记和技术思考，用文字沉淀知识。</p>
          <div class="author-links">
            <a href="https://github.com" target="_blank" rel="noopener" class="social-link">
              <svg class="icon" viewBox="0 0 24 24" fill="currentColor">
                <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
              </svg>
              <span>GitHub</span>
            </a>
          </div>
        </div>
      </div>
    </div>

    <div class="about-content">
      <div class="about-section">
        <h2>技术栈</h2>
        <ul class="tech-list">
          <li>前端：Vue 3 + Vue Router</li>
          <li>后端：Go + Gin</li>
          <li>样式：纯 CSS，支持明暗主题</li>
        </ul>
      </div>

      <div class="about-section">
        <h2>联系方式</h2>
        <p>如果你对文章内容有疑问或建议，欢迎交流。</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.about-header {
  margin-bottom: 32px;
}

.page-title {
  font-size: 24px;
  font-weight: 700;
  margin: 0 0 24px;
  color: var(--text);
}

.author-card {
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 20px;
  border: 1px solid var(--border);
  border-radius: 12px;
  background: var(--bg-secondary);
}

.author-avatar {
  flex-shrink: 0;
  width: 64px;
  height: 64px;
  border-radius: 50%;
  background: var(--bg-secondary);
  border: 1px solid var(--border);
  display: grid;
  place-items: center;
}

.avatar-text {
  font-size: 20px;
  font-weight: 700;
  color: var(--text);
}

.author-info {
  min-width: 0;
}

.author-name {
  margin: 0 0 6px;
  font-size: 18px;
  font-weight: 600;
  color: var(--text);
}

.author-bio {
  margin: 0 0 12px;
  font-size: 14px;
  color: var(--text-secondary);
  line-height: 1.6;
}

.author-links {
  display: flex;
  gap: 12px;
}

.social-link {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: var(--text-muted);
  text-decoration: none;
  transition: color 0.2s ease;
}

.social-link:hover {
  color: var(--link);
}

.icon {
  width: 16px;
  height: 16px;
}

.about-content {
  margin-bottom: 48px;
}

.about-section {
  margin-bottom: 28px;
}

.about-section h2 {
  font-size: 18px;
  font-weight: 600;
  color: var(--text);
  margin: 0 0 12px;
}

.about-section p {
  color: var(--text-secondary);
  line-height: 1.8;
  margin: 0;
}

.tech-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.tech-list li {
  padding: 8px 0;
  color: var(--text-secondary);
}

@media (max-width: 600px) {
  .page-title {
    font-size: 24px;
  }

  .author-card {
    gap: 16px;
  }

  .author-avatar {
    width: 52px;
    height: 52px;
  }

  .avatar-text {
    font-size: 18px;
  }
}
</style>
