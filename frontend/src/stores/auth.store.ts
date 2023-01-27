import type { IUser } from '@/api/types';
import { defineStore } from 'pinia';

export type AuthStoreState = {
  authUser: IUser | null;
};

export const useAuthStore = defineStore({
  id: 'authStore',
  state: () =>
    ({
      authUser: JSON.parse(localStorage.getItem('user') || '')
    } as AuthStoreState),
  getters: {},
  actions: {
    setAuthUser(user: IUser | null) {
      this.authUser = user;
      // store user details and jwt in local storage to keep user logged in between page refreshes
      localStorage.setItem('user', JSON.stringify(this.authUser));
    }
  }
});
