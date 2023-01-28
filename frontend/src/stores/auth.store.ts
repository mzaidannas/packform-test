import type { IUser } from '@/api/types';
import { defineStore } from 'pinia';

export type AuthStoreState = {
  authUser: IUser | null;
  authToken: string | null;
};

export const useAuthStore = defineStore({
  id: 'authStore',
  state: (): AuthStoreState => ({
    authUser: JSON.parse(localStorage.getItem('user') || 'null'),
    authToken: localStorage.getItem('user_token') || null
  }),
  getters: {
    getAuthUser(): IUser | null {
      return this.authUser;
    },
    getAuthToken(): string | null {
      return this.authToken;
    }
  },
  actions: {
    async setAuthUser(user: IUser | null) {
      this.authUser = user;
      // store user details and jwt in local storage to keep user logged in between page refreshes
      localStorage.setItem('user', JSON.stringify(this.authUser));
    },
    async setAuthToken(token: string | null) {
      this.authToken = token;
      localStorage.setItem('user_token', this.authToken || '');
    }
  }
});
