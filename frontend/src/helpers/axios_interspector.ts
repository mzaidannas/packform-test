import axios from 'axios';
import { useAuthStore } from '@/stores/auth.store';
import router from '@/router/index';

export default {
  install: app => {
    const authStore = useAuthStore();
    app.config.globalProperties.$http = axios;

    axios.defaults.headers.common['Accept'] = 'application/json';
    axios.defaults.headers.common['Content-Type'] = 'application/json';
    axios.defaults.headers.common['Authorization'] = `Bearer ${authStore.authUser}`;

    app.provide('http', app.config.globalProperties.$http);

    axios.interceptors.request.use(request => {
      request.headers.set('Authorization', `Bearer ${authStore.authUser}`);
      return request;
    });

    axios.interceptors.response.use(
      response => {
        return response;
      },
      async error => {
        const errCode = error.response.status as number;
        if ([400, 401].includes(errCode)) {
          router.push('/login');
        }
        return Promise.reject(error);
      }
    );
  }
};
