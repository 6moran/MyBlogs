<template>
  <div class="comment-section">
    <div class="comment-header">
      <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>
      <span class="comment-title">评论区</span>
      <span class="comment-count">共 {{ allComments.length }} 条评论</span>
    </div>

    <!-- 登录提示 -->
    <div v-if="!isLoggedIn" class="login-card">
      <div class="login-card-content">
        <svg class="github-icon" viewBox="0 0 24 24" width="32" height="32">
          <path fill="currentColor" d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
        </svg>
        <h3>登录后即可评论</h3>
        <p>使用 GitHub 账号快速登录</p>
        <button class="github-login-btn" @click="handleGitHubLogin">
          <svg viewBox="0 0 24 24" width="18" height="18">
            <path fill="currentColor" d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
          </svg>
          使用 GitHub 登录
        </button>
      </div>
    </div>

    <!-- 默认评论输入框 -->
    <div v-if="isLoggedIn && !replyingTo" class="comment-input-section">
      <div class="current-user">
        <img v-if="userInfo?.avatar" :src="userInfo.avatar" class="user-avatar" />
        <div v-else class="user-avatar-placeholder">
          {{ userInfo?.nickname?.charAt(0) || 'U' }}
        </div>
        <span class="user-name">{{ userInfo?.nickname || userInfo?.username }}</span>
        <button class="logout-btn" @click="handleLogout">退出登录</button>
      </div>
      <div class="comment-input-wrapper">
        <textarea
          v-model="newComment"
          class="comment-input"
          placeholder="写下你的评论..."
          rows="3"
        ></textarea>
        <div class="comment-input-footer">
          <span class="comment-hint">支持 Markdown 语法</span>
          <button
            class="submit-btn"
            :disabled="!newComment.trim() || submitting"
            @click="handleSubmitComment"
          >
            {{ submitting ? '发送中...' : '发表评论' }}
          </button>
        </div>
      </div>
    </div>

    <!-- 评论列表 -->
    <div v-if="loading" class="loading">加载评论中...</div>
    <div v-else-if="groupedComments.length === 0" class="empty-comments">
      <p>暂无评论，来发表第一条评论吧</p>
    </div>
    <div v-else class="comment-list">
      <template v-for="group in groupedComments" :key="group.parent.id">
        <div class="comment-item">
          <img v-if="group.parent.avatar" :src="group.parent.avatar" class="comment-avatar" />
          <div v-else class="comment-avatar-placeholder">
            {{ group.parent.nickname?.charAt(0) || group.parent.username?.charAt(0) || 'U' }}
          </div>
          <div class="comment-body">
            <div class="comment-meta">
              <span class="comment-author">{{ group.parent.nickname || group.parent.username }}</span>
              <span v-if="group.parent.userID === userInfo?.id" class="owner-badge">博主</span>
            </div>
            <p class="comment-content">{{ group.parent.content }}</p>
            <div class="comment-footer">
              <time>{{ formatTime(group.parent.createdAt) }}</time>
              <span class="dot">·</span>
              <button v-if="isLoggedIn" class="reply-btn" @click="handleReply(group.parent)">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>
                回复
              </button>
            </div>
          </div>
        </div>

        <!-- 回复输入框 -->
        <div v-if="replyingTo?.id === group.parent.id" class="reply-input-section">
          <div class="reply-input-wrapper">
            <textarea
              v-model="newComment"
              class="comment-input"
              :placeholder="`回复 ${replyingTo.nickname || replyingTo.username}...`"
              rows="2"
            ></textarea>
            <div class="comment-input-footer">
              <button class="cancel-btn" @click="cancelReply">取消</button>
              <button
                class="submit-btn"
                :disabled="!newComment.trim() || submitting"
                @click="handleSubmitComment"
              >
                {{ submitting ? '发送中...' : '回复' }}
              </button>
            </div>
          </div>
        </div>

        <div v-if="group.replies.length" class="reply-group">
          <template v-for="reply in group.replies" :key="reply.id">
            <div class="reply-item">
              <img v-if="reply.avatar" :src="reply.avatar" class="comment-avatar small" />
              <div v-else class="comment-avatar-placeholder small">
                {{ reply.nickname?.charAt(0) || reply.username?.charAt(0) || 'U' }}
              </div>
              <div class="comment-body">
                <div class="comment-meta">
                  <span class="comment-author">{{ reply.nickname || reply.username }}</span>
                  <span v-if="reply.userID === userInfo?.id" class="owner-badge">博主</span>
                </div>
                <p class="comment-content">
                  <span v-if="reply.replyTo" class="reply-to">@{{ reply.replyTo }} </span>{{ reply.content }}
                </p>
                <div class="comment-footer">
                  <time>{{ formatTime(reply.createdAt) }}</time>
                  <span class="dot">·</span>
                  <button v-if="isLoggedIn" class="reply-btn" @click="handleReply(reply)">
                    <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>
                    回复
                  </button>
                </div>
              </div>
            </div>
            <div v-if="replyingTo?.id === reply.id" class="inline-reply-input">
              <textarea
                v-model="newComment"
                class="comment-input"
                :placeholder="`回复 ${replyingTo.nickname || replyingTo.username}...`"
                rows="2"
              ></textarea>
              <div class="comment-input-footer">
                <button class="cancel-btn" @click="cancelReply">取消</button>
                <button
                  class="submit-btn"
                  :disabled="!newComment.trim() || submitting"
                  @click="handleSubmitComment"
                >
                  {{ submitting ? '发送中...' : '回复' }}
                </button>
              </div>
            </div>
          </template>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, nextTick } from 'vue'
