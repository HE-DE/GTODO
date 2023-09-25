import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import 'virtual:windi.css'
import axios from '~/plugins/axiosInterfaces.js'


const pinia = createPinia()
const app = createApp(App)

app.use(pinia)
app.use(router)
app.use(ElementPlus)


app.mount('#app')
app.config.globalProperties.$axios=axios