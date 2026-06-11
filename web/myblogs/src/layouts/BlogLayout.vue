<template>
  <div class="blog-layout">
    <NavBar />
    <main class="blog-main">
      <router-view />
    </main>
    <footer class="blog-footer">
      <p>&copy; 2026 HwBlogs</p>
    </footer>
  </div>
</template>

<script setup>
import NavBar from '@/components/blog/NavBar.vue'
</script>

<style scoped>
.blog-layout {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.blog-main {
  flex: 1;
  max-width: 900px;
  width: 100%;
  margin: 0 auto;
  padding: 40px 20px;
}

.blog-footer {
  border-top: 1px solid var(--border);
  padding: 24px 20px;
  text-align: center;
  color: var(--text-secondary);
  font-size: 14px;
}

.blog-footer p {
  margin: 0;
}
</style>
