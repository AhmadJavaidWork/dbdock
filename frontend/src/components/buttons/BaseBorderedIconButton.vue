<script setup lang="ts">
import { ClassValue, ref, Ref } from "vue";

const {
  disabled = false,
  btnClass = "",
  loading = false,
} = defineProps<{
  disabled?: boolean;
  variant?: "primary" | "danger";
  btnClass?: ClassValue;
  loading?: boolean;
}>();

const btnRef = ref<HTMLButtonElement | null>(null);

const emit = defineEmits<{
  (e: "click", event: MouseEvent): void;
  (e: "focus", event: FocusEvent): void;
  (e: "blur", event: FocusEvent): void;
  (e: "mouseenter", event: MouseEvent): void;
  (e: "mouseleave", event: MouseEvent): void;
  (e: "keydown", event: KeyboardEvent): void;
  (e: "keyup", event: KeyboardEvent): void;
}>();

defineExpose<{ btnRef: Ref<HTMLButtonElement | null> }>({ btnRef });
</script>

<template>
  <button
    ref="btnRef"
    v-bind="$attrs"
    :disabled="disabled"
    @click="emit('click', $event)"
    @focus="emit('focus', $event)"
    @blur="emit('blur', $event)"
    @mouseenter="emit('mouseenter', $event)"
    @mouseleave="emit('mouseleave', $event)"
    @keydown="emit('keydown', $event)"
    @keyup="emit('keyup', $event)"
    :class="[
      'w-[40px] h-[40px] flex items-center justify-center font-medium rounded-full transition-colors duration-200 dark:hover:opacity-90 dark:focus-visible:opacity-90 disabled:cursor-not-allowed disabled:pointer-events-none text-text-light bg-white border-2 border-primary hover:bg-primary-light focus-visible:bg-primary-light dark:text-text-dark dark:border-0 dark:bg-secondary dark:hover:bg-secondary-dark-hover dark:focus-visible:bg-secondary-dark-hover',
      btnClass,
      {
        'opacity-70': loading,
      },
    ]"
  >
    <slot />
  </button>
</template>
