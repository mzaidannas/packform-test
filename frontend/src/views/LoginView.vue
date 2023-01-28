<template>
  <section>
    <div class="container mx-auto">
      <div class="min-h-screen flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
        <div class="max-w-md w-full space-y-8">
          <div>
            <img src="/frontend/src/assets/logo.svg" class="mx-auto h-12 w-auto" />
          </div>
          <div class="mt-8 space-y-6 text-white">
            <form @submit="onSubmit">
              <div class="rounded-md shadow-sm">
                <input type="text" placeholder="Username" autofocus="true" autocomplete="username" v-model="username"
                  class="block w-full rounded-2xl appearance-none focus:outline-none py-2 px-4" id="username" />
                <span class="text-red-500 text-xs pt-1 block">{{ errors.username }}</span>
              </div>
              <div class="rounded-md shadow-sm">
                <input v-model="password" type="password" autocomplete="current-password" placeholder="Password"
                  class="block w-full rounded-2xl appearance-none focus:outline-none py-2 px-4" id="password" />
                <span class="text-red-500 text-xs pt-1 block">{{ errors.password }}</span>
              </div>
              <div class="mt-8">
                <LoadingButton :loading="isLoading">
                  <span class="absolute left-0 inset-y-0 flex items-center pl-3">
                    <IconLock />
                  </span>
                  Sign In
                </LoadingButton>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { onBeforeUpdate } from 'vue';
import { Form, useField, useForm } from 'vee-validate';
import { toFormValidator } from '@vee-validate/zod';
import { z } from 'zod';
import { useMutation, useQuery, useQueryClient } from 'vue-query';
import { getMeFn, loginUserFn } from '@/api/auth';
import type { ILoginInput, ILoginResponse, IUserResponse } from '@/api/types';
import { createToast } from 'mosha-vue-toastify';
import router from '@/router';
import { useAuthStore } from '@/stores/auth.store';
import LoadingButton from '@/components/buttons/LoadingButton.vue';
import IconLock from '@/components/icons/IconLock.vue';

const authStore = useAuthStore();

const loginSchema = toFormValidator(
  z.object({
    username: z.string().min(1, 'Username is required').max(50, 'Username should be less than 50 characters'),
    password: z
      .string()
      .min(1, 'Password is required')
      .min(8, 'Password must be more than 8 characters')
      .max(32, 'Password must be less than 32 characters')
  })
);

const { handleSubmit, errors, resetForm } = useForm({
  validationSchema: loginSchema
});

const { value: username } = useField('username');
const { value: password } = useField('password');

const authResult = useQuery('authUser', () => getMeFn(), {
  enabled: false,
  retry: 1,
  onError: error => {
    if (Array.isArray((error as any).response.data.error)) {
      (error as any).response.data.error.forEach((el: any) =>
        createToast(el.message, {
          position: 'top-right',
          type: 'warning'
        })
      );
    } else {
      createToast((error as any).response.data.message, {
        position: 'top-right',
        type: 'danger'
      });
    }
  },
  onSuccess: async (response: IUserResponse) => {
    await authStore.setAuthUser(response.data);
  }
});

const queryClient = useQueryClient();

const { isLoading, mutate } = useMutation((credentials: ILoginInput) => loginUserFn(credentials), {
  onError: error => {
    if (Array.isArray((error as any).response.data.error)) {
      (error as any).response.data.error.forEach((el: any) =>
        createToast(el.message, {
          position: 'top-right',
          type: 'warning'
        })
      );
    } else {
      createToast((error as any).response.data.message, {
        position: 'top-right',
        type: 'danger'
      });
    }
  },
  onSuccess: async data => {
    createToast('Successfully logged in', {
      position: 'top-right'
    });
    await authStore.setAuthToken(data.access_token);
    await queryClient.refetchQueries('authUser');
    await router.push({ name: 'reports' });
  }
});

const onSubmit = handleSubmit(values => {
  mutate({
    username: values.username,
    password: values.password
  });
  resetForm();
});

onBeforeUpdate(async () => {
  if (authResult.isSuccess.value) {
    const authUser = Object.assign({}, authResult.data.value?.data);
    await authStore.setAuthUser(authUser);
  }
});
</script>
