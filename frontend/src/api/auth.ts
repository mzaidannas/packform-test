import axios from 'axios';
import type { GenericResponse, ILoginInput, ILoginResponse, IUserResponse, ISignUpInput } from '@/api/types';

const BASE_URL = 'http://localhost:8000/api/';

const authApi = axios.create({
  baseURL: BASE_URL,
  withCredentials: true
});

authApi.defaults.headers.common['Content-Type'] = 'application/json';

export const refreshAccessTokenFn = async () => {
  const response = await authApi.get<ILoginResponse>('auth/refresh');
  return response.data;
};

authApi.interceptors.response.use(
  response => {
    return response;
  },
  async error => {
    const originalRequest = error.config;
    const errMessage = error.response.data.message as string;
    if (errMessage.includes('not logged in') && !originalRequest._retry) {
      originalRequest._retry = true;
      await refreshAccessTokenFn();
      return authApi(originalRequest);
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
