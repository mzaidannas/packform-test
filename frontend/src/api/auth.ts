import axios from 'axios';
import type { GenericResponse, ILoginInput, ILoginResponse, IUserResponse, ISignUpInput } from '@/api/types';

const BASE_URL = `http://${import.meta.env.VITE_BASE_HOST}/api/`;

const authApi = axios.create({
  baseURL: BASE_URL,
  withCredentials: true,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
    Authorization: `Bearer ${localStorage.getItem('user_token')}`
  }
});

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
  const response = await authApi.delete<GenericResponse>('auth/logout');
  return response.data;
};

export const getMeFn = async () => {
  const response = await authApi.get<IUserResponse>('user');
  return response.data;
};
