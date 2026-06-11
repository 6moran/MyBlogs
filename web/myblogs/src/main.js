import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import './styles/blog.css'
import './styles/admin.css'

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.mount('#app')
