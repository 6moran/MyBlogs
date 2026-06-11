<template>
  <div class="post-list">
    <RouterLink
      v-for="post in posts"
      :key="post.id"
      :to="`/posts/${post.id}`"
      class="post-card"
    >
      <div class="post-body">
        <h2 class="post-title">{{ post.title }}</h2>
        <p class="post-excerpt">{{ post.excerpt }}</p>
        <div class="post-meta">
          <span class="meta-date">{{ post.date }}</span>
          <span class="meta-likes">{{ post.likes }} 点赞</span>
          <span class="meta-comments">{{ post.comments }} 评论</span>
        </div>
      </div>
      <div v-if="post.cover" class="post-cover">
        <img :src="post.cover" :alt="post.title" loading="lazy" />
      </div>
    </RouterLink>
  </div>
</template>

<script setup>
defineProps({
  posts: {
    type: Array,
    required: true,
  },
})
</script>

<style scoped>
.post-list {
  display: flex;
  flex-direction: column;
}

.post-card {
  display: flex;
  gap: 20px;
  padding: 20px 0;
  align-items: center;
  border-bottom: 1px solid var(--border);
  text-decoration: none;
  color: var(--text);
  transition: color 0.2s;
  height: 180px;
}

.post-card:last-child {
  border-bottom: none;
}

.post-card:hover .post-title {
  color: var(--link);
}

.post-body {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  min-width: 0;
}

.post-title {
  margin: 0 0 12px;
  font-size: 18px;
  font-weight: 600;
  line-height: 1.4;
  color: var(--text);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;

}

.post-excerpt {
  flex: 1;
  margin: 0;
  font-size: 14px;
  line-height: 1.8;
  color: var(--text-secondary);
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.post-meta {
  margin-top: 12px;
  display: flex;
  gap: 16px;
  font-size: 13px;
  color: var(--text-muted);

}

.post-cover {
  flex-shrink: 0;
  width: 180px;
  height: 120px;
  border-radius: var(--radius-lg);
  overflow: hidden;
}

.post-cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

@media (max-width: 600px) {
  .post-card {
    flex-direction: column-reverse;
    min-height: auto;
    height: auto;
  }

  .post-cover {
    width: 100%;
    height: 160px;
  }

  .post-title {
    white-space: normal;
  }
}
</style>
