<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{
  type?: 'button' | 'submit' | 'reset';
  disabled?: boolean;
  variant?: 'primary' | 'danger';
  class?: string;
}>();

const variant = computed(() => props.variant || 'primary');

const emit = defineEmits<{
  (e: 'click', event: MouseEvent): void;
  (e: 'focus', event: FocusEvent): void;
  (e: 'blur', event: FocusEvent): void;
  (e: 'mouseenter', event: MouseEvent): void;
  (e: 'mouseleave', event: MouseEvent): void;
  (e: 'keydown', event: KeyboardEvent): void;
  (e: 'keyup', event: KeyboardEvent): void;
}>();

const buttonClasses = computed(() => {
  const base = `px-[20px] py-[8px] font-medium rounded transition-colors duration-200 dark:text-dark dark:hover:opacity-90`;

  if (variant.value === 'primary') {
    return `${base} text-text-light bg-white border border-primary hover:bg-primary-light dark:text-text-dark dark:bg-primary dark:hover:bg-primary-dark-hover`;
  } else if (variant.value === 'danger') {
    return `${base} text-text-light bg-white border border-danger hover:bg-danger-light dark:text-text-dark dark:bg-danger dark:hover:bg-danger-dark-hover`;
  }

  return base;
});
</script>

<template>
  <button
    v-bind="$attrs"
    @click="emit('click', $event)"
    @focus="emit('focus', $event)"
    @blur="emit('blur', $event)"
    @mouseenter="emit('mouseenter', $event)"
    @mouseleave="emit('mouseleave', $event)"
    @keydown="emit('keydown', $event)"
    @keyup="emit('keyup', $event)"
    :class="[buttonClasses, props.class]"
  >
    <slot />
  </button>
</template>
