<script setup lang="ts">
import { useLoaderStore } from "@/stores/loader";
</script>

<template>
  <Transition name="fade">
    <div v-if="useLoaderStore().isLoading" class="fixed inset-0 z-[9999]">
      <div
        class="absolute inset-0 flex items-center justify-center z-[9999] bg-black/30 backdrop-blur-sm"
      >
        <div
          class="w-10 h-10 border-4 border-primary border-t-transparent rounded-full animate-spin"
        ></div>
      </div>
    </div>
  </Transition>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
