import { getMeFn } from '@/api/auth';
import type { NavigationGuardNext } from 'vue-router';
import router from '..';

export default async function requireAuth({ next, authStore }: { next: NavigationGuardNext; authStore: any }) {
  try {
    const response = await getMeFn();
    const user = response.data;
    await authStore.setAuthUser(user);

    if (!user) {
      return next({
        name: 'login'
      });
    }
  } catch (error) {
    router.push('/login');
  }

  return next();
}