import { useUserStore } from '@/stores/user'
import { getComments, createComment } from '@/services/commentService'
import { getGitHubLoginURL } from '@/services/authService'

const props = defineProps({
  articleID: { type: Number, required: true }
})

const userStore = useUserStore()
const isLoggedIn = computed(() => userStore.isLoggedIn)
const userInfo = computed(() => userStore.userInfo)

const allComments = ref([])
const newComment = ref('')
const submitting = ref(false)
const replyingTo = ref(null)
const loading = ref(true)

// 从后端数据映射为前端格式
function mapComments(list) {
  return list.map(c => ({
    id: c.id,
    articleID: props.articleID,
    userID: c.user_id,
    username: c.username,
    nickname: c.nickname,
    avatar: c.avatar,
    content: c.content,
    createdAt: c.created_at,
    parentID: c.parent_id,
    rootID: c.root_id,
    replyTo: null
  }))
}

// 计算分组评论
const groupedComments = computed(() => {
  const parents = allComments.value.filter(c => c.parentID === 0)
  return parents.map(parent => {
    const replies = allComments.value.filter(c => c.parentID === parent.id && c.id !== parent.id)
    // 为回复填充 replyTo 字段
    replies.forEach(r => {
      const target = allComments.value.find(c => c.id === r.parentID)
      if (target) r.replyTo = target.nickname || target.username
    })
    return { parent, replies }
  })
})

onMounted(() => {
  loadComments()
})

async function loadComments() {
  loading.value = true
  try {
    const res = await getComments(props.articleID)
    const data = res.data
    allComments.value = mapComments(data.comments || [])
  } catch {
    allComments.value = []
  } finally {
    loading.value = false
  }
}

function handleGitHubLogin() {
  window.location.href = getGitHubLoginURL()
}

function handleReply(comment) {
  replyingTo.value = comment
  newComment.value = ''
  nextTick(() => {
    const inputs = document.querySelectorAll('.reply-input-section textarea, .inline-reply-input textarea')
    const last = inputs[inputs.length - 1]
    if (last) last.focus()
  })
}

function cancelReply() {
  replyingTo.value = null
  newComment.value = ''
}

function handleLogout() {
  userStore.logout()
  replyingTo.value = null
  newComment.value = ''
}

async function handleSubmitComment() {
  if (!newComment.value.trim() || submitting.value) return
  submitting.value = true
  try {
    const parentID = replyingTo.value ? replyingTo.value.parentID || replyingTo.value.id : 0
    await createComment(props.articleID, newComment.value, parentID)
    newComment.value = ''
    replyingTo.value = null
    await loadComments()
  } catch (e) {
    alert(e.message || '评论失败')
  } finally {
    submitting.value = false
  }
}

function formatTime(time) {
  if (!time) return ''
  const date = new Date(time)
  const now = new Date()
  const diff = now - date

  if (diff < 60000) return '刚刚'
  if (diff < 3600000) return `${Math.floor(diff / 60000)}分钟前`
  if (diff < 86400000) return `${Math.floor(diff / 3600000)}小时前`
  if (diff < 2592000000) return `${Math.floor(diff / 86400000)}天前`

  return date.toLocaleDateString('zh-CN')
}
</script>

<style scoped>
.comment-section {
  margin-top: 48px;
  padding: 24px 28px;
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  position: relative;
}

.comment-section::before {
  content: '';
  position: absolute;
  top: -25px;
  left: 0;
  right: 0;
  height: 1px;
  background: var(--border);
}

.comment-header {
  display: flex;
  align-items: center;
  gap: 8px;
  padding-bottom: 20px;
  border-bottom: 1px solid var(--border-light);
  margin-bottom: 4px;
  color: var(--text);
}

.comment-title {
  font-size: 16px;
  font-weight: 600;
}

.comment-count {
  font-size: 13px;
  color: var(--text-muted);
}

.login-card {
  margin: 16px 0;
  padding: 24px;
  background: var(--bg-secondary);
  border: 1px dashed var(--border);
  border-radius: 8px;
  text-align: center;
}

.login-card-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
}

.login-card .github-icon {
  color: #24292e;
  opacity: 0.8;
}

