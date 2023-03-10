import {createApp} from 'vue'
import './styles/style.scss'
import App from './App.vue'
import {createPinia} from 'pinia'
import {createRouter, createWebHistory} from 'vue-router'


const pinia = createPinia()
const router = createRouter({
    history: createWebHistory('/'),
    routes: [
        {
            path: "/login",
            redirect: "/auth/login"
        },
        {
            path: "/auth",
            name: "Auth",
            redirect: "/auth/login",
            component: () => import('./views/Auth.vue'),
            children: [
                {
                    path: "login",
                    name: "Login",
                    component: () => import('./views/Login.vue'),
                },
                {
                    path: "resetpassword",
                    name: "PasswordReset",
                    component: () => import('./views/PasswordReset.vue'),
                },
            ],
        },
        {
            path: "/",
            redirect: '/monitoring',
            name: "Home",
            component: () => import('./views/Home.vue'),
            children: [
                {
                    path: 'monitoring',
                    name: 'Monitoring',
                    component: () => import('./views/Monitoring.vue'),
                }
            ]
        },
        {
            path: "/:pathMatch(.*)*",
            name: "404",
            component: () => import('./views/404.vue'),
        }
    ]
})

const app = createApp(App)
app.use(pinia)
app.use(router)

app.mount('#app')
