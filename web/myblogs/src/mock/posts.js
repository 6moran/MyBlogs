export const posts = [
  {
    slug: 'vue3-composition-api',
    title: 'Vue 3 Composition API 实战指南',
    date: '2025-12-15',
    excerpt: 'Vue 3 引入的 Composition API 是一次重大的范式转变，它让我们能够更灵活地组织组件逻辑，通过 composables 函数实现逻辑的灵活组合。',
    cover: 'https://picsum.photos/seed/vue3/800/400',
    likes: 36,
    comments: 12,
    tags: ['vue', 'javascript', 'state-management'],
    content: `# Vue 3 Composition API 实战指南\n\nVue 3 引入的 Composition API 是一次重大的范式转变...\n\n## 基础用法\n\n\`\`\`javascript\nimport { ref, computed } from 'vue'\nconst count = ref(0)\n\`\`\``
  },
  {
    slug: 'go-gin-restful-api',
    title: '使用 Go + Gin 构建 RESTful API',
    date: '2025-11-20',
    excerpt: 'Gin 是 Go 语言中最流行的 Web 框架之一，以高性能和简洁的 API 著称，本文介绍如何用它构建生产级 RESTful 服务。',
    cover: 'https://picsum.photos/seed/golang/800/400',
    likes: 24,
    comments: 8,
    tags: ['go', 'gin', 'restful', 'api-design'],
    content: `# 使用 Go + Gin 构建 RESTful API\n\nGin 是 Go 语言中最流行的 Web 框架之一...\n\n\`\`\`go\nr := gin.Default()\nr.GET("/api/articles", func(c *gin.Context) {\n    c.JSON(200, gin.H{"message": "ok"})\n})\n\`\`\``
  },
  {
    slug: 'docker-multi-stage-build',
    title: 'Docker 多阶段构建最佳实践',
    date: '2025-10-08',
    excerpt: '多阶段构建是 Docker 17.05 引入的特性，它可以显著减小镜像体积，是现代 Docker 镜像构建的标准做法。',
    cover: '',
    likes: 15,
    comments: 5,
    tags: ['docker', 'linux'],
    content: `# Docker 多阶段构建最佳实践\n\n多阶段构建可以显著减小镜像体积。\n\n\`\`\`dockerfile\nFROM golang:1.21-alpine AS builder\nWORKDIR /app\nCOPY . .\nRUN go build -o server .\n\nFROM alpine:latest\nCOPY --from=builder /app/server .\nCMD ["./server"]\n\`\`\``
  },
  {
    slug: 'css-grid-layout',
    title: 'CSS Grid 布局完全指南',
    date: '2025-09-22',
    excerpt: 'CSS Grid 是一个二维布局系统，它彻底改变了我们构建网页布局的方式，配合 minmax() 和 auto-fill 可以实现完美的响应式布局。',
    cover: 'https://picsum.photos/seed/cssgrid/800/400',
    likes: 42,
    comments: 15,
    tags: ['css', 'responsive'],
    content: `# CSS Grid 布局完全指南\n\n\`\`\`css\n.container {\n  display: grid;\n  grid-template-columns: repeat(3, 1fr);\n  gap: 20px;\n}\n\`\`\``
  },
  {
    slug: 'git-workflow-tips',
    title: 'Git 工作流实用技巧',
    date: '2025-08-30',
    excerpt: '掌握 Git 的高级用法可以显著提升团队协作效率，包括分支策略、交互式 rebase 和 Commit 规范。',
    cover: '',
    likes: 8,
    comments: 3,
    tags: ['git'],
    content: `# Git 工作流实用技巧\n\n## 分支策略\n\n- main：生产环境\n- develop：开发分支\n- feature/*：功能分支\n\n## 实用命令\n\n\`\`\`bash\ngit rebase -i HEAD~3\ngit stash\ngit reset --soft HEAD~1\n\`\`\``
  },
  {
    slug: 'linux-command-line',
    title: 'Linux 命令行进阶技巧',
    date: '2025-08-10',
    excerpt: '熟练使用命令行可以大幅提升开发效率，本文涵盖管道重定向、文本处理三剑客和进程管理等实用技巧。',
    cover: 'https://picsum.photos/seed/linux/800/400',
    likes: 19,
    comments: 7,
    tags: ['linux', 'docker'],
    content: `# Linux 命令行进阶技巧\n\n## 管道与重定向\n\n\`\`\`bash\ncat app.log | grep "error" | wc -l\n\`\`\`\n\n## grep / sed / awk\n\n\`\`\`bash\ngrep -r "TODO" --include="*.go" .\nsed -i 's/old/new/g' file.txt\nawk '{print $1, $3}' data.txt\n\`\`\``
  },
  {
    slug: 'restful-api-design',
    title: 'RESTful API 设计最佳实践',
    date: '2025-07-18',
    excerpt: '良好的 API 设计是后端开发的基础，本文介绍 URL 设计、状态码使用、响应格式和分页设计等核心原则。',
    cover: '',
    likes: 28,
    comments: 9,
    tags: ['restful', 'api-design', 'go'],
    content: `# RESTful API 设计最佳实践\n\n## URL 设计\n\n\`\`\`\nGET    /api/articles\nGET    /api/articles/:id\nPOST   /api/articles\nPUT    /api/articles/:id\nDELETE /api/articles/:id\n\`\`\``
  },
  {
    slug: 'typescript-generics',
    title: 'TypeScript 泛型深入理解',
    date: '2025-06-25',
    excerpt: '泛型是 TypeScript 中最强大的特性之一，它让代码更加灵活和类型安全，掌握泛型是深入使用 TypeScript 的关键一步。',
    cover: 'https://picsum.photos/seed/typescript/800/400',
    likes: 17,
    comments: 6,
    tags: ['typescript', 'javascript', 'generics'],
    content: `# TypeScript 泛型深入理解\n\n\`\`\`typescript\nfunction identity<T>(arg: T): T {\n  return arg;\n}\n\`\`\`\n\n## 泛型约束\n\n\`\`\`typescript\ninterface HasLength { length: number }\nfunction logLength<T extends HasLength>(arg: T) {}\n\`\`\``
  },
  {
    slug: 'pinia-state-management',
    title: 'Pinia 状态管理实战',
    date: '2025-05-12',
    excerpt: 'Pinia 是 Vue 官方推荐的状态管理库，比 Vuex 更简洁更灵活，支持组合式 API 风格和状态持久化。',
    cover: '',
    likes: 11,
    comments: 4,
    tags: ['pinia', 'vue', 'state-management'],
    content: `# Pinia 状态管理实战\n\n\`\`\`typescript\nexport const useUserStore = defineStore('user', {\n  state: () => ({ name: '', token: '' }),\n  actions: {\n    async login(username, password) { /* ... */ }\n  }\n})\n\`\`\``
  },
  {
    slug: 'vue-router-guard',
    title: 'Vue Router 导航守卫详解',
    date: '2025-04-20',
    excerpt: '导航守卫是 Vue Router 中控制路由访问的核心机制，本文详解全局守卫、路由独享守卫和组件内守卫。',
    cover: 'https://picsum.photos/seed/vuerouter/800/400',
    likes: 33,
    comments: 11,
    tags: ['vue-router', 'vue', 'navigation-guards'],
    content: `# Vue Router 导航守卫详解\n\n\`\`\`javascript\nrouter.beforeEach((to, from, next) => {\n  if (to.meta.requiresAuth && !isAuthenticated()) {\n    next('/login')\n  } else {\n    next()\n  }\n})\n\`\`\``
  },
  {
    slug: 'element-plus-tips',
    title: 'Element Plus 高级使用技巧',
    date: '2025-03-15',
    excerpt: 'Element Plus 是 Vue 3 生态中最流行的 UI 组件库，本文介绍按需引入、自定义主题和表单验证等高级技巧。',
    cover: '',
    likes: 6,
    comments: 2,
    tags: ['element-plus', 'vue', 'javascript'],
    content: `# Element Plus 高级使用技巧\n\n## 按需引入\n\n\`\`\`javascript\nimport Components from 'unplugin-vue-components/vite'\nimport { ElementPlusResolver } from 'unplugin-vue-components/resolvers'\n\`\`\``
  },
  {
    slug: 'markdown-rendering-in-vue',
    title: '在 Vue 中渲染 Markdown 内容',
    date: '2025-02-28',
    excerpt: '在博客系统中 Markdown 渲染是核心功能，本文介绍如何使用 marked 库配合 highlight.js 实现语法高亮。',
    cover: 'https://picsum.photos/seed/markdown/800/400',
    likes: 29,
    comments: 10,
    tags: ['markdown', 'vue', 'highlight'],
    content: `# 在 Vue 中渲染 Markdown 内容\n\n\`\`\`javascript\nimport { marked } from 'marked'\nconst html = marked('# Hello World')\n\`\`\`\n\n## 安全注意事项\n\n使用 DOMPurify 过滤 XSS 风险。`
  }
]

export default posts
