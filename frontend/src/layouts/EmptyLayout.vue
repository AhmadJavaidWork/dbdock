<script setup lang="ts">
import AppLoader from "@/components/common/AppLoader.vue";
</script>

<template>
  <div
    class="max-h-screen flex items-center justify-center dark:bg-background-dark text-text-light dark:text-text-dark"
  >
    <RouterView />
    <AppLoader />
  </div>
</template>
