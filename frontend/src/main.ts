import { createApp } from 'vue'
import { createPinia } from 'pinia'
const pinia = createPinia()
const app = createApp(App)
app.use(pinia)

import { router } from './routes/routes'

import './styles/style.scss'
import App from './App.vue'


import Toast from 'vue-toastification'
import 'vue-toastification/dist/index.css'
import 'vue-advanced-cropper/dist/style.css';
import { usePreloaderStore } from './store/preloader.store'






app.use(router)

app.use(Toast, {
  position: 'bottom-right',
})

app.mount('#app')
