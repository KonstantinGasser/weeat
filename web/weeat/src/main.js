import { createApp } from 'vue'
import App from './App.vue'
import moshaToast from 'mosha-vue-toastify'
import 'mosha-vue-toastify/dist/style.css'
import router from './router'

createApp(App).use(router).use(moshaToast).mount('#app')
 