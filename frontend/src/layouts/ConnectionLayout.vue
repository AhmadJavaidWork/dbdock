<script setup lang="ts">
import TableList from "@/components/sidebars/TableList.vue";
</script>

<template>
  <div
    class="h-screen flex bg-background-light dark:bg-background-dark text-text-light dark:text-text-dark"
  >
    <TableList />
    <RouterView />
  </div>
</template>
