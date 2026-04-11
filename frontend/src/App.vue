<script setup lang="ts">
import AppLoader from "@/components/common/AppLoader.vue";
import BaseContextMenu from "@/components/common/BaseContextMenu.vue";
import BasePrompt from "@/components/modals/BasePrompt.vue";
import { onUnmounted } from "vue";
import { Reload } from "~/wailsjs/go/main/App";

function handleKeydown(e: KeyboardEvent): void {
  if (e.ctrlKey && e.key.toLowerCase() === "r") {
    e.stopImmediatePropagation();
    Reload();
  }
}

function handleContextMenu(e: MouseEvent): void {
  e.preventDefault();
}

window.addEventListener("keydown", handleKeydown, true);
window.addEventListener("contextmenu", handleContextMenu);

onUnmounted(function (): void {
  window.removeEventListener("keydown", handleKeydown, true);
  window.removeEventListener("contextmenu", handleContextMenu);
});
</script>

<template>
  <AppLoader />
  <main class="flex-1 overflow-hidden flex flex-col">
    <RouterView />
  </main>
  <BasePrompt />
  <BaseContextMenu />
</template>
