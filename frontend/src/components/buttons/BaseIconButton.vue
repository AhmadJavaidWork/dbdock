<script setup lang="ts">
import { ClassValue, computed } from "vue";

const {
  disabled = false,
  variant = "primary",
  btnClass = "",
  loading = false,
} = defineProps<{
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
  const base = `py-auto px-[7px] font-medium rounded transition-colors duration-200 dark:text-dark dark:hover:opacity-90 disabled:cursor-not-allowed disabled:pointer-events-none`;

  if (variant === "primary") {
    return `${base} text-text-light bg-white border border-primary hover:bg-primary-light dark:text-text-dark dark:border-0 dark:bg-secondary dark:hover:bg-secondary-dark-hover`;
  } else if (variant === "danger") {
    return `${base} text-text-light bg-white border border-danger hover:bg-danger-light dark:text-text-dark dark:bg-danger dark:hover:bg-danger-dark-hover`;
  }

  return base;
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
    :class="[
      buttonClasses,
      btnClass,
      {
        'opacity-70': loading,
      },
    ]"
  >
    <slot />
  </button>
</template>
