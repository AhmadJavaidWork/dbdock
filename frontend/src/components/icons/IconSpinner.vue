<script setup lang="ts">
const { color = 'white' } = defineProps<{
  color?: string
}>()
</script>
>

<template>
  <div
    :class="[
      'w-4 h-4 border-2 border-solid border-t-transparent rounded-full animate-spin',
      `border-${color}`
    ]"
    v-bind="$attrs"
  ></div>
</template>
