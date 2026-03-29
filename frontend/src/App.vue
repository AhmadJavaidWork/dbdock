<script setup lang="ts">
import BasePrimaryButton from "@/components/buttons/BasePrimaryButton.vue";
import AppLoader from "@/components/common/AppLoader.vue";
import BasePrompt from "@/components/modals/BasePrompt.vue";
import { useThemeStore } from "@/stores/theme.store";
</script>

<template>
  <AppLoader />
  <BasePrimaryButton @click="useThemeStore().toggleTheme">Toggle Theme</BasePrimaryButton>
  <div class="flex flex-col items-center">
    <RouterView />
    <BasePrompt />
  </div>
</template>
