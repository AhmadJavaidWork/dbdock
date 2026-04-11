<script setup lang="ts">
import { ClassValue, watch } from "vue";

const {
  isOpen = false,
  title = "",
  titleClase = "",
  bodyClass = "",
  footerClass = "",
} = defineProps<{
  isOpen: boolean;
  title: string;
  titleClase?: ClassValue;
  bodyClass?: ClassValue;
  footerClass?: ClassValue;
}>();

const emit = defineEmits<{
  close: [];
}>();

function handleKeydown(e: KeyboardEvent) {
  if (e.key === "Escape" && isOpen) {
    e.stopImmediatePropagation();
    emit("close");
  }
}

watch(
  () => isOpen,
  function (newVal) {
    if (newVal) {
      window.addEventListener("keydown", handleKeydown, true);
    } else {
      window.removeEventListener("keydown", handleKeydown, true);
    }
  }
);
</script>

<template>
  <transition name="modal-fade">
    <div
      v-if="isOpen"
      class="md:flex fixed inset-0 flex items-center justify-center bg-black/50 z-[1000]"
      @click="emit('close')"
    >
      <div
        class="text-text-light bg-background-light dark:bg-background-dark dark:text-text-dark rounded-prompt shadow-lg max-w-md"
        @click.stop
      >
        <div :class="['px-[30px] py-4', titleClase]">
          <h3 class="text-lg font-semibold">{{ title }}</h3>
        </div>
        <div
          :class="[
            'px-[30px] py-4 border-t border-b border-textfield-border-light dark:border-textfield-border-dark',
            bodyClass,
          ]"
        >
          <slot name="modal-body"></slot>
        </div>
        <div :class="['px-[30px] py-4', footerClass]">
          <slot name="modal-footer"></slot>
        </div>
      </div>
    </div>
  </transition>
</template>

<style scoped>
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.2s ease;
}
.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}
</style>
