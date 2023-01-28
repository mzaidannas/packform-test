<template>
  <section>
    <div class="container mx-auto my-10">
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
                <input type="email" placeholder="Email address" autofocus="true" autocomplete="email" v-model="email"
                  class="block w-full rounded-2xl appearance-none focus:outline-none py-2 px-4" id="email" />
                <span class="text-red-500 text-xs pt-1 block">{{ errors.email }}</span>
              </div>
              <div class="rounded-md shadow-sm">
                <input type="password" autocomplete="current-password" placeholder="Password" v-model="password"
                  class="block w-full rounded-2xl appearance-none focus:outline-none py-2 px-4" id="password" />
                <span class="text-red-500 text-xs pt-1 block">{{ errors.password }}</span>
              </div>
              <div class="rounded-md shadow-sm">
                <input type="text" autocomplete="name" placeholder="Name" v-model="name"
                  class="block w-full rounded-2xl appearance-none focus:outline-none py-2 px-4" id="name" />
                <span class="text-red-500 text-xs pt-1 block">{{ errors.name }}</span>
              </div>
              <div class="mt-8">
                <LoadingButton :loading="isLoading">
                  <span class="absolute left-0 inset-y-0 flex items-center pl-3">
                    <IconLock />
                  </span>
                  Sign Up
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
import { useField, useForm } from 'vee-validate';
import { toFormValidator } from '@vee-validate/zod';
import { z } from 'zod';
import { useMutation } from 'vue-query';
import { signUpUserFn } from '@/api/auth';
import type { ISignUpInput } from '@/api/types';
import { createToast } from 'mosha-vue-toastify';
import router from '@/router';
import { useAuthStore } from '@/stores/auth.store';
import LoadingButton from '@/components/buttons/LoadingButton.vue';
import IconLock from '@/components/icons/IconLock.vue';

const authStore = useAuthStore();

const registerSchema = toFormValidator(
  z.object({
    username: z.string().min(1, 'Username is required'),
    email: z.string().min(1, 'Email address is required').email('Email Address is invalid'),
    password: z
      .string()
      .min(1, 'Password is required')
      .min(8, 'Password must be more than 8 characters')
      .max(32, 'Password must be less than 32 characters'),
    name: z.string().min(1, 'Name is required').max(50, 'Name must be less than 50 characters')
  })
);

const { handleSubmit, errors, resetForm } = useForm({
  validationSchema: registerSchema
});

const { value: username } = useField('username');
const { value: email } = useField('email');
const { value: password } = useField('password');
const { value: name } = useField('name');

const { isLoading, mutate } = useMutation((credentials: ISignUpInput) => signUpUserFn(credentials), {
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
  onSuccess: () => {
    createToast('Successfully registered user', {
      position: 'top-right'
    });
    router.push('/login');
  }
});

const onSubmit = handleSubmit(values => {
  mutate({
    username: values.username,
    email: values.email,
    password: values.password,
    name: values.name
  });
  resetForm();
});
</script>
