<script setup lang="ts">
import { IconEmits } from "@/types/common.types";

const emit = defineEmits<IconEmits>();
</script>

<template>
  <svg v-bind="$attrs" fill="currentColor" viewBox="0 0 24 24">
    @click="emit('click', $event)"
    <path
      d="M12 2C6.5 2 2 6.5 2 12S6.5 22 12 22 22 17.5 22 12 17.5 2 12 2M10 17L5 12L6.41 10.59L10 14.17L17.59 6.58L19 8L10 17Z"
    ></path>
  </svg>
</template>
