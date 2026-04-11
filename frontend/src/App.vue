<script setup lang="ts">
import AppLoader from "@/components/common/AppLoader.vue";
import BasePrompt from "@/components/modals/BasePrompt.vue";
import { onUnmounted } from "vue";
import { Reload } from "~/wailsjs/go/main/App";

function handleKeydown(e: KeyboardEvent) {
  if (e.ctrlKey && e.key.toLowerCase() === "r") {
    e.stopImmediatePropagation();
    Reload();
  }
}

window.addEventListener("keydown", handleKeydown, true);

onUnmounted(function () {
  window.addEventListener("keydown", handleKeydown, true);
});
</script>

<template>
  <AppLoader />
  <main class="flex-1 overflow-hidden flex flex-col">
    <RouterView />
  </main>
  <BasePrompt />
</template>
