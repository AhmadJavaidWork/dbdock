<script setup lang="ts">
import { IconEmits } from "@/types/common.types";

const emit = defineEmits<IconEmits>();
</script>

<template>
  <svg v-bind="$attrs" fill="currentColor" viewBox="0 0 24 24" @click="emit('click', $event)">
    <path d="M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z"></path>
  </svg>
</template>
