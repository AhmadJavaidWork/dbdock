<script setup lang="ts">
import AppLoader from "@/components/common/AppLoader.vue";
import BasePrompt from "@/components/modals/BasePrompt.vue";
</script>

<template>
  <AppLoader />
  <main class="flex-1 overflow-hidden flex flex-col">
    <RouterView />
  </main>
  <BasePrompt />
</template>
