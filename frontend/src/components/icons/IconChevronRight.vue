<script setup lang="ts">
import { IconEmits } from "@/types/common.types";

const emit = defineEmits<IconEmits>();
</script>

<template>
  <svg v-bind="$attrs" fill="currentColor" @click="emit('click', $event)" viewBox="0 0 24 24">
    <path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z"></path>
  </svg>
</template>
