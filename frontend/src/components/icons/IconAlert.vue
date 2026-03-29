<script setup lang="ts">
import { IconEmits } from "@/types/common.types";

const emit = defineEmits<IconEmits>();
</script>

<template>
  <svg fill="currentColor" v-bind="$attrs" @click="emit('click', $event)" viewBox="0 0 24 24">
    <path d="M11 7h2v6h-2zM11 15h2v2h-2z"></path>
  </svg>
</template>
