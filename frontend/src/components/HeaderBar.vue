<template>
  <header class="bg-gray-800 h-16">
    <nav class="h-full flex justify-between container items-center">
      <div>
        <RouterLink to="/" class="text-white font-semibold">
          <img src="/frontend/src/assets/logo.svg" class="h-12" />
        </RouterLink>
      </div>
      <ul class="flex items-center gap-4">
        <li v-for="(link, idx) in filteredLinks" :key="idx">
          <RouterLink :to="link.link"
            class="text-white rounded-md focus:bg-gray-600 transition duration-150 ease-in-out py-2 pt-1 px-4"
            :class="link.link == currentLink ? 'bg-gray-900' : ''">{{ link.name }}</RouterLink>
        </li>
        <li v-if="getAuthUser"
          class="cursor-pointer text-white rounded-md focus:bg-gray-600 transition duration-150 ease-in-out"
          @click="handleLogout">Logout</li>
      </ul>
    </nav>
  </header>
</template>

<script setup lang="ts">
import { useAuthStore } from '@/stores/auth.store';
import { logoutUserFn } from '@/api/auth';
import { useMutation } from 'vue-query';
import { createToast } from 'mosha-vue-toastify';
import router from '@/router';
import { computed, reactive, ref } from 'vue';
import { storeToRefs } from 'pinia';

const authStore = useAuthStore();

const { getAuthUser } = storeToRefs(authStore);

const { mutate: logoutUser } = useMutation(() => logoutUserFn(), {
  onSuccess: async () => {
    await authStore.setAuthUser(null);
    await authStore.setAuthToken(null);
    router.push('/login');
  },
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
  }
});

const handleLogout = () => {
  logoutUser();
};

const links = reactive([
  { name: 'Home', link: '/' },
  { name: 'Sign Up', link: '/register' },
  { name: 'Login', link: '/login' },
  { name: 'Reports', link: '/reports' }
]);

const currentLink = ref(router.currentRoute.value.path);

const restrictedLinks = new Set(['Reports']);
const anonymousLinks = new Set(['Sign Up', 'Login']);

const filteredLinks = computed(() => {
  return links.filter(link => (getAuthUser.value && restrictedLinks.has(link.name)) || (!getAuthUser.value && anonymousLinks.has(link.name)));
});
</script>
