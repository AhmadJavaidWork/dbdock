<script setup lang="ts">
import BasePrimaryButton from "@/components/buttons/BasePrimaryButton.vue";
import BasePrompt from "@/components/modals/BasePrompt.vue";
import { useThemeStore } from "@/stores/theme.store";
</script>

<template>
  <BasePrimaryButton @click="useThemeStore().toggleTheme">Toggle Theme</BasePrimaryButton>
  <div class="flex flex-col items-center">
    <img id="logo" alt="Wails logo" src="@/assets/images/logo.svg" class="" />
    <RouterView />
    <BasePrompt />
  </div>
</template>
