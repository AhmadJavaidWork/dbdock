<script setup lang="ts">
import { ClassValue, computed } from "vue";

const {
  type = "button",
  disabled = false,
  variant = "primary",
  btnClass = "",
  loading = false,
} = defineProps<{
  type?: "button" | "submit" | "reset";
  disabled?: boolean;
  variant?: "primary" | "danger";
  btnClass?: ClassValue;
  loading?: boolean;
}>();

const emit = defineEmits<{
  (e: "click", event: MouseEvent): void;
  (e: "focus", event: FocusEvent): void;
  (e: "blur", event: FocusEvent): void;
  (e: "mouseenter", event: MouseEvent): void;
  (e: "mouseleave", event: MouseEvent): void;
  (e: "keydown", event: KeyboardEvent): void;
  (e: "keyup", event: KeyboardEvent): void;
}>();

const buttonClasses = computed(() => {
  const base = `px-[20px] py-[8px] font-medium rounded transition-colors duration-200 text-white dark:text-gray-100 disabled:cursor-not-allowed disabled:pointer-events-none`;

  if (variant === "primary") {
    return `${base} bg-primary hover:bg-primary-hover dark:bg-primary-dark dark:hover:bg-primary-dark-hover`;
  } else if (variant === "danger") {
    return `${base} bg-danger hover:bg-danger-hover dark:bg-danger-dark dark:hover:bg-danger-darker`;
  }

  return `${base} bg-gray-200 hover:bg-gray-300 dark:bg-gray-700 dark:hover:bg-gray-600`;
});
</script>

<template>
  <button
    v-bind="$attrs"
    :disabled="disabled"
    @click="emit('click', $event)"
    @focus="emit('focus', $event)"
    @blur="emit('blur', $event)"
    @mouseenter="emit('mouseenter', $event)"
    @mouseleave="emit('mouseleave', $event)"
    @keydown="emit('keydown', $event)"
    @keyup="emit('keyup', $event)"
    :class="[buttonClasses, btnClass, { 'opacity-70': loading }]"
  >
    <slot />
  </button>
</template>
