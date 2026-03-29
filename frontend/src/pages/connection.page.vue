<script setup lang="ts">
import ConnectionForm from "@/components/forms/ConnectionForm.vue";
</script>

<template>
  <div class="p-6 flex flex-col items-center">
    <div class="text-2xl font-bold mb-[10px]">
      <span class="text text-[#60A5FA]">DB</span>
      <span class="text-primary">Dock</span>
    </div>
    <img id="logo" alt="Wails logo" src="@/assets/images/logo.svg" />
    <ConnectionForm />
  </div>
</template>
