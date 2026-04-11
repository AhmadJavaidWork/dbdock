<script setup lang="ts">
import { useContextMenu } from "@/composables/useContextMenu";
import { watch } from "vue";

const { show, options, onSelect, onClose } = useContextMenu();

watch(show, function (): void {
  if (show.value) {
    window.addEventListener("click", onClose);
  } else {
    window.removeEventListener("click", onClose);
  }
});
</script>

<template>
  <div
    v-show="show"
    :style="{ top: `${options.position.y}px`, left: `${options.position.x}px` }"
    class="fixed bg-white dark:bg-gray-800 shadow-lg rounded z-50 w-[300px]"
  >
    <div
      v-for="option in options.options"
      :key="`context-menu-options-${option}`"
      :class="[
        'py-2 px-5 flex-wrap cursor-pointer font-medium text-select-options-text-light focus-visible:bg-select-options-background-hovered-light hover:bg-select-options-background-hovered-light border-b border-textfield-border-light dark:text-select-options-text-dark dark:focus-visible:bg-select-options-background-hovered-dark dark:hover:bg-select-options-background-hovered-dark dark:border-textfield-border-dark last:border-b-0',
      ]"
      @click="onSelect(option)"
    >
      {{ option }}
    </div>
  </div>
</template>
