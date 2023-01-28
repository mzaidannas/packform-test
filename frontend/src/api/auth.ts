import axios from 'axios';
import type { GenericResponse, ILoginInput, ILoginResponse, IUserResponse, ISignUpInput } from '@/api/types';
// import { useAuthStore } from '@/stores/auth.store';
import router from '@/router';

const BASE_URL = `http://${import.meta.env.VITE_BASE_HOST}/api/`;
// const authStore = useAuthStore();

const authApi = axios.create({
  baseURL: BASE_URL,
  withCredentials: true
});

authApi.defaults.headers.common['Accept'] = 'application/json';
authApi.defaults.headers.common['Content-Type'] = 'application/json';

// export const refreshAccessTokenFn = async () => {
//   const response = await authApi.get<ILoginResponse>('auth/refresh');
//   return response.data;
// };

// authApi.interceptors.request.use(request => {
//   request.headers.set('Authorization', 'Bearer ' + authStore.authToken);
//   return request;
// });

authApi.interceptors.response.use(
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

export const signUpUserFn = async (user: ISignUpInput) => {
  const response = await authApi.post<GenericResponse>('auth/register', user);
  return response.data;
};

export const loginUserFn = async (user: ILoginInput) => {
  const response = await authApi.post<ILoginResponse>('auth/login', user);
  authApi.defaults.headers.common['Authorization'] = `Bearer ${response.data.access_token}`;
  return response.data;
};

export const logoutUserFn = async () => {
  const response = await authApi.get<GenericResponse>('auth/logout');
  return response.data;
};

export const getMeFn = async () => {
  const response = await authApi.get<IUserResponse>('user');
  return response.data;
};
