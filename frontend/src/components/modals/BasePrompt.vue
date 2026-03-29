<script setup lang="ts">
import IconAlertCircle from "@/components/icons/IconAlertCircle.vue";
import IconCheckCircle from "@/components/icons/IconCheckCircle.vue";
import IconCloseCircle from "@/components/icons/IconCloseCircle.vue";
import IconInfoCircle from "@/components/icons/IconInfoCircle.vue";
import { usePrompt } from "@/composables/usePrompt";
import { watch } from "vue";

const { isOpen, options, onClose } = usePrompt();

const handleKeydown = (e: KeyboardEvent) => {
  console.log(e.key);
  if (e.key === "Escape" && isOpen.value) {
    onClose();
  }
};

watch(isOpen, function (newVal) {
  if (newVal) {
    window.addEventListener("keydown", handleKeydown);
  } else {
    window.removeEventListener("keydown", handleKeydown);
  }
});
</script>

<template>
  <div
    v-if="isOpen"
    class="md:flex fixed inset-0 flex items-center justify-center bg-black/50 z-[1000]"
    @click="onClose"
  >
    <div
      class="text-text-light bg-background-light dark:bg-background-dark dark:text-text-dark rounded-prompt shadow-lg max-w-sm"
      @click.stop
    >
      <div class="p-4">
        <h3 class="text-lg font-semibold">{{ options.title }}</h3>
      </div>
      <div
        class="p-4 border-t border-t-gray-light border-b border-b-gray-light flex flex-col items-center gap-[15px]"
      >
        <IconCheckCircle
          v-if="options.type === 'success'"
          class="text-text-success dark:text-text-success-darker w-[44px] h-[44px]"
        />
        <IconInfoCircle
          v-else-if="options.type === 'info'"
          class="text-text-info dark:text-text-info-darker w-[44px] h-[44px]"
        />
        <IconAlertCircle
          v-else-if="options.type === 'warn'"
          class="text-text-warning dark:text-text-warning-darker w-[44px] h-[44px]"
        />
        <IconCloseCircle
          v-else-if="options.type === 'error'"
          class="text-text-danger dark:text-text-danger-darker w-[44px] h-[44px]"
        />

        <p class="font-medium text-center">{{ options.message }}</p>
      </div>
      <div>
        <button
          class="font-medium px-[20px] py-[8px] w-full bg-background-light dark:bg-background-dark text-text-light dark:text-text-dark rounded-e-prompt rounded-b-prompt hover:bg-select-options-background-hovered-light dark:hover:bg-select-options-background-hovered-dark"
          @click="onClose"
        >
          {{ options.closeBtnTxt }}
        </button>
      </div>
    </div>
  </div>
</template>
