export async function saveArticle(articleData) {
  // eslint-disable-next-line no-console
  console.log('模拟发送到后端的数据:', articleData)

  return new Promise((resolve, reject) => {
    setTimeout(() => {
      const shouldFail = false

      if (shouldFail) {
        reject(new Error('模拟请求失败，请稍后重试'))
        return
      }

      resolve({
        success: true,
        message: '操作成功',
        articleId: 'mock_article_id_123'
      })
    }, 2000)
  })
}

export async function uploadImage(file, maxRetries = 2) {
  if (!file || !file.type.startsWith('image/')) {
    throw new Error('不支持的文件类型，请上传图片文件')
  }

  const allowedTypes = ['image/jpeg', 'image/png', 'image/gif', 'image/webp', 'image/svg+xml']
  if (!allowedTypes.includes(file.type)) {
    throw new Error('不支持的图片格式，请上传 JPG、PNG、GIF、WebP 或 SVG 文件')
  }

  if (file.size > 10 * 1024 * 1024) {
    throw new Error('图片大小不能超过 10MB')
  }

  const formData = new FormData()
  formData.append('image', file, file.name)

  let lastError = null

  for (let attempt = 0; attempt <= maxRetries; attempt++) {
    try {
      const response = await fetch('/api/articles/images', {
        method: 'POST',
        body: formData
      })

      if (!response.ok) {
        const errorText = await response.text().catch(() => '')
        throw new Error(errorText || `上传失败（HTTP ${response.status}）`)
      }

      const result = await response.json()

      if (!result.url && !result.path && !result.data) {
        throw new Error('服务器未返回图片路径')
      }

      return result.url || result.path || result.data
    } catch (error) {
      lastError = error
      if (attempt < maxRetries) {
        await new Promise(resolve => setTimeout(resolve, 1000 * (attempt + 1)))
      }
    }
  }

  throw lastError || new Error('图片上传失败，请稍后重试')
}
