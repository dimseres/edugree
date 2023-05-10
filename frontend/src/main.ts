import { createApp } from 'vue'
import './styles/style.scss'
import App from './App.vue'
import { createPinia } from 'pinia'
import { router } from './routes/routes'
import Toast from 'vue-toastification'
import 'vue-toastification/dist/index.css'
import 'vue-advanced-cropper/dist/style.css';


const pinia = createPinia()


const app = createApp(App)
app.use(pinia)
app.use(router)

app.use(Toast, {
  position: 'bottom-right',
})

app.mount('#app')
