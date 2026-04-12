<script setup lang="ts">
import AppLoader from "@/components/common/AppLoader.vue";
import BaseContextMenu from "@/components/common/BaseContextMenu.vue";
import BaseConfirm from "@/components/modals/BaseConfirm.vue";
import BasePrompt from "@/components/modals/BasePrompt.vue";
import { useConnectionStore } from "@/stores/connection.store";
import { onMounted, onUnmounted } from "vue";
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

function handleDragAndDrop(e: DragEvent): void {
  e.preventDefault();
}

window.addEventListener("keydown", handleKeydown, true);
window.addEventListener("contextmenu", handleContextMenu);
window.addEventListener("dragover", handleDragAndDrop, false);
window.addEventListener("drop", handleDragAndDrop, false);

onMounted(function () {
  useConnectionStore().getConnections();
  useConnectionStore().getActiveConnections();
});

onUnmounted(function (): void {
  window.removeEventListener("keydown", handleKeydown, true);
  window.removeEventListener("contextmenu", handleContextMenu);
  window.removeEventListener("dragover", handleDragAndDrop, false);
  window.removeEventListener("drop", handleDragAndDrop, false);
});
</script>

<template>
  <AppLoader />
  <main class="flex-1 overflow-hidden flex flex-col">
    <RouterView />
  </main>
  <BasePrompt />
  <BaseContextMenu />
  <BaseConfirm />
</template>
