import { createRouter, createWebHistory, NavigationGuardNext, RouteLocationNormalized } from 'vue-router'
import { useToast } from 'vue-toastification'

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
      redirect: '/monitoring',
      name: 'Home',
      component: () => import('../views/Home.vue'),
      children: [
        {
          path: 'monitoring',
          name: 'Monitoring',
          component: () => import('../views/Monitoring.vue'),
        },
      ],
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
      return next()
    }
  }
  return next()
})