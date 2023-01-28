import { createRouter, createWebHistory, type NavigationGuardNext, type RouteLocationNormalized } from 'vue-router';
import HomeView from '@/views/HomeView.vue';
import { useAuthStore } from '@/stores/auth.store';
import requireAuth from '@/router/middleware/requireAuth';
import middlewarePipeline from './middlewarePipeline';

const routes = [
  {
    name: 'home',
    path: '/',
    component: HomeView
  },
  {
    name: 'register',
    path: '/register',
    component: () => import('@/views/RegisterView.vue')
  },
  {
    name: 'login',
    path: '/login',
    component: () => import('@/views/LoginView.vue')
  },
  {
    name: 'reports',
    path: '/reports',
    component: () => import('@/views/ReportsView.vue'),
    meta: {
      middleware: [requireAuth]
    }
  }
];

const router = createRouter({
  history: createWebHistory('/'),
  routes
});

router.beforeEach((to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
  const authStore = useAuthStore();

  if (!to.meta.middleware) {
    return next();
  }
  const middleware = to.meta.middleware as any;

  const context = {
    to,
    from,
    next,
    authStore
  };

  return middleware[0]({
    ...context,
    next: middlewarePipeline(context, middleware, 1)
  });
});

export default router;
