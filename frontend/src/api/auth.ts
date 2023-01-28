import axios from 'axios';
import type { GenericResponse, ILoginInput, ILoginResponse, IUserResponse, ISignUpInput, IReportResponse } from '@/api/types';

const BASE_URL = `http://${import.meta.env.VITE_BASE_HOST}/api/`;
const UNAUTHORIZED_CODES = [400, 401];

const authApi = axios.create({
  baseURL: BASE_URL,
  withCredentials: true,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
    Authorization: `Bearer ${localStorage.getItem('user_token')}`
  }
});

authApi.interceptors.response.use(
  response => {
    return response;
  },
  async error => {
    if (UNAUTHORIZED_CODES.includes(error.status)) {
      window.location.href = '/login';
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
  authApi.defaults.headers['Authorization'] = `Bearer ${response.data.access_token}`;
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

export const getReports = async (search: string, start_date: Date, end_date: Date, orderCol: string, order: string, page: Number, limit: Number) => {
  const response = await authApi.get<IReportResponse>('report', {
    params: {
      search: search,
      start_date: start_date,
      end_date: end_date,
      order_col: orderCol,
      order: order,
      page: page,
      limit: limit
    }
  });
  return response.data;
};
