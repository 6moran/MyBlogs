# MyBlogs 后台新建文章页面

这是一个基于 Vue 3 的后台管理子模块示例，当前专注实现“新建文章”功能。

## 功能概览

- 独立的后台管理页面，入口路由为 `/admin/articles/new`
- 文章标题输入区
- 支持富文本与 Markdown 双模式切换的内容编辑器
- Markdown 模式下实时预览
- “保存草稿”和“发布文章”两个操作按钮
- 前端模拟 `saveArticle` API 调用逻辑，包含 2 秒延迟与控制台日志输出
- 页面内成功/失败反馈提示

## 运行方式

```bash
npm install
npm run serve
```

本地开发服务启动后，打开浏览器访问：

```text
http://localhost:8080/admin/articles/new
```

如果项目已经安装过依赖，也可以直接执行：

```bash
npm run serve
```

## 模拟 API 调用说明

模拟 API 位于 `src/services/articleService.js`，核心函数如下：

- `saveArticle(articleData)`

这个函数会：

1. 接收文章数据对象，例如标题、内容、状态和更新时间
2. 使用 `console.log` 打印即将发送到后端的数据
3. 通过 `setTimeout` 模拟 2 秒网络请求延迟
4. 返回一个模拟成功响应：

```javascript
{
  success: true,
  message: '操作成功',
  articleId: 'mock_article_id_123'
}
```

当用户点击：

- “保存草稿”时，会以 `status: 'draft'` 调用 `saveArticle`
- “发布文章”时，会以 `status: 'published'` 调用 `saveArticle`

## 编辑器实现说明

- 富文本模式：基于 `@tiptap/vue-3`
- Markdown 同步：基于 `tiptap-markdown`
- Markdown 实时预览：基于 `marked`

页面内部以 Markdown 内容作为两种编辑模式之间的共享数据格式，因此切换模式时内容会尽量保持一致。

## 目录说明

```text
src/
  components/
    ArticleEditor.vue
    MarkdownWorkspace.vue
    RichTextEditor.vue
  services/
    articleService.js
  views/
    AdminArticleCreateView.vue
  router/
    index.js
  App.vue
  main.js
```

## 可扩展方向

- 接入真实后端 API
- 增加文章分类、标签、封面图字段
- 增加自动保存
- 增加草稿列表和文章管理页
