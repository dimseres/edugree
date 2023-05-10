import { createRouter, createWebHistory, NavigationGuardNext, RouteLocationNormalized } from 'vue-router'
import { useToast } from 'vue-toastification'
import { AuthGuard } from '../services/guards/auth.guard'

export const router = createRouter({
    history: createWebHistory('/'),
    routes: [
        {
            path: '/login',
            redirect: '/auth/login',
        },
        {
            path: '/auth',
            name: 'Auth',
            redirect: '/auth/login',
            component: () => import('../views/Auth.vue'),
            children: [
                {
                    path: 'login',
                    name: 'Login',
                    component: () => import('../views/Login.vue'),
                },
                {
                    path: 'resetpassword',
                    name: 'PasswordReset',
                    component: () => import('../views/PasswordReset.vue'),
                },
                {
                    path: 'signin',
                    name: 'Registration',
                    component: () => import('../views/Registration.vue'),
                },
            ],
        },
        {
            path: '/',
            name: 'Home',
            component: () => import('../views/Home.vue'),
            meta: {
                guard: [new AuthGuard()],
            },
            children: [
                {
                    path: '/users',
                    name: 'Users',
                    meta: {
                        guard: [new AuthGuard(['owner', 'administrator', 'moderator'])],
                    },
                    component: () => import('../views/Users.vue'),
                },
                {
                    path: '/organizations',
                    name: 'Organizations',
                    component: () => import('../views/Organization.vue')
                },
                {
                    path: '/organizations/create',
                    name: 'OrganizationCreate',
                    component: () => import('../views/OrganizationCreate.vue')
                },
            ],
        },
        {
            path: '/choose',
            name: 'OrganizationChoose',
            component: () => import('../views/OrganizationChoose.vue'),
            meta: {
                guard: [new AuthGuard()],
            },
        },
        {
            path: '/:pathMatch(.*)*',
            name: '404',
            component: () => import('../views/404.vue'),
        },
    ],
})

export type GuardResponse = {
    success: boolean;
    message?: string;
    role_failed?: boolean;
    nextRouteName?: string;
};

export interface IRouteGuard {
    routeAllowed(route: RouteLocationNormalized): Promise<GuardResponse>;
}

declare module 'vue-router' {
    interface RouteMeta {
        // must be declared by every route
        guard?: Array<IRouteGuard>;
    }
}

router.beforeEach(async (to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
    const toasted = useToast()
    if (to.meta.guard) {
        for (const guard of to.meta.guard) {
            const { success, message, nextRouteName, role_failed } = await guard.routeAllowed(to)
            console.log(success, message, nextRouteName, role_failed)
            if (!success) {
                if (role_failed) {
                    toasted.error('недостаточно прав')
                    return next(from)
                }
                return next({ name: nextRouteName ?? 'Login' })
            }
        }
    }
    return next()
})