<script setup lang="ts">
import BaseBorderButton from '@/components/buttons/BaseBorderButton.vue';
import { ref } from 'vue';
import { Greet } from '~/wailsjs/go/main/App';

const name = ref<string>('');
const greeting = ref<string>('');

function greet() {
  Greet(name.value).then((result) => {
    greeting.value = result;
  });
}
</script>

<template>
  <div
    class="p-6 rounded-lg-background-light dark:bg-background-dark text-text-light dark:text-text-dark"
  >
    <h1 class="text-3xl font-bold mb-4">Welcome to DBDock</h1>
    <BaseBorderButton>Click me</BaseBorderButton>
    <BaseBorderButton variant="danger">Click me</BaseBorderButton>
    <p class="mt-4 text-text-secondary dark:text-secondary-dark">
      This text changes automatically in dark mode.
    </p>
  </div>
</template>
