<script setup lang="ts">
import { IconEmits, IconProps } from "@/types/common.types";

const { rotate = false } = defineProps<IconProps>();
const emit = defineEmits<IconEmits>();
</script>

<template>
  <svg
    v-bind="$attrs"
    :class="{ 'rotate-180': rotate }"
    fill="none"
    stroke="currentColor"
    viewBox="0 0 24 24"
    @click="emit('click', $event)"
  >
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
  </svg>
</template>