.login-card h3 {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: var(--text);
}

.login-card p {
  margin: 0;
  color: var(--text-secondary);
  font-size: 14px;
}

.github-login-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  background: #24292e;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  cursor: pointer;
  transition: background 0.2s;
}

.github-login-btn:hover {
  background: #1a1e22;
}

.comment-input-section {
  margin-bottom: 24px;
}

.current-user {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 12px;
}

.user-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  object-fit: cover;
}

.user-avatar-placeholder {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background: var(--link);
  color: white;
  display: grid;
  place-items: center;
  font-size: 14px;
  font-weight: 600;
}

.user-name {
  font-size: 14px;
  font-weight: 500;
  color: var(--text);
}

.logout-btn {
  margin-left: auto;
  padding: 4px 8px;
  background: transparent;
  color: var(--text-muted);
  border: 1px solid var(--border);
  border-radius: 4px;
  font-size: 12px;
  cursor: pointer;
  transition: all 0.2s;
}

.logout-btn:hover {
  color: var(--text);
  border-color: var(--text-muted);
}

.comment-input-wrapper,
.reply-input-wrapper {
  border: 1px solid var(--border);
  border-radius: 8px;
  overflow: hidden;
}

.reply-input-section {
  margin: 12px 0;
}

.reply-input-section .reply-input-wrapper {
  width: 100%;
}

.inline-reply-input {
  margin-top: 12px;
  border: 1px solid var(--border);
  border-radius: 8px;
  overflow: hidden;
}

.reply-item:last-child {
  border-bottom: none;
}

.comment-input {
  width: 100%;
  padding: 12px;
  border: none;
  resize: vertical;
  font-size: 14px;
  line-height: 1.6;
  color: var(--text);
  background: var(--bg);
  box-sizing: border-box;
}

.comment-input:focus {
  outline: none;
}

.comment-input::placeholder {
  color: var(--text-muted);
}

.comment-input-footer {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  background: var(--bg-secondary);
  border-top: 1px solid var(--border-light);
}

.comment-hint {
  margin-right: auto;
  font-size: 12px;
  color: var(--text-muted);
}

.cancel-btn {
  padding: 6px 12px;
  background: transparent;
  color: var(--text-muted);
  border: 1px solid var(--border);
  border-radius: 4px;
  font-size: 13px;
  cursor: pointer;
  transition: all 0.2s;
}

.cancel-btn:hover {
  color: var(--text);
  border-color: var(--text-muted);
}

.submit-btn {
  padding: 6px 16px;
  background: var(--link);
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 13px;
  cursor: pointer;
  transition: opacity 0.2s;
}

.submit-btn:hover:not(:disabled) {
  opacity: 0.9;
}

.submit-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.empty-comments {
  padding: 40px 0;
  text-align: center;
  color: var(--text-muted);
  font-size: 14px;
}

.loading {
  text-align: center;
  padding: 20px 0;
  color: var(--text-muted);
}

.comment-list {
  display: flex;
  flex-direction: column;
}

.comment-item {
  display: flex;
  gap: 12px;
  padding: 20px 0 12px;
}

.comment-avatar {
  flex-shrink: 0;
  width: 36px;
  height: 36px;
  border-radius: 50%;
  object-fit: cover;
}

.comment-avatar.small {
  width: 30px;
  height: 30px;
}

.comment-avatar-placeholder {
  flex-shrink: 0;
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background: var(--link);
  color: white;
  display: grid;
  place-items: center;
  font-size: 14px;
  font-weight: 600;
}

.comment-avatar-placeholder.small {
  width: 30px;
  height: 30px;
  font-size: 12px;
}

.comment-body {
  flex: 1;
  min-width: 0;
}

.comment-meta {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 4px;
}

.comment-author {
  font-size: 14px;
  font-weight: 600;
  color: var(--text);
}

.owner-badge {
  font-size: 11px;
  padding: 1px 6px;
  border-radius: var(--radius);
  background: var(--link);
  color: #fff;
  line-height: 1.5;
}

.comment-content {
  margin: 0 0 6px;
  font-size: 14px;
  line-height: 1.7;
  color: var(--text);
  white-space: pre-wrap;
}

.reply-to {
  color: var(--link);
}

.comment-footer {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: var(--text-muted);
}

.dot {
  color: var(--border);
}

.reply-btn {
  display: inline-flex;
  align-items: center;
  gap: 3px;
  padding: 0;
  background: none;
  border: none;
  font-size: 13px;
  color: var(--text-muted);
  cursor: pointer;
  transition: color 0.2s;
}

.reply-btn:hover {
  color: var(--link);
}

.reply-group {
  margin: 0 0 0 48px;
  padding-left: 14px;
  border-left: 2px solid var(--border-light);
}

.reply-item {
  display: flex;
  gap: 10px;
  padding: 14px 0;
}
</style>
