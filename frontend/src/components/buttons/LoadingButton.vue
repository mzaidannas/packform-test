<template>
  <button type="submit"
    class="group relative w-full flex justify-center py-2 px-4 border border-transparent font-semibold rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2"
    :disabled="loading" :class="[loading ? 'bg-blue-300' : '']">
    <div v-if="loading" class="flex items-center gap-3">
      <IconSpinner />
      <span class="text-white">Loading...</span>
    </div>
    <span v-else>
      <slot></slot>
    </span>
  </button>
</template>

<script setup lang="ts">
import IconSpinner from '@/components/icons/IconSpinner.vue';
defineProps<{
  loading: boolean;
}>();
</script>
