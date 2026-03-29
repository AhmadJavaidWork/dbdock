<script setup lang="ts">
import { IconEmits } from "@/types/common.types";

const emit = defineEmits<IconEmits>();
</script>

<template>
  <svg fill="currentColor" v-bind="$attrs" @click="emit('click', $event)" viewBox="0 0 24 24">
    <path d="M11 9H13V7H11V9M11 17H13V11H11V17Z"></path>
  </svg>
</template>
